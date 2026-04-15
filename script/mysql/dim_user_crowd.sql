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

 Date: 17/05/2025 17:06:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
USE fcas_service;

-- ----------------------------
-- Table structure for dim_user_crowd
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd`;
CREATE TABLE `dim_user_crowd`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户群ID',
  `crowd_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户群名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd
-- ----------------------------
INSERT INTO `dim_user_crowd` VALUES (1, '内容平台', '');
INSERT INTO `dim_user_crowd` VALUES (3, '企业用户分区', '');
INSERT INTO `dim_user_crowd` VALUES (4, 'VOD统一出口', '');
INSERT INTO `dim_user_crowd` VALUES (5, '郊县NGB设备', '');
INSERT INTO `dim_user_crowd` VALUES (6, 'IDC', '');
INSERT INTO `dim_user_crowd` VALUES (7, 'IPv6', '');
INSERT INTO `dim_user_crowd` VALUES (8, '徐汇', '');
INSERT INTO `dim_user_crowd` VALUES (9, '宝山', '');
INSERT INTO `dim_user_crowd` VALUES (10, '崇明', '');
INSERT INTO `dim_user_crowd` VALUES (11, '奉贤', '');
INSERT INTO `dim_user_crowd` VALUES (12, '虹口', '');
INSERT INTO `dim_user_crowd` VALUES (13, '嘉定', '');
INSERT INTO `dim_user_crowd` VALUES (14, '金山', '');
INSERT INTO `dim_user_crowd` VALUES (15, '闵行', '');
INSERT INTO `dim_user_crowd` VALUES (16, '浦东', '');
INSERT INTO `dim_user_crowd` VALUES (17, '浦东子公司', '');
INSERT INTO `dim_user_crowd` VALUES (18, '普陀', '');
INSERT INTO `dim_user_crowd` VALUES (19, '市中', '');
INSERT INTO `dim_user_crowd` VALUES (20, '松江', '');
INSERT INTO `dim_user_crowd` VALUES (21, '杨浦', '');
INSERT INTO `dim_user_crowd` VALUES (22, '闸北', '');
INSERT INTO `dim_user_crowd` VALUES (23, '长宁', '');
INSERT INTO `dim_user_crowd` VALUES (24, '青浦', '');
INSERT INTO `dim_user_crowd` VALUES (25, '网宿', '');
INSERT INTO `dim_user_crowd` VALUES (26, '网宿村居', '');
INSERT INTO `dim_user_crowd` VALUES (27, '瑞高', '');
INSERT INTO `dim_user_crowd` VALUES (28, '中广', '');
INSERT INTO `dim_user_crowd` VALUES (29, '速亨', '');
INSERT INTO `dim_user_crowd` VALUES (30, '未分配企业用户', '');
INSERT INTO `dim_user_crowd` VALUES (31, '未分配内容平台', '');
INSERT INTO `dim_user_crowd` VALUES (32, '未分配个人', '');
INSERT INTO `dim_user_crowd` VALUES (33, '个人用户（vbras）', '');

SET FOREIGN_KEY_CHECKS = 1;
