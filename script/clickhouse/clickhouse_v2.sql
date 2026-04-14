-- DROP database if exists bigdata_fcas_v2;
-- USE default;
-- CREATE DATABASE IF NOT EXISTS bigdata_fcas;

CREATE DATABASE IF NOT EXISTS bigdata_fcas_v2;
USE bigdata_fcas_v2;

CREATE TABLE IF NOT EXISTS ods_gen_traffic
(
    `o_start_time` DateTime64(3),
    `record_time`  DateTime64(3) DEFAULT now(),
    `src_ip`       String,
    `dst_ip`       String,
    `app_type`     UInt32,
    `app_id`       UInt32,
    `link_id`      UInt32,
    `host`         String,
    `src_area_id`  UInt32        DEFAULT 0,
    `dst_area_id`  UInt32        DEFAULT 0,
    `user_id`      UInt32        DEFAULT 0,
    `d_user_id`    UInt32        DEFAULT 0,
    `dst_province` String,
    `isp`          String,
    `d_isp`        String,
    `bytes_up`     UInt64,
    `bytes_dn`     UInt64
)
ENGINE = MergeTree
PARTITION BY toYYYYMMDD(o_start_time)
PRIMARY KEY (o_start_time, link_id, user_id, d_user_id,isp, d_isp, app_type, app_id, dst_ip,src_ip)
ORDER BY (o_start_time, link_id, user_id, d_user_id,isp, d_isp, app_type, app_id, dst_ip,src_ip)
TTL toDate(o_start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;
CREATE TABLE IF NOT EXISTS ods_gen_traffic_distributed AS ods_gen_traffic ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2', 'ods_gen_traffic', rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_app_type_10m_local
(
    `start_time` DateTime,
    `link_id` UInt32,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
ORDER BY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_app_type_10m_mv TO bigdata_fcas_v2.dws_app_type_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       link_id,
       user_id,
       d_user_id,
       isp,
       d_isp,
       app_type,
       dst_province,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, link_id, user_id, d_user_id, isp, d_isp, app_type, dst_province);
CREATE TABLE IF NOT EXISTS dws_app_type_10m AS bigdata_fcas_v2.dws_app_type_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_app_type_10m_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_app_type_1h_local
(
    `start_time` DateTime,
    `link_id` UInt32,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
ORDER BY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_app_type_1h_mv to bigdata_fcas_v2.dws_app_type_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       link_id,
       user_id,
       d_user_id,
       isp,
       d_isp,
       app_type,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_user_id, isp, d_isp, app_type, dst_province);
CREATE TABLE IF NOT EXISTS dws_app_type_1h AS bigdata_fcas_v2.dws_app_type_1h_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_app_type_1h_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_app_type_1d_local
(
    `start_time` DateTime,
    `link_id` UInt32,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
ORDER BY (start_time,link_id,user_id,d_user_id,  isp, d_isp, app_type,  dst_province)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_app_type_1d_mv to bigdata_fcas_v2.dws_app_type_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       link_id,
       user_id,
       d_user_id,
       isp,
       d_isp,
       app_type,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_user_id, isp, d_isp, app_type, dst_province);
CREATE TABLE IF NOT EXISTS dws_app_type_1d AS bigdata_fcas_v2.dws_app_type_1d_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_app_type_1d_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_10m_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_id` UInt32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id, app_id, d_user_id,  isp, dst_ip)
ORDER BY (start_time, link_id, app_id, d_user_id,  isp, dst_ip)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_10m_mv to bigdata_fcas_v2.dws_dstip_top_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       app_id,
       dst_ip,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_10m AS bigdata_fcas_v2.dws_dstip_top_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_dstip_top_10m_local', rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_10m_no_params_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id,  d_user_id, isp, dst_ip)
ORDER BY (start_time, link_id,  d_user_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_10m_no_params_mv TO bigdata_fcas_v2.dws_dstip_top_10m_no_params_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       dst_ip,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_10m_no_params ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_dstip_top_10m_no_params_local',rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_1h_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_id` UInt32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id,app_id, d_user_id, isp, dst_ip)
ORDER BY (start_time, link_id, app_id, d_user_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_1h_mv TO bigdata_fcas_v2.dws_dstip_top_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       app_id,
       dst_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_1h AS bigdata_fcas_v2.dws_dstip_top_1h_local  ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_dstip_top_1h_local', rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_1h_no_params_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, d_user_id, link_id, isp, dst_ip)
ORDER BY (start_time, d_user_id, link_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_1h_no_params_mv TO bigdata_fcas_v2.dws_dstip_top_1h_no_params_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       dst_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_1h_no_params ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_dstip_top_1h_no_params_local', rand());



CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_1d_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_id` UInt32,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id,app_id, d_user_id, isp, dst_ip)
ORDER BY (start_time, link_id, app_id, d_user_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_1d_mv TO bigdata_fcas_v2.dws_dstip_top_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       app_id,
       dst_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_1d AS bigdata_fcas_v2.dws_dstip_top_1d_local  ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_dstip_top_1d_local', rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_1d_no_params_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, d_user_id, link_id, isp, dst_ip)
ORDER BY (start_time, d_user_id, link_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_1d_no_params_mv TO bigdata_fcas_v2.dws_dstip_top_1d_no_params_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       dst_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_1d_no_params AS bigdata_fcas_v2.dws_dstip_top_1d_no_params_local ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_dstip_top_1d_no_params_local', rand());


















CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_gen_traffic_1d_local
(
    `start_time` DateTime,
    `src_ip` String,
    `dst_ip` String,
    `app_id` UInt32,
    `host` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, src_ip, dst_ip, app_id, host)
ORDER BY (start_time, src_ip, dst_ip, app_id, host)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_gen_traffic_1d_mv TO bigdata_fcas_v2.dws_gen_traffic_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       src_ip,
       dst_ip,
       app_id,
       host,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, src_ip, dst_ip, app_id, host);
CREATE TABLE IF NOT EXISTS dws_gen_traffic_1d AS bigdata_fcas_v2.dws_gen_traffic_1d_local ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_gen_traffic_1d_local', rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_dstip_top_1d_no_params_local
(
    `start_time` DateTime,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `dst_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, d_user_id, link_id, isp, dst_ip)
ORDER BY (start_time, d_user_id, link_id, isp, dst_ip)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_dstip_top_1d_no_params_mv TO bigdata_fcas_v2.dws_dstip_top_1d_no_params_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       d_user_id,
       link_id,
       isp,
       dst_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, d_user_id, link_id, isp, dst_ip);
