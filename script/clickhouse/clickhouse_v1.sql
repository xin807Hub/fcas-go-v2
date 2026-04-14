CREATE DATABASE IF NOT EXISTS bigdata_fcas;
USE bigdata_fcas;

CREATE TABLE IF NOT EXISTS dws_app_type_10min
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_app_type_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_app_type_10min_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (user_id, d_user_id, link_id, isp, d_isp, app_type, start_time, dst_province)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (user_id, d_user_id, link_id, isp, d_isp, app_type, start_time, dst_province);

CREATE TABLE IF NOT EXISTS dws_app_type_hour
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_app_type_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_app_type_hour_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (user_id, d_user_id, link_id, isp, d_isp, app_type, start_time)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (user_id, d_user_id, link_id, isp, d_isp, app_type, start_time, dst_province);

CREATE TABLE IF NOT EXISTS dws_dstip_top_10min
(
    `start_time` DateTime64(3),
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `app_id` Int32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_dstip_top_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_10min_aggr
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `app_id` Int32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_id, d_user_id, link_id, isp, dst_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    d_user_id,
    link_id,
    isp,
    app_id,
    dst_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, d_user_id, link_id, isp, dst_ip);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_10min_aggr_no_params
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, d_user_id, link_id, isp, dst_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    d_user_id,
    link_id,
    isp,
    dst_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);

CREATE TABLE IF NOT EXISTS dws_dstip_top_10min_no_params
(
    `start_time` DateTime64(3),
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_dstip_top_10min_aggr_no_params', rand());

CREATE TABLE IF NOT EXISTS dws_dstip_top_hour
(
    `start_time` DateTime64(3),
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `app_id` Int32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_dstip_top_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_hour_aggr
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `app_id` Int32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_id, d_user_id, link_id, isp, dst_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    d_user_id,
    link_id,
    isp,
    app_id,
    dst_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, d_user_id, link_id, isp, dst_ip);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_hour_aggr_no_params
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, d_user_id, link_id, isp, dst_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    d_user_id,
    link_id,
    isp,
    dst_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);

CREATE TABLE IF NOT EXISTS dws_dstip_top_hour_no_params
(
    `start_time` DateTime64(3),
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_dstip_top_hour_aggr_no_params', rand());

CREATE TABLE IF NOT EXISTS dws_gen_traffic_hour
(
    `start_time` DateTime64(3),
    `src_ip` String,
    `dst_ip` String,
    `src_port` UInt32,
    `dst_port` UInt32,
    `protocol` Int32,
    `app_id` Int32,
    `host` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_gen_traffic_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_gen_traffic_hour_aggr
(
    `start_time` DateTime,
    `src_ip` String,
    `dst_ip` String,
    `src_port` UInt32,
    `dst_port` UInt32,
    `protocol` Int32,
    `app_id` Int32,
    `host` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, src_ip, dst_ip, src_port, dst_port, protocol, app_id, host)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    src_ip,
    dst_ip,
    src_port,
    dst_port,
    protocol,
    app_id,
    host,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, src_ip, dst_ip, src_port, dst_port, protocol, app_id, host);

CREATE TABLE IF NOT EXISTS dws_isp_10min
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `is_oversea` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_isp_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_10min_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `is_oversea` Int8,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, user_id, d_user_id, link_id, isp, d_isp, is_oversea)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    is_oversea,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, d_user_id, link_id, isp, d_isp, is_oversea, dst_province);

CREATE TABLE IF NOT EXISTS dws_isp_hour
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `is_oversea` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_isp_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_hour_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `is_oversea` Int8,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, user_id, d_user_id, link_id, isp, d_isp, is_oversea)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    is_oversea,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, d_user_id, link_id, isp, d_isp, is_oversea, dst_province);

