#!/bin/bash
# 启动每秒执行 run.sh 的后台进程
nohup /bin/bash -c 'while true; do sh copy_file.sh; sleep 1; done' > /dev/null 2>&1 &

#mkdir -p  /home/work/copy_file/
echo $! > /home/copy_file_run.pid
