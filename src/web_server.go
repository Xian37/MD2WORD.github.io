package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// 全局正則表達式定義
var (
	reCiteFull = regexp.MustCompile(`\[cite_start\][\s\S]*?\[cite_end\]`)
	reCiteAny  = regexp.MustCompile(`\[cite[:\s_][^\]]*\]`)
	reRef      = regexp.MustCompile(`\[\d+\]`)
	reHeading  = regexp.MustCompile(`(?m)^#{1,6} `)
)

// 轉換請求結構
type ConvertRequest struct {
	Content  string `json:"content"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Filename string `json:"filename"`
}

// 轉換響應結構
type ConvertResponse struct {
	Success      bool   `json:"success"`
	Message      string `json:"message"`
	OutputFile   string `json:"outputFile"`
	CiteCount    int    `json:"citeCount"`
	OriginalSize int    `json:"originalSize"`
	CleanedSize  int    `json:"cleanedSize"`
}

// removeCiteTags 移除所有 cite 標籤
func removeCiteTags(content string) string {
	content = reCiteFull.ReplaceAllString(content, "")
	content = reCiteAny.ReplaceAllString(content, "")
	content = reRef.ReplaceAllString(content, "")
	return content
}

// extractTitle 從內容中提取標題
func extractTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "文檔"
}

// preprocessContent 預處理內容，移除非報告部分
func preprocessContent(content string) string {
	lines := strings.Split(content, "\n")
	var result []string
	var inReport bool

	// 常見的報告開始標記
	reportMarkers := []string{"#", "##", "###", "1.", "1)", "(1)"}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳過空行在報告開始前
		if !inReport && line == "" {
			continue
		}

		// 檢查是否開始報告內容
		for _, marker := range reportMarkers {
			if strings.HasPrefix(line, marker) && strings.ContainsAny(marker, "#12()") {
				inReport = true
				break
			}
		}

		// 跳過常見的非報告文本
		if strings.Contains(line, "Deep dive") || strings.Contains(line, "In-depth analysis") ||
			strings.Contains(line, "此為學術報告") || strings.Contains(line, "DUE DATE") ||
			strings.Contains(line, "Due date") || strings.Contains(line, "Abstract:") && len(line) < 50 {
			continue
		}

		if inReport || strings.HasPrefix(line, "#") {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

// addTitleWithDate 添加標題和日期
func addTitleWithDate(content string, title string, date string) string {
	if strings.HasPrefix(content, "# ") {
		return content
	}

	if title == "" {
		title = extractTitle(content)
	}

	if date == "" {
		now := time.Now()
		date = fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())
	}

	titleSection := fmt.Sprintf("# %s\n\n**DATE：%s**\n\n---\n\n", title, date)
	return titleSection + content
}

// handleConvert 處理轉換請求
func handleConvert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ConvertRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 計算 cite 標籤數量
	citeMatches := reCiteAny.FindAllString(req.Content, -1)
	citeCount := len(citeMatches)

	originalSize := len(req.Content)

	// 清理內容
	cleanedContent := removeCiteTags(req.Content)
	// 預處理，移除非報告部分
	cleanedContent = preprocessContent(cleanedContent)
	cleanedSize := len(cleanedContent)

	// 添加標題和日期
	if req.Title == "" {
		req.Title = extractTitle(cleanedContent)
		if req.Title == "文檔" {
			// 從文件名提取標題
			filename := strings.TrimSuffix(req.Filename, ".md")
			req.Title = filename
		}
	}

	processedContent := addTitleWithDate(cleanedContent, req.Title, req.Date)

	// 創建 output 資料夾
	outputDir := "output"
	os.MkdirAll(outputDir, 0755)

	// 生成輸出文件名
	timestamp := time.Now().Unix()
	baseName := strings.TrimSuffix(req.Filename, ".md")
	tempMdFilename := fmt.Sprintf("%s_%d_temp.md", baseName, timestamp)
	outputFilename := fmt.Sprintf("%s_%d.docx", baseName, timestamp)
	tempMdPath := filepath.Join(outputDir, tempMdFilename)
	outputPath := filepath.Join(outputDir, outputFilename)

	// 寫入處理後的臨時 Markdown 文件
	err = os.WriteFile(tempMdPath, []byte(processedContent), 0644)
	if err != nil {
		response := ConvertResponse{
			Success: false,
			Message: fmt.Sprintf("寫入臨時文件失敗: %v", err),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// 將臨時 Markdown 轉換為 DOCX
	err = convertMarkdownToDocx(tempMdPath, outputPath)
	if err != nil {
		os.Remove(tempMdPath)
		response := ConvertResponse{
			Success: false,
			Message: fmt.Sprintf("轉換為 DOCX 失敗: %v", err),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	os.Remove(tempMdPath)

	// 設置 DOCX 字體
	err = setDocxFonts(outputPath)
	if err != nil {
		// 字體設置失敗不中斷流程，只記錄
		fmt.Printf("⚠️  字體設置失敗: %v\n", err)
	}

	// 返回成功響應
	response := ConvertResponse{
		Success:      true,
		Message:      "轉換成功",
		OutputFile:   outputPath,
		CiteCount:    citeCount,
		OriginalSize: originalSize,
		CleanedSize:  cleanedSize,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func convertMarkdownToDocx(inputFile, outputFile string) error {
	absInputFile, err := filepath.Abs(inputFile)
	if err != nil {
		return fmt.Errorf("無法解析輸入路徑: %v", err)
	}

	absOutputFile, err := filepath.Abs(outputFile)
	if err != nil {
		return fmt.Errorf("無法解析輸出路徑: %v", err)
	}

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	md2wordPath := filepath.Join(filepath.Dir(exePath), "md2word.exe")
	if _, err := os.Stat(md2wordPath); os.IsNotExist(err) {
		md2wordPath = "md2word.exe"
		if _, err := os.Stat(md2wordPath); os.IsNotExist(err) {
			return fmt.Errorf("找不到 md2word.exe")
		}
	}

	cmd := exec.Command(md2wordPath, "-input", absInputFile)
	cmd.Dir = filepath.Dir(md2wordPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("轉換失敗: %v\n%s", err, string(output))
	}

	docxFile := strings.TrimSuffix(absInputFile, filepath.Ext(absInputFile)) + ".docx"
	if docxFile != absOutputFile {
		if err := os.Rename(docxFile, absOutputFile); err != nil {
			return fmt.Errorf("重命名 DOCX 失敗: %v", err)
		}
	}

	return nil
}

// setDocxFonts 設置 DOCX 字體離
// \u4e2d文：標楷體 (KaiTi)
// \u82f1\u6587：imes New Roman
func setDocxFonts(docxPath string) error {
	// \u6a94案不\u5b58\u5728或\u4e0d\u662f DOCX
	if _, err := os.Stat(docxPath); err != nil {
		return fmt.Errorf("DOCX \u6a94\u6848\u4e0d\u5b58\u5728: %v", err)
	}

	// \u958b\u555f DOCX \u4f5c爲 ZIP
	zipReader, err := zip.OpenReader(docxPath)
	if err != nil {
		return fmt.Errorf("\u7121\u6cd5\u6253\u958b DOCX: %v", err)
	}
	defer zipReader.Close()

	// \u65b0\u5efa\u4e00\u500b \u6df7\u8a18\u7684\u5185\u5bb9
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, file := range zipReader.File {
		contents, err := file.Open()
		if err != nil {
			zipWriter.Close()
			return err
		}

		data, _ := io.ReadAll(contents)
		contents.Close()

		// \u53ea\u4fee\u6539 document.xml
		if file.Name == "word/document.xml" {
			data = modifyDocumentFonts(data)
		}

		// \u5beb\u5165\u65b0 ZIP
		w, err := zipWriter.Create(file.Name)
		if err != nil {
			zipWriter.Close()
			return err
		}
		w.Write(data)
	}

	zipWriter.Close()

	// \u5099\u4efd\u7136\u4f8b\u6beb\u5c40\u6246\u6b76
	backupPath := docxPath + ".backup"
	os.Rename(docxPath, backupPath)

	// \u5beb\u5165\u65b0 DOCX
	err = os.WriteFile(docxPath, buf.Bytes(), 0644)
	if err != nil {
		os.Rename(backupPath, docxPath)
		return fmt.Errorf("\u5beb\u5165 DOCX \u5931\u6557: %v", err)
	}

	os.Remove(backupPath)
	return nil
}

// modifyDocumentFonts 修改 document.xml 中的字體設置
func modifyDocumentFonts(data []byte) []byte {
	docContent := string(data)

	// 替換所有 Calibri 為標楷體（中文）和 Times New Roman（英文）
	// Word 字體屬性通常包含: rFonts, ascii, hAnsi, eastAsia, cs

	// 替換基礎字體屬性
	docContent = strings.ReplaceAll(docContent, `"Calibri"`, `"標楷體"`)

	// 然後分別設置不同類型的文字
	docContent = strings.ReplaceAll(docContent, `w:ascii="標楷體"`, `w:ascii="Times New Roman"`)
	docContent = strings.ReplaceAll(docContent, `w:hAnsi="標楷體"`, `w:hAnsi="Times New Roman"`)
	docContent = strings.ReplaceAll(docContent, `w:eastAsia="標楷體"`, `w:eastAsia="標楷體"`)

	fmt.Println("✅ 字體修改已應用: 中文→標楷體, 英文→Times New Roman")
	return []byte(docContent)
}

// handleDownload 下載文件
func handleDownload(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		http.Error(w, "File not specified", http.StatusBadRequest)
		return
	}

	// 檢查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// 設置響應頭
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	w.Header().Set("Content-Type", "application/octet-stream")

	// 發送文件
	http.ServeFile(w, r, filePath)
}

// handleCORS 處理 CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// 路由設置
	mux := http.NewServeMux()

	// API 端點
	mux.HandleFunc("/api/convert", handleConvert)
	mux.HandleFunc("/api/download", handleDownload)

	// 靜態文件服務
	mux.Handle("/", http.FileServer(http.Dir(".")))

	// 應用 CORS 中間件
	handler := corsMiddleware(mux)

	// 啟動服務器
	port := ":8888"
	fmt.Printf("🚀 伺服器啟動在 http://localhost%s\n", port)
	fmt.Printf("📂 請在浏覽器中打開 http://localhost%s/web_gui.html\n", port)

	err := http.ListenAndServe(port, handler)
	if err != nil {
		fmt.Printf("❌ 伺服器錯誤: %v\n", err)
		os.Exit(1)
	}
}
