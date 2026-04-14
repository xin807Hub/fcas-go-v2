#!/bin/bash

# 定义源目录和目的目录
src_dir="/home/ftpuser1"
dest_dir="/home/ftpuser"

# 检查源目录是否存在
if [ ! -d "$src_dir" ]; then
    echo "源目录不存在: $src_dir"
    exit 1
fi

# 检查目的目录是否存在，如果不存在则创建
if [ ! -d "$dest_dir" ]; then
    echo "目的目录不存在，正在创建: $dest_dir"
    mkdir -p "$dest_dir"
fi

# 逐个复制文件
for file in "$src_dir"/*; do
    # 检查是否为文件
    if [ -f "$file" ]; then
        # 提取文件名
        filename=$(basename "$file")
        # 复制文件到目的目录
        cp "$file" "$dest_dir/$filename"
        echo "已复制: $filename"
    fi
done

echo "所有文件复制完成。"


