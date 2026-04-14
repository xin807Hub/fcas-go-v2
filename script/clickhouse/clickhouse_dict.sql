CREATE DATABASE IF NOT EXISTS bigdata_fcas_v2;
USE bigdata_fcas_v2;

CREATE DICTIONARY IF NOT EXISTS link_dict
(
    `id`   UInt64,
    `name` String
)
    PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'link_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS app_type_dict
(
    `id`   UInt64,
    `name` String
)
    PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'app_type_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS app_id_dict
(
    `id`   UInt64,
    `name` String
)
    PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'app_id_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS isp_dict
(
    `id`   UInt64,
    `name` String
) PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'isp_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS isp_oversea_dict
(
    id   UInt8,
    name String
) PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'dim_isp_oversea' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS province_dict
(
    `id`   UInt64,
    `name` String
) PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'province_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS city_dict
(
    `id`   UInt64,
    `name` String
)
    PRIMARY KEY id
    SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'city_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());

CREATE DICTIONARY IF NOT EXISTS isp_oversea_dict
(
    `id`   UInt64,
    `is_oversea` UInt8
)
    PRIMARY KEY id
	SOURCE (MYSQL(
            DB 'fcas_service'
            TABLE 'is_oversea_view' -- 使用视图名
            USER 'root'
            PASSWORD '123456'
            HOST '127.0.0.1'
            PORT 3306))
    LIFETIME (MIN 0 MAX 300)
    LAYOUT (HASHED());