package main

import (
	"flag"
	"fmt"
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

// removeCiteTags 移除所有 cite 標籤
func removeCiteTags(content string) string {
	content = reCiteFull.ReplaceAllString(content, "")
	content = reCiteAny.ReplaceAllString(content, "")
	content = reRef.ReplaceAllString(content, "")
	return content
}

// removeDialogue 移除不必要的對話文本
func removeDialogue(content string) string {
	// 移除AI助理的開頭介紹
	reAIIntro := regexp.MustCompile(`^[\s\S]*?非常抱歉[\s\S]*?PDF\s*即可。\s*\n*---\s*\n+`)
	content = reAIIntro.ReplaceAllString(content, "")

	// 移除"### 📄 檔案一..."這類指示行
	reFileIndicator := regexp.MustCompile(`(?m)^###\s*📄\s*.*?\.pdf[\s\S]*?報告中\)\s*\n+`)
	content = reFileIndicator.ReplaceAllString(content, "")

	// 移除結尾的AI道歉和說明
	reAIOutro := regexp.MustCompile(`\n*再次為[\s\S]*?請隨時跟我說[!！]\s*$`)
	content = reAIOutro.ReplaceAllString(content, "")

	// 移除MATLAB放置指示
	reMatLabNote := regexp.MustCompile(`\*\(請在此處貼上[\s\S]*?\)\*`)
	content = reMatLabNote.ReplaceAllString(content, "")

	// 移除檔案說明（粗體的檔案指示）
	reCopyNote := regexp.MustCompile(`\*\*[\s\S]*?HW[\d][\s\S]*?\*\*`)
	content = reCopyNote.ReplaceAllString(content, "")

	// 清理因移除導致的多餘空行
	reExtraBlank := regexp.MustCompile(`\n{3,}`)
	content = reExtraBlank.ReplaceAllString(content, "\n\n")

	return strings.TrimSpace(content)
}

// extractTitle 從內容中提取標題
func extractTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "學術報告"
}

// addTitleWithDate 在內容開始添加標題和日期
func addTitleWithDate(content string, title string, date string) string {
	// 檢查是否已有頂級標題
	if strings.HasPrefix(content, "# ") {
		return content
	}

	// 如果沒有提供標題，自動提取或使用默認值
	if title == "" {
		title = extractTitle(content)
	}

	// 格式化日期為中文格式：2026年4月27日
	if date == "" {
		now := time.Now()
		date = fmt.Sprintf("%d年%d月%d日", now.Year(), now.Month(), now.Day())
	}

	// 組成新的標題部分
	titleSection := fmt.Sprintf("# %s\n\n**日期：%s**\n\n---\n\n", title, date)
	return titleSection + content
}

// preprocessContent 進行完整的前處理
func preprocessContent(input []byte) string {
	content := string(input)

	// 移除引用標籤
	content = removeCiteTags(content)

	// 移除不該有的對答
	content = removeDialogue(content)

	// 從第一個標題開始保留（如果有的話）
	loc := reHeading.FindStringIndex(content)
	if loc != nil {
		content = content[loc[0]:]
	}

	return strings.TrimSpace(content)
}

// convertToWord 調用 md2word.exe 進行轉換
func convertToWord(mdFile string) (string, error) {
	// 獲取md2word.exe的路徑
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	md2wordPath := filepath.Join(filepath.Dir(exePath), "md2word.exe")

	// 檢查md2word.exe是否存在
	if _, err := os.Stat(md2wordPath); os.IsNotExist(err) {
		// 嘗試在當前目錄查找
		md2wordPath = "md2word.exe"
		if _, err := os.Stat(md2wordPath); os.IsNotExist(err) {
			return "", fmt.Errorf("找不到 md2word.exe")
		}
	}

	// 執行md2word轉換
	cmd := exec.Command(md2wordPath, "-input", mdFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("轉換失敗: %v\n%s", err, string(output))
	}

	fmt.Println(string(output))

	// 確定輸出檔案名稱
	ext := filepath.Ext(mdFile)
	docxFile := strings.TrimSuffix(mdFile, ext) + ".docx"
	return docxFile, nil
}

// convertMarkdownToWord 完整的轉換流程
func convertMarkdownToWord(inputFile, outputFile, title, date string) error {
	// 讀取輸入檔案
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("讀取檔案失敗: %v", err)
	}

	text := string(content)
	originalLen := len(text)

	fmt.Println("🔄 開始轉換流程...")
	fmt.Printf("📥 輸入檔案: %s (%d 字元)\n", filepath.Base(inputFile), originalLen)

	// 步驟1: 前處理（移除cite標籤等）
	fmt.Println("\n[1/4] 正在清理Markdown（移除cite標籤）...")
	cleanedContent := preprocessContent(content)
	cleanedLen := len(cleanedContent)

	citeFull := len(reCiteFull.FindAllString(text, -1))
	citeAny := len(reCiteAny.FindAllString(text, -1))
	totalCite := citeFull + citeAny

	fmt.Printf("      ✓ 移除了 %d 個 cite 標籤\n", totalCite)
	fmt.Printf("      ✓ 字元數: %d → %d (%d 字元)\n", originalLen, cleanedLen, cleanedLen-originalLen)

	// 步驟2: 添加標題和日期
	fmt.Println("\n[2/4] 正在添加標題和日期...")
	contentWithTitle := addTitleWithDate(cleanedContent, title, date)

	// 使用date參數（如果為空則自動生成）
	if date == "" {
		now := time.Now()
		date = fmt.Sprintf("%d年%d月%d日", now.Year(), now.Month(), now.Day())
	}
	fmt.Printf("      ✓ 標題: %s\n", title)
	fmt.Printf("      ✓ 日期: %s\n", date)

	// 步驟3: 寫入臨時Markdown檔案
	fmt.Println("\n[3/4] 正在轉換為Word...")
	tempMdFile := strings.TrimSuffix(outputFile, ".docx") + "_temp.md"
	err = os.WriteFile(tempMdFile, []byte(contentWithTitle), 0644)
	if err != nil {
		return fmt.Errorf("寫入臨時檔案失敗: %v", err)
	}

	// 步驟4: 轉換為Word
	docxFile, err := convertToWord(tempMdFile)
	if err != nil {
		os.Remove(tempMdFile)
		return err
	}

	// 如果需要重命名輸出檔案
	if docxFile != outputFile {
		err = os.Rename(docxFile, outputFile)
		if err != nil {
			return fmt.Errorf("重命名檔案失敗: %v", err)
		}
	}

	// 清理臨時檔案
	os.Remove(tempMdFile)

	// 顯示完成資訊
	fileInfo, _ := os.Stat(outputFile)
	fmt.Printf("\n✅ 轉換完成！\n")
	fmt.Printf("📤 輸出檔案: %s (%d bytes)\n", filepath.Base(outputFile), fileInfo.Size())

	return nil
}

func main() {
	inputFile := flag.String("i", "homework.md", "輸入 Markdown 檔案路徑")
	outputFile := flag.String("o", "homework_final.docx", "輸出 Word 檔案路徑")
	title := flag.String("title", "", "Word文檔標題（若不指定則自動提取）")
	date := flag.String("date", "", "文檔日期（若不指定則使用當前日期）")
	flag.Parse()

	// 檢查輸入檔案是否存在
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		fmt.Printf("❌ 錯誤: 找不到檔案 %s\n", *inputFile)
		os.Exit(1)
	}

	// 執行轉換
	if err := convertMarkdownToWord(*inputFile, *outputFile, *title, *date); err != nil {
		fmt.Printf("❌ 錯誤: %v\n", err)
		os.Exit(1)
	}
}