CREATE VIEW IF NOT EXISTS  dws_link_10min
(
    `start_time` DateTime64(3),
    `link_id` Int32,
    `app_traffic_up` UInt64,
    `app_traffic_dn` UInt64,
    `app_traffic_total` UInt64,
    `app_traffic_up_speed` UInt64,
    `app_traffic_dn_speed` UInt64,
    `app_traffic_total_speed` UInt64,
    `date_10min` UInt32,
    `date_hour` UInt32,
    `date_day` UInt32
) AS
SELECT
    start_time,
    link_id,
    toUInt64(sumMerge(bytes_up_view) / 1000000) AS app_traffic_up,
    toUInt64(sumMerge(bytes_dn_view) / 1000000) AS app_traffic_dn,
    app_traffic_up + app_traffic_dn AS app_traffic_total,
    toUInt64(app_traffic_up / 75) AS app_traffic_up_speed,
    toUInt64(app_traffic_dn / 75) AS app_traffic_dn_speed,
    app_traffic_up_speed + app_traffic_dn_speed AS app_traffic_total_speed,
    toUInt32(toYYYYMMDDhhmmss(start_time) / 10000) AS date_10min,
    toUInt32(toYYYYMMDDhhmmss(start_time) / 100) AS date_hour,
    toYYYYMMDD(start_time) AS date_day
FROM dws_link_10min_view
GROUP BY (start_time, link_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_link_10min_aggr
(
    `start_time` DateTime,
    `link_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, link_id)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    link_id,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, link_id);

CREATE TABLE IF NOT EXISTS dws_link_10min_view
(
    `start_time` DateTime64(3),
    `link_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_link_10min_aggr', rand());

CREATE VIEW IF NOT EXISTS  dws_link_app_1min
(
    `start_time` DateTime64(3),
    `link_id` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `app_traffic_up` UInt64,
    `app_traffic_dn` UInt64,
    `app_traffic_total` UInt64,
    `app_traffic_up_speed` UInt64,
    `app_traffic_dn_speed` UInt64,
    `app_traffic_total_speed` UInt64
) AS
SELECT
    start_time,
    link_id,
    app_type,
    app_id,
    toUInt64(sumMerge(bytes_up_view) / 1000) AS app_traffic_up,
    toUInt64(sumMerge(bytes_dn_view) / 1000) AS app_traffic_dn,
    app_traffic_up + app_traffic_dn AS app_traffic_total,
    toUInt64((app_traffic_up / 300) * 8) AS app_traffic_up_speed,
    toUInt64((app_traffic_dn / 300) * 8) AS app_traffic_dn_speed,
    app_traffic_up_speed + app_traffic_dn_speed AS app_traffic_total_speed
FROM dws_link_app_view
GROUP BY (start_time, link_id, app_type, app_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_link_app_aggr
(
    `start_time` DateTime,
    `link_id` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, link_id, app_type, app_id)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 300) AS start_time,
    link_id,
    app_type,
    app_id,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, link_id, app_type, app_id);

CREATE TABLE IF NOT EXISTS dws_link_app_view
(
    `start_time` DateTime64(3),
    `link_id` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_link_app_aggr', rand());

CREATE TABLE IF NOT EXISTS dws_server_10min
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `host` String,
    `src_ip` String,
    `dst_ip` String,
    `dst_area_id` Int32,
    `src_area_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_server_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_10min_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `host` String,
    `src_ip` String,
    `dst_ip` String,
    `dst_area_id` Int32,
    `src_area_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, host, src_ip, dst_ip, dst_area_id, src_area_id)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    app_id,
    host,
    src_ip,
    dst_ip,
    dst_area_id,
    src_area_id,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, host, src_ip, dst_ip, dst_area_id, src_area_id);

CREATE TABLE IF NOT EXISTS dws_server_hour
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `host` String,
    `src_ip` String,
    `dst_ip` String,
    `dst_area_id` Int32,
    `src_area_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_server_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_hour_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `host` String,
    `src_ip` String,
    `dst_ip` String,
    `dst_area_id` Int32,
    `src_area_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, host, src_ip, dst_ip, dst_area_id, src_area_id)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    app_id,
    host,
    src_ip,
    dst_ip,
    dst_area_id,
    src_area_id,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, host, src_ip, dst_ip, dst_area_id, src_area_id);

