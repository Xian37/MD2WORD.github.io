@echo off
setlocal enabledelayedexpansion

if "%1"=="" (
    echo 用法: convert.bat input.md "文檔標題" [output.docx]
    echo 示例: convert.bat homework.md "GaN驅動器IC設計報告"
    exit /b 1
)

set INPUT=%1
set TITLE=%2
if "%TITLE%"=="" set TITLE=學術報告

if "%3"=="" (
    set OUTPUT=%~n1_final.docx
) else (
    set OUTPUT=%3
)

echo.
echo ============================================
echo 開始轉換 Word 檔案...
echo ============================================
echo 輸入: %INPUT%
echo 標題: %TITLE%
echo 輸出: %OUTPUT%
echo.

md2word_full.exe -i "%INPUT%" -o "%OUTPUT%" -title "%TITLE%"

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ============================================
    echo 成功！最終Word檔案已生成: %OUTPUT%
    echo ============================================
    echo.
    exit /b 0
) else (
    echo.
    echo 轉換失敗！
    exit /b 1
)
