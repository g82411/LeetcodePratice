@echo off
REM 檢查是否有足夠的參數
if "%~1"=="" goto usage
if "%~2"=="" goto usage

REM 取得參數
set "name=%~1"
set "link=%~2"

REM 處理名稱，將空格和點換成底線
set "new_name=%name:. =_%"
set "new_name=%new_name: =_%"
set "new_name=%new_name:.=_%"

REM 更新 README.md
echo - **[#%new_name%](./problems/%new_name%/solution.go)** - [LeetCode Link](%link%) >> README.md

REM 創建資料夾
if not exist ".\problems\%new_name%" mkdir ".\problems\%new_name%"

REM 創建 solution.go 並寫入包名（如果包名以數字開頭則在前面加底線）
echo package _%new_name% > ".\problems\%new_name%\solution.go"
REM 創建README.md
echo # %name% >> ".\problems\%new_name%\README.md"

echo 完成。已更新 README 並創建了包含解決方案文件的文件夾。
goto end

:usage
echo 用法： %0 [name] [link]
goto end

:end
