#! /bin/bash

clickhouse-client -udefault --password=BoyDB2022 --multiquery < /root/clickhouse_v2.sql
clickhouse-client -udefault --password=BoyDB2022 --multiquery < /root/clickhouse_dict.sql

