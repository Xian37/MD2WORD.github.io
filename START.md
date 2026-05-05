# 🎉 Markdown to Word 自動轉換系統

**版本**: 3.0 (Web GUI)  
**狀態**: ✅ 已就緒  
**組織**: ✨ 已整理  
**日期**: 2026-04-27

---

## ⚡ 立即開始（四種方式）

### 1️⃣ 最簡單 - 雙擊 Web GUI
```
📁 md2word/
└── 🌐 web_gui.bat  ← 點擊這裡！（推薦）
```

### 2️⃣ 簡單 - 雙擊 GUI 菜單
```
📁 md2word/
└── 🚀 gui.bat  ← 點擊這裡！
```

### 3️⃣ 命令行
```powershell
bin\md2word_full.exe -i your_file.md -title "標題"
```

### 4️⃣ 快速批次檔
```powershell
bin\convert.bat your_file.md "標題"
```

---

## 📂 項目結構

```
md2word/
├── 🚀 gui.bat                ← GUI菜單（推薦）
├── bin/                       可執行工具
├── src/                       源代碼
├── docs/                      說明文檔
├── samples/                   樣例檔案
├── lib/                       庫檔案
└── output/                    輸出目錄
```

詳見: [`docs/PROJECT_STRUCTURE.md`](docs/PROJECT_STRUCTURE.md)

---

## 🎯 主要功能

- ✅ **自動清理**: 移除 120+ 個 cite 標籤
- ✅ **添加標題**: 用戶指定的標題
- ✅ **自動日期**: 使用系統日期（可自訂）
- ✅ **轉換到 Word**: 生成 .docx 檔案
- ✅ **保留格式**: 完整保留原始內容格式

---

## 📖 文檔導航

| 文檔 | 說明 | 適合 |
|-----|-----|-----|
| [README.md](docs/README.md) | 項目概述 | 所有人 |
| [QUICKSTART.md](docs/QUICKSTART.md) | 快速開始 | 新手 |
| [COMPLETE_GUIDE.md](docs/COMPLETE_GUIDE.md) | 完整指南 | 進階用戶 |
| [CLEANING_GUIDE.md](docs/CLEANING_GUIDE.md) | 工具詳解 | 開發者 |
| [PROJECT_STRUCTURE.md](docs/PROJECT_STRUCTURE.md) | 結構說明 | 所有人 |

---

## 🌐 Web GUI 功能介紹

全新升級的 **Web 介面** 提供現代化的視覺化操作體驗：

### 🎯 五大功能頁面

#### 📄 轉換頁面
- **拖拽上傳**: 支援拖拽檔案到瀏覽器
- **即時預覽**: 顯示檔案資訊和大小
- **多模板選擇**: 預設、學術論文、技術報告、簡潔風格
- **智慧清理**: 自動移除 120+ 個 cite 標籤

#### 📦 批量處理
- **多檔案上傳**: 一次處理多個 Markdown 檔案
- **檔案管理**: 可個別移除不需要的檔案
- **統一設置**: 為所有檔案應用相同模板和日期
- **進度追蹤**: 即時顯示處理狀態

#### ⚙️ 設置頁面
- **自動清理**: 控制是否自動移除 cite 標籤
- **智慧標題**: 根據檔案名自動生成標題
- **壓縮優化**: 減小輸出檔案大小
- **歷史記錄**: 保存轉換歷史
- **桌面通知**: 轉換完成後顯示通知

#### 📊 歷史記錄
- **完整記錄**: 保存所有轉換操作
- **狀態追蹤**: 成功/失敗狀態顯示
- **詳細資訊**: 檔案名稱、模板、時間戳記
- **匯出功能**: 可匯出歷史記錄為 JSON

#### ❓ 說明頁面
- **統計數據**: 顯示使用統計和效能指標
- **功能介紹**: 詳細說明各項功能
- **使用指南**: 步驟化操作指導
- **技術支援**: 系統需求和支援資訊

### ⌨️ 鍵盤快捷鍵
- `Ctrl+1`: 切換到轉換頁面
- `Ctrl+2`: 切換到批量處理
- `Ctrl+3`: 切換到設置頁面
- `Ctrl+4`: 切換到歷史記錄
- `Ctrl+5`: 切換到說明頁面

