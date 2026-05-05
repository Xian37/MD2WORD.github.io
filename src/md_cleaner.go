package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 全局正則表達式定義
var (
	reCiteFull    = regexp.MustCompile(`\[cite_start\][\s\S]*?\[cite_end\]`)
	reCiteAny     = regexp.MustCompile(`\[cite[:\s_][^\]]*\]`)
	reRef         = regexp.MustCompile(`\[\d+\]`)
	reHeading     = regexp.MustCompile(`(?m)^#{1,6} `)
	reMath        = regexp.MustCompile(`(?m)(^\$\$[\s\S]*?\$\$)`)
	reMultiNL     = regexp.MustCompile(`\n{3,}`)
	// 移除AI助理的開頭道歉及說明段落
	reAIIntro     = regexp.MustCompile(`^[\s\S]*?非常抱歉[\s\S]*?PDF\s*即可。\s*\n*---\s*\n+`)
	// 移除結尾的AI道歉和說明
	reAIOutro     = regexp.MustCompile(`\n*再次為[\s\S]*?請隨時跟我說[!！]\s*$`)
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
	content = reAIIntro.ReplaceAllString(content, "")

	// 移除"### 📄 檔案一..."這類指示行
	reFileIndicator := regexp.MustCompile(`(?m)^###\s*📄\s*.*?\.pdf[\s\S]*?報告中\)\s*\n+`)
	content = reFileIndicator.ReplaceAllString(content, "")

	// 移除結尾的AI道歉和說明
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

// preprocessContent 進行完整的前處理
func preprocessContent(input []byte) string {
	content := string(input)

	// 移除引用標籤
	content = removeCiteTags(content)

	// 移除不該有的對答
	content = removeDialogue(content)

	// 從第一個標題開始保留
	loc := reHeading.FindStringIndex(content)
	if loc != nil {
		content = content[loc[0]:]
	}

	return strings.TrimSpace(content)
}

// cleanMarkdownFile 清理Markdown檔案
func cleanMarkdownFile(inputFile, outputFile string) error {
	// 讀取輸入檔案
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("讀取檔案失敗: %v", err)
	}

	text := string(content)
	originalLen := len(text)

	// 進行前處理
	cleanedContent := preprocessContent(content)
	cleanedLen := len(cleanedContent)

	// 計算移除統計
	citeFull := len(reCiteFull.FindAllString(text, -1))
	citeAny := len(reCiteAny.FindAllString(text, -1))
	totalCite := citeFull + citeAny

	// 驗證清理後沒有cite標籤
	citeAnyAfter := len(reCiteAny.FindAllString(cleanedContent, -1))

	// 寫入輸出檔案
	err = os.WriteFile(outputFile, []byte(cleanedContent), 0644)
	if err != nil {
		return fmt.Errorf("寫入檔案失敗: %v", err)
	}

	// 顯示統計結果
	fmt.Println("✅ 清理完成！")
	fmt.Println("=== 檔案資訊 ===")
	fmt.Printf("輸入: %s (%d 字元)\n", filepath.Base(inputFile), originalLen)
	fmt.Printf("輸出: %s (%d 字元)\n", filepath.Base(outputFile), cleanedLen)
	fmt.Printf("移除: %d 字元 (%.1f%%)\n", originalLen-cleanedLen, float64(originalLen-cleanedLen)/float64(originalLen)*100)
	
	fmt.Println("\n=== 移除的 cite 標籤 ===")
	fmt.Printf("  - [cite_start/end] 塊: %d\n", citeFull)
	fmt.Printf("  - [cite:...] 標籤: %d\n", citeAny)
	fmt.Printf("  - 共計: %d 個\n", totalCite)

	if citeAnyAfter > 0 {
		fmt.Printf("\n⚠️ 警告: 清理後仍有 %d 個 cite 標籤未被移除\n", citeAnyAfter)
	} else {
		fmt.Println("\n✓ 驗證: 所有 cite 標籤已成功移除！")
	}

	return nil
}

func main() {
	// 定義命令行標誌
	inputFile := flag.String("i", "homework.md", "輸入 Markdown 檔案路徑")
	outputFile := flag.String("o", "homework_cleaned.md", "輸出 Markdown 檔案路徑")
	inplace := flag.Bool("inplace", false, "直接修改原檔案（慎用）")
	flag.Parse()

	// 如果指定了 -inplace，則覆蓋原檔案
	if *inplace {
		*outputFile = *inputFile
	}

	// 檢查輸入檔案是否存在
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		fmt.Printf("❌ 錯誤: 找不到檔案 %s\n", *inputFile)
		os.Exit(1)
	}

	// 執行清理
	if err := cleanMarkdownFile(*inputFile, *outputFile); err != nil {
		fmt.Printf("❌ 錯誤: %v\n", err)
		os.Exit(1)
	}
}
