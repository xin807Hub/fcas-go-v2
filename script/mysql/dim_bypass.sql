USE fcas_service;

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
-- Records of dim_bypass
-- ----------------------------
INSERT INTO `dim_bypass` VALUES (11, 1, '张江出口CR16K HGE2/0/3-龙南NCS Hu0/0/1/1', '10.104.1.163', 8080, '龙南');
INSERT INTO `dim_bypass` VALUES (12, 1, '张江出口CR16K HGE2/0/2-张江NCS Hu0/0/1/1', '10.104.1.93', 8080, '张江');
INSERT INTO `dim_bypass` VALUES (13, 1, '张江出口CR16K HGE3/0/2-广纪NE40 50|100GE4/1/0', '10.104.1.203', 8080, '广纪');
INSERT INTO `dim_bypass` VALUES (14, 1, '张江出口CR16K HGE3/0/3-临空NE40 50|100GE4/1/0', '10.104.1.219', 8080, '临空');
INSERT INTO `dim_bypass` VALUES (16, 2, '广电出口CR16K HGE3/0/3-临空NE40 50|100GE5/1/0', '10.104.1.219', 8080, '临空');
INSERT INTO `dim_bypass` VALUES (17, 2, '广电出口CR16K HGE3/0/2-广纪NE40 50|100GE5/1/0', '10.104.1.203', 8080, '广纪');
INSERT INTO `dim_bypass` VALUES (18, 2, '广电出口CR16K HGE2/0/2-张江NCS Hu0/2/1/1', '10.104.1.93', 8080, '张江');
INSERT INTO `dim_bypass` VALUES (19, 2, '广电出口CR16K HGE2/0/3-龙南NCS Hu0/2/1/1', '10.104.1.163', 8080, '龙南');
