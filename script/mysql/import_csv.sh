#!/bin/bash
# 脚本用途：将一批 CSV 文件导入指定的 Docker MySQL 容器中
# 使用方式：./import_csv.sh [-c container_name] [-h db_host] [-P db_port] [-u user] [-p password] [-d database]

# ========= 默认配置 ============
CONTAINER_NAME="mysql"
DB_HOST="127.0.0.1"
DB_PORT=3306
DB_USER="root"
DB_ROOT_PASSWORD="Mysql@2o20..."
DATABASE_NAME=""

# CSV 文件配置
# 格式："CSV文件路径:目标表名:数据库名"（数据库名可选，不填则使用默认数据库）
CSV_FILES=(
    "./biz_ip_address.csv:biz_ip_address"
    "./biz_isp.csv:biz_isp"
    # 添加更多CSV文件...
)

# CSV 导入选项配置
CSV_FIELD_SEPARATOR=","           # 字段分隔符
CSV_LINE_SEPARATOR="\n"           # 行分隔符
CSV_FIELD_ENCLOSURE=""            # 字段包围符（可选，如 '"'，只支持双引号或不包含引号）
CSV_ESCAPE_CHAR="\\\\"            # 转义字符
CSV_IGNORE_LINES=1                # 跳过开头行数（通常是表头）

# 解析命令行参数
while getopts ":c:h:P:u:p:d:" opt; do
  case $opt in
    c) CONTAINER_NAME="$OPTARG" ;;
    h) DB_HOST="$OPTARG" ;;
    P) DB_PORT="$OPTARG" ;;
    u) DB_USER="$OPTARG" ;;
    p) DB_ROOT_PASSWORD="$OPTARG" ;;
    d) DATABASE_NAME="$OPTARG" ;;
    *)
      echo "用法: $0 [-c 容器名] [-h 主机] [-P 端口] [-u 用户] [-p 密码] [-d 数据库]"
      echo "示例: $0 -c mysql-container -d mydb"
      exit 1
      ;;
  esac
done

# 通用函数，检查命令执行结果
check_succ(){
    if [ $? -ne 0 ]; then
        echo -e "[失败] $1\n"
        exit 1
    else
        echo -e "[成功] $1\n"
    fi
}

# 检查容器是否运行
check_container() {
    docker ps | grep -q "$CONTAINER_NAME"
    check_succ "检查容器 $CONTAINER_NAME 是否运行"
}

# 测试数据库连接
test_connection() {
    docker exec "$CONTAINER_NAME" mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_ROOT_PASSWORD" -e "SELECT 1;" > /dev/null 2>&1
    check_succ "检查数据库连接"
}

# 获取CSV文件行数（用于统计）
get_csv_row_count() {
    local csv_file="$1"
    local skip_lines="$2"
    
    if [ ! -f "$csv_file" ]; then
        echo "0"
        return
    fi
    
    # 获取总行数，减去跳过的行数
    local total_lines=$(wc -l < "$csv_file")
    local data_lines=$((total_lines - skip_lines))
    
    # 确保不为负数
    if [ $data_lines -lt 0 ]; then
        data_lines=0
    fi
    
    echo "$data_lines"
}

