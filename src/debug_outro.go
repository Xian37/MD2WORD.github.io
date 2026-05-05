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
	
	// 显示最后500个字符
	runes := []rune(text)
	fmt.Println("=== 文件最后500个字符 ===")
	if len(runes) > 500 {
		fmt.Println(string(runes[len(runes)-500:]))
	} else {
		fmt.Println(text)
	}
	
	fmt.Println("\n=== 尝试匹配结尾对话 ===")
	
	// 测试当前的正则
	reAIOutro := regexp.MustCompile(`\n*再次為[\s\S]*?請隨時跟我說[!！]\s*$`)
	matches := reAIOutro.FindAllString(text, -1)
	fmt.Printf("当前正则匹配数: %d\n", len(matches))
	if len(matches) > 0 {
		fmt.Printf("匹配内容长度: %d\n", len(matches[0]))
	}
	
	// 尝试更强大的正则
	fmt.Println("\n=== 尝试更强大的正则 ===")
	reAIOutro2 := regexp.MustCompile(`再次為[\s\S]*?請隨時跟我說[!！]`)
	matches2 := reAIOutro2.FindAllString(text, -1)
	fmt.Printf("更强大正则匹配数: %d\n", len(matches2))
	if len(matches2) > 0 {
		fmt.Printf("匹配内容: %s\n", matches2[0])
	}
}
