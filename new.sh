#!/bin/bash

# 檢查是否有足夠的參數
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 [name] [link]"
    exit 1
fi

# 取得參數
name="$1"
link="$2"

# 處理名稱，將空格和點換成底線
# 首先處理 '. ' 的情況，然後處理單獨的空格和點
new_name=$(echo "$name" | sed 's/\. /_/g' | sed 's/[ .]/_/g')

# 更新 README.md
echo "- **[#${new_name}](./problems/${new_name}/solution.go)** - [LeetCode Link]($link)" >> README.md

# 創建資料夾
mkdir -p "./problems/${new_name}"

# 創建 solution.go 並寫入包名（如果包名以數字開頭則在前面加底線）
echo "package _${new_name}" > "./problems/${new_name}/solution.go"

echo "Done. Updated README and created folder with solution file."