# 导入单个 CSV 文件
import_single_csv() {
    local csv_config="$1"
    
    # 解析配置：CSV文件路径:目标表名:数据库名
    IFS=':' read -r csv_file table_name db_name <<< "$csv_config"
    
    if [ ! -f "$csv_file" ]; then
        echo "CSV 文件 $csv_file 不存在，跳过。"
        return 1
    fi

    # 如果没有指定数据库，使用默认数据库
    if [ -z "$db_name" ]; then
        db_name="$DATABASE_NAME"
    fi

    # 获取预期导入行数
    local expected_rows=$(get_csv_row_count "$csv_file" "$CSV_IGNORE_LINES")
    
    echo -e "正在导入 $csv_file 到表 $table_name (数据库: $db_name)..."
    echo -e "预期导入行数: $expected_rows\n"

    local start_time=$(date +%s%3N)

    # 复制 CSV 文件到容器
    local container_csv_path="/tmp/$(basename "$csv_file")"
    docker cp "$csv_file" "$CONTAINER_NAME:$container_csv_path"
    check_succ "复制 CSV 文件 $csv_file 到容器"

    # 构建 LOAD DATA INFILE SQL
    local load_sql="USE $db_name; LOAD DATA LOCAL INFILE '$container_csv_path' INTO TABLE $table_name"
    load_sql="$load_sql FIELDS TERMINATED BY '$CSV_FIELD_SEPARATOR'"
    
    # 添加字段包围符（如果配置了）
    if [ -n "$CSV_FIELD_ENCLOSURE" ]; then
        load_sql="$load_sql ENCLOSED BY '$CSV_FIELD_ENCLOSURE'"
    fi
    
    load_sql="$load_sql ESCAPED BY '$CSV_ESCAPE_CHAR'"
    load_sql="$load_sql LINES TERMINATED BY '$CSV_LINE_SEPARATOR'"
    
    # 跳过表头行
    if [ "$CSV_IGNORE_LINES" -gt 0 ]; then
        load_sql="$load_sql IGNORE $CSV_IGNORE_LINES ROWS"
    fi
    
    load_sql="$load_sql;"

    # 临时开启 LOAD DATA INFILE 功能
    docker exec -i "$CONTAINER_NAME" mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_ROOT_PASSWORD" -e "SET GLOBAL local_infile=1;"
    check_succ "临时开启 LOAD DATA INFILE 功能"

    # 执行导入
    echo "$load_sql" | docker exec -i "$CONTAINER_NAME" mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_ROOT_PASSWORD" --local-infile=1 --default-character-set=utf8mb4
    local import_status=$?
    
    # 清理容器中的临时文件
    docker exec "$CONTAINER_NAME" rm "$container_csv_path" > /dev/null 2>&1

    local end_time=$(date +%s%3N)
    local duration_sec=$(awk -v start="$start_time" -v end="$end_time" 'BEGIN { printf "%.2f", (end - start) / 1000 }')
    
    if [ $import_status -ne 0 ]; then
        echo -e "[错误] 导入 CSV 文件 $csv_file 到表 $table_name （耗时 ${duration_sec}s）\n"
        return 1
    else
        # 验证导入结果
        local imported_count=$(docker exec "$CONTAINER_NAME" mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_ROOT_PASSWORD" -sN -e "USE $db_name; SELECT COUNT(*) FROM $table_name;" 2>/dev/null)
        
        if [ -n "$imported_count" ]; then
            echo -e "[成功] 导入 CSV 文件 $csv_file 到表 $table_name （表中总记录数: $imported_count, 耗时 ${duration_sec}s）\n"
        else
            echo -e "[成功] 导入 CSV 文件 $csv_file 到表 $table_name （耗时 ${duration_sec}s）\n"
        fi
        return 0
    fi
}

# 导入所有 CSV 文件
import_csv_files() {
    echo -e "开始导入 CSV 文件...\n"

    if [ ${#CSV_FILES[@]} -eq 0 ]; then
        echo "没有配置 CSV 文件，请在脚本中配置 CSV_FILES 数组。"
        return 0
    fi

    local total_start_time=$(date +%s%3N)
    local success_count=0
    local fail_count=0

    for csv_config in "${CSV_FILES[@]}"; do
        if import_single_csv "$csv_config"; then
            ((success_count++))
        else
            ((fail_count++))
        fi
    done

    local total_end_time=$(date +%s%3N)
    local total_duration_sec=$(awk -v start="$total_start_time" -v end="$total_end_time" 'BEGIN { printf "%.2f", (end - start) / 1000 }')

    echo "========== 导入完成 =========="
    echo "成功导入: $success_count 个文件"
    echo "导入失败: $fail_count 个文件"
    echo "总耗时: ${total_duration_sec}s"
    
    if [ $fail_count -gt 0 ]; then
        echo -e "\n[警告] 有 $fail_count 个文件导入失败，请检查日志"
        exit 1
    fi
}

# 显示配置信息
show_config() {
    echo "========== 配置信息 =========="
    echo "容器名称: $CONTAINER_NAME"
    echo "数据库主机: $DB_HOST"
    echo "数据库端口: $DB_PORT"
    echo "数据库用户: $DB_USER"
    echo "默认数据库: $DATABASE_NAME"
    echo "字段分隔符: '$CSV_FIELD_SEPARATOR'"
    echo "行分隔符: '$CSV_LINE_SEPARATOR'"
    echo "跳过行数: $CSV_IGNORE_LINES"
    echo "待导入文件数: ${#CSV_FILES[@]}"
    echo "=============================="
    echo ""
}

main() {
    show_config
    check_container
    test_connection
    import_csv_files
}

main