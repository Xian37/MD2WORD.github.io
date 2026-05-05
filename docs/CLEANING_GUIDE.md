# Markdown cite 標籤清理工具使用指南

## 問題描述
轉換後的Word檔中仍然保留了許多 `[cite_start]`、`[cite_end]` 和 `[cite: XX]` 這樣的引用標籤，需要在轉換前移除這些標籤。

## 解決方案

### 方式一：使用快速清理工具（推薦）

#### 步驟1: 清理Markdown檔案
```powershell
.\md_cleaner.exe -i homework.md -o homework_cleaned.md
```

**參數說明:**
- `-i homework.md` : 指定輸入檔案（默認值）
- `-o homework_cleaned.md` : 指定輸出檔案名稱（默認值）
- `-inplace` : 直接修改原檔案（危險操作，不推薦）

**輸出示例:**
```
✅ 清理完成！
=== 檔案資訊 ===
輸入: homework.md (11853 字元)
輸出: homework_cleaned.md (10138 字元)
移除: 1715 字元 (14.5%)

=== 移除的 cite 標籤 ===
  - [cite_start/end] 塊: 0
  - [cite:...] 標籤: 120
  - 共計: 120 個

✓ 驗證: 所有 cite 標籤已成功移除！
```

#### 步驟2: 使用清理後的檔案進行Word轉換
將 `homework_cleaned.md` 用於您的md2word轉換工具。

### 方式二：清理其他Markdown檔案
```powershell
.\md_cleaner.exe -i your_file.md -o your_file_cleaned.md
```

## 清理的內容

該工具會自動移除以下內容：

1. **所有 cite 標籤**
   - `[cite_start]...[cite_end]` 塊
   - `[cite: XX]` 格式的標籤
   - `[cite: XX, YY, ZZ]` 多引用標籤

2. **AI助理的對話文本**（可選）
   - 開頭的道歉和說明段落
   - 結尾的道歉和說明
   - 檔案指示符（`### 📄 檔案...`）
   - MATLAB/HW相關說明

3. **多餘空行**
   - 三個以上連續空行會被合併為兩個

## 清理效果

### 清理前後對比

**清理前:**
```markdown
* [cite_start]**GaN 物理特性與應用優勢：** GaN 是一種... [cite: 33][cite_start]。
  相較於傳統矽 (Si) 元件，GaN 能在... [cite: 34, 35]。
```

**清理後:**
```markdown
* **GaN 物理特性與應用優勢：** GaN 是一種...。
  相較於傳統矽 (Si) 元件，GaN 能在...。
```

## 常用命令速查

| 用途 | 命令 |
|-----|-----|
| 清理 homework.md | `.\md_cleaner.exe` |
| 自訂輸入輸出 | `.\md_cleaner.exe -i input.md -o output.md` |
| 清理多個檔案 | 逐個執行上述命令 |

## 文件清單

- `md_cleaner.exe` - 可執行的清理工具
- `md_cleaner.go` - 清理工具的源代碼
- `homework_cleaned.md` - 已清理的Markdown檔案（如果已運行清理工具）

## 技術細節

### 使用的正則表達式

```go
reCiteFull    = `\[cite_start\][\s\S]*?\[cite_end\]`    // cite_start/end 塊
reCiteAny     = `\[cite[:\s_][^\]]*\]`                  // 所有 cite 標籤
reRef         = `\[\d+\]`                                // 數字引用
reAIIntro     = `^[\s\S]*?非常抱歉[\s\S]*?PDF\s*即可。` // AI開頭
reAIOutro     = `\n*再次為[\s\S]*?請隨時跟我說[!！]`    // AI結尾
```

## 故障排除

### Q: 執行md_cleaner.exe時出現「找不到檔案」錯誤
**A:** 確保homework.md檔案在同一目錄中。

### Q: 清理後的檔案大小沒有變化
**A:** 檢查原始檔案中是否確實包含cite標籤。

### Q: 清理後仍有cite標籤遺留
**A:** 您的cite標籤可能採用了不同的格式，請報告該格式以便更新工具。

## 進階用法

### 在Go程序中使用清理函數

```go
import "os"

// 讀取檔案
content, _ := os.ReadFile("homework.md")

// 進行清理
cleanedContent := preprocessContent(content)

// 寫入新檔案
os.WriteFile("homework_cleaned.md", []byte(cleanedContent), 0644)
```

## 建議工作流程

1. **原始Markdown** → `homework.md` (包含cite標籤)
2. **運行清理工具** → `.\md_cleaner.exe`
3. **清理後Markdown** → `homework_cleaned.md` (無cite標籤)
4. **轉換為Word** → 使用md2word或其他工具
5. **最終Word檔** → `homework.docx` (無cite標籤)

---

**更新日期:** 2026-04-27
**版本:** 1.0
