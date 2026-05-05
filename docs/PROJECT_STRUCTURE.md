# 📁 項目結構說明

```
md2word/
│
├── 🚀 gui.bat                 【GUI菜單 - 推薦使用】
│                              點擊這個檔案打開視覺化菜單
│
├── 📁 bin/                    【可執行檔】
│   ├── md2word.exe           原始轉換工具
│   ├── md2word_full.exe      ⭐ 完整轉換工具（推薦）
│   ├── md_cleaner.exe        獨立清理工具
│   └── convert.bat           命令行快速轉換
│
├── 📁 src/                    【源代碼】
│   ├── md2word_full.go       完整轉換程序
│   ├── md_cleaner.go         清理程序
│   └── test_*.go             測試文件
│
├── 📁 docs/                   【說明文檔】
│   ├── README.md             📖 項目概述
│   ├── QUICKSTART.md         ⚡ 快速開始
│   ├── COMPLETE_GUIDE.md     📚 完整指南
│   └── CLEANING_GUIDE.md     🔧 工具說明
│
├── 📁 samples/                【樣例檔案】
│   ├── homework.md           樣例 Markdown
│   ├── homework_*.docx       轉換結果範例
│   └── test_*.md, .docx      測試檔案
│
├── 📁 lib/                    【庫檔案】
│   ├── template.docx         Word 範本
│   ├── go.mod                Go 模塊配置
│   └── go.sum                Go 依賴清單
│
└── 📁 output/                 【輸出目錄】
    轉換後的 Word 檔案保存位置
```

## 🎯 快速開始

### 方式一：GUI 菜單（最簡單 ⭐）
1. 雙擊 **gui.bat**
2. 選擇菜單選項
3. 完成！

### 方式二：命令行
```powershell
bin\md2word_full.exe -i your_file.md -o output.docx -title "標題"
```

### 方式三：批次檔
```powershell
bin\convert.bat your_file.md "標題"
```

## 📂 各目錄用途

| 目錄 | 用途 | 內容 |
|------|------|------|
| **bin** | 執行工具 | .exe 和 .bat 文件 |
| **src** | 源代碼 | Go 語言程序 |
| **docs** | 文檔 | 使用說明和指南 |
| **samples** | 範例 | 樣例 Markdown 和轉換結果 |
| **lib** | 庫檔案 | 範本和配置 |
| **output** | 輸出 | 生成的 Word 檔案 |

## 🚀 使用流程

```
step 1: 打開 gui.bat
         ↓
step 2: 選擇 "1. 一鍵轉換"
         ↓
step 3: 輸入 Markdown 檔案路徑
         ↓
step 4: 輸入文檔標題
         ↓
step 5: 等待轉換完成
         ↓
step 6: 在輸出目錄查看結果
```

## 💡 常用操作

### 查看文檔
在 GUI 菜單選擇：選項 4 > 選擇文檔

### 查看樣例
在 GUI 菜單選擇：選項 2 (打開 samples 目錄)

### 轉換檔案
在 GUI 菜單選擇：選項 1 > 輸入檔案和標題

### 編輯源代碼
在 GUI 菜單選擇：選項 5 > 選擇 src 資料夾

## 📖 文檔速查表

| 文檔 | 適合人群 | 內容 |
|-----|---------|------|
| README.md | 所有人 | 項目概述 |
| QUICKSTART.md | 新手 | 30秒快速開始 |
| COMPLETE_GUIDE.md | 進階用戶 | 完整功能說明 |
| CLEANING_GUIDE.md | 開發者 | 工具技術細節 |

## 🔄 工作流程

```
Markdown 檔案（含 cite）
         ↓
   [GUI 菜單]
         ↓
   輸入檔案和標題
         ↓
   點擊轉換
         ↓
自動清理 → 添加標題/日期 → 轉換為 Word
         ↓
生成最終 Word 檔案
```

## ✨ 功能清單

- ✅ 自動移除 cite 標籤
- ✅ 自動添加標題
- ✅ 自動添加日期
- ✅ 支援自訂輸出檔名
- ✅ 保留原始格式
- ✅ 快速轉換

## 🛠️ 常見操作

### 我想轉換一個新的 Markdown 檔案
→ 打開 gui.bat → 選 1 → 輸入檔案路徑

### 我想查看說明文檔
→ 打開 gui.bat → 選 4 → 選擇文檔

### 我想看到轉換結果
→ 打開 gui.bat → 選 3 (打開 output 目錄)

### 我想編輯或編譯源代碼
→ 打開 gui.bat → 選 5 → 打開 src 資料夾

## 📋 檔案大小參考

| 檔案 | 大小 |
|-----|------|
| md2word_full.exe | ~3.3 MB |
| md_cleaner.exe | ~3.0 MB |
| md2word.exe | ~3.3 MB |
| template.docx | ~15 KB |

## 🎓 學習資源

- 使用 GUI：打開 gui.bat
- 查看文檔：在 GUI 選 4
- 查看源代碼：在 GUI 選 5

---

**版本**: 2.0
**狀態**: ✅ 已整理和優化
**最後更新**: 2026-04-27
