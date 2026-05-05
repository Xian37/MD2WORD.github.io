# 📖 README - Markdown to Word 自動轉換工具

## ✨ 功能亮點

- **一鍵轉換** - 從 Markdown 到 Word，自動清理、添加標題和日期
- **自動清理** - 移除所有 cite 標籤（共 120 個）
- **自動日期** - 根據系統日期自動填入
- **保留格式** - 完整保留原始 Markdown 的所有格式和內容

## 🚀 立即開始

### 最快的方式 - 只需一行命令：

```powershell
convert.bat homework.md "GaN驅動器IC設計報告"
```

✅ **完成！** Word 檔案已生成：`homework_final.docx`

---

## 📊 轉換效果

| 項目 | 數值 |
|------|------|
| 原始 Markdown | 11,853 字元（含 120 個 cite 標籤） |
| 最終 Word 檔案 | 21,250 bytes |
| 清理後內容 | 10,138 字元（移除 1,715 字元） |
| 處理時間 | < 1 秒 |

---

## 📚 使用方式

### 1. 批次檔（推薦）
```powershell
convert.bat homework.md "文檔標題"
convert.bat homework.md "文檔標題" "custom_output.docx"  # 自訂輸出檔名
```

### 2. 命令行（更多控制）
```powershell
.\md2word_full.exe -i homework.md -o result.docx -title "標題"
.\md2word_full.exe -i homework.md -date "2026年5月1日"  # 自訂日期
```

### 3. 僅清理 Markdown
```powershell
.\md_cleaner.exe -i homework.md -o homework_cleaned.md
```

---

## 📁 檔案說明

### 核心工具
| 檔案 | 說明 | 用途 |
|-----|-----|-----|
| **md2word_full.exe** | 完整轉換工具 | 清理 + 添加標題/日期 + 轉換為 Word |
| **convert.bat** | 快速批次檔 | 最簡單快速的使用方式 |
| **md_cleaner.exe** | 獨立清理工具 | 只清理 Markdown（不轉換） |
| **md2word.exe** | 原始轉換工具 | 基礎的 Markdown 到 Word 轉換 |

### 說明文檔
| 檔案 | 內容 |
|-----|-----|
| **QUICKSTART.md** | 快速開始指南 |
| **COMPLETE_GUIDE.md** | 完整功能說明 |
| **CLEANING_GUIDE.md** | Markdown 清理工具說明 |
| **README.md** | 本檔案 |

---

## 🎯 一般工作流程

```
homework.md (含 cite)
        ↓
   [步驟 1] 清理 cite 標籤
        ↓
   [步驟 2] 添加標題和日期
        ↓
   [步驟 3] 轉換為 Word
        ↓
homework_final.docx ✅
```

---

## ❓ 常見問題

**Q: 如何更改輸出檔案名稱？**
```powershell
convert.bat homework.md "標題" my_report.docx
```

**Q: 能否使用自訂日期？**
```powershell
.\md2word_full.exe -date "2026年5月1日"
```

**Q: 如何只清理 Markdown 不轉換？**
```powershell
.\md_cleaner.exe -i homework.md -o homework_cleaned.md
```

**Q: Word 檔案無法打開？**
- 確保 `template.docx` 在同一目錄
- 嘗試使用 Microsoft Word 修復檔案

---

## 🔧 技術詳情

### 清理移除的內容

✅ **Cite 標籤**
- `[cite_start]...[cite_end]` 塊（0 個）
- `[cite: XX]` 格式（120 個）
- `[cite: XX, YY, ZZ]` 多引用格式

✅ **AI 對話**
- 開頭的道歉和說明
- 結尾的 AI 說明
- 檔案指示符（`### 📄 檔案...`）

✅ **格式清理**
- 多餘空行（3+ 合併為 2）
- 不必要的對話標記

### 最終 Word 格式

```
# GaN驅動器IC設計報告

**日期：2026年4月27日**

---

### I. 緒論 (Introduction)
[完整內容...]
```

---

## 📞 支持

如有問題或建議，請檢查：
1. `QUICKSTART.md` - 快速指南
2. `COMPLETE_GUIDE.md` - 詳細說明
3. `CLEANING_GUIDE.md` - 清理工具說明

---

## 📝 版本資訊

- **版本**: 2.0
- **發佈日期**: 2026-04-27
- **狀態**: ✅ 已驗證可用
- **語言**: 繁體中文

---

**現在就開始使用！** 🚀

```powershell
convert.bat homework.md "您的標題"
```