CREATE TABLE IF NOT EXISTS dws_srcip_top_10min
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `app_id` Int32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_srcip_top_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_10min_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `app_id` Int32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_id, user_id, link_id, d_isp, src_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    link_id,
    d_isp,
    app_id,
    src_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, user_id, link_id, d_isp, src_ip);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_10min_aggr_no_params
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, user_id, link_id, d_isp, src_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    link_id,
    d_isp,
    src_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_10min_no_params
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_srcip_top_10min_aggr_no_params', rand());

CREATE TABLE IF NOT EXISTS dws_srcip_top_hour
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `app_id` Int32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_srcip_top_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_hour_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `app_id` Int32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_id, user_id, link_id, d_isp, src_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    link_id,
    d_isp,
    app_id,
    src_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, user_id, link_id, d_isp, src_ip);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_hour_aggr_no_params
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, user_id, link_id, d_isp, src_ip)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    link_id,
    d_isp,
    src_ip,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_hour_no_params
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `link_id` Int32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_srcip_top_hour_aggr_no_params', rand());

CREATE TABLE IF NOT EXISTS dws_traffic_alarm_log
(
    `start_time` DateTime,
    `policy_id` UInt32,
    `link_id` Int32,
    `app_type_id` Int32,
    `app_id` Int32,
    `app_traffic_total_speed` UInt64,
    `last_period_total_speed` UInt64
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY start_time
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS dws_user_10min
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_user_10min_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_user_10min_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 600) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    app_id,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, dst_province);

CREATE TABLE IF NOT EXISTS dws_user_hour
(
    `start_time` DateTime64(3),
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'dws_user_hour_aggr', rand());

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_user_hour_aggr
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` Int32,
    `isp` String,
    `d_isp` String,
    `app_type` Int32,
    `app_id` Int32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMM(start_time)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
SETTINGS index_granularity = 8192 AS
SELECT
    dateTrunc('hour', o_start_time) AS start_time,
    user_id,
    d_user_id,
    link_id,
    isp,
    d_isp,
    app_type,
    app_id,
    dst_province,
    sumState(bytes_up) AS bytes_up_view,
    sumState(bytes_dn) AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, dst_province);

CREATE TABLE IF NOT EXISTS ods_gen_traffic
(
    `o_start_time` DateTime64(3),
    `user_id` UInt32 DEFAULT 0,
    `d_user_id` UInt32 DEFAULT 0,
    `isp` String,
    `d_isp` String,
    `src_ip` String,
    `dst_ip` String,
    `src_port` UInt32,
    `dst_port` UInt32,
    `protocol` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `bytes_up` UInt64,
    `bytes_dn` UInt64,
    `src_area_id` Int32 DEFAULT 0,
    `dst_area_id` Int32 DEFAULT 0,
    `link_id` Int32,
    `is_oversea` Int8,
    `dst_province` String,
    `host` String
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(o_start_time)
PRIMARY KEY src_ip
ORDER BY src_ip
TTL toDate(o_start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE TABLE IF NOT EXISTS ods_gen_traffic_view
(
    `o_start_time` DateTime64(3),
    `user_id` UInt32 DEFAULT 0,
    `d_user_id` UInt32 DEFAULT 0,
    `isp` String,
    `d_isp` String,
    `src_ip` String,
    `dst_ip` String,
    `src_port` UInt32,
    `dst_port` UInt32,
    `protocol` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `bytes_up` UInt64,
    `bytes_dn` UInt64,
    `src_area_id` Int32 DEFAULT 0,
    `dst_area_id` Int32 DEFAULT 0,
    `link_id` Int32,
    `is_oversea` Int8,
    `dst_province` String,
    `host` String
)
ENGINE = Distributed('fcas_cluster', 'bigdata_fcas', 'ods_gen_traffic', rand());