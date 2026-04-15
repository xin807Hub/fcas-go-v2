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

 Date: 17/05/2025 17:06:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
USE fcas_service;

-- ----------------------------
-- Table structure for dim_user_crowd_group_relation
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_group_relation`;
CREATE TABLE `dim_user_crowd_group_relation`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '关联id',
  `crowd_id` int NULL DEFAULT NULL COMMENT '用户群id',
  `group_id` int NULL DEFAULT NULL COMMENT '用户群组id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_group_relation
-- ----------------------------
INSERT INTO `dim_user_crowd_group_relation` VALUES (1, 6, 1);
INSERT INTO `dim_user_crowd_group_relation` VALUES (2, 7, 2);
INSERT INTO `dim_user_crowd_group_relation` VALUES (3, 4, 3);
INSERT INTO `dim_user_crowd_group_relation` VALUES (4, 5, 4);
INSERT INTO `dim_user_crowd_group_relation` VALUES (10, 3, 6);
INSERT INTO `dim_user_crowd_group_relation` VALUES (11, 1, 7);
INSERT INTO `dim_user_crowd_group_relation` VALUES (13, 31, 9);
INSERT INTO `dim_user_crowd_group_relation` VALUES (14, 13, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (15, 11, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (16, 9, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (17, 10, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (18, 19, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (19, 8, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (20, 18, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (21, 21, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (22, 20, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (23, 16, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (24, 17, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (25, 12, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (26, 14, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (27, 23, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (28, 15, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (29, 22, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (30, 24, 10);
INSERT INTO `dim_user_crowd_group_relation` VALUES (31, 32, 11);
INSERT INTO `dim_user_crowd_group_relation` VALUES (32, 30, 8);
INSERT INTO `dim_user_crowd_group_relation` VALUES (33, 28, 5);
INSERT INTO `dim_user_crowd_group_relation` VALUES (34, 27, 5);
INSERT INTO `dim_user_crowd_group_relation` VALUES (35, 25, 5);
INSERT INTO `dim_user_crowd_group_relation` VALUES (36, 26, 5);
INSERT INTO `dim_user_crowd_group_relation` VALUES (37, 29, 5);
INSERT INTO `dim_user_crowd_group_relation` VALUES (38, 33, 5);

SET FOREIGN_KEY_CHECKS = 1;
