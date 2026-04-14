#!/bin/bash

# ClickHouse 服务器的地址、用户名和密码
CLICKHOUSE_HOST="localhost"
CLICKHOUSE_PORT=9000
CLICKHOUSE_USER="default"
CLICKHOUSE_PASSWORD="RootSi314"
LIMIT=1000000

# 获取今天的0点时间
START_TIME=$(date +%Y-%m-%d)

# 给出SQL语句的模板
QUERY_TEMPLATE="INSERT INTO bigdata_fcas.ods_gen_traffic (o_start_time, user_id, d_user_id, isp, d_isp, src_ip, dst_ip, src_port, dst_port, protocol, app_type, app_id, bytes_up, bytes_dn, src_area_id, dst_area_id, link_id, is_oversea, dst_province, host)
               SELECT '\$DATE_TIME', user_id, d_user_id, isp, d_isp, src_ip, dst_ip, src_port, dst_port, protocol, app_type, app_id, bytes_up, bytes_dn, src_area_id, dst_area_id, link_id, is_oversea, dst_province, host FROM bigdata_fcas.ods_gen_traffic LIMIT $LIMIT"

# 循环执行24次
for i in {0..23}; do
    # 计算当前时间
    CURRENT_MIN_SEC=$(date +%M:%S)
    CURRENT_HOUR=$(printf "%02d" $i)
    CURRENT_TIME=$(date -d "$START_TIME $CURRENT_HOUR:$CURRENT_MIN_SEC" +"%Y-%m-%d %H:%M:%S")
    # 打印当前时间
    echo "==========================================================================================="
    echo "i=$i"
    echo "START_TIME=$START_TIME"
    echo "CURRENT_TIME=$CURRENT_TIME"

    # 替换模板中的时间变量
    QUERY=$(echo "$QUERY_TEMPLATE" | sed "s/\\\$DATE_TIME/$CURRENT_TIME/")

    # 打印替换模板中的时间变量
#    echo "QUERY= $QUERY"

    # 构建完整的 clickhouse-client 命令
    COMMAND="clickhouse-client --host $CLICKHOUSE_HOST --port $CLICKHOUSE_PORT -u $CLICKHOUSE_USER --password $CLICKHOUSE_PASSWORD -d bigdata_fcas --query \"$QUERY\""

    # 打印将要执行的命令
#    echo "Executing command: $COMMAND"

    # 执行查询并将结果输出到标准输出
    eval $COMMAND

    # 可选：打印当前循环次数和时间戳
    # echo "Iteration $i at $(date)"

    # 等待7min（420秒）
    sleep 420
done

# 获取当前日期和时间
current_datetime=$(date +"%Y-%m-%d %H:%M:%S")

# 向文件中写入 "finish" 和当前日期时间，如果文件不存在则创建
echo "finish insert $LIMIT data to ck....  [$current_datetime]" >> finish.txt