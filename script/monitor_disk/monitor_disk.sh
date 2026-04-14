#!/bin/bash

# 配置参数
MOUNT_PATH="/home"                # 监控的挂载路径
THRESHOLD="85"                    # 磁盘使用量阈值（百分比）
TABLES=("dws_app_type_10min_local" "dws_app_type_hour_local" "dws_dstip_top_10min_local" "dws_dstip_top_10min_no_params_local" "dws_dstip_top_hour_local" "dws_dstip_top_hour_no_params_local" "ods_xxx_table7" "dws_gen_traffic_hour_local" "dws_isp_10min_local" "dws_isp_hour_local" "dws_link_10min_local" "dws_link_app_local" "dws_server_10min_local" "dws_server_hour_local" "dws_srcip_top_10min_local" "dws_srcip_top_10min_no_params_local" "dws_srcip_top_hour_local" "dws_srcip_top_hour_no_params_local" "dws_user_10min_local" "dws_user_hour_local" "ods_traffic_alarm_log")  # 需要删除数据的ClickHouse表列表
CHECK_INTERVAL="3600*3"                # 检查频率（秒）
IPADDRS=("127.0.0.1" "192.168.5.246")  # 包含多个IP地址的数组
DATABASE="bigdata_fcas_v2"
USERNAME="default"
PASSWORD="123456"  #RootSi314

# 获取当前磁盘使用量
get_disk_usage() {
    df -h "$MOUNT_PATH" | awk 'NR==2 {print $5}' | sed 's/%//'
}

# 查询最老的分区
get_oldest_partition() {
    local ip=$1
    local table_name=$2
    clickhouse-client -h"$ip" -u"$USERNAME" --password="$PASSWORD" -d "$DATABASE" --query "SELECT partition FROM system.parts WHERE database = '$DATABASE' AND table = '$table_name' ORDER BY partition ASC LIMIT 1"
}

# 删除最老的分区
delete_oldest_partition() {
    local ip=$1
    local table_name=$2
    local oldest_partition=$(get_oldest_partition "$ip" "$table_name")
    
    if [ -n "$oldest_partition" ]; then
        clickhouse-client -h"$ip" -u"$USERNAME" --password="$PASSWORD" -d "$DATABASE" --query "ALTER TABLE $table_name DROP PARTITION '$oldest_partition'"
        echo "最老分区数据： $oldest_partition ,表格是： $table_name ,服务器是： $ip"
    else
        echo "该表格：$table_name ,无分区数据，服务器IP为： $ip"
    fi
}

# 主循环
while true; do
    # 获取当前磁盘使用量
    current_usage=$(get_disk_usage)

    # 检查是否超过阈值
    if [ "$current_usage" -ge "$THRESHOLD" ]; then
        echo "该挂载目录：$MOUNT_PATH 的磁盘使用量为：$current_usage%, 已经超过了设定的使用阈值： $THRESHOLD%. 开始删除全部表格的最老的分区数据..."
        
        # 遍历每台大数据服务器
        for ip in "${IPADDRS[@]}"; do
            echo "当前正在执行的服务器IP为： $ip..."
            
            # 遍历每个表并删除最老的分区
            for table in "${TABLES[@]}"; do
                echo "准备删除表格为：$table 的分区数据, 服务器为：$ip..."
                delete_oldest_partition "$ip" "$table"
            done
        done
    else
        echo "该挂载目录：$MOUNT_PATH 的磁盘使用量为：$current_usage%, 目前还未超过阈值： $THRESHOLD%."
    fi

    # 休眠一段时间后继续检查
    sleep "$CHECK_INTERVAL"
done

