#! /bin/bash

clickhouse-client -udefault --password=123456 --multiquery < /root/clickhouse_v2.sql 
clickhouse-client -udefault --password=123456 --multiquery < /root/clickhouse_dict.sql

