@echo off
chcp 65001 >nul
title Markdown to Word 轉換工具 - Web GUI

echo.
echo ============================================
echo    🌐 Markdown to Word Web 介面
echo ============================================
echo.
echo 正在啟動 Web 服務器...
echo.
echo 📱 功能頁面：
echo    • 轉換 - 單檔案處理
echo    • 批量處理 - 多檔案批量轉換
echo    • 設置 - 自訂選項配置
echo    • 歷史記錄 - 轉換記錄管理
echo    • 說明 - 詳細使用指南
echo.
echo 🌐 訪問網址：
echo http://localhost:8000/web_gui.html
echo.
echo ⌨️  鍵盤快捷鍵：
echo    Ctrl+1: 轉換頁面
echo    Ctrl+2: 批量處理
echo    Ctrl+3: 設置頁面
echo    Ctrl+4: 歷史記錄
echo    Ctrl+5: 說明頁面
echo.
echo 按 Ctrl+C 停止服務器
echo ============================================
echo.

cd /d "%~dp0"

REM 檢查 Python 是否可用
python --version >nul 2>&1
if %errorlevel% equ 0 (
    echo 正在使用 Python 啟動服務器...
    python -m http.server 8000 --bind 127.0.0.1
) else (
    REM 如果沒有 Python，嘗試使用 Node.js
    node --version >nul 2>&1
    if %errorlevel% equ 0 (
        echo 正在使用 Node.js 啟動服務器...
        npx http-server -p 8000 -a 127.0.0.1
    ) else (
        echo 錯誤：需要安裝 Python 或 Node.js 來運行 Web 服務器
        echo.
        echo 請安裝 Python: https://www.python.org/downloads/
        echo 或安裝 Node.js: https://nodejs.org/
        echo.
        pause
        exit /b 1
    )
)

pause