@echo off
cd /d "%~dp0"

echo.
echo ========================================
echo   Markdown to Word 轉換工具
echo ========================================
echo.
echo 正在啟動伺服器...
echo.

REM 啟動後端伺服器
start cmd /k "cd /d "%~dp0" && .\bin\web_server.exe"

REM 等待 1 秒，確保伺服器啟動
timeout /t 1 /nobreak

REM 用預設瀏覽器打開 Web GUI
start http://localhost:8888/web_gui.html

echo.
echo ✅ 伺服器已啟動！
echo 🌐 Web 介面已在瀏覽器中打開
echo 📂 如果頁面沒有打開，請手動訪問: http://localhost:8888/web_gui.html
echo.
echo ⚠️  轉換後的檔案將保存在 output 資料夾中
echo.
