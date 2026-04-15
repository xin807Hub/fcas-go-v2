/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.3.232-mysql
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : 192.168.3.232:3306
 Source Schema         : fcas_service

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 05/03/2025 12:11:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS fcas_service;
CREATE DATABASE IF NOT EXISTS fcas_service;
USE fcas_service;

-- ----------------------------
-- Table structure for biz_ip_address
-- ----------------------------
DROP TABLE IF EXISTS `biz_ip_address`;
CREATE TABLE `biz_ip_address`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip_start` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `ip_end` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `country` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `unknown` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `province` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `city` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `isp` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `isp_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for dim_app_classify
-- ----------------------------
DROP TABLE IF EXISTS `dim_app_classify`;
CREATE TABLE `dim_app_classify`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `app_type_id` int NOT NULL,
  `app_type_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `app_id` int NOT NULL,
  `app_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for biz_isp
-- ----------------------------
DROP TABLE IF EXISTS `biz_isp`;
CREATE TABLE `biz_isp`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `isp_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_oversea` tinyint(1) NOT NULL COMMENT '1国外；0国内',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_app_classify
-- ----------------------------
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (1, 2, '视频业务', 20164, '抖音短视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (2, 19, '浏览下载', 190002, '腾讯门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (3, 22, '其他网络应用', 229999, '其他', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (4, 1, '即时通信', 10002, '微信', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (5, 12, '应用商店', 120005, '苹果', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (6, 15, '直播业务', 159999, '抖音直播', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (7, 13, '网上商城', 130008, '淘宝', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (8, 2, '视频业务', 20049, '快手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (9, 19, '浏览下载', 190017, '今日头条', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (10, 19, '浏览下载', 190682, '字节跳动科技', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (11, 2, '视频业务', 20052, '哔哩哔哩', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (12, 2, '视频业务', 20001, '腾讯', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (13, 13, '网上商城', 130016, '小红书', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (14, 17, '云盘服务', 170023, '华为', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (15, 2, '视频业务', 20003, '爱奇艺PPS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (16, 19, '浏览下载', 190071, '百度门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (17, 13, '网上商城', 130024, '拼多多', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (18, 2, '视频业务', 20002, '优酷土豆视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (19, 22, '其他网络应用', 220123, 'Microsoft', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (20, 15, '直播业务', 150021, '小米直播', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (21, 22, '其他网络应用', 220033, 'JPush', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (22, 1, '即时通信', 10001, 'QQ', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (23, 22, '其他网络应用', 220036, 'AdMaster', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (24, 19, '浏览下载', 190370, 'VIVO', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (25, 22, '其他网络应用', 220090, 'echo', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (26, 22, '其他网络应用', 220210, '迅游', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (27, 22, '其他网络应用', 220091, 'discard', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (28, 22, '其他网络应用', 220092, 'daytime', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (29, 3, 'P2P业务', 30003, 'BT', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (30, 22, '其他网络应用', 220225, '爱贝云计费', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (31, 22, '其他网络应用', 220230, 'SNTP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (32, 22, '其他网络应用', 220080, 'Rsync', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (33, 1, '即时通信', 10013, '陌陌', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (34, 22, '其他网络应用', 220161, 'Gopher', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (35, 22, '其他网络应用', 220138, '多盟', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (36, 22, '其他网络应用', 220012, 'TalKingData', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (37, 22, '其他网络应用', 220089, 'whois', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (38, 22, '其他网络应用', 220001, 'FTP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (39, 22, '其他网络应用', 220086, 'BGP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (40, 22, '其他网络应用', 220208, '美洽', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (41, 22, '其他网络应用', 220155, 'Ident', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (42, 16, '财经支付', 160003, '支付宝', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (43, 22, '其他网络应用', 220079, 'DHCP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (44, 22, '其他网络应用', 220242, 'MQTT', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (45, 22, '其他网络应用', 220030, 'Speedtest', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (46, 22, '其他网络应用', 220031, 'CNZZ统计', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (47, 22, '其他网络应用', 220040, '微小宝', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (48, 22, '其他网络应用', 220014, 'inmobi', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (49, 2, '视频业务', 20018, '芒果tv', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (50, 22, '其他网络应用', 220172, 'Time', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (51, 17, '云盘服务', 170006, '金山', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (52, 13, '网上商城', 130002, '大众点评美团', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (53, 22, '其他网络应用', 220142, 'smb', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (54, 2, '视频业务', 20149, '阳光宽频网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (55, 22, '其他网络应用', 220238, 'MediaV聚胜万合', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (56, 22, '其他网络应用', 220068, 'Exchange', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (57, 22, '其他网络应用', 220082, 'NETBIOS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (58, 22, '其他网络应用', 220071, 'Telnet', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (59, 22, '其他网络应用', 220171, 'RIP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (60, 2, '视频业务', 20193, 'CIBN互联网电视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (61, 22, '其他网络应用', 220224, '乐变游戏服务平台', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (62, 12, '应用商店', 120026, '小米商店', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (63, 22, '其他网络应用', 220019, 'Reddit', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (64, 22, '其他网络应用', 220129, '友盟', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (65, 22, '其他网络应用', 220206, 'WPSOffice', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (66, 11, '音乐', 110005, '喜马拉雅FM', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (67, 22, '其他网络应用', 220130, 'flurryAPP分析', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (68, 9, '微博社区', 90001, '新浪微博论坛', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (69, 15, '直播业务', 150002, '斗鱼', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (70, 22, '其他网络应用', 220231, 'NETCONF', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (71, 22, '其他网络应用', 220151, 'Modbus', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (72, 22, '其他网络应用', 220093, 'finger', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (73, 22, '其他网络应用', 220133, '扫描全能王', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (74, 22, '其他网络应用', 220029, '米柚', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (75, 2, '视频业务', 20051, '移动视频业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (76, 13, '网上商城', 130003, '京东', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (77, 22, '其他网络应用', 220067, 'TFTP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (78, 22, '其他网络应用', 220002, 'WiFI万能钥匙', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (79, 22, '其他网络应用', 220084, 'SSDP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (80, 22, '其他网络应用', 220174, 'LPD', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (81, 22, '其他网络应用', 220190, '安沃广告', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (82, 22, '其他网络应用', 220096, 'syslog', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (83, 22, '其他网络应用', 220223, '金山电池医生', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (84, 22, '其他网络应用', 220173, 'iSCSI', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (85, 5, '游戏业务', 50048, 'QQ', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (86, 17, '云盘服务', 170001, '百度网盘', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (87, 22, '其他网络应用', 220070, 'SSH', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (88, 22, '其他网络应用', 220065, 'DNS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (89, 3, 'P2P业务', 30001, '迅雷', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (90, 22, '其他网络应用', 220085, 'SNMP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (91, 22, '其他网络应用', 220095, 'MDNS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (92, 22, '其他网络应用', 220191, 'AdMob广告', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (93, 10, '地图出行', 100003, '高德导航', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (94, 22, '其他网络应用', 220245, '生日管家', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (95, 22, '其他网络应用', 220187, '百度移动广告联盟', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (96, 22, '其他网络应用', 220177, 'Rsh', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (97, 1, '即时通信', 10005, 'MSN', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (98, 13, '网上商城', 130014, '饿了么', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (99, 19, '浏览下载', 190294, '中国联通手机营业厅', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (100, 17, '云盘服务', 170021, '阿里云', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (101, 19, '浏览下载', 190024, '阿里巴巴旗下业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (102, 11, '音乐', 110006, '网易云音乐', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (103, 22, '其他网络应用', 220175, 'Chargen', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (104, 22, '其他网络应用', 220026, '触宝输入法', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (105, 1, '即时通信', 10048, 'viber', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (106, 16, '财经支付', 160167, '东方财富旗下产品', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (107, 18, '安全杀毒', 180001, '360安全', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (108, 2, '视频业务', 20124, '央视新闻', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (109, 17, '云盘服务', 170035, 'iCloud', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (110, 15, '直播业务', 150022, '虎牙直播', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (111, 8, '阅读学习', 80002, '闲鱼', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (112, 2, '视频业务', 20209, '创维酷开电视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (113, 19, '浏览下载', 190656, 'Ksosoft', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (114, 11, '音乐', 110001, '腾讯', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (115, 22, '其他网络应用', 220088, 'NNTP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (116, 2, '视频业务', 20165, '西瓜视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (117, 2, '视频业务', 20090, '华数TV', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (118, 19, '浏览下载', 190093, '谷歌门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (119, 19, '浏览下载', 190049, '360门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (120, 2, '视频业务', 20006, 'PPTV', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (121, 19, '浏览下载', 190043, '新浪门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (122, 5, '游戏业务', 59006, 'PlayStation', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (123, 4, 'VoIP业务', 40041, '钉钉', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (124, 1, '即时通信', 10042, '快信', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (125, 8, '阅读学习', 80022, '掌阅', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (126, 5, '游戏业务', 59025, '和平精英UDP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (127, 5, '游戏业务', 59007, '原神', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (128, 22, '其他网络应用', 230001, 'OPPO应用', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (129, 2, '视频业务', 20061, '网易视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (130, 8, '阅读学习', 80183, '一点资讯', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (131, 5, '游戏业务', 50086, 'JJ', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (132, 11, '音乐', 110013, '酷狗音乐', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (133, 22, '其他网络应用', 220021, 'Xbox_Live', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (134, 1, '即时通信', 10008, 'YY', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (135, 4, 'VoIP业务', 40015, 'NET2PHONE', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (136, 19, '浏览下载', 190302, '联想', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (137, 14, '影像处理', 140003, '美图科技', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (138, 19, '浏览下载', 190225, '最美天气', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (139, 7, '电信自营业务', 70008, '天翼网盘(天翼云)', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (140, 19, '浏览下载', 190001, '汽车之家', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (141, 16, '财经支付', 160001, '同花顺', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (142, 5, '游戏业务', 50022, '王者荣耀', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (143, 16, '财经支付', 160009, '招商', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (144, 17, '云盘服务', 170002, '115网盘', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (145, 12, '应用商店', 120006, '智汇云(华为应用市场)', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (146, 11, '音乐', 110007, '全民K歌', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (147, 19, '浏览下载', 190041, '网易门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (148, 2, '视频业务', 20199, 'GiTV影视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (149, 13, '网上商城', 130058, '三星', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (150, 2, '视频业务', 20004, '乐视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (151, 10, '地图出行', 100004, '携程网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (152, 1, '即时通信', 10020, 'BLUED', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (153, 12, '应用商店', 120001, '360手机助手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (154, 8, '阅读学习', 80001, '知乎', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (155, 19, '浏览下载', 190051, '小米旗下业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (156, 10, '地图出行', 100010, '12306订票网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (157, 10, '地图出行', 100007, '腾讯地图', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (158, 2, '视频业务', 20122, '百视通', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (159, 22, '其他网络应用', 220334, '互传', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (160, 8, '阅读学习', 80201, '南方周末', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (161, 5, '游戏业务', 59031, '叠纸游戏科技', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (162, 2, '视频业务', 20102, '微视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (163, 16, '财经支付', 160020, '工商银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (164, 5, '游戏业务', 50530, 'TapTap', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (165, 12, '应用商店', 120045, '联想乐商店', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (166, 19, '浏览下载', 190045, '搜狐门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (167, 16, '财经支付', 160040, '腾讯', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (168, 16, '财经支付', 160013, '农业银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (169, 12, '应用商店', 120003, '百度手机助手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (170, 1, '即时通信', 10047, 'twitter', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (171, 2, '视频业务', 20141, 'Gif快手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (172, 5, '游戏业务', 50004, '魔兽世界', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (173, 12, '应用商店', 120013, '应用宝', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (174, 5, '游戏业务', 50503, 'Steam平台', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (175, 2, '视频业务', 20005, '搜狐视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (176, 1, '即时通信', 10038, '腾讯通', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (177, 13, '网上商城', 130013, '唯品会', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (178, 8, '阅读学习', 80087, '新东方在线', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (179, 19, '浏览下载', 190144, '搜狗旗下产品', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (180, 14, '影像处理', 140007, 'faceu', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (181, 6, '邮件服务', 60014, 'IMAPS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (182, 16, '财经支付', 160016, '交通银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (183, 16, '财经支付', 160064, '第一财经', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (184, 5, '游戏业务', 50233, '九游', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (185, 19, '浏览下载', 190007, '掌上链家', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (186, 16, '财经支付', 160149, '浦发银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (187, 19, '浏览下载', 190009, '墨迹天气', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (188, 2, '视频业务', 20047, '央视影音', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (189, 12, '应用商店', 120067, 'pp助手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (190, 18, '安全杀毒', 180009, '金山旗下产品', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (191, 16, '财经支付', 160002, '大智慧旗下产品', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (192, 16, '财经支付', 160159, '国泰君安国际', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (193, 1, '即时通信', 10016, 'WhatsApp', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (194, 8, '阅读学习', 80145, '互动百科', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (195, 19, '浏览下载', 190118, '搜房网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (196, 2, '视频业务', 20021, 'CNTV', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (197, 19, '浏览下载', 190267, '华为天气', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (198, 2, '视频业务', 20067, '优米网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (199, 5, '游戏业务', 59060, '奇妙游乐园世界', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (200, 10, '地图出行', 100045, '华住', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (201, 11, '音乐', 110002, '酷我音乐', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (202, 2, '视频业务', 20016, '风行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (203, 19, '浏览下载', 190042, '凤凰门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (204, 15, '直播业务', 150001, '六间房', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (205, 10, '地图出行', 100100, '哈罗单车Hellobike', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (206, 19, '浏览下载', 190150, '寻医问药', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (207, 2, '视频业务', 20117, '360影视大全', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (208, 2, '视频业务', 20119, '影视大全', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (209, 16, '财经支付', 160004, '建设银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (210, 9, '微博社区', 90002, '豆瓣网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (211, 6, '邮件服务', 60004, 'QQ邮箱', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (212, 18, '安全杀毒', 180026, '迈克菲杀毒', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (213, 5, '游戏业务', 50038, '迷你世界', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (214, 19, '浏览下载', 190019, '作业帮', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (215, 11, '音乐', 110077, '阿基米德FM', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (216, 5, '游戏业务', 50032, '腾讯游戏', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (217, 5, '游戏业务', 50019, '开心消消乐', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (218, 9, '微博社区', 90006, '腾讯', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (219, 19, '浏览下载', 190014, 'Keep', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (220, 16, '财经支付', 160015, '快钱', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (221, 5, '游戏业务', 50327, '糖果苏打传奇', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (222, 19, '浏览下载', 190094, '必应', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (223, 8, '阅读学习', 80011, '移动手机阅读&和阅读&咪咕阅读', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (224, 10, '地图出行', 100005, '滴滴出行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (225, 17, '云盘服务', 170025, '七牛云存储', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (226, 8, '阅读学习', 80009, '起点', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (227, 13, '网上商城', 130009, '苏宁易购', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (228, 8, '阅读学习', 80014, '书旗网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (229, 19, '浏览下载', 190383, '网宿科技', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (230, 8, '阅读学习', 80031, '美篇', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (231, 13, '网上商城', 130004, '亚马逊', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (232, 6, '邮件服务', 60030, 'Outlook', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (233, 5, '游戏业务', 50493, '崩坏3', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (234, 10, '地图出行', 100072, '东航移动E', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (235, 19, '浏览下载', 190082, '世纪佳缘', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (236, 5, '游戏业务', 50392, '碧蓝航线', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (237, 8, '阅读学习', 80073, '金山词霸', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (238, 10, '地图出行', 100002, '去哪儿', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (239, 6, '邮件服务', 60013, 'SMTPS', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (240, 15, '直播业务', 150025, '直播吧', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (241, 5, '游戏业务', 59026, '233乐园', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (242, 19, '浏览下载', 190004, '58同城旗下业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (243, 16, '财经支付', 160005, '银联', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (244, 4, 'VoIP业务', 40013, 'SIP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (245, 19, '浏览下载', 190226, '虎扑体育', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (246, 13, '网上商城', 130063, '海信', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (247, 12, '应用商店', 120035, '搜狗', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (248, 15, '直播业务', 150008, '火山直播', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (249, 19, '浏览下载', 190027, '中国移动门户业务', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (250, 8, '阅读学习', 80221, '多看阅读', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (251, 11, '音乐', 110008, '蜻蜓FM', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (252, 16, '财经支付', 160008, '中国银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (253, 18, '安全杀毒', 180022, '卡巴斯基杀毒', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (254, 16, '财经支付', 160093, '益盟操盘手', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (255, 9, '微博社区', 90012, '人民网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (256, 6, '邮件服务', 60003, '网易邮箱', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (257, 1, '即时通信', 10057, '探探', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (258, 2, '视频业务', 20178, '彩视', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (259, 9, '微博社区', 90039, 'LOFTER', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (260, 5, '游戏业务', 50075, '完美世界', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (261, 15, '直播业务', 150020, '咪咕直播', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (262, 19, '浏览下载', 190556, '学而思培优', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (263, 11, '音乐', 110061, '荔枝FM', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (264, 17, '云盘服务', 170003, 'QQ', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (265, 13, '网上商城', 130072, '华为商城', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (266, 3, 'P2P业务', 30002, '电驴', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (267, 5, '游戏业务', 50081, '暴雪战网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (268, 13, '网上商城', 130071, '索尼商城', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (269, 19, '浏览下载', 190299, '汽车报价大全', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (270, 19, '浏览下载', 190005, '安居客', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (271, 19, '浏览下载', 190028, '趣头条', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (272, 6, '邮件服务', 60012, 'POP3S', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (273, 5, '游戏业务', 59033, '白日梦游戏', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (274, 16, '财经支付', 160075, '财联社', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (275, 19, '浏览下载', 190176, '天气网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (276, 2, '视频业务', 20101, 'CC视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (277, 19, '浏览下载', 190399, '中国青年网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (278, 19, '浏览下载', 190269, '讯飞输入法', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (279, 7, '电信自营业务', 70015, '电信营业厅', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (280, 13, '网上商城', 130019, '考拉海购', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (281, 5, '游戏业务', 50090, '盛大', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (282, 18, '安全杀毒', 180031, '鲁大师', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (283, 13, '网上商城', 130087, '亲宝宝', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (284, 5, '游戏业务', 59013, '香肠派对', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (285, 2, '视频业务', 20143, '萤石云视频', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (286, 6, '邮件服务', 60011, 'IMAP', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (287, 8, '阅读学习', 80018, '微信读书', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (288, 19, '浏览下载', 190167, '当乐网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (289, 8, '阅读学习', 80124, '中国军网', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (290, 15, '直播业务', 150041, '映客', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (291, 5, '游戏业务', 59003, '三七玩', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (292, 16, '财经支付', 160010, '邮政储蓄银行', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (293, 2, '视频业务', 20216, 'GZSTV', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (294, 5, '游戏业务', 50093, '巨人征途', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (295, 2, '视频业务', 20020, '酷6', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (296, 11, '音乐', 110023, '口袋故事听听', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (297, 19, '浏览下载', 190026, '交管12123', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (298, 8, '阅读学习', 80198, '澎湃新闻', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (299, 2, '视频业务', 20014, '皮皮', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (300, 12, '应用商店', 120044, '魅族应用中心', NULL);
INSERT INTO `dim_app_classify` (`id`, `app_type_id`, `app_type_name`, `app_id`, `app_name`, `deleted_at`) VALUES (301, 20, 'VPN类应用', 200002, 'PPTP', NULL);


-- ----------------------------
-- Table structure for dim_bypass
-- ----------------------------
DROP TABLE IF EXISTS `dim_bypass`;
CREATE TABLE `dim_bypass`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `olp_id` int NULL DEFAULT NULL COMMENT 'bypass编号',
  `bypass_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分流器名称',
  `bypass_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '分流器ip',
  `bypass_port` int NULL DEFAULT NULL COMMENT '分流器交互端口',
  `remark` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- Table structure for dim_control_policy
-- ----------------------------
DROP TABLE IF EXISTS `dim_control_policy`;
CREATE TABLE `dim_control_policy`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '策略名称',
  `user_type` tinyint NULL DEFAULT NULL COMMENT '1：用户群组 2：用户群 3：用户',
  `user_crowd_group_id` int NULL DEFAULT NULL COMMENT '用户群组ID',
  `user_crowd_id` int NULL DEFAULT NULL COMMENT '用户群ID',
  `user_id` int NULL DEFAULT NULL COMMENT '用户ID',
  `ul_flow_rate` int NULL DEFAULT NULL COMMENT '上行限速Mbps',
  `dl_flow_rate` int NULL DEFAULT NULL COMMENT '下行限速Mbps',
  `start_time` datetime NOT NULL COMMENT '生效时间',
  `end_time` datetime NOT NULL COMMENT '失效时间',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '备注',
  `flow_ctrl_type` int NOT NULL COMMENT '流控类型，1：应用大小类 2：目的地址',
  `app_type_id` int NULL DEFAULT NULL COMMENT '应用大类ID',
  `app_id` int NULL DEFAULT NULL COMMENT '应用小类ID',
  `ip_data`             json         null,
  `period_type` int NULL DEFAULT NULL COMMENT '周期类型，1：每日 2：每周',
  `policy_period` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '周期时间，多个以逗号分隔，开始和结束之间以-分隔',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `link_ids` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_control_policy
-- ----------------------------

-- ----------------------------
-- Table structure for dim_control_policy_action
-- ----------------------------
DROP TABLE IF EXISTS `dim_control_policy_action`;
CREATE TABLE `dim_control_policy_action`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `policy_id` int NULL DEFAULT NULL,
  `vlan_id` int NULL DEFAULT NULL,
  `shunt_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '分流器IP',
  `upload_action_id` int NULL DEFAULT NULL COMMENT '上行actionID',
  `download_action_id` int NULL DEFAULT NULL COMMENT '下行actionID',
  `upload_device_id` int NULL DEFAULT NULL COMMENT '上报设备ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for dim_control_policy_log
-- ----------------------------
DROP TABLE IF EXISTS `dim_control_policy_log`;
CREATE TABLE `dim_control_policy_log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `record_time` timestamp NULL DEFAULT NULL,
  `policy_id` int NULL DEFAULT NULL COMMENT '策略ID',
  `up_traffic` bigint NULL DEFAULT NULL,
  `dn_traffic` bigint NULL DEFAULT NULL,
  `up_pass` bigint NULL DEFAULT NULL,
  `dn_pass` bigint NULL DEFAULT NULL,
  `up_discard` bigint NULL DEFAULT NULL,
  `dn_discard` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_control_policy_log
-- ----------------------------

-- ----------------------------
-- Table structure for dim_device_info
-- ----------------------------
DROP TABLE IF EXISTS `dim_device_info`;
CREATE TABLE `dim_device_info`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '设备ID',
  `device_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备名称',
  `device_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备IP',
  `snmp_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'SNMP读写团体名',
  `udp_port` int NULL DEFAULT NULL COMMENT 'UDP端口号',
  `is_snmp_trap` int NULL DEFAULT NULL COMMENT '开启SNMP Trap上报(0:关闭，1:开启)',
  `report_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '上报服务器地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '设备备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_device_info
-- ----------------------------

-- ----------------------------
-- Table structure for dim_isp_oversea
-- ----------------------------
DROP TABLE IF EXISTS `dim_isp_oversea`;
CREATE TABLE `dim_isp_oversea`  (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '运营商-国内外字典' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_isp_oversea
-- ----------------------------
INSERT INTO `dim_isp_oversea` VALUES (0, '国内');
INSERT INTO `dim_isp_oversea` VALUES (1, '国外');
INSERT INTO `dim_isp_oversea` VALUES (2, '未知');

-- ----------------------------
-- Table structure for dim_line_info
-- ----------------------------
DROP TABLE IF EXISTS `dim_line_info`;
CREATE TABLE `dim_line_info`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '线路ID',
  `line_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '线路名称',
  `line_num` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '线路编号',
  `line_vlan` int NULL DEFAULT NULL COMMENT 'vlan',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '线路备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for dim_traffic_alarm_config
-- ----------------------------
DROP TABLE IF EXISTS `dim_traffic_alarm_config`;
CREATE TABLE `dim_traffic_alarm_config` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `start_time` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `end_time` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `app_type_id` int NOT NULL,
    `app_id` int NOT NULL,
    `link_ids` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `base_value` int DEFAULT NULL,
    `increase_ratio` double NOT NULL,
    `decrease_ratio` double NOT NULL,
    `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `increase_base_value` int DEFAULT NULL,
    `decrease_base_value` int DEFAULT NULL,
    `deleted` tinyint DEFAULT '0',
    UNIQUE KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for dim_user_crowd
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd`;
CREATE TABLE `dim_user_crowd`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `crowd_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户群名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户组' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd
-- ----------------------------

-- ----------------------------
-- Table structure for dim_user_crowd_group
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_group`;
CREATE TABLE `dim_user_crowd_group`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户群组名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户组群' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_group
-- ----------------------------

-- ----------------------------
-- Table structure for dim_user_crowd_group_relation
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_group_relation`;
CREATE TABLE `dim_user_crowd_group_relation`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_id` int NULL DEFAULT NULL COMMENT '用户群组id',
  `crowd_id` int NULL DEFAULT NULL COMMENT '用户群id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户组群和用户组的关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_group_relation
-- ----------------------------

-- ----------------------------
-- Table structure for dim_user_crowd_relation
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_relation`;
CREATE TABLE `dim_user_crowd_relation`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `crowd_id` int NULL DEFAULT NULL COMMENT '用户群ID',
  `user_id` int NULL DEFAULT NULL COMMENT '用户ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户组和用户的关系' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_relation
-- ----------------------------

-- ----------------------------
-- Table structure for dim_user_info
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_info`;
CREATE TABLE `dim_user_info`  (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY  COMMENT '用户ID',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名称',
  `user_type` tinyint NULL DEFAULT NULL COMMENT '0： 正常 1：监测用户',
  `ip_address` json NULL COMMENT 'ip地址段',
  `ip_address_num` varchar(255) NULL COMMENT '可用IP地址数量',
  `remark` varchar(255)  NULL COMMENT '备注'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_info
-- ----------------------------

-- ----------------------------
-- Table structure for dim_white_policy
-- ----------------------------
DROP TABLE IF EXISTS `dim_white_policy`;
CREATE TABLE `dim_white_policy`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '策略名称',
  `user_type` tinyint NULL DEFAULT NULL COMMENT '1：用户群组 2：用户群 3：用户',
  `user_crowd_group_id` int NULL DEFAULT NULL COMMENT '用户群组ID',
  `user_crowd_id` int NULL DEFAULT NULL COMMENT '用户群ID',
  `user_id` int NULL DEFAULT NULL COMMENT '用户ID',
  `ul_tos` int NULL DEFAULT NULL COMMENT '上行tos',
  `dl_tos` int NULL DEFAULT NULL COMMENT '下行tos',
  `app_type_id` int NULL DEFAULT NULL COMMENT '应用大类ID',
  `app_id` int NULL DEFAULT NULL COMMENT '应用小类ID',
  `start_time` datetime NOT NULL COMMENT '生效时间',
  `end_time` datetime NOT NULL COMMENT '失效时间',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_white_policy
-- ----------------------------

-- ----------------------------
-- Table structure for dws_traffic_alarm_log
-- ----------------------------
DROP TABLE IF EXISTS `dws_traffic_alarm_log`;
CREATE TABLE `dws_traffic_alarm_log`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `start_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `policy_id` int NOT NULL,
  `policy_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `link_id` int NULL DEFAULT NULL,
  `link_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `app_id` int NULL DEFAULT NULL,
  `app_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `app_type_id` int NULL DEFAULT NULL,
  `app_type_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `alarm_type` tinyint NULL DEFAULT NULL COMMENT '1-上浮；2-下降',
  `app_traffic_total_speed` int NULL DEFAULT NULL,
  `last_period_total_speed` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 55 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dws_traffic_alarm_log
-- ----------------------------

-- ----------------------------
-- Table structure for biz_isp
-- ----------------------------
DROP TABLE IF EXISTS `biz_isp`;
CREATE TABLE `biz_isp`  (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `isp_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
                            `is_oversea` tinyint(1) NOT NULL COMMENT '1国外；0国内',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for biz_ip_address
-- ----------------------------
DROP TABLE IF EXISTS `biz_ip_address`;
CREATE TABLE `biz_ip_address`  (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `ip_start` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `ip_end` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `country` varchar(128) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `unknown` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `province` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `city` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `isp` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
                                   `isp_id` int NULL DEFAULT NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;


-- ----------------------------
-- View structure for app_id_view
-- ----------------------------
DROP VIEW IF EXISTS `app_id_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `app_id_view` AS select `dim_app_classify`.`app_id` AS `id`,`dim_app_classify`.`app_name` AS `name` from `dim_app_classify`;

-- ----------------------------
-- View structure for app_type_view
-- ----------------------------
DROP VIEW IF EXISTS `app_type_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `app_type_view` AS select `dim_app_classify`.`app_type_id` AS `id`,`dim_app_classify`.`app_type_name` AS `name` from `dim_app_classify` group by `dim_app_classify`.`app_type_id`,`dim_app_classify`.`app_type_name`;

-- ----------------------------
-- View structure for isp_view
-- ----------------------------
DROP VIEW IF EXISTS `isp_view`;
CREATE VIEW `isp_view` AS select `a`.`id` AS `id`,if((`b`.`is_oversea` = 1),'国外',if((`a`.`isp` not in ('东方有线','移动','电信','国外','腾讯','联通','阿里巴巴','百度','华为','教育网','铁通','金山云','斐讯','广电','中华电信','美团','字节跳动','世纪互联','鹏博士')),'其他',`a`.`isp`)) AS `name`,`b`.`is_oversea` AS `is_oversea`,`a`.`province` AS `province` from (`biz_ip_address` `a` left join `biz_isp` `b` on((`a`.`isp_id` = `b`.`id`)));
-- ----------------------------
-- View structure for is_oversea_view
-- ----------------------------
DROP VIEW IF EXISTS `is_oversea_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `is_oversea_view` AS select `a`.`id` AS `id`,`b`.`is_oversea` AS `is_oversea` from (`biz_ip_address` `a` left join `biz_isp` `b` on((`a`.`isp_id` = `b`.`id`)));

-- ----------------------------
-- View structure for link_view
-- ----------------------------
DROP VIEW IF EXISTS `link_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `link_view` AS select `dim_line_info`.`line_vlan` AS `id`,`dim_line_info`.`line_name` AS `name` from `dim_line_info` union select 0 AS `id`,'未知链路' AS `name`;

-- ----------------------------
-- View structure for province_view
-- ----------------------------
DROP VIEW IF EXISTS `province_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `province_view` AS select `biz_ip_address`.`id` AS `id`,`biz_ip_address`.`province` AS `name` from `biz_ip_address`;

-- ----------------------------
-- View structure for user_ip_view
-- ----------------------------
DROP VIEW IF EXISTS `user_ip_view`;
CREATE OR REPLACE DEFINER = root@localhost VIEW user_ip_view AS
SELECT  jt.ip_address AS ip_address,
        id
FROM dim_user_info,
     JSON_TABLE(
         dim_user_info.ip_address,
         '$[*]' COLUMNS (
             ip_address VARCHAR(255) PATH '$'
             )
         ) AS jt;

DROP VIEW IF EXISTS `ip_address_view`;
CREATE VIEW `ip_address_view` AS select `biz_ip_address`.`id` AS `id`,concat_ws(',',if((`biz_ip_address`.`country` in ('','0')),'未知',`biz_ip_address`.`country`),if((`biz_ip_address`.`province` in ('','0')),'未知',`biz_ip_address`.`province`),if((`biz_ip_address`.`city` in ('','0')),'未知',`biz_ip_address`.`city`),if((`biz_ip_address`.`isp` in ('','0')),'未知',`biz_ip_address`.`isp`)) AS `location` from `biz_ip_address`;

SET FOREIGN_KEY_CHECKS = 1;
