package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 复制全局正则表达式定义
var (
	reCiteFull = regexp.MustCompile(`\[cite_start\][\s\S]*?\[cite_end\]`)
	reCiteAny  = regexp.MustCompile(`\[cite[:\s_][^\]]*\]`)
	reRef      = regexp.MustCompile(`\[\d+\]`)
	reHeading  = regexp.MustCompile(`(?m)^#{1,6} `)
	reMath     = regexp.MustCompile(`(?m)(^\$\$[\s\S]*?\$\$)`)
	reMultiNL  = regexp.MustCompile(`\n{3,}`)
	// 移除AI助理的開頭道歉及說明段落
	reAIIntro = regexp.MustCompile(`^[\s\S]*?非常抱歉[\s\S]*?PDF\s*即可。\s*\n*---\s*\n+`)
	// 移除結尾的AI道歉和說明
	reAIOutro = regexp.MustCompile(`\n*再次為[\s\S]*?請隨時跟我說[!！]\s*$`)
)

func removeCiteTags(content string) string {
	content = reCiteFull.ReplaceAllString(content, "")
	content = reCiteAny.ReplaceAllString(content, "")
	content = reRef.ReplaceAllString(content, "")
	return content
}

func removeDialogue(content string) string {
	// 移除AI助理的開頭介紹
	content = reAIIntro.ReplaceAllString(content, "")

	// 移除"### 📄 檔案一..."這類指示行
	reFileIndicator := regexp.MustCompile(`(?m)^###\s*📄\s*.*?\.pdf[\s\S]*?報告中\)\s*\n+`)
	content = reFileIndicator.ReplaceAllString(content, "")

	// 移除結尾的AI道歉和說明
	fmt.Println("[DEBUG] 应用reAIOutro之前...")
	content = reAIOutro.ReplaceAllString(content, "")
	fmt.Println("[DEBUG] 应用reAIOutro之后")

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

func extractTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "Unknown Course—Homework"
}

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

func main() {
	// 读取原始文件
	content, err := os.ReadFile("homework.md")
	if err != nil {
		fmt.Println("❌ 讀取檔案失敗:", err)
		return
	}

	text := string(content)
	
	// 调用preprocessContent来处理
	cleanedContent := preprocessContent(content)
	
	// 寫入清理後的檔案
	outputFile := "homework_cleaned.md"
	err = os.WriteFile(outputFile, []byte(cleanedContent), 0644)
	if err != nil {
		fmt.Println("❌ 寫入檔案失敗:", err)
		return
	}
	
	// 統計結果
	citeFull := len(reCiteFull.FindAllString(text, -1))
	citeAny := len(reCiteAny.FindAllString(text, -1))
	totalCite := citeFull + citeAny
	
	citeAnyAfter := len(reCiteAny.FindAllString(cleanedContent, -1))
	
	fmt.Println("\n✅ 清理完成！")
	fmt.Println("=== 移除統計 ===")
	fmt.Printf("原始檔案: %s (%d 字元)\n", "homework.md", len(text))
	fmt.Printf("清理後檔案: %s (%d 字元)\n", outputFile, len(cleanedContent))
	fmt.Printf("移除字元數: %d\n", len(text)-len(cleanedContent))
	fmt.Printf("\n移除的 cite 標籤數:\n")
	fmt.Printf("  - [cite_start/end] 塊: %d\n", citeFull)
	fmt.Printf("  - [cite:...] 標籤: %d\n", citeAny)
	fmt.Printf("  - 共計: %d 個\n", totalCite)
	
	if citeAnyAfter > 0 {
		fmt.Printf("\n⚠️ 警告: 清理後仍有 %d 個 cite 標籤未被移除\n", citeAnyAfter)
	} else {
		fmt.Println("\n✓ 所有 cite 標籤已成功移除！")
	}
}
