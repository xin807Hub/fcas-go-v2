/*
 Navicat Premium Dump SQL

 Source Server         : 10.104.6.66
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : 10.104.6.66:3306
 Source Schema         : fcas_service

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 17/05/2025 17:06:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
USE fcas_service;

-- ----------------------------
-- Table structure for dim_user_crowd_group
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_group`;
CREATE TABLE `dim_user_crowd_group`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户群组名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_group
-- ----------------------------
INSERT INTO `dim_user_crowd_group` VALUES (1, 'IDC', '');
INSERT INTO `dim_user_crowd_group` VALUES (2, 'IPv6', '');
INSERT INTO `dim_user_crowd_group` VALUES (3, 'VOD统一出口', '');
INSERT INTO `dim_user_crowd_group` VALUES (4, '郊县NGB设备', '');
INSERT INTO `dim_user_crowd_group` VALUES (5, '个人用户', '');
INSERT INTO `dim_user_crowd_group` VALUES (6, '企业用户', '');
INSERT INTO `dim_user_crowd_group` VALUES (7, '内容平台', '');
INSERT INTO `dim_user_crowd_group` VALUES (8, '未分配-企业', '');
INSERT INTO `dim_user_crowd_group` VALUES (9, '未分配-内容平台', '');
INSERT INTO `dim_user_crowd_group` VALUES (10, '个人用户-行政区', '');
INSERT INTO `dim_user_crowd_group` VALUES (11, '未分配个人', '');

SET FOREIGN_KEY_CHECKS = 1;
