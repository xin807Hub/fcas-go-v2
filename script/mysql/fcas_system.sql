/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.4.146
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 192.168.4.146:3306
 Source Schema         : fcas_system

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 15/11/2024 17:00:28
*/
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS fcas_system;
USE fcas_system;
-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint unsigned NOT NULL,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 710 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (610, 'p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (611, 'p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (616, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (619, 'p', '888', '/api/enterSyncApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (614, 'p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (615, 'p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (618, 'p', '888', '/api/getApiGroups', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (613, 'p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (620, 'p', '888', '/api/ignoreApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (617, 'p', '888', '/api/syncApi', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (612, 'p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (621, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (622, 'p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (623, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (625, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (626, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (624, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (691, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (690, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (689, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (662, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (659, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (656, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (664, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (668, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (658, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (654, 'p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (665, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (663, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (667, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (655, 'p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (660, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (657, 'p', '888', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (661, 'p', '888', '/autoCode/pubPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (666, 'p', '888', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (628, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (627, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (701, 'p', '888', '/configuration/dimdeviceinfo/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (703, 'p', '888', '/configuration/dimlineinfo/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (651, 'p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (652, 'p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (650, 'p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (649, 'p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (653, 'p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (687, 'p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (688, 'p', '888', '/email/sendEmail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (639, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (640, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (643, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (644, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (638, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (645, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (641, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (642, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (599, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (629, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (637, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (631, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (633, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (635, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (630, 'p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (636, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (634, 'p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (632, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (708, 'p', '888', '/policy/alarmConfig/delete', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (707, 'p', '888', '/policy/alarmConfig/page', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (709, 'p', '888', '/policy/alarmConfig/saveOrUpdate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (704, 'p', '888', '/policy/whitePolicy/page', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (705, 'p', '888', '/policy/whitePolicy/saveOrUpdate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (685, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (686, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (684, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (674, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (675, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (677, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (678, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (676, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (670, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (671, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (672, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (673, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (669, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (692, 'p', '888', '/sysExportTemplate/createSysExportTemplate', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (693, 'p', '888', '/sysExportTemplate/deleteSysExportTemplate', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (694, 'p', '888', '/sysExportTemplate/deleteSysExportTemplateByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (698, 'p', '888', '/sysExportTemplate/exportExcel', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (699, 'p', '888', '/sysExportTemplate/exportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (696, 'p', '888', '/sysExportTemplate/findSysExportTemplate', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (697, 'p', '888', '/sysExportTemplate/getSysExportTemplateList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (700, 'p', '888', '/sysExportTemplate/importExcel', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (695, 'p', '888', '/sysExportTemplate/updateSysExportTemplate', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (679, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (682, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (683, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (680, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (681, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (646, 'p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (647, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (648, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (706, 'p', '888', '/traffic/alarmLog/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (702, 'p', '888', '/traffic/dwsTotalTraffic/trend', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (601, 'p', '888', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (607, 'p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (600, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (605, 'p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (602, 'p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (609, 'p', '888', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (604, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (606, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (608, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (603, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', '8881', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', '8881', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (144, 'p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (147, 'p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (179, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (170, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (169, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (177, 'p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (175, 'p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (176, 'p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (174, 'p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (178, 'p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (167, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (168, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (166, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (165, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (171, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (155, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (157, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (159, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (161, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (156, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (158, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (160, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (172, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (173, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', '9528', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (162, 'p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (180, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (163, 'p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (164, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '');

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 112 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (3, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (6, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (8, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (19, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/syncApi', '获取待同步API', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (20, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/getApiGroups', '获取路由组', 'api', 'GET');
INSERT INTO `sys_apis` VALUES (21, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/enterSyncApi', '确认同步API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/api/ignoreApi', '忽略API', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (27, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES (41, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (48, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (51, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (52, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (53, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (54, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (55, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (56, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (58, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (59, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (60, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (61, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (62, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/pubPlug', '打包插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (65, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (66, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (67, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (68, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (69, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (70, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (71, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (72, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (73, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (74, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (75, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (76, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (77, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (78, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (79, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (80, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (81, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (82, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (83, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (84, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (85, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (86, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES (87, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (88, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (89, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (90, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/email/sendEmail', '发送邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (91, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (92, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (93, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (94, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/createSysExportTemplate', '新增导出模板', '表格模板', 'POST');
INSERT INTO `sys_apis` VALUES (95, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/deleteSysExportTemplate', '删除导出模板', '表格模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (96, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/deleteSysExportTemplateByIds', '批量删除导出模板', '表格模板', 'DELETE');
INSERT INTO `sys_apis` VALUES (97, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/updateSysExportTemplate', '更新导出模板', '表格模板', 'PUT');
INSERT INTO `sys_apis` VALUES (98, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/findSysExportTemplate', '根据ID获取导出模板', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (99, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/getSysExportTemplateList', '获取导出模板列表', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (100, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/exportExcel', '导出Excel', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (101, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/exportTemplate', '下载模板', '表格模板', 'GET');
INSERT INTO `sys_apis` VALUES (102, '2024-10-15 13:05:01.289', '2024-10-15 13:05:01.289', NULL, '/sysExportTemplate/importExcel', '导入Excel', '表格模板', 'POST');
INSERT INTO `sys_apis` VALUES (103, '2024-10-15 20:39:44.103', '2024-10-15 20:39:44.103', NULL, '/configuration/dimdeviceinfo/list', '获取设备列表信息', ' 设备管理', 'GET');
INSERT INTO `sys_apis` VALUES (104, '2024-10-29 17:58:17.380', '2024-10-29 17:58:33.892', NULL, '/traffic/dwsTotalTraffic/trend', '流量趋势图', '首页', 'POST');
INSERT INTO `sys_apis` VALUES (105, '2024-10-30 09:45:38.700', '2024-10-30 09:45:38.700', NULL, '/configuration/dimlineinfo/list', '获取链路列表', '链路配置', 'GET');
INSERT INTO `sys_apis` VALUES (106, '2024-10-30 09:45:38.700', '2024-10-30 09:45:38.700', NULL, '/policy/whitePolicy/page', '获取优先策略', '优先策略', 'POST');
INSERT INTO `sys_apis` VALUES (107, '2024-10-30 09:45:38.700', '2024-10-30 09:45:38.700', NULL, '/policy/whitePolicy/saveOrUpdate', '增加优先策略', '优先策略', 'POST');
INSERT INTO `sys_apis` VALUES (108, '2024-11-14 14:22:34.000', '2024-11-14 14:22:37.000', NULL, '/traffic/alarmLog/list', '获取流量告警日志', '流量告警', 'GET');
INSERT INTO `sys_apis` VALUES (109, '2024-11-14 14:31:11.000', '2024-11-14 14:31:14.000', NULL, '/policy/alarmConfig/page', '业务流量告警分页查询', '流量告警', 'POST');
INSERT INTO `sys_apis` VALUES (110, '2024-11-14 14:31:36.000', '2024-11-14 14:31:40.000', NULL, '/policy/alarmConfig/delete', '删除业务流量告警', '流量告警', 'GET');
INSERT INTO `sys_apis` VALUES (111, '2024-11-14 14:33:11.000', '2024-11-14 14:33:14.000', NULL, '/policy/alarmConfig/saveOrUpdate', '新增或修改业务流量告警', '流量告警', 'POST');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `uni_sys_authorities_authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9529 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2024-10-15 13:05:01.317', '2024-11-14 09:44:15.114', NULL, 888, '普通用户', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2024-10-15 13:05:01.317', '2024-10-15 13:05:01.559', NULL, 8881, '普通用户子角色', 888, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2024-10-15 13:05:01.317', '2024-11-14 09:44:20.655', NULL, 9528, '测试角色', 0, 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint unsigned COMMENT '角色ID',
  `sys_menu_id` bigint unsigned COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, 888);
INSERT INTO `sys_authority_menus` VALUES (1, 8881);
INSERT INTO `sys_authority_menus` VALUES (1, 9528);
INSERT INTO `sys_authority_menus` VALUES (3, 888);
INSERT INTO `sys_authority_menus` VALUES (3, 9528);
INSERT INTO `sys_authority_menus` VALUES (4, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 8881);
INSERT INTO `sys_authority_menus` VALUES (4, 9528);
INSERT INTO `sys_authority_menus` VALUES (5, 888);
INSERT INTO `sys_authority_menus` VALUES (5, 8881);
INSERT INTO `sys_authority_menus` VALUES (5, 9528);
INSERT INTO `sys_authority_menus` VALUES (6, 888);
INSERT INTO `sys_authority_menus` VALUES (6, 8881);
INSERT INTO `sys_authority_menus` VALUES (6, 9528);
INSERT INTO `sys_authority_menus` VALUES (7, 888);
INSERT INTO `sys_authority_menus` VALUES (7, 8881);
INSERT INTO `sys_authority_menus` VALUES (7, 9528);
INSERT INTO `sys_authority_menus` VALUES (8, 888);
INSERT INTO `sys_authority_menus` VALUES (8, 8881);
INSERT INTO `sys_authority_menus` VALUES (9, 888);
INSERT INTO `sys_authority_menus` VALUES (9, 8881);
INSERT INTO `sys_authority_menus` VALUES (9, 9528);
INSERT INTO `sys_authority_menus` VALUES (10, 888);
INSERT INTO `sys_authority_menus` VALUES (10, 8881);
INSERT INTO `sys_authority_menus` VALUES (10, 9528);
INSERT INTO `sys_authority_menus` VALUES (11, 8881);
INSERT INTO `sys_authority_menus` VALUES (15, 8881);
INSERT INTO `sys_authority_menus` VALUES (31, 888);
INSERT INTO `sys_authority_menus` VALUES (31, 9528);
INSERT INTO `sys_authority_menus` VALUES (32, 888);
INSERT INTO `sys_authority_menus` VALUES (32, 9528);
INSERT INTO `sys_authority_menus` VALUES (33, 888);
INSERT INTO `sys_authority_menus` VALUES (33, 9528);
INSERT INTO `sys_authority_menus` VALUES (34, 888);
INSERT INTO `sys_authority_menus` VALUES (34, 9528);
INSERT INTO `sys_authority_menus` VALUES (35, 888);
INSERT INTO `sys_authority_menus` VALUES (35, 9528);
INSERT INTO `sys_authority_menus` VALUES (36, 888);
INSERT INTO `sys_authority_menus` VALUES (36, 9528);
INSERT INTO `sys_authority_menus` VALUES (37, 888);
INSERT INTO `sys_authority_menus` VALUES (37, 9528);
INSERT INTO `sys_authority_menus` VALUES (38, 888);
INSERT INTO `sys_authority_menus` VALUES (38, 9528);
INSERT INTO `sys_authority_menus` VALUES (39, 888);
INSERT INTO `sys_authority_menus` VALUES (39, 9528);
INSERT INTO `sys_authority_menus` VALUES (40, 888);
INSERT INTO `sys_authority_menus` VALUES (40, 9528);
INSERT INTO `sys_authority_menus` VALUES (41, 888);
INSERT INTO `sys_authority_menus` VALUES (41, 9528);
INSERT INTO `sys_authority_menus` VALUES (42, 888);
INSERT INTO `sys_authority_menus` VALUES (42, 9528);
INSERT INTO `sys_authority_menus` VALUES (43, 888);
INSERT INTO `sys_authority_menus` VALUES (43, 9528);
INSERT INTO `sys_authority_menus` VALUES (44, 888);
INSERT INTO `sys_authority_menus` VALUES (44, 9528);
INSERT INTO `sys_authority_menus` VALUES (45, 888);
INSERT INTO `sys_authority_menus` VALUES (45, 9528);
INSERT INTO `sys_authority_menus` VALUES (46, 888);
INSERT INTO `sys_authority_menus` VALUES (46, 9528);
INSERT INTO `sys_authority_menus` VALUES (47, 888);
INSERT INTO `sys_authority_menus` VALUES (47, 9528);
INSERT INTO `sys_authority_menus` VALUES (48, 888);
INSERT INTO `sys_authority_menus` VALUES (48, 9528);
INSERT INTO `sys_authority_menus` VALUES (49, 888);
INSERT INTO `sys_authority_menus` VALUES (49, 9528);
INSERT INTO `sys_authority_menus` VALUES (50, 888);
INSERT INTO `sys_authority_menus` VALUES (50, 9528);
INSERT INTO `sys_authority_menus` VALUES (51, 888);
INSERT INTO `sys_authority_menus` VALUES (51, 9528);
INSERT INTO `sys_authority_menus` VALUES (52, 888);
INSERT INTO `sys_authority_menus` VALUES (52, 9528);
INSERT INTO `sys_authority_menus` VALUES (53, 888);
INSERT INTO `sys_authority_menus` VALUES (53, 9528);

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned,
  `parent_id` bigint unsigned COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(0) DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2024-10-15 13:05:01.384', '2024-10-29 15:49:07.535', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '首页', 'home-filled', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:34.182', 0, 0, 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2024-10-15 13:05:01.384', '2024-10-29 15:50:57.963', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 5, '', 0, 0, '系统管理', 'setting', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2024-10-15 13:05:01.384', '2024-10-29 15:34:01.556', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 2, '', 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2024-10-15 13:05:01.384', '2024-10-29 15:34:29.345', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 3, '', 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2024-10-15 13:05:01.384', '2024-10-29 15:41:23.812', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 6, '', 1, 0, 'api管理', 'aim', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2024-10-15 13:05:01.384', '2024-10-29 15:34:20.847', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', 1, '', 0, 0, '管理员列表', 'user-filled', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2024-10-15 13:05:01.384', '2024-10-29 15:34:36.452', NULL, 0, 3, 'dictionary', 'dictionary', 1, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (9, '2024-10-15 13:05:01.384', '2024-10-29 15:39:04.018', NULL, 0, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 7, '', 0, 0, '系统日志', 'document-copy', 0);
INSERT INTO `sys_base_menus` VALUES (10, '2024-10-15 13:05:01.384', '2024-10-29 15:06:13.808', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', 100, '', 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (13, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:39:12.273', 0, 11, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (15, '2024-10-15 13:05:01.384', '2024-10-29 15:06:56.284', NULL, 0, 0, 'systemTools', 'systemTools', 1, 'view/systemTools/index.vue', 50, '', 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (16, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:10.867', 0, 15, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:13.072', 0, 15, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', NULL, 0, 15, 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:08.393', 0, 15, 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:05.281', 0, 15, 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (21, '2024-10-15 13:05:01.384', '2024-10-15 13:05:01.384', '2024-10-15 13:42:03.087', 0, 15, 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (22, '2024-10-15 13:05:01.384', '2024-10-29 15:07:02.586', NULL, 0, 0, 'state', 'state', 0, 'view/system/state.vue', 51, '', 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (31, '2024-10-29 15:03:42.609', '2024-10-29 15:50:37.870', NULL, 0, 0, 'deviceUser', 'deviceUser', 0, 'view/deviceUser', 2, '', 0, 0, '设备和用户管理', 'monitor', 0);
INSERT INTO `sys_base_menus` VALUES (32, '2024-10-29 15:05:08.302', '2024-10-31 15:43:22.511', NULL, 0, 0, 'flowMonitor', 'flowMonitor', 0, 'view/flowMonitor', 3, '', 0, 0, '流量监测分析', 'histogram', 0);
INSERT INTO `sys_base_menus` VALUES (33, '2024-10-29 15:05:44.697', '2024-10-29 15:05:44.697', NULL, 0, 0, 'policyManage', 'policyManage', 0, 'view/policyManage', 4, '', 0, 0, '策略管控', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (34, '2024-10-29 15:07:47.617', '2024-10-29 15:09:43.358', NULL, 0, 0, 'customConfig', 'customConfig', 0, 'view/customConfig', 6, '', 0, 0, '对象自定义', 'help-filled', 0);
INSERT INTO `sys_base_menus` VALUES (35, '2024-10-29 15:12:14.667', '2024-10-29 15:12:14.667', NULL, 0, 31, 'deviceManage', 'deviceManage', 0, 'view/deviceUser/deviceManage.vue', 1, '', 0, 0, '设备管理', 'setting', 0);
INSERT INTO `sys_base_menus` VALUES (36, '2024-10-29 15:13:20.363', '2024-10-29 15:13:20.363', NULL, 0, 31, 'linkManage', 'linkManage', 0, 'view/deviceUser/linkManage.vue', 2, '', 0, 0, '链路管理', 'link', 0);
INSERT INTO `sys_base_menus` VALUES (37, '2024-10-29 15:13:50.605', '2024-10-29 15:13:50.605', NULL, 0, 31, 'userManage', 'userManage', 0, 'view/deviceUser/userManage.vue', 3, '', 0, 0, '用户管理', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (38, '2024-10-29 15:15:58.276', '2024-10-30 18:10:36.345', NULL, 0, 31, 'crowdManage', 'crowdManage', 0, 'view/deviceUser/crowdManage.vue', 4, '', 0, 0, '用户群管理', 'user-filled', 0);
INSERT INTO `sys_base_menus` VALUES (39, '2024-10-29 15:17:13.052', '2024-10-30 18:10:47.973', NULL, 0, 31, 'crowdGroupManage', 'crowdGroupManage', 0, 'view/deviceUser/crowdGroupManage.vue', 5, '', 0, 0, '用户群组管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (40, '2024-10-29 15:17:47.624', '2024-10-29 15:17:47.624', NULL, 0, 31, 'bypassManage', 'bypassManage', 0, 'view/deviceUser/bypassManage.vue', 6, '', 0, 0, 'Bypass管理', 'setting', 0);
INSERT INTO `sys_base_menus` VALUES (41, '2024-10-29 15:20:10.699', '2024-11-14 09:43:09.178', NULL, 0, 32, 'ispRank', 'ispRank', 0, 'view/flowMonitor/ispRank.vue', 1, '', 0, 0, '运营商流量排名分析', 'data-analysis', 0);
INSERT INTO `sys_base_menus` VALUES (42, '2024-10-29 15:21:24.068', '2024-11-11 10:02:14.193', NULL, 0, 32, 'appType', 'appType', 0, 'view/flowMonitor/appType.vue', 2, '', 0, 0, '大类业务流量排名分析', 'data-line', 0);
INSERT INTO `sys_base_menus` VALUES (43, '2024-10-29 15:21:58.696', '2024-11-11 10:02:26.046', NULL, 0, 32, 'appId', 'appId', 0, 'view/flowMonitor/appId.vue', 3, '', 0, 0, '小类业务流量排名分析', 'data-line', 0);
INSERT INTO `sys_base_menus` VALUES (44, '2024-10-29 15:23:00.140', '2024-10-29 15:25:10.793', NULL, 0, 32, 'userRank', 'userRank', 0, 'view/flowMonitor/userRank.vue', 4, '', 0, 0, '用户排名分析', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (45, '2024-10-29 15:23:44.216', '2024-11-11 10:19:41.928', NULL, 0, 32, 'userCrowd', 'userCrowd', 0, 'view/flowMonitor/userCrowd.vue', 5, '', 0, 0, '用户群排名分析', 'user-filled', 0);
INSERT INTO `sys_base_menus` VALUES (46, '2024-10-29 15:24:23.007', '2024-11-11 10:19:50.069', NULL, 0, 32, 'userCrowdGroup', 'userCrowdGroup', 0, 'view/flowMonitor/userCrowdGroup.vue', 6, '', 0, 0, '用户群组排名分析', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (47, '2024-10-29 15:26:27.206', '2024-11-12 15:34:06.108', NULL, 0, 32, 'userAction', 'userAction', 0, 'view/flowMonitor/userAction.vue', 7, '', 0, 0, '用户行为分析', 'menu', 0);
INSERT INTO `sys_base_menus` VALUES (48, '2024-10-29 15:27:07.215', '2024-10-29 15:27:07.215', NULL, 0, 32, 'alarmLog', 'alarmLog', 0, 'view/flowMonitor/alarmLog.vue', 8, '', 0, 0, '流量告警日志', 'list', 0);
INSERT INTO `sys_base_menus` VALUES (49, '2024-10-29 15:29:44.437', '2024-10-29 16:02:29.015', NULL, 0, 33, 'forward', 'forward', 0, 'view/policyManage/forward.vue', 1, '', 0, 0, '优先转发策略配置', 'position', 0);
INSERT INTO `sys_base_menus` VALUES (50, '2024-10-29 15:30:58.337', '2024-10-29 15:30:58.337', NULL, 0, 33, 'control', 'control', 0, 'view/policyManage/control.vue', 2, '', 0, 0, '策略管控配置', 'connection', 0);
INSERT INTO `sys_base_menus` VALUES (51, '2024-10-29 15:31:59.377', '2024-10-29 15:31:59.377', NULL, 0, 33, 'flowAlarm', 'flowAlarm', 0, 'view/policyManage/flowAlarm.vue', 3, '', 0, 0, '业务流量告警设置', 'setting', 0);
INSERT INTO `sys_base_menus` VALUES (52, '2024-10-29 15:36:22.063', '2024-10-29 15:36:22.063', NULL, 0, 3, 'task', 'task', 0, 'view/superAdmin/task/task.vue', 4, '', 0, 0, '定时任务', 'timer', 0);
INSERT INTO `sys_base_menus` VALUES (53, '2024-10-29 15:37:50.842', '2024-10-29 15:37:50.842', NULL, 0, 34, 'app', 'app', 0, 'view/customConfig/app.vue', 1, '', 0, 0, '应用分类与自定义', 'setting', 0);

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (888, 888);
INSERT INTO `sys_data_authority_id` VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` VALUES (9528, 9528);

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.338', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.344', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.351', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.357', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.364', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2024-10-15 13:05:01.333', '2024-10-15 13:05:01.373', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(0) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2024-10-15 13:05:01.340', '2024-10-15 13:05:01.340', NULL, '男', '1', '', 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2024-10-15 13:05:01.340', '2024-10-15 13:05:01.340', NULL, '女', '2', '', 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'smallint', '1', 'mysql', 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'mediumint', '2', 'mysql', 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'int', '3', 'mysql', 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'bigint', '4', 'mysql', 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'int2', '5', 'pgsql', 1, 5, 2);
INSERT INTO `sys_dictionary_details` VALUES (8, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'int4', '6', 'pgsql', 1, 6, 2);
INSERT INTO `sys_dictionary_details` VALUES (9, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'int6', '7', 'pgsql', 1, 7, 2);
INSERT INTO `sys_dictionary_details` VALUES (10, '2024-10-15 13:05:01.345', '2024-10-15 13:05:01.345', NULL, 'int8', '8', 'pgsql', 1, 8, 2);
INSERT INTO `sys_dictionary_details` VALUES (11, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'date', '', '', 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'time', '1', 'mysql', 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (13, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'year', '2', 'mysql', 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (14, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'datetime', '3', 'mysql', 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (15, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'timestamp', '5', 'mysql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (16, '2024-10-15 13:05:01.352', '2024-10-15 13:05:01.352', NULL, 'timestamptz', '6', 'pgsql', 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (17, '2024-10-15 13:05:01.357', '2024-10-15 13:05:01.357', NULL, 'float', '', '', 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (18, '2024-10-15 13:05:01.357', '2024-10-15 13:05:01.357', NULL, 'double', '1', 'mysql', 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (19, '2024-10-15 13:05:01.357', '2024-10-15 13:05:01.357', NULL, 'decimal', '2', 'mysql', 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (20, '2024-10-15 13:05:01.357', '2024-10-15 13:05:01.357', NULL, 'numeric', '3', 'pgsql', 1, 3, 4);
INSERT INTO `sys_dictionary_details` VALUES (21, '2024-10-15 13:05:01.357', '2024-10-15 13:05:01.357', NULL, 'smallserial', '4', 'pgsql', 1, 4, 4);
INSERT INTO `sys_dictionary_details` VALUES (22, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'char', '', '', 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'varchar', '1', 'mysql', 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'tinyblob', '2', 'mysql', 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'tinytext', '3', 'mysql', 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (26, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'text', '4', 'mysql', 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (27, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'blob', '5', 'mysql', 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (28, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'mediumblob', '6', 'mysql', 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (29, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'mediumtext', '7', 'mysql', 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (30, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'longblob', '8', 'mysql', 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (31, '2024-10-15 13:05:01.366', '2024-10-15 13:05:01.366', NULL, 'longtext', '9', 'mysql', 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (32, '2024-10-15 13:05:01.374', '2024-10-15 13:05:01.374', NULL, 'tinyint', '1', 'mysql', 1, 0, 6);
INSERT INTO `sys_dictionary_details` VALUES (33, '2024-10-15 13:05:01.374', '2024-10-15 13:05:01.374', NULL, 'bool', '2', 'pgsql', 1, 0, 6);

-- ----------------------------
-- Table structure for sys_ignore_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_ignore_apis`;
CREATE TABLE `sys_ignore_apis`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_ignore_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_ignore_apis
-- ----------------------------
INSERT INTO `sys_ignore_apis` VALUES (1, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/swagger/*any', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (2, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/api/freshCasbin', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (3, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/uploads/file/*filepath', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (4, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/health', 'GET');
INSERT INTO `sys_ignore_apis` VALUES (5, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/uploads/file/*filepath', 'HEAD');
INSERT INTO `sys_ignore_apis` VALUES (6, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/autoCode/llmAuto', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (7, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/system/reloadSystem', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (8, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/base/login', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (9, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/base/captcha', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (10, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/init/initdb', 'POST');
INSERT INTO `sys_ignore_apis` VALUES (11, '2024-10-15 13:05:01.301', '2024-10-15 13:05:01.301', NULL, '/init/checkdb', 'POST');

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint(0) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(0) DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 8881);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (2, 888);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `authority_id` bigint unsigned COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(0) DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid`) USING BTREE,
  INDEX `idx_sys_users_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2024-10-15 13:05:01.501', '2024-10-15 13:05:01.506', NULL, '40f79bf4-9df4-4d61-b521-31311e473023', 'admin', '$2a$10$mUtjmdxOM9iBIw9yXW5UY.19nAMTkXtrFEb8UssPDLzEsnYtIhZLO', '超级管理员', 'dark', '', '#fff', 888, '17611111111', '333333333@qq.com', 1);
# INSERT INTO `sys_users` VALUES (2, '2024-10-15 13:05:01.501', '2024-10-15 13:05:01.513', NULL, 'a51e08a6-351b-44cb-8e55-100f731e179f', 'a303176530', '$2a$10$S14lJ8v8ORRfn75sU7vdGuDOMiSdGDIHnWdQTm65N.AgRpKr3EMam', '用户1', 'dark', '', '#fff', 9528, '17611111111', '333333333@qq.com', 1);

SET FOREIGN_KEY_CHECKS = 1;
