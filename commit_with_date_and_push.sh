#!/bin/bash

git pull

# 現在の日付を取得して、yyyyMMdd形式にフォーマット
current_date=$(date +%Y%m%d)

# コミットするファイルをすべてステージング
git add . 

# コミットメッセージに日付を使用
git commit -m "$current_date"

# コミットが成功した場合のメッセージ
if [ $? -eq 0 ]; then
    echo "Committed with message: $current_date"
else
    echo "Failed to commit. Please check for errors."
fi

git push origin main
