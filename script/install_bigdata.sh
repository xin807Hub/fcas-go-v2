#!/bin/bash

workdir=$(pwd)

function update_etl() {
    echo  "更新etl 开始"
    cd $workdir/etl/
    sh install.sh
    echo "更新etl 完成"
}

function update_ck() {
    echo "更新clickhouse 开始"
    clickhouse-client -udefault --password=BoyDB2022 --multiquery < $workdir/clickhouse/clickhouse_dict.sql
    clickhouse-client -udefault --password=BoyDB2022 --multiquery < $workdir/clickhouse/clickhouse_v2.sql
    echo "更新clickhouse 完成"
}

update_etl
update_ck
