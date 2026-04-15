#!/bin/bash
# 脚本用途：将一批 SQL 文件导入指定的 Docker MySQL 容器中
# 使用方式：./import_sql.sh [-c container_name] [-h db_host] [-P db_port] [-u user] [-p password]

# ========= 默认配置 ============
CONTAINER_NAME="mysql"
DB_HOST="127.0.0.1"
DB_PORT=3306
DB_USER="root"
DB_ROOT_PASSWORD="Mysql@2o20..."
SQL_FILES=(
    "./fcas_system.sql"
    "./fcas_service.sql"
#    "./dim_user_info.sql"
#    "./dim_user_crowd.sql"
#    "./dim_user_crowd_group.sql"
#    "./dim_user_crowd_relation.sql"
#    "./dim_user_crowd_group_relation.sql"
#    "./dim_bypass.sql"
)

# 解析命令行参数
while getopts ":c:h:P:u:p:" opt; do
  case $opt in
    c) CONTAINER_NAME="$OPTARG" ;;
    h) DB_HOST="$OPTARG" ;;
    P) DB_PORT="$OPTARG" ;;
    u) DB_USER="$OPTARG" ;;
    p) DB_ROOT_PASSWORD="$OPTARG" ;;
    *)
      echo "用法: $0 [-c 容器名] [-h 主机] [-P 端口] [-u 用户] [-p 密码]"
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

# 导入 SQL 文件
import_sql_files() {
    echo -e "开始导入 SQL 文件...\n"

    total_start_time=$(date +%s%3N)

    for sql_file in "${SQL_FILES[@]}"; do
        if [ ! -f "$sql_file" ]; then
            echo "SQL 文件 $sql_file 不存在，跳过。"
            continue
        fi

        echo -e "正在导入 $sql_file ...\n"

        start_time=$(date +%s%3N) # 毫秒时间戳

        # 导入 SQL 文件
        docker exec -i "$CONTAINER_NAME" mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_ROOT_PASSWORD" --default-character-set=utf8mb4 < "$sql_file"
        status=$? # 保存命令执行结果状态

        end_time=$(date +%s%3N) # 毫秒时间戳

        # 获取命令耗时
        duration_sec=$(awk -v start="$start_time" -v end="$end_time" 'BEGIN { printf "%.2f", (end - start) / 1000 }')

        # 检查命令执行结果
        if [ $status -ne 0 ]; then
            echo -e "[错误] 导入 SQL 文件 $sql_file （耗时 ${duration_sec}s）\n"
            exit 1
        else
            echo -e "[成功] 导入 SQL 文件 $sql_file （耗时 ${duration_sec}s）\n"
        fi

    done

    total_end_time=$(date +%s%3N)
    total_duration_sec=$(awk -v start="$total_start_time" -v end="$total_end_time" 'BEGIN { printf "%.2f", (end - start) / 1000 }')


    echo "所有 SQL 文件导入完成，总耗时：${total_duration_sec}s"
}

main() {
    import_sql_files
}

main
