cd /home/work/

上传 insert_ods_data_per_day.sh 至 "/home/work/"

chmod +x insert_ods_data_per_day.sh

crontab -e

# 每日凌晨1点执行一次
0 1 * * * /home/work/insert_ods_data_per_day.sh >> /home/work/insert_ods_data_per_day.log 2>&1

# 查看
crontab -l
