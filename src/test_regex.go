package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	// 读取原始文件
	content, err := os.ReadFile("homework.md")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)
	fmt.Println("=== 原始文件长度 ===")
	fmt.Println("字符数:", len(text))
	fmt.Println()

	// 测试 reAIIntro
	reAIIntro := regexp.MustCompile(`^[\s\S]*?非常抱歉[\s\S]*?PDF\s*即可。\s*\n*---\s*\n+`)
	matches := reAIIntro.FindAllString(text, -1)
	fmt.Println("=== reAIIntro 匹配结果 ===")
	fmt.Println("匹配数:", len(matches))
	if len(matches) > 0 {
		fmt.Println("第一个匹配长度:", len(matches[0]))
		fmt.Println("前100个字符:", string([]rune(matches[0])[:min(100, len([]rune(matches[0])))]))
	}
	fmt.Println()

	// 测试 reAIOutro
	reAIOutro := regexp.MustCompile(`\n*再次為[\s\S]*?請隨時跟我說[!！]\s*$`)
	matches2 := reAIOutro.FindAllString(text, -1)
	fmt.Println("=== reAIOutro 匹配结果 ===")
	fmt.Println("匹配数:", len(matches2))
	if len(matches2) > 0 {
		fmt.Println("第一个匹配长度:", len(matches2[0]))
		fmt.Println("最后100个字符:", string([]rune(matches2[0])[max(0, len([]rune(matches2[0]))-100):]))
	}
	fmt.Println()

	// 应用移除
	text = reAIIntro.ReplaceAllString(text, "")

	// 移除"### 📄 檔案一..."这类指示行
	reFileIndicator := regexp.MustCompile(`(?m)^###\s*📄\s*.*?\.pdf[\s\S]*?報告中\)\s*\n+`)
	text = reFileIndicator.ReplaceAllString(text, "")

	text = reAIOutro.ReplaceAllString(text, "")

	fmt.Println("=== 处理后的文件长度 ===")
	fmt.Println("字符数:", len(text))
	fmt.Println()

	fmt.Println("=== 处理后的前500个字符 ===")
	if len(text) > 500 {
		fmt.Println(text[:500])
	} else {
		fmt.Println(text)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
