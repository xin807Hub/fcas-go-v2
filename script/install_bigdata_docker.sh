#!/bin/bash

workdir=$(pwd)
container="clickhouse"

function update_etl() {
    echo "更新etl 开始"
    sh $workdir/etl/install.sh
    echo "更新etl 完成"
}

function update_ck() {
    echo "更新clickhouse 开始"
    docker cp $workdir/clickhouse/clickhouse_dict.sql $container:/root/
    docker cp $workdir/clickhouse/clickhouse_v2.sql $container:/root/
    docker cp $workdir/clickhouse/load.sh $container:/root/
    docker exec clickhouse /root/load.sh
    echo "更新clickhouse 完成"
}

update_etl
update_ck
