#!/bin/bash

workdir=$(pwd)

install -D go_etl config.yaml /home/etl/ # 工作目录兼容老项目

if [ -d "/home/etl/" ];then
    current_time=$(date '+%Y%m%d%H%M%S')
    echo  "备份旧的etl"
    mv /home/etl/ /home/etl_$current_time
fi

mkdir /home/etl/

\cp -f $workdir/go_etl /home/etl/
\cp -f $workdir/config.yaml /home/etl/
\cp -rf $workdir/dict /home/etl/

chmod +x /home/etl/go_etl
\cp -rf $workdir/etl.service /usr/lib/systemd/system/
# \cp -rf $workdir/libzmq.so.5.0.0 /usr/lib64/
# ln -s /usr/lib64/libzmq.so.5.0.0 /usr/lib64/libzmq.so.5

systemctl daemon-reload
systemctl enable etl
systemctl restart etl
systemctl status etl
