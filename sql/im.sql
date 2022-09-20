-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: go_im_test
-- ------------------------------------------------------
-- Server version	5.7.38-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `im_app`
--

DROP TABLE IF EXISTS `im_app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_app` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `app_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '平台ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '平台名称',
  `uid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '平台名称',
  `license` json DEFAULT NULL COMMENT '营业相关信息',
  `status` int(6) DEFAULT NULL COMMENT '状态: 1: 待审核; 2: 审核成功; 3: 被禁用;',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `logo` varchar(266) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号码',
  `email` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱账户',
  `host` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设定的域名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='平台应用表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_article`
--

DROP TABLE IF EXISTS `im_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `app_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '平台ID',
  `uid` int(11) NOT NULL COMMENT '发布人',
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
  `publish_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `weight` int(11) DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='帮助中心文章';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_category`
--

DROP TABLE IF EXISTS `im_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_category` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `app_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '平台ID',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `icon` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_category_user`
--

DROP TABLE IF EXISTS `im_category_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_category_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `app_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '平台ID',
  `category_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '关联分类ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uid` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '关联的用户ID',
  `form` int(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户 => 分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_chat_message`
--

DROP TABLE IF EXISTS `im_chat_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_chat_message` (
  `m_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `session_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `cli_seq` bigint(20) NOT NULL,
  `from` bigint(20) NOT NULL,
  `to` bigint(20) NOT NULL,
  `type` int(11) NOT NULL,
  `send_at` bigint(20) NOT NULL,
  `create_at` bigint(20) NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`m_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1783 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_collect_data`
--

DROP TABLE IF EXISTS `im_collect_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_collect_data` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ',
  `app_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT 'app id',
  `uid` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1' COMMENT '用户ID',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名称',
  `region` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '地区',
  `browser` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '浏览器',
  `device` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '设备(端)',
  `origin` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '来源',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=334 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户数据';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_contacts`
--

DROP TABLE IF EXISTS `im_contacts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_contacts` (
  `fid` char(254) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uid` bigint(20) NOT NULL,
  `id` bigint(20) NOT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `type` int(11) NOT NULL,
  `last_mid` int(11) DEFAULT NULL COMMENT '最后一次聊天更新的ID',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '联系人状态: 1: 正常关系; 2: 等待同意; 3: 双方不存在好友关系;',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`fid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_member_model`
--

DROP TABLE IF EXISTS `im_group_member_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_member_model` (
  `mb_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `gid` bigint(20) DEFAULT NULL,
  `uid` bigint(20) DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  `type` bigint(20) DEFAULT NULL,
  `remark` char(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`mb_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_member_msg_state`
--

DROP TABLE IF EXISTS `im_group_member_msg_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_member_msg_state` (
  `mb_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `g_id` bigint(20) DEFAULT NULL,
  `uid` bigint(20) DEFAULT NULL,
  `last_ack_m_id` bigint(20) DEFAULT NULL,
  `last_ack_seq` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`mb_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_message`
--

DROP TABLE IF EXISTS `im_group_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_message` (
  `m_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `seq` bigint(20) NOT NULL,
  `to` bigint(20) NOT NULL,
  `from` bigint(20) NOT NULL,
  `type` bigint(20) NOT NULL,
  `send_at` bigint(20) NOT NULL,
  `content` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int(11) NOT NULL,
  `recall_by` int(11) NOT NULL,
  PRIMARY KEY (`m_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_message_state`
--

DROP TABLE IF EXISTS `im_group_message_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_message_state` (
  `gid` bigint(20) NOT NULL AUTO_INCREMENT,
  `last_m_id` bigint(20) DEFAULT NULL,
  `last_seq` bigint(20) DEFAULT NULL,
  `last_msg_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`gid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_model`
--

DROP TABLE IF EXISTS `im_group_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_model` (
  `gid` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `mute` tinyint(1) DEFAULT NULL,
  `flag` int(11) DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`gid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_msg_seq`
--

DROP TABLE IF EXISTS `im_group_msg_seq`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_msg_seq` (
  `gid` bigint(20) NOT NULL AUTO_INCREMENT,
  `seq` bigint(20) DEFAULT NULL,
  `step` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`gid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_offline_message`
--

DROP TABLE IF EXISTS `im_offline_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_offline_message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `m_id` bigint(20) DEFAULT NULL,
  `uid` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_user`
--

DROP TABLE IF EXISTS `im_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_user` (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `fingerprint_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '指纹ID(guest账户登录使用)',
  `update_at` bigint(20) DEFAULT NULL,
  `create_at` bigint(20) DEFAULT NULL,
  `app_id` int(11) DEFAULT NULL,
  `role` int(11) DEFAULT '1' COMMENT '角色: 1: 客户; 2:访问用户;',
  `account` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=543843 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-20 12:58:12