CREATE TABLE IF NOT EXISTS dws_dstip_top_1d_no_params AS bigdata_fcas_v2.dws_dstip_top_1d_no_params_local ENGINE = Distributed('fcas_cluster', 'bigdata_fcas_v2', 'dws_dstip_top_1d_no_params_local', rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_isp_10m_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, isp)
ORDER BY (start_time, user_id, link_id, isp)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_10m_src_mv to bigdata_fcas_v2.dws_isp_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       d_user_id as user_id,
       link_id,
       isp,
       dst_province,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id,link_id, isp, dst_province);
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_10m_dst_mv to bigdata_fcas_v2.dws_isp_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp as isp,
       dst_province,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id,link_id, isp, dst_province);
CREATE TABLE IF NOT EXISTS dws_isp_10m AS bigdata_fcas_v2.dws_isp_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_isp_10m_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_isp_1h_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `is_oversea` UInt8,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, isp, is_oversea)
ORDER BY (start_time, user_id, link_id, isp, is_oversea)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_1h_src_mv to bigdata_fcas_v2.dws_isp_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       d_user_id as user_id,
       link_id,
       isp,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, isp, dst_province);
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_1h_dst_mv to bigdata_fcas_v2.dws_isp_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp as isp,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, isp, dst_province);
CREATE TABLE IF NOT EXISTS dws_isp_1h AS  bigdata_fcas_v2.dws_isp_1h_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_isp_1h_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_isp_1d_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `is_oversea` UInt8,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, isp, is_oversea)
ORDER BY (start_time, user_id, link_id, isp, is_oversea)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_1d_src_mv to bigdata_fcas_v2.dws_isp_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       d_user_id as user_id,
       link_id,
       isp,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, isp, dst_province);
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_isp_1d_dst_mv to bigdata_fcas_v2.dws_isp_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp as isp,
       dst_province,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, isp, dst_province);
