#!/bin/bash

workdir=$(pwd)

function install_web() {
    echo  "安装web......开始"
    cp -f  $workdir/web/nginx.conf /etc/nginx/
    rm -rf /home/fcas/web/*
    unzip -o -q $workdir/web/*.zip -d /home/fcas/web/
    systemctl restart nginx
    echo "安装web......完成"
}

function install_server() {
    echo "安装fcas_v2_server 开始"
    # chmod +x yaml_tool
    if [ -f /home/fcas/server/config.yaml ]; then
      # 如果文件存在
      # ./yaml_tool --merge backend/config.yaml /home/fcas/server/config.yaml config.yaml
        \cp -f server/config.yaml /home/fcas/server/config.yaml
    else
        mkdir -p /home/fcas/server
        mkdir -p /home/fcas/web
        \cp -rf  server/config.yaml /home/fcas/server/config.yaml
    fi

    \cp -f server/fcas_server /home/fcas/server/
    \cp -f server/fcas_server_d /home/fcas/server/
    \cp -f server/fcas_server.service /usr/lib/systemd/system/
    systemctl daemon-reload
    systemctl enable fcas_server
    systemctl restart fcas_server

    \cp -f ./BUILD_VERSION /home/fcas
    current_time=$(date '+%Y-%m-%d %H:%M:%S')
    echo "$current_time install successful!" >> /home/install_fcas_server_status

    echo "安装fcas_v2_server 完成"
}

function update_mysql() {
  echo  "更新mysql 开始"
  cd $workdir/mysql/
  sh $workdir/mysql/import_sql.sh -c mysql8 -p Mysql@2o20...
  sh $workdir/mysql/import_csv.sh -c mysql8 -p Mysql@2o20... -d fcas_service
  cd -
  echo  "更新mysql 完成"
}

install_server
install_web
update_mysql