### 🎨 設計特色
- **現代UI**: 採用 Material Design 風格
- **響應式設計**: 支援各種螢幕尺寸
- **流暢動畫**: 平滑的頁面切換效果
- **直觀操作**: 拖拽、點擊、鍵盤操作一應俱全

---

## 🚀 快速使用

### 方式 A: GUI 菜單（推薦）✨

1. **雙擊** `gui.bat`
2. **選擇** "1. 一鍵轉換"
3. **輸入** Markdown 檔案路徑
4. **輸入** 文檔標題
5. **完成！** 檔案已生成
### 方式 A+: Web GUI（推薦）

1. **雙擊** `web_gui.bat`
2. **瀏覽器自動開啟** http://localhost:8000/web_gui.html
3. **拖拽** Markdown 檔案
4. **輸入** 文檔標題
5. **點擊** "轉換為 Word"
6. **完成！** 檔案已生成
### 方式 B: 命令行

```powershell
# 基本用法
bin\md2word_full.exe -i homework.md -title "報告標題"

# 自訂日期
bin\md2word_full.exe -i homework.md -title "標題" -date "2026年5月1日"

# 自訂輸出
bin\md2word_full.exe -i homework.md -o custom_name.docx -title "標題"
```

### 方式 C: 批次檔

```powershell
bin\convert.bat homework.md "報告標題"
```

---

## 📊 轉換效果示例

**輸入**: homework.md (11,853 字元，包含 120 個 cite 標籤)  
**輸出**: homework_final.docx (21,250 bytes)

**清理統計**:
- 移除 cite 標籤: 120 個
- 移除字元: 1,715 (14.5%)
- 最終字元: 10,138

---

## 🛠️ 工具說明

### bin/ - 可執行工具

| 工具 | 說明 |
|-----|------|
| **md2word_full.exe** | ⭐ 推薦使用 - 完整功能 |
| md2word.exe | 基礎轉換 |
| md_cleaner.exe | 僅清理 Markdown |
| convert.bat | 快速批次轉換 |
| **gui.bat** | ✨ GUI 菜單（主程序） |

### src/ - 源代碼

可自行編譯或修改：
```powershell
go build -o bin\md2word_full.exe src\md2word_full.go
```

### samples/ - 樣例

包含轉換前後的範例檔案，可用於測試。

---

## ❓ 常見問題

### Q: 怎樣使用最簡單？
**A**: 雙擊 `gui.bat`，在菜單中選擇選項即可。

### Q: 轉換出來的 Word 檔在哪裡？
**A**: 預設在 `output/` 目錄，或根據您指定的路徑。

### Q: 可以自訂日期嗎？
**A**: 可以，使用 `-date "2026年5月1日"` 參數。

### Q: 可以只清理 Markdown 不轉換嗎？
**A**: 可以，使用 `bin\md_cleaner.exe`。

### Q: 支援哪些 cite 格式？
**A**: 支援所有格式：`[cite_start]...[cite_end]`、`[cite: XX]`、`[cite: XX, YY, ZZ]`。

---

## 📋 系統需求

- Windows 7 或更高版本
- Microsoft Word 或 Word Online（查看結果）
- 不需要額外安裝軟體

---

## 🎓 下一步

1. **了解**: 閱讀 [QUICKSTART.md](docs/QUICKSTART.md)
2. **嘗試**: 雙擊 `gui.bat` 體驗
3. **轉換**: 轉換您的第一個檔案
4. **深入**: 查看 [COMPLETE_GUIDE.md](docs/COMPLETE_GUIDE.md) 了解更多

---

## 🔗 快速連結

- 🚀 **GUI 菜單**: `gui.bat`
- 📚 **文檔**: `docs/`
- 💾 **工具**: `bin/`
- 📦 **樣例**: `samples/`
- 📤 **輸出**: `output/`

---

## ✨ 特色

- 🎨 簡潔的 GUI 界面
- ⚡ 快速轉換（< 1秒）
- 🔧 完全可自訂
- 📊 詳細的轉換統計
- 🛡️ 安全可靠
- 💬 繁體中文支援

---

## 📞 需要幫助？

1. 查看相關文檔
2. 檢查常見問題
3. 查看樣例檔案

---

**祝您使用愉快！** 🎉

*Markdown to Word 自動轉換系統 v2.0*  
*最後更新: 2026-04-27*