CREATE TABLE IF NOT EXISTS dws_isp_1d AS bigdata_fcas_v2.dws_isp_1d_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_isp_1d_local',rand());


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_link_10m_local
(
    `start_time` DateTime,
    `link_id` UInt32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id)
ORDER BY (start_time, link_id)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS bigdata_fcas_v2.dws_link_10m_mv TO dws_link_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       link_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, link_id);

CREATE TABLE IF NOT EXISTS dws_link_10m AS bigdata_fcas_v2.dws_link_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_link_10m_local',rand());

CREATE VIEW IF NOT EXISTS dws_link_10m_view
AS
SELECT start_time,
       link_id,
       toUInt64(sumMerge(bytes_up_view)) AS up_byte,
       toUInt64(sumMerge(bytes_dn_view)) AS dn_byte,
       up_byte + dn_byte                 AS total_byte,
       toUInt64(up_byte * 8 / 600)       AS up_bps,
       toUInt64(dn_byte * 8 / 600)       AS dn_bps,
       up_bps + dn_bps                   AS total_bps
FROM bigdata_fcas_v2.dws_link_10m
GROUP BY (start_time, link_id);

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_link_app_local
(
    `start_time` DateTime,
    `link_id` Int32,
    `app_type` Int32,
    `app_id` Int32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, link_id, app_type, app_id)
ORDER BY (start_time, link_id, app_type, app_id)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_link_app_mv TO bigdata_fcas_v2.dws_link_app_local
AS
SELECT dateTrunc('second', o_start_time) - (toUInt64(o_start_time) % 300) AS start_time, -- 5min
       link_id,
       app_type,
       app_id,
       sumState(bytes_up)                                                 AS bytes_up_view,
       sumState(bytes_dn)                                                 AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, link_id, app_type, app_id);

CREATE TABLE IF NOT EXISTS dws_link_app_view AS bigdata_fcas_v2.dws_link_app_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_link_app_local',rand());

-- 用于告警的业务
CREATE VIEW IF NOT EXISTS dws_link_app_5min AS
SELECT start_time,
       link_id,
       app_type,
       app_id,
       toUInt64(sumMerge(bytes_up_view) / 1000)    AS app_traffic_up,
       toUInt64(sumMerge(bytes_dn_view) / 1000)    AS app_traffic_dn,
       app_traffic_up + app_traffic_dn             AS app_traffic_total,
       toUInt64((app_traffic_up / 300) * 8)        AS app_traffic_up_speed,
       toUInt64((app_traffic_dn / 300) * 8)        AS app_traffic_dn_speed,
       app_traffic_up_speed + app_traffic_dn_speed AS app_traffic_total_speed
FROM bigdata_fcas_v2.dws_link_app_view
GROUP BY (start_time, link_id, app_type, app_id);


CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_server_10m_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `host` String,
    `dst_ip` String,
    `dst_area_id` UInt32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id,user_id, link_id,isp, host,dst_ip,dst_area_id)
ORDER BY (start_time, app_type, app_id,user_id, link_id,isp, host,dst_ip,dst_area_id)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_10m_src_mv TO bigdata_fcas_v2.dws_server_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       d_user_id AS user_id,
       link_id,
       isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp,host, dst_ip, dst_area_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_10m_dst_mv TO bigdata_fcas_v2.dws_server_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp AS isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp, host, dst_ip, dst_area_id);

