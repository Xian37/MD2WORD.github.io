# Markdown 到 Word 自動轉換完整指南

## 🎯 功能概述

這個工具能一次性完成：
1. ✅ **清理 Markdown** - 自動移除所有 `cite` 標籤（共 120 個）
2. ✅ **添加標題和日期** - 自動根據現在日期填入標題
3. ✅ **轉換為 Word** - 生成最終的 `.docx` 檔案

## 🚀 快速開始

### 方式一：使用批次檔（推薦 - 最簡單）

```powershell
convert.bat homework.md "GaN驅動器IC設計報告"
```

**就是這樣！** 
- 輸入：`homework.md`
- 輸出：`homework_final.docx`（自動生成，包含今日日期）

### 方式二：使用命令行（需要更多控制）

```powershell
.\md2word_full.exe -i homework.md -o result.docx -title "您的標題"
```

### 方式三：使用 PowerShell 指令碼

```powershell
# 創建一個 PowerShell 指令碼檔案：convert.ps1
.\md2word_full.exe -i $args[0] -o ("$(Split-Path $args[0] -LeafBase)_final.docx") -title $args[1]
```

然後運行：
```powershell
.\convert.ps1 homework.md "GaN驅動器IC設計報告"
```

## 📊 工作流示例

### 示例 1：基本使用
```
原始檔案: homework.md (11,853 字元)
↓
1. 清理 cite 標籤
2. 添加標題：GaN驅動器IC設計報告
3. 添加日期：2026年4月27日
4. 轉換為 Word
↓
最終檔案: homework_final.docx ✅
```

### 示例 2：自訂輸出檔案名稱

```powershell
.\md2word_full.exe -i homework.md -o my_report.docx -title "報告標題"
```

## 📝 參數說明

### md2word_full.exe 參數

| 參數 | 說明 | 預設值 | 範例 |
|-----|-----|-------|------|
| `-i` | 輸入 Markdown 檔案 | `homework.md` | `-i input.md` |
| `-o` | 輸出 Word 檔案 | `homework_final.docx` | `-o report.docx` |
| `-title` | 文檔標題 | 自動提取或預設值 | `-title "我的報告"` |
| `-date` | 文檔日期 | 當前日期 | `-date "2026年4月27日"` |

### 例子

```powershell
# 基本用法（使用所有預設值）
.\md2word_full.exe

# 自訂標題
.\md2word_full.exe -title "GaN驅動器設計"

# 完整參數
.\md2word_full.exe -i homework.md -o final_report.docx -title "電路設計報告" -date "2026年4月27日"

# 使用批次檔
convert.bat homework.md "電路設計報告"

# 批次檔自訂輸出
convert.bat homework.md "電路設計報告" custom_output.docx
```

## 📂 檔案清單

### 核心檔案
- **md2word_full.exe** - 主程序（所有功能整合）
- **md2word_full.go** - 主程序源代碼
- **convert.bat** - 便捷批次檔

### 輔助工具
- **md_cleaner.exe** - 獨立的 Markdown 清理工具
- **md_cleaner.go** - 清理工具源代碼
- **md2word.exe** - 原始的 Markdown 到 Word 轉換工具

### 說明檔案
- **COMPLETE_GUIDE.md** - 完整使用指南（本檔案）
- **CLEANING_GUIDE.md** - Markdown 清理工具說明

## 🔍 清理效果

### 清理統計
- 🔴 **原始檔案**：11,853 字元（包含 120 個 cite 標籤）
- 🟢 **清理後**：10,138 字元
- 📊 **移除**：1,715 字元（14.5%）

### 清理內容

#### 1. Cite 標籤
- `[cite_start]...[cite_end]` 塊
- `[cite: XX]` 格式
- `[cite: XX, YY]` 多引用格式

#### 2. AI 對話（自動移除）
- 開頭的道歉和說明段落
- 結尾的 AI 說明
- 檔案指示符
- MATLAB/HW 相關說明

#### 3. 多餘空行
- 三個以上連續空行 → 合併為兩個

## 🎨 Word 檔案格式

生成的 Word 檔案包含：

