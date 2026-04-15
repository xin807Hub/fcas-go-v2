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

 Date: 17/05/2025 17:06:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
USE fcas_service;

-- ----------------------------
-- Table structure for dim_user_crowd_relation
-- ----------------------------
DROP TABLE IF EXISTS `dim_user_crowd_relation`;
CREATE TABLE `dim_user_crowd_relation`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '管理id',
  `user_id` int NULL DEFAULT NULL COMMENT '用户Id',
  `crowd_id` int NULL DEFAULT NULL COMMENT '用户群id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1095 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dim_user_crowd_relation
-- ----------------------------
INSERT INTO `dim_user_crowd_relation` VALUES (62, 1, 4);
INSERT INTO `dim_user_crowd_relation` VALUES (63, 148, 5);
INSERT INTO `dim_user_crowd_relation` VALUES (64, 62, 6);
INSERT INTO `dim_user_crowd_relation` VALUES (66, 67, 8);
INSERT INTO `dim_user_crowd_relation` VALUES (67, 66, 8);
INSERT INTO `dim_user_crowd_relation` VALUES (68, 65, 8);
INSERT INTO `dim_user_crowd_relation` VALUES (69, 64, 8);
INSERT INTO `dim_user_crowd_relation` VALUES (70, 63, 8);
INSERT INTO `dim_user_crowd_relation` VALUES (71, 82, 9);
INSERT INTO `dim_user_crowd_relation` VALUES (72, 81, 9);
INSERT INTO `dim_user_crowd_relation` VALUES (73, 80, 9);
INSERT INTO `dim_user_crowd_relation` VALUES (74, 79, 9);
INSERT INTO `dim_user_crowd_relation` VALUES (75, 78, 9);
INSERT INTO `dim_user_crowd_relation` VALUES (76, 87, 10);
INSERT INTO `dim_user_crowd_relation` VALUES (77, 86, 10);
INSERT INTO `dim_user_crowd_relation` VALUES (78, 85, 10);
INSERT INTO `dim_user_crowd_relation` VALUES (79, 84, 10);
INSERT INTO `dim_user_crowd_relation` VALUES (80, 83, 10);
INSERT INTO `dim_user_crowd_relation` VALUES (81, 92, 11);
INSERT INTO `dim_user_crowd_relation` VALUES (82, 91, 11);
INSERT INTO `dim_user_crowd_relation` VALUES (83, 90, 11);
INSERT INTO `dim_user_crowd_relation` VALUES (84, 89, 11);
INSERT INTO `dim_user_crowd_relation` VALUES (85, 88, 11);
INSERT INTO `dim_user_crowd_relation` VALUES (86, 97, 12);
INSERT INTO `dim_user_crowd_relation` VALUES (87, 96, 12);
INSERT INTO `dim_user_crowd_relation` VALUES (88, 95, 12);
INSERT INTO `dim_user_crowd_relation` VALUES (89, 94, 12);
INSERT INTO `dim_user_crowd_relation` VALUES (90, 93, 12);
INSERT INTO `dim_user_crowd_relation` VALUES (91, 102, 13);
INSERT INTO `dim_user_crowd_relation` VALUES (92, 101, 13);
INSERT INTO `dim_user_crowd_relation` VALUES (93, 100, 13);
INSERT INTO `dim_user_crowd_relation` VALUES (94, 99, 13);
INSERT INTO `dim_user_crowd_relation` VALUES (95, 98, 13);
INSERT INTO `dim_user_crowd_relation` VALUES (96, 107, 14);
INSERT INTO `dim_user_crowd_relation` VALUES (97, 106, 14);
INSERT INTO `dim_user_crowd_relation` VALUES (98, 105, 14);
INSERT INTO `dim_user_crowd_relation` VALUES (99, 104, 14);
INSERT INTO `dim_user_crowd_relation` VALUES (100, 103, 14);
INSERT INTO `dim_user_crowd_relation` VALUES (101, 77, 15);
INSERT INTO `dim_user_crowd_relation` VALUES (102, 76, 15);
INSERT INTO `dim_user_crowd_relation` VALUES (103, 75, 15);
INSERT INTO `dim_user_crowd_relation` VALUES (104, 74, 15);
INSERT INTO `dim_user_crowd_relation` VALUES (105, 73, 15);
INSERT INTO `dim_user_crowd_relation` VALUES (106, 112, 16);
INSERT INTO `dim_user_crowd_relation` VALUES (107, 111, 16);
INSERT INTO `dim_user_crowd_relation` VALUES (108, 110, 16);
INSERT INTO `dim_user_crowd_relation` VALUES (109, 109, 16);
INSERT INTO `dim_user_crowd_relation` VALUES (110, 108, 16);
INSERT INTO `dim_user_crowd_relation` VALUES (111, 122, 17);
INSERT INTO `dim_user_crowd_relation` VALUES (112, 121, 17);
INSERT INTO `dim_user_crowd_relation` VALUES (113, 120, 17);
INSERT INTO `dim_user_crowd_relation` VALUES (114, 119, 17);
INSERT INTO `dim_user_crowd_relation` VALUES (115, 118, 17);
INSERT INTO `dim_user_crowd_relation` VALUES (116, 117, 18);
INSERT INTO `dim_user_crowd_relation` VALUES (117, 116, 18);
INSERT INTO `dim_user_crowd_relation` VALUES (118, 115, 18);
INSERT INTO `dim_user_crowd_relation` VALUES (119, 114, 18);
INSERT INTO `dim_user_crowd_relation` VALUES (120, 113, 18);
INSERT INTO `dim_user_crowd_relation` VALUES (121, 127, 19);
INSERT INTO `dim_user_crowd_relation` VALUES (122, 126, 19);
INSERT INTO `dim_user_crowd_relation` VALUES (123, 125, 19);
INSERT INTO `dim_user_crowd_relation` VALUES (124, 124, 19);
INSERT INTO `dim_user_crowd_relation` VALUES (125, 123, 19);
INSERT INTO `dim_user_crowd_relation` VALUES (126, 132, 20);
INSERT INTO `dim_user_crowd_relation` VALUES (127, 131, 20);
INSERT INTO `dim_user_crowd_relation` VALUES (128, 130, 20);
INSERT INTO `dim_user_crowd_relation` VALUES (129, 129, 20);
INSERT INTO `dim_user_crowd_relation` VALUES (130, 128, 20);
INSERT INTO `dim_user_crowd_relation` VALUES (131, 72, 21);
INSERT INTO `dim_user_crowd_relation` VALUES (132, 71, 21);
INSERT INTO `dim_user_crowd_relation` VALUES (133, 70, 21);
INSERT INTO `dim_user_crowd_relation` VALUES (134, 69, 21);
INSERT INTO `dim_user_crowd_relation` VALUES (135, 68, 21);
INSERT INTO `dim_user_crowd_relation` VALUES (136, 137, 22);
INSERT INTO `dim_user_crowd_relation` VALUES (137, 136, 22);
INSERT INTO `dim_user_crowd_relation` VALUES (138, 135, 22);
INSERT INTO `dim_user_crowd_relation` VALUES (139, 134, 22);
INSERT INTO `dim_user_crowd_relation` VALUES (140, 133, 22);
INSERT INTO `dim_user_crowd_relation` VALUES (141, 142, 23);
INSERT INTO `dim_user_crowd_relation` VALUES (142, 141, 23);
INSERT INTO `dim_user_crowd_relation` VALUES (143, 140, 23);
INSERT INTO `dim_user_crowd_relation` VALUES (144, 139, 23);
INSERT INTO `dim_user_crowd_relation` VALUES (145, 138, 23);
INSERT INTO `dim_user_crowd_relation` VALUES (146, 147, 24);
INSERT INTO `dim_user_crowd_relation` VALUES (147, 146, 24);
INSERT INTO `dim_user_crowd_relation` VALUES (148, 145, 24);
INSERT INTO `dim_user_crowd_relation` VALUES (149, 144, 24);
INSERT INTO `dim_user_crowd_relation` VALUES (150, 143, 24);
INSERT INTO `dim_user_crowd_relation` VALUES (151, 143, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (152, 138, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (153, 133, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (154, 128, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (155, 123, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (156, 118, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (157, 113, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (158, 108, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (159, 103, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (160, 98, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (161, 93, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (162, 88, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (163, 83, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (164, 78, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (165, 73, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (166, 68, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (167, 63, 25);
INSERT INTO `dim_user_crowd_relation` VALUES (168, 147, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (169, 142, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (170, 137, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (171, 132, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (172, 127, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (173, 122, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (174, 117, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (175, 112, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (176, 107, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (177, 102, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (178, 97, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (179, 92, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (180, 87, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (181, 82, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (182, 77, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (183, 72, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (184, 67, 26);
INSERT INTO `dim_user_crowd_relation` VALUES (185, 145, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (186, 140, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (187, 135, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (188, 130, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (189, 125, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (190, 120, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (191, 115, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (192, 110, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (193, 105, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (194, 100, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (195, 95, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (196, 90, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (197, 85, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (198, 80, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (199, 75, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (200, 70, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (201, 65, 27);
INSERT INTO `dim_user_crowd_relation` VALUES (202, 146, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (203, 141, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (204, 136, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (205, 131, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (206, 126, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (207, 121, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (208, 116, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (209, 111, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (210, 106, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (211, 101, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (212, 96, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (213, 91, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (214, 86, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (215, 81, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (216, 76, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (217, 71, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (218, 66, 28);
INSERT INTO `dim_user_crowd_relation` VALUES (219, 144, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (220, 139, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (221, 134, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (222, 129, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (223, 124, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (224, 119, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (225, 114, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (226, 109, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (227, 104, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (228, 99, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (229, 94, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (230, 89, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (231, 84, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (232, 79, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (233, 74, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (234, 69, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (235, 64, 29);
INSERT INTO `dim_user_crowd_relation` VALUES (237, 155, 31);
INSERT INTO `dim_user_crowd_relation` VALUES (238, 160, 32);
INSERT INTO `dim_user_crowd_relation` VALUES (552, 335, 30);
INSERT INTO `dim_user_crowd_relation` VALUES (765, 13, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (766, 12, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (767, 11, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (768, 10, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (769, 9, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (770, 8, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (771, 7, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (772, 6, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (773, 5, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (774, 4, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (775, 3, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (776, 2, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (777, 339, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (778, 338, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (779, 337, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (780, 336, 1);
INSERT INTO `dim_user_crowd_relation` VALUES (965, 375, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (966, 374, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (967, 373, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (968, 372, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (969, 371, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (970, 369, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (971, 368, 33);
INSERT INTO `dim_user_crowd_relation` VALUES (1030, 353, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1031, 352, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1032, 351, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1033, 350, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1034, 348, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1035, 377, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1036, 376, 7);
INSERT INTO `dim_user_crowd_relation` VALUES (1037, 370, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1038, 366, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1039, 354, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1040, 176, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1041, 175, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1042, 170, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1043, 167, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1044, 61, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1045, 60, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1046, 59, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1047, 58, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1048, 57, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1049, 56, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1050, 54, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1051, 53, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1052, 52, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1053, 51, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1054, 50, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1055, 49, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1056, 48, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1057, 47, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1058, 46, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1059, 45, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1060, 44, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1061, 43, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1062, 42, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1063, 41, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1064, 40, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1065, 39, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1066, 38, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1067, 37, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1068, 36, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1069, 35, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1070, 34, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1071, 33, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1072, 32, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1073, 31, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1074, 30, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1075, 29, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1076, 28, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1077, 27, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1078, 26, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1079, 25, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1080, 24, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1081, 23, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1082, 22, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1083, 21, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1084, 20, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1085, 19, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1086, 18, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1087, 17, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1088, 16, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1089, 15, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1090, 14, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1091, 1, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1092, 384, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1093, 383, 3);
INSERT INTO `dim_user_crowd_relation` VALUES (1094, 382, 3);

SET FOREIGN_KEY_CHECKS = 1;