CREATE TABLE IF NOT EXISTS dws_server_10m as bigdata_fcas_v2.dws_server_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_server_10m_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_server_1h_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `host` String,
    `dst_ip` String,
    `dst_area_id` UInt32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id,user_id,link_id,isp, host,dst_ip,dst_area_id)
ORDER BY (start_time, app_type, app_id,user_id,link_id,isp, host,dst_ip,dst_area_id)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_1h_src_mv TO bigdata_fcas_v2.dws_server_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       d_user_id AS user_id,
       link_id,
       isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp,host, dst_ip, dst_area_id);

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_1h_dst_mv TO bigdata_fcas_v2.dws_server_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp AS isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp, host, dst_ip, dst_area_id);

CREATE TABLE IF NOT EXISTS dws_server_1h AS bigdata_fcas_v2.dws_server_1h_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_server_1h_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_server_1d_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `dst_area_id` UInt32,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id,user_id, link_id,isp, dst_area_id)
ORDER BY (start_time, app_type, app_id,user_id, link_id,isp, dst_area_id)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_1d_src_mv TO bigdata_fcas_v2.dws_server_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       d_user_id AS user_id,
       link_id,
       isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp,host, dst_ip, dst_area_id);
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_server_1h_dst_mv TO bigdata_fcas_v2.dws_server_1h_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp AS isp,
       app_type,
       app_id,
       host,
       dst_ip,
       dst_area_id,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, link_id, isp, host, dst_ip, dst_area_id);
CREATE TABLE IF NOT EXISTS dws_server_1d AS bigdata_fcas_v2.dws_server_1d_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_server_1d_local',rand());
CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_10m_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `app_id` UInt32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_id, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, app_id, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_10m_mv to bigdata_fcas_v2.dws_srcip_top_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       app_id,
       src_ip,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, user_id, link_id, d_isp, src_ip);
CREATE TABLE IF NOT EXISTS dws_srcip_top_10m As bigdata_fcas_v2.dws_srcip_top_10m_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_10m_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_10m_no_params_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_10m_no_params_mv to bigdata_fcas_v2.dws_srcip_top_10m_no_params_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       src_ip,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_10m_no_params AS bigdata_fcas_v2.dws_srcip_top_10m_no_params_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_10m_no_params_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_1h_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `app_id` UInt32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_id, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, app_id, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_1h_mv TO bigdata_fcas_v2.dws_srcip_top_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       app_id,
       src_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_1h as bigdata_fcas_v2.dws_srcip_top_1h_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_1h_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_1h_no_params_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_1h_no_params_mv to bigdata_fcas_v2.dws_srcip_top_1h_no_params_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       src_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_1h_no_params as bigdata_fcas_v2.dws_srcip_top_1h_no_params_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_1h_no_params_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_1d_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `app_id` UInt32,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_id, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, app_id, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_1d_mv TO bigdata_fcas_v2.dws_srcip_top_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       app_id,
       src_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_id, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_1d as bigdata_fcas_v2.dws_srcip_top_1d_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_1d_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_srcip_top_1d_no_params_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `link_id` UInt32,
    `d_isp` String,
    `src_ip` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, user_id, link_id, d_isp, src_ip)
ORDER BY (start_time, user_id, link_id, d_isp, src_ip)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_srcip_top_1d_mv_no_params to bigdata_fcas_v2.dws_srcip_top_1d_no_params_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       user_id,
       link_id,
       d_isp,
       src_ip,
       sumState(bytes_up)          AS bytes_up_view,
       sumState(bytes_dn)          AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, user_id, link_id, d_isp, src_ip);

CREATE TABLE IF NOT EXISTS dws_srcip_top_1d_no_params as bigdata_fcas_v2.dws_srcip_top_1d_no_params_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_srcip_top_1d_no_params_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_user_10m_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
TTL toDate(start_time) + toIntervalDay(7)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_user_10m_mv to bigdata_fcas_v2.dws_user_10m_local
AS
SELECT toStartOfTenMinutes(o_start_time) AS start_time,
       user_id,
       d_user_id,
       link_id,
       isp,
       d_isp,
       app_type,
       app_id,
       dst_province,
       sumState(bytes_up)                AS bytes_up_view,
       sumState(bytes_dn)                AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, dst_province);

