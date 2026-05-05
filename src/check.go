package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ 請先設定 GEMINI_API_KEY 環境變數")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("❌ 無法初始化客戶端: %v", err)
	}
	defer client.Close()

	fmt.Println("🔍 正在連線至 Google 伺服器，查詢您專屬的可用模型清單...")

	iter := client.ListModels(ctx)
	found := false
	for {
		m, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("❌ 查詢失敗: %v", err)
		}

		// 我們只需要支援文字生成 (generateContent) 的模型
		supportedMethods := strings.Join(m.SupportedGenerationMethods, ", ")
		if strings.Contains(supportedMethods, "generateContent") {
			// Google 回傳的名稱會帶有 "models/" 前綴，我們將其去掉
			name := strings.TrimPrefix(m.Name, "models/")
			fmt.Printf("✅ 發現可用模型: \"%s\"\n", name)
			found = true
		}
	}

	if !found {
		fmt.Println("⚠️ 您的 API 金鑰目前似乎沒有任何支援文字生成的模型權限。")
	} else {
		fmt.Println("\n🎯 解決方案：請複製上面其中一個綠色打勾的「模型名稱」，貼回 main.go 中替換！")
	}
}
