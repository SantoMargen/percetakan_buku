/*
 Navicat Premium Data Transfer

 Source Server         : 172.18.135.223-LAN
 Source Server Type    : MySQL
 Source Server Version : 80020 (8.0.20)
 Source Host           : 172.18.135.223:3306
 Source Schema         : rkf

 Target Server Type    : MySQL
 Target Server Version : 80020 (8.0.20)
 File Encoding         : 65001

 Date: 19/11/2024 08:19:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for data_rkf
-- ----------------------------
DROP TABLE IF EXISTS `data_rkf`;
CREATE TABLE `data_rkf` (
  `id_rkf` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `periode` int DEFAULT NULL,
  `is_revisi` int DEFAULT NULL,
  `kostl` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `kostl_tx` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `entry_user` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `entry_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `entry_time` datetime DEFAULT NULL,
  `updated_user` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `updated_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `updated_time` datetime DEFAULT NULL,
  `approval_posisi_divisi` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `approval_list_divisi` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `status` int DEFAULT NULL,
  `file` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `catatan_divisi` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `catatan_tolakan_detail` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `catatan_maker_validasi` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `catatan_validasi` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `catatan_tolakan_detail_validasi` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `entry_user_ppm` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `entry_name_ppm` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `entry_time_ppm` datetime DEFAULT NULL,
  `updated_user_ppm` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `updated_name_ppm` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `updated_time_ppm` datetime DEFAULT NULL,
  `approval_posisi_ppm` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '',
  `approval_list_ppm` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  PRIMARY KEY (`id_rkf`) USING BTREE,
  KEY `IDX_ID_PERIODE` (`periode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