CREATE TABLE IF NOT EXISTS dws_user_10m AS bigdata_fcas_v2.dws_user_10m_local ENGINE  = Distributed('fcas_cluster','bigdata_fcas_v2','dws_user_10m_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_user_1h_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
TTL toDate(start_time) + toIntervalMonth(1)
SETTINGS index_granularity = 8192;

CREATE MATERIALIZED VIEW IF NOT EXISTS dws_user_1h_mv TO bigdata_fcas_v2.dws_user_1h_local
AS
SELECT toStartOfHour(o_start_time) AS start_time,
       user_id,
       d_user_id,
       link_id,
       isp,
       d_isp,
       app_type,
       app_id,
       dst_province,
       sumState(bytes_up)              AS bytes_up_view,
       sumState(bytes_dn)              AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, dst_province);

CREATE TABLE IF NOT EXISTS dws_user_1h AS bigdata_fcas_v2.dws_user_1h_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_user_1h_local',rand());

CREATE TABLE IF NOT EXISTS bigdata_fcas_v2.dws_user_1d_local
(
    `start_time` DateTime,
    `user_id` UInt32,
    `d_user_id` UInt32,
    `link_id` UInt32,
    `isp` String,
    `d_isp` String,
    `app_type` UInt32,
    `app_id` UInt32,
    `dst_province` String,
    `bytes_up_view` AggregateFunction(sum, UInt64),
    `bytes_dn_view` AggregateFunction(sum, UInt64)
)
ENGINE = AggregatingMergeTree
PARTITION BY toYYYYMMDD(start_time)
PRIMARY KEY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
ORDER BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp)
TTL toDate(start_time) + toIntervalMonth(6)
SETTINGS index_granularity = 8192;
CREATE MATERIALIZED VIEW IF NOT EXISTS dws_user_1d_mv TO bigdata_fcas_v2.dws_user_1d_local
AS
SELECT toStartOfDay(o_start_time) AS start_time,
       user_id,
       d_user_id,
       link_id,
       isp,
       d_isp,
       app_type,
       app_id,
       dst_province,
       sumState(bytes_up)              AS bytes_up_view,
       sumState(bytes_dn)              AS bytes_dn_view
FROM ods_gen_traffic
GROUP BY (start_time, app_type, app_id, user_id, d_user_id, link_id, isp, d_isp, dst_province);
CREATE TABLE IF NOT EXISTS dws_user_1d AS bigdata_fcas_v2.dws_user_1d_local ENGINE = Distributed('fcas_cluster','bigdata_fcas_v2','dws_user_1d_local',rand());

-- BoyDB2022
-- CREATE DICTIONARY IF NOT EXISTS bigdata_fcas_v2.app_type_dict
-- (
--
--     `id` UInt8,
--
--     `name` String
-- )
-- PRIMARY KEY id
-- SOURCE(MYSQL(PORT 3306 USER 'root' PASSWORD '123456' HOST '127.0.0.1' DB 'fcas_service' TABLE 'app_type_view'))
-- LIFETIME(MIN 1 MAX 3)
-- LAYOUT(HASHED());
--
-- -- BoyDB2022
-- CREATE DICTIONARY IF NOT EXISTS bigdata_fcas_v2.app_id_dict
-- (
--     `id` UInt8,
--     `name` String
-- )
-- PRIMARY KEY id
-- SOURCE(MYSQL(PORT 3306 USER 'root' PASSWORD '123456' HOST '127.0.0.1' DB 'fcas_service' TABLE 'app_id_view'))
-- LIFETIME(MIN 1 MAX 3)
-- LAYOUT(HASHED());
