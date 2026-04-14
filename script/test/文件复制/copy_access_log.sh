#!/bin/bash

src_dir=/home/dpi_log/local_access_log  # local_info_safe_log
dst_dir=/home/dpi_log/access_log  # info_safe_log

# 复制文件并创建 .ok 文件
for file in "$src_dir"/*; do
    if [ -f "$file" ]; then
        # filename=$(basename "$file")
	
        # 分离文件名和扩展名
	filename=$(basename "$file")
        #extension="${filename##*.}"
	extension="tar.gz"
        name="${filename%.*}"

        # 如果扩展名中还有点 则继续处理
        while [[ "$extension" == *.? ]]; do
            name="${filename%.*}"
	    #extension="${filename##*.}"
	    filename="$name"
        done
        
        #dst_file="$dst_dir/$filename"
        # timestamp=$(date +%s)
	# 获取当前时间戳（精确到微秒）
	timestamp=$(date +%s.%N | cut -d '.' -f 1,2)  # 取前6位作为微秒

        # 生成新的文件名
	only_name="${name}_${timestamp}"
        new_filename="${only_name}.${extension}"
        dst_file="$dst_dir/$new_filename"
        
        # 复制文件
        cp "$file" "$dst_file"
        
        # 检查复制操作是否成功
        if [ $? -eq 0 ]; then
            echo "File '$filename' has been successfully copied."
            
            # 创建 .ok 文件
	    ok_file="${dst_dir}/${only_name}.ok"
            touch "$ok_file"
            echo "Created .ok file for '$new_filename'."
        else
            echo "Failed to copy file '$filename'."
        fi
    fi
done



