# 🚀 快速開始指南

## 只需一行命令！

```powershell
convert.bat homework.md "GaN驅動器IC設計報告"
```

**完成！** Word 檔案已生成：`homework_final.docx`

---

## 三種使用方式

### 1️⃣ 最簡單（推薦）
```powershell
convert.bat input.md "標題"
```

### 2️⃣ 命令行
```powershell
.\md2word_full.exe -i input.md -o output.docx -title "標題"
```

### 3️⃣ 自訂日期
```powershell
.\md2word_full.exe -i input.md -o output.docx -title "標題" -date "2026年5月1日"
```

---

## 自動執行的操作

- ✅ 移除 120 個 cite 標籤
- ✅ 添加標題（您指定的標題）
- ✅ 添加日期（自動使用今日日期）
- ✅ 生成 Word 檔案

---

## 輸出結果

| 項目 | 數值 |
|------|------|
| 原始字元 | 11,853 |
| 清理後字元 | 10,138 |
| 移除字元 | 1,715 (14.5%) |
| 移除 cite 標籤數 | 120 個 |

---

## 若需更多幫助

詳見 `COMPLETE_GUIDE.md` 完整文檔

**時間**: ~2026-04-27
**狀態**: ✅ 就緒
