/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : morty

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 11/09/2020 15:07:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_user
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `last_login` datetime(6) NULL DEFAULT NULL,
  `username` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `uid` int(11) NULL DEFAULT NULL COMMENT '对应Linux系统中UID',
  `gid` int(11) NULL DEFAULT NULL COMMENT '对应Linux系统中GID',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int(9) UNSIGNED ZEROFILL NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `upated_at` timestamp(0) NULL DEFAULT NULL,
  `describption` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `group_id` int(11) NULL DEFAULT NULL COMMENT '所属用户组id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `user_group`(`group_id`) USING BTREE,
  CONSTRAINT `user_group` FOREIGN KEY (`group_id`) REFERENCES `auth_group` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_user
-- ----------------------------
INSERT INTO `auth_user` VALUES (1, 'duliang', '2020-04-19 02:32:10.178691', 'admin', 'liangdu1992@gmail.com', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `auth_user` VALUES (2, 'pbkdf2_sha256$180000$8Nj0Px0cTfYR$/iIAdZiRzT4o+qKk8M7QNE+ZIrMM8pH/qUQMxnzVREY=', NULL, 'djangouser', '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `auth_user` VALUES (3, 'pbkdf2_sha256$180000$5dLhD1prXMC1$ajllEl2xrDX7q3maf+rIJqf1r1LaEXyDug17PAglRpo=', NULL, 'duliang', '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
