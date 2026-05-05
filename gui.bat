@echo off
chcp 65001 >nul
cd /d "%~dp0"

:menu
cls
echo ==========================================
echo  Markdown to Word 轉換工具 - GUI 菜單
echo ==========================================
echo.
echo  1. 一鍵轉換 ^(輸入Markdown檔案和標題^)
echo  2. 打開檔案資源管理器 - 樣例目錄
echo  3. 打開檔案資源管理器 - 輸出目錄
echo  4. 查看說明文檔
echo  5. 打開工具資料夾
echo  6. 退出
echo.
set /p choice="請選擇 ^(1-6^): "

if "%choice%"=="1" goto convert
if "%choice%"=="2" goto samples_dir
if "%choice%"=="3" goto output_dir
if "%choice%"=="4" goto docs
if "%choice%"=="5" goto folders
if "%choice%"=="6" goto end
goto menu

:convert
cls
echo.
echo 一鍵轉換 Markdown 到 Word
echo ======================================
echo.
set /p mdfile="輸入 Markdown 檔案路徑: "
set /p title="輸入文檔標題: "

if "%mdfile%"=="" (
    echo 錯誤: 請輸入檔案路徑
    pause
    goto menu
)

if "%title%"=="" (
    echo 錯誤: 請輸入標題
    pause
    goto menu
)

echo.
echo 正在轉換...
echo.

bin\md2word_full.exe -i "%mdfile%" -title "%title%"

echo.
echo 轉換完成！ 按任意鍵返回菜單...
pause >nul
goto menu

:samples_dir
start explorer samples
goto menu

:output_dir
start explorer output
goto menu

:docs
cls
echo.
echo 選擇文檔查看
echo ======================================
echo.
echo 1. README - 項目概述
echo 2. QUICKSTART - 快速開始
echo 3. COMPLETE_GUIDE - 完整指南
echo 4. CLEANING_GUIDE - 清理工具說明
echo.
set /p docch="選擇 ^(1-4^): "

if "%docch%"=="1" start notepad docs\README.md
if "%docch%"=="2" start notepad docs\QUICKSTART.md
if "%docch%"=="3" start notepad docs\COMPLETE_GUIDE.md
if "%docch%"=="4" start notepad docs\CLEANING_GUIDE.md

goto menu

:folders
cls
echo.
echo 工具資料夾
echo ======================================
echo.
echo 1. bin - 可執行檔和工具
echo 2. src - 源代碼
echo 3. docs - 說明文檔
echo 4. samples - 樣例檔案
echo 5. lib - 庫檔案
echo 6. output - 輸出目錄
echo.
set /p folch="選擇 ^(1-6^): "

if "%folch%"=="1" start explorer bin
if "%folch%"=="2" start explorer src
if "%folch%"=="3" start explorer docs
if "%folch%"=="4" start explorer samples
if "%folch%"=="5" start explorer lib
if "%folch%"=="6" start explorer output

goto menu

:end
echo.
echo 感謝使用！
echo.
exit /b
