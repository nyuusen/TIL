#!/bin/bash

git pull

current_date=$(date +%Y%m%d)

git add . 

git commit -m "$current_date"

if [ $? -eq 0 ]; then
    echo "Committed with message: $current_date"
else
    echo "Failed to commit. Please check for errors."
fi

git push origin main
