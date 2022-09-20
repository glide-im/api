-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: new_im
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
-- Table structure for table `im_friend_records`
--

DROP TABLE IF EXISTS `im_friend_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_friend_records` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) NOT NULL,
  `to_id` int(11) NOT NULL,
  `status` tinyint(1) DEFAULT NULL COMMENT '0 等待通过 1 已通过 2 已拒绝',
  `created_at` timestamp NULL DEFAULT NULL,
  `information` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_friends`
--

DROP TABLE IF EXISTS `im_friends`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_friends` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) DEFAULT NULL,
  `to_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `note` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `top_time` datetime DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0' COMMENT '0.未置顶 1.已置顶',
  `uid` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_messages`
--

DROP TABLE IF EXISTS `im_group_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` json NOT NULL COMMENT '消息实体',
  `send_time` bigint(20) DEFAULT NULL COMMENT '消息添加时间',
  `message_id` int(11) DEFAULT NULL COMMENT '服务端消息id',
  `client_message_id` int(11) DEFAULT NULL COMMENT '客户端消息id',
  `form_id` int(11) DEFAULT NULL COMMENT '消息发送者id',
  `group_id` int(11) DEFAULT NULL COMMENT '群聊id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_user_messages`
--

DROP TABLE IF EXISTS `im_group_user_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_user_messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0' COMMENT '0 未读 1 已读',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_group_users`
--

DROP TABLE IF EXISTS `im_group_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_group_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `group_id` int(11) DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=179 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_groups`
--

DROP TABLE IF EXISTS `im_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `group_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `info` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `group_avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_messages`
--

DROP TABLE IF EXISTS `im_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `msg` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `form_id` int(11) DEFAULT NULL,
  `to_id` int(11) DEFAULT NULL,
  `is_read` tinyint(1) DEFAULT NULL COMMENT '0 未读 1已读',
  `msg_type` tinyint(1) DEFAULT '1',
  `status` tinyint(1) DEFAULT NULL,
  `data` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=591 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_offline_messages`
--

DROP TABLE IF EXISTS `im_offline_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_offline_messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` json DEFAULT NULL COMMENT '消息体',
  `send_time` int(11) DEFAULT NULL COMMENT '消息接收时间',
  `status` tinyint(1) DEFAULT NULL COMMENT '消息状态 0.未推送 1.已推送',
  `receive_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_sessions`
--

DROP TABLE IF EXISTS `im_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_sessions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '会话表',
  `form_id` int(11) NOT NULL,
  `to_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `top_status` tinyint(1) DEFAULT '0' COMMENT '0.否 1.是',
  `top_time` timestamp NULL DEFAULT NULL,
  `note` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `channel_type` tinyint(1) DEFAULT '0' COMMENT '0.单聊 1.群聊',
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '会话名称',
  `avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '会话头像',
  `status` tinyint(1) DEFAULT '0' COMMENT '会话状态 0.正常 1.禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `im_users`
--

DROP TABLE IF EXISTS `im_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `im_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `oauth_id` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方id',
  `bound_oauth` tinyint(1) DEFAULT '0' COMMENT '1\\github 2\\gitee',
  `oauth_type` tinyint(1) DEFAULT NULL COMMENT '1.微博 2.github',
  `status` tinyint(1) DEFAULT '0' COMMENT '0 离线 1 在线',
  `bio` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户简介',
  `sex` tinyint(1) DEFAULT '0' COMMENT '0 未知 1.男 2.女',
  `client_type` tinyint(1) DEFAULT NULL COMMENT '1.web 2.pc 3.app',
  `age` int(3) DEFAULT NULL,
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `uid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'uid 关联',
  `user_json` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-20 12:27:50