```
┌─────────────────────────────────────┐
│  # GaN驅動器IC設計報告              │
│                                     │
│  **日期：2026年4月27日**            │
│                                     │
│  ─────────────────────────────────  │
│                                     │
│  ### I. 緒論 (Introduction)         │
│  本章節探討...                      │
│  ...                                │
│                                     │
└─────────────────────────────────────┘
```

## 🛠️ 常用命令速查

| 場景 | 命令 |
|------|------|
| 快速轉換（默認） | `convert.bat homework.md "標題"` |
| 自訂輸出名稱 | `.\md2word_full.exe -i homework.md -o report.docx -title "標題"` |
| 自訂日期 | `.\md2word_full.exe -date "2026年5月1日"` |
| 僅清理 Markdown | `.\md_cleaner.exe -i homework.md -o clean.md` |

## ⚙️ 進階用法

### 批量轉換

創建 `batch_convert.ps1`：
```powershell
$files = @("homework1.md", "homework2.md", "homework3.md")
$title = "課程作業"
$date = (Get-Date).ToString("yyyy年MM月dd日")

foreach ($file in $files) {
    $output = [System.IO.Path]::GetFileNameWithoutExtension($file) + "_final.docx"
    .\md2word_full.exe -i $file -o $output -title "$title - $([System.IO.Path]::GetFileNameWithoutExtension($file))" -date $date
}
```

運行：
```powershell
.\batch_convert.ps1
```

### 管道處理

```powershell
# 查找所有 .md 檔案並轉換
Get-ChildItem *.md | ForEach-Object {
    .\md2word_full.exe -i $_.Name -o "$($_.BaseName)_final.docx" -title $_.BaseName
}
```

## 📋 故障排除

### Q: 執行 convert.bat 時出現「找不到檔案」
**A:** 確保：
1. 已將 `.md` 檔案放在同一目錄
2. `md2word_full.exe` 和 `md2word.exe` 都在同一目錄

### Q: Word 檔案已生成但無法打開
**A:** 
- 確認 `template.docx` 在同一目錄中
- 嘗試用 Microsoft Word 修復檔案

### Q: 標題或日期沒有出現在 Word 檔案中
**A:**
- 檢查 `-title` 參數是否正確傳入
- 嘗試移除 `template.docx`，讓工具使用預設格式

### Q: cite 標籤仍未完全移除
**A:** 
- 檢查原始檔案中是否有其他非標準格式
- 聯繫技術支持以更新清理規則

## 📞 常見問題

### Q: 能否只清理 Markdown 而不轉換為 Word？
**A:** 可以，使用 `md_cleaner.exe`：
```powershell
.\md_cleaner.exe -i homework.md -o homework_cleaned.md
```

### Q: 能否保留原始的 cite 標籤？
**A:** 可以，跳過清理工具，直接使用 `md2word.exe`：
```powershell
.\md2word.exe -input homework.md
```

### Q: 如何修改 Word 檔案的樣式？
**A:** 編輯 `template.docx`，該檔案作為所有 Word 轉換的範本

### Q: 能否用其他日期格式？
**A:** 可以，使用 `-date` 參數：
```powershell
.\md2word_full.exe -date "April 27, 2026"
.\md2word_full.exe -date "27/04/2026"
```

## 📚 技術細節

### 使用的正則表達式

```go
// 移除 cite 標籤
reCiteFull = `\[cite_start\][\s\S]*?\[cite_end\]`
reCiteAny  = `\[cite[:\s_][^\]]*\]`

// 其他清理
reRef      = `\[\d+\]`                          // 數字引用
```

### 轉換流程

```
輸入 Markdown
     ↓
[Step 1] 移除 cite 標籤
     ↓
[Step 2] 移除 AI 對話
     ↓
[Step 3] 添加標題和日期
     ↓
[Step 4] 調用 md2word.exe
     ↓
輸出 Word 檔案
```

## 🎓 學習資源

- [Go 語言官網](https://golang.org)
- [Markdown 語法](https://commonmark.org)
- [Microsoft Word 格式規範](https://docs.microsoft.com/en-us/office/open-xml/wordprocessingml-structure)

---

**版本**: 2.0  
**最後更新**: 2026-04-27  
**語言**: 繁體中文  
**狀態**: ✅ 已驗證可用
