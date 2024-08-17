/*!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-11.4.2-MariaDB, for osx10.19 (arm64)
--
-- Host: localhost    Database: paopao
-- ------------------------------------------------------
-- Server version	11.4.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `p_attachment`
--

DROP TABLE IF EXISTS `p_attachment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_attachment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT 0,
  `file_size` bigint(20) NOT NULL,
  `img_width` bigint(20) NOT NULL DEFAULT 0,
  `img_height` bigint(20) NOT NULL DEFAULT 0,
  `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '1图片，2视频，3其他附件',
  `content` varchar(255) NOT NULL DEFAULT '',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_attachment_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100041 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='附件';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_attachment`
--

LOCK TABLES `p_attachment` WRITE;
/*!40000 ALTER TABLE `p_attachment` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_attachment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_captcha`
--

DROP TABLE IF EXISTS `p_captcha`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_captcha` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '验证码ID',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号',
  `captcha` varchar(16) NOT NULL DEFAULT '' COMMENT '验证码',
  `use_times` int(11) NOT NULL DEFAULT 0 COMMENT '使用次数',
  `expired_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '过期时间',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_captcha_phone` (`phone`) USING BTREE,
  KEY `idx_captcha_expired_on` (`expired_on`) USING BTREE,
  KEY `idx_captcha_use_times` (`use_times`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1021 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='手机验证码';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_captcha`
--

LOCK TABLES `p_captcha` WRITE;
/*!40000 ALTER TABLE `p_captcha` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_captcha` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_comment`
--

DROP TABLE IF EXISTS `p_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'POST ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `is_essence` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否精选',
  `reply_count` int(11) NOT NULL DEFAULT 0 COMMENT '回复数',
  `thumbs_up_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `thumbs_down_count` int(11) NOT NULL DEFAULT 0 COMMENT '点踩数',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_post_id` (`post_id`) USING BTREE,
  KEY `idx_comment_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6001736 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_comment`
--

LOCK TABLES `p_comment` WRITE;
/*!40000 ALTER TABLE `p_comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_comment_content`
--

DROP TABLE IF EXISTS `p_comment_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_comment_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `comment_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '评论ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `content` varchar(4000) NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint(4) NOT NULL DEFAULT 2 COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址',
  `sort` bigint(20) NOT NULL DEFAULT 100 COMMENT '排序，越小越靠前',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_content_comment_id` (`comment_id`) USING BTREE,
  KEY `idx_comment_content_user_id` (`user_id`) USING BTREE,
  KEY `idx_comment_content_type` (`type`) USING BTREE,
  KEY `idx_comment_content_sort` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11001738 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_comment_content`
--

LOCK TABLES `p_comment_content` WRITE;
/*!40000 ALTER TABLE `p_comment_content` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_comment_content` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_comment_metric`
--

DROP TABLE IF EXISTS `p_comment_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_comment_metric` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `comment_id` bigint(20) NOT NULL,
  `rank_score` bigint(20) NOT NULL DEFAULT 0,
  `incentive_score` int(11) NOT NULL DEFAULT 0,
  `decay_factor` int(11) NOT NULL DEFAULT 0,
  `motivation_factor` int(11) NOT NULL DEFAULT 0,
  `is_del` tinyint(4) NOT NULL DEFAULT 0,
  `created_on` bigint(20) NOT NULL DEFAULT 0,
  `modified_on` bigint(20) NOT NULL DEFAULT 0,
  `deleted_on` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_metric_comment_id_rank_score` (`comment_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_comment_metric`
--

LOCK TABLES `p_comment_metric` WRITE;
/*!40000 ALTER TABLE `p_comment_metric` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_comment_metric` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_comment_reply`
--

DROP TABLE IF EXISTS `p_comment_reply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_comment_reply` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '回复ID',
  `comment_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '评论ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `at_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '@用户ID',
  `content` varchar(4000) NOT NULL DEFAULT '' COMMENT '内容',
  `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `thumbs_up_count` int(11) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `thumbs_down_count` int(11) NOT NULL DEFAULT 0 COMMENT '点踩数',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_reply_comment_id` (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12000015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论回复';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_comment_reply`
--

LOCK TABLES `p_comment_reply` WRITE;
/*!40000 ALTER TABLE `p_comment_reply` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_comment_reply` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_community`
--

DROP TABLE IF EXISTS `p_community`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_community` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `banner_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_community`
--

LOCK TABLES `p_community` WRITE;
/*!40000 ALTER TABLE `p_community` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_community` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_community_members`
--

DROP TABLE IF EXISTS `p_community_members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_community_members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(11) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `role` enum('member','moderator','admin') NOT NULL,
  `joined_at` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `community_id` (`community_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `p_community_members_ibfk_1` FOREIGN KEY (`community_id`) REFERENCES `p_community` (`id`),
  CONSTRAINT `p_community_members_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `p_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_community_members`
--

LOCK TABLES `p_community_members` WRITE;
/*!40000 ALTER TABLE `p_community_members` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_community_members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_community_tags`
--

DROP TABLE IF EXISTS `p_community_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_community_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(11) NOT NULL,
  `tag_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `community_id` (`community_id`),
  CONSTRAINT `p_community_tags_ibfk_1` FOREIGN KEY (`community_id`) REFERENCES `p_community` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_community_tags`
--

LOCK TABLES `p_community_tags` WRITE;
/*!40000 ALTER TABLE `p_community_tags` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_community_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_contact`
--

DROP TABLE IF EXISTS `p_contact`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_contact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `friend_id` bigint(20) NOT NULL COMMENT '好友ID',
  `group_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '好友分组ID:默认为0无分组',
  `remark` varchar(32) NOT NULL DEFAULT '' COMMENT '好友备注',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友',
  `is_top` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否置顶, 0否, 1是',
  `is_black` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否为黑名单, 0否, 1是',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除好友, 0否, 1是',
  `notice_enable` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否有消息提醒, 0否, 1是',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_contact_user_friend` (`user_id`,`friend_id`) USING BTREE,
  KEY `idx_contact_user_friend_status` (`user_id`,`friend_id`,`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_contact`
--

LOCK TABLES `p_contact` WRITE;
/*!40000 ALTER TABLE `p_contact` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_contact` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_contact_group`
--

DROP TABLE IF EXISTS `p_contact_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_contact_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分组名称',
  `is_del` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否删除, 0否, 1是',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人分组';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_contact_group`
--

LOCK TABLES `p_contact_group` WRITE;
/*!40000 ALTER TABLE `p_contact_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_contact_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_following`
--

DROP TABLE IF EXISTS `p_following`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_following` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `follow_id` bigint(20) NOT NULL,
  `is_del` tinyint(4) NOT NULL DEFAULT 0,
  `created_on` bigint(20) NOT NULL DEFAULT 0,
  `modified_on` bigint(20) NOT NULL DEFAULT 0,
  `deleted_on` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_following_user_follow` (`user_id`,`follow_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_following`
--

LOCK TABLES `p_following` WRITE;
/*!40000 ALTER TABLE `p_following` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_following` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_message`
--

DROP TABLE IF EXISTS `p_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '消息通知ID',
  `sender_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '发送方用户ID',
  `receiver_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '接收方用户ID',
  `type` tinyint(4) NOT NULL DEFAULT 1 COMMENT '通知类型，1动态，2评论，3回复，4私信，99系统通知',
  `brief` varchar(255) NOT NULL DEFAULT '' COMMENT '摘要说明',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '详细内容',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '动态ID',
  `comment_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '评论ID',
  `reply_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '回复ID',
  `is_read` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否已读',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_message_receiver_user_id` (`receiver_user_id`) USING BTREE,
  KEY `idx_message_is_read` (`is_read`) USING BTREE,
  KEY `idx_message_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16000033 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='消息通知';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_message`
--

LOCK TABLES `p_message` WRITE;
/*!40000 ALTER TABLE `p_message` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_post`
--

DROP TABLE IF EXISTS `p_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主题ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `comment_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '评论数',
  `collection_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '收藏数',
  `upvote_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `share_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '分享数',
  `visibility` tinyint(4) NOT NULL DEFAULT 0 COMMENT '可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开',
  `is_top` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `is_essence` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否精华',
  `is_lock` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否锁定',
  `latest_replied_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '最新回复时间',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
  `attachment_price` bigint(20) NOT NULL DEFAULT 0 COMMENT '附件价格(分)',
  `ip` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `community_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_user_id` (`user_id`) USING BTREE,
  KEY `idx_post_visibility` (`visibility`) USING BTREE,
  KEY `community_id` (`community_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1080017989 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post`
--

LOCK TABLES `p_post` WRITE;
/*!40000 ALTER TABLE `p_post` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_post_attachment_bill`
--

DROP TABLE IF EXISTS `p_post_attachment_bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post_attachment_bill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '购买记录ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'POST ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `paid_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '支付金额',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_attachment_bill_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_attachment_bill_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5000002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章附件账单';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post_attachment_bill`
--

LOCK TABLES `p_post_attachment_bill` WRITE;
/*!40000 ALTER TABLE `p_post_attachment_bill` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_attachment_bill` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `p_post_by_comment`
--

DROP TABLE IF EXISTS `p_post_by_comment`;
/*!50001 DROP VIEW IF EXISTS `p_post_by_comment`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `p_post_by_comment` AS SELECT
 1 AS `id`,
  1 AS `user_id`,
  1 AS `comment_count`,
  1 AS `collection_count`,
  1 AS `upvote_count`,
  1 AS `share_count`,
  1 AS `visibility`,
  1 AS `is_top`,
  1 AS `is_essence`,
  1 AS `is_lock`,
  1 AS `latest_replied_on`,
  1 AS `tags`,
  1 AS `attachment_price`,
  1 AS `ip`,
  1 AS `ip_loc`,
  1 AS `created_on`,
  1 AS `modified_on`,
  1 AS `deleted_on`,
  1 AS `is_del`,
  1 AS `comment_user_id` */;
SET character_set_client = @saved_cs_client;

--
-- Temporary table structure for view `p_post_by_media`
--

DROP TABLE IF EXISTS `p_post_by_media`;
/*!50001 DROP VIEW IF EXISTS `p_post_by_media`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `p_post_by_media` AS SELECT
 1 AS `id`,
  1 AS `user_id`,
  1 AS `comment_count`,
  1 AS `collection_count`,
  1 AS `upvote_count`,
  1 AS `share_count`,
  1 AS `visibility`,
  1 AS `is_top`,
  1 AS `is_essence`,
  1 AS `is_lock`,
  1 AS `latest_replied_on`,
  1 AS `tags`,
  1 AS `attachment_price`,
  1 AS `ip`,
  1 AS `ip_loc`,
  1 AS `created_on`,
  1 AS `modified_on`,
  1 AS `deleted_on`,
  1 AS `is_del` */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `p_post_collection`
--

DROP TABLE IF EXISTS `p_post_collection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post_collection` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'POST ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_collection_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_collection_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000012 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章收藏';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post_collection`
--

LOCK TABLES `p_post_collection` WRITE;
/*!40000 ALTER TABLE `p_post_collection` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_collection` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_post_content`
--

DROP TABLE IF EXISTS `p_post_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'POST ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `content` varchar(4000) NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint(4) NOT NULL DEFAULT 2 COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源',
  `sort` int(11) NOT NULL DEFAULT 100 COMMENT '排序，越小越靠前',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_content_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_content_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=180022546 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post_content`
--

LOCK TABLES `p_post_content` WRITE;
/*!40000 ALTER TABLE `p_post_content` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_content` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_post_metric`
--

DROP TABLE IF EXISTS `p_post_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post_metric` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL,
  `rank_score` bigint(20) NOT NULL DEFAULT 0,
  `incentive_score` int(11) NOT NULL DEFAULT 0,
  `decay_factor` int(11) NOT NULL DEFAULT 0,
  `motivation_factor` int(11) NOT NULL DEFAULT 0,
  `is_del` tinyint(4) NOT NULL DEFAULT 0,
  `created_on` bigint(20) NOT NULL DEFAULT 0,
  `modified_on` bigint(20) NOT NULL DEFAULT 0,
  `deleted_on` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_metric_post_id_rank_score` (`post_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post_metric`
--

LOCK TABLES `p_post_metric` WRITE;
/*!40000 ALTER TABLE `p_post_metric` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_metric` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_post_star`
--

DROP TABLE IF EXISTS `p_post_star`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_post_star` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'POST ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_star_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_star_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000028 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章点赞';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_post_star`
--

LOCK TABLES `p_post_star` WRITE;
/*!40000 ALTER TABLE `p_post_star` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_star` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_tag`
--

DROP TABLE IF EXISTS `p_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_tag` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `tag` varchar(255) NOT NULL COMMENT '标签名',
  `quote_num` bigint(20) NOT NULL DEFAULT 0 COMMENT '引用数',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_tag_tag` (`tag`) USING BTREE,
  KEY `idx_tag_user_id` (`user_id`) USING BTREE,
  KEY `idx_tag_quote_num` (`quote_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9000065 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='标签';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_tag`
--

LOCK TABLES `p_tag` WRITE;
/*!40000 ALTER TABLE `p_tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_topic_user`
--

DROP TABLE IF EXISTS `p_topic_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_topic_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `topic_id` bigint(20) NOT NULL COMMENT '标签ID',
  `user_id` bigint(20) NOT NULL COMMENT '创建者ID',
  `alias_name` varchar(255) DEFAULT NULL COMMENT '别名',
  `remark` varchar(512) DEFAULT NULL COMMENT '备注',
  `quote_num` bigint(20) DEFAULT NULL COMMENT '引用数',
  `is_top` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否置顶 0 为未置顶、1 为已置顶',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  `reserve_a` varchar(255) DEFAULT NULL COMMENT '保留字段a',
  `reserve_b` varchar(255) DEFAULT NULL COMMENT '保留字段b',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_topic_user_uid_tid` (`topic_id`,`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户话题';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_topic_user`
--

LOCK TABLES `p_topic_user` WRITE;
/*!40000 ALTER TABLE `p_topic_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_topic_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_tweet_comment_thumbs`
--

DROP TABLE IF EXISTS `p_tweet_comment_thumbs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_tweet_comment_thumbs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'thumbs ID',
  `user_id` bigint(20) NOT NULL,
  `tweet_id` bigint(20) NOT NULL COMMENT '推文ID',
  `comment_id` bigint(20) NOT NULL COMMENT '评论ID',
  `reply_id` bigint(20) DEFAULT NULL COMMENT '评论回复ID',
  `comment_type` tinyint(4) NOT NULL DEFAULT 0 COMMENT '评论类型 0为推文评论、1为评论回复',
  `is_thumbs_up` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否点赞',
  `is_thumbs_down` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否点踩',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tweet_comment_thumbs_uid_tid` (`user_id`,`tweet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推文评论点赞';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_tweet_comment_thumbs`
--

LOCK TABLES `p_tweet_comment_thumbs` WRITE;
/*!40000 ALTER TABLE `p_tweet_comment_thumbs` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_tweet_comment_thumbs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_user`
--

DROP TABLE IF EXISTS `p_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nickname` varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT 'MD5密码',
  `salt` varchar(16) NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态，1正常，2停用',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `balance` bigint(20) NOT NULL COMMENT '用户余额（分）',
  `is_admin` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否管理员',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_username` (`username`) USING BTREE,
  KEY `idx_user_phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100058 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_user`
--

LOCK TABLES `p_user` WRITE;
/*!40000 ALTER TABLE `p_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_user_metric`
--

DROP TABLE IF EXISTS `p_user_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_user_metric` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `tweets_count` int(11) NOT NULL DEFAULT 0,
  `latest_trends_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '最新动态时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0,
  `created_on` bigint(20) NOT NULL DEFAULT 0,
  `modified_on` bigint(20) NOT NULL DEFAULT 0,
  `deleted_on` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_metric_user_id_tweets_count_trends` (`user_id`,`tweets_count`,`latest_trends_on`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_user_metric`
--

LOCK TABLES `p_user_metric` WRITE;
/*!40000 ALTER TABLE `p_user_metric` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_user_metric` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `p_user_relation`
--

DROP TABLE IF EXISTS `p_user_relation`;
/*!50001 DROP VIEW IF EXISTS `p_user_relation`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `p_user_relation` AS SELECT
 1 AS `user_id`,
  1 AS `he_uid`,
  1 AS `style` */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `p_wallet_recharge`
--

DROP TABLE IF EXISTS `p_wallet_recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_wallet_recharge` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '充值ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '充值金额',
  `trade_no` varchar(64) NOT NULL DEFAULT '' COMMENT '支付宝订单号',
  `trade_status` varchar(32) NOT NULL DEFAULT '' COMMENT '交易状态',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_wallet_recharge_user_id` (`user_id`) USING BTREE,
  KEY `idx_wallet_recharge_trade_no` (`trade_no`) USING BTREE,
  KEY `idx_wallet_recharge_trade_status` (`trade_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10023 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_wallet_recharge`
--

LOCK TABLES `p_wallet_recharge` WRITE;
/*!40000 ALTER TABLE `p_wallet_recharge` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_wallet_recharge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `p_wallet_statement`
--

DROP TABLE IF EXISTS `p_wallet_statement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `p_wallet_statement` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账单ID',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `change_amount` bigint(20) NOT NULL DEFAULT 0 COMMENT '变动金额',
  `balance_snapshot` bigint(20) NOT NULL DEFAULT 0 COMMENT '资金快照',
  `reason` varchar(255) NOT NULL COMMENT '变动原因',
  `post_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '关联动态',
  `created_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` bigint(20) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_wallet_statement_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `p_wallet_statement`
--

LOCK TABLES `p_wallet_statement` WRITE;
/*!40000 ALTER TABLE `p_wallet_statement` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_wallet_statement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES
(19,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `p_post_by_comment`
--

/*!50001 DROP VIEW IF EXISTS `p_post_by_comment`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_post_by_comment` AS select `P`.`id` AS `id`,`P`.`user_id` AS `user_id`,`P`.`comment_count` AS `comment_count`,`P`.`collection_count` AS `collection_count`,`P`.`upvote_count` AS `upvote_count`,`P`.`share_count` AS `share_count`,`P`.`visibility` AS `visibility`,`P`.`is_top` AS `is_top`,`P`.`is_essence` AS `is_essence`,`P`.`is_lock` AS `is_lock`,`P`.`latest_replied_on` AS `latest_replied_on`,`P`.`tags` AS `tags`,`P`.`attachment_price` AS `attachment_price`,`P`.`ip` AS `ip`,`P`.`ip_loc` AS `ip_loc`,`P`.`created_on` AS `created_on`,`P`.`modified_on` AS `modified_on`,`P`.`deleted_on` AS `deleted_on`,`P`.`is_del` AS `is_del`,`c`.`user_id` AS `comment_user_id` from ((select `p_comment`.`post_id` AS `post_id`,`p_comment`.`user_id` AS `user_id` from `p_comment` where `p_comment`.`is_del` = 0 union select `COMMENT`.`post_id` AS `post_id`,`reply`.`user_id` AS `user_id` from (`p_comment_reply` `reply` join `p_comment` `COMMENT` on(`reply`.`comment_id` = `COMMENT`.`id`)) where `reply`.`is_del` = 0 and `COMMENT`.`is_del` = 0) `C` join `p_post` `P` on(`c`.`post_id` = `P`.`id`)) where `P`.`is_del` = 0 */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `p_post_by_media`
--

/*!50001 DROP VIEW IF EXISTS `p_post_by_media`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_post_by_media` AS select `post`.`id` AS `id`,`post`.`user_id` AS `user_id`,`post`.`comment_count` AS `comment_count`,`post`.`collection_count` AS `collection_count`,`post`.`upvote_count` AS `upvote_count`,`post`.`share_count` AS `share_count`,`post`.`visibility` AS `visibility`,`post`.`is_top` AS `is_top`,`post`.`is_essence` AS `is_essence`,`post`.`is_lock` AS `is_lock`,`post`.`latest_replied_on` AS `latest_replied_on`,`post`.`tags` AS `tags`,`post`.`attachment_price` AS `attachment_price`,`post`.`ip` AS `ip`,`post`.`ip_loc` AS `ip_loc`,`post`.`created_on` AS `created_on`,`post`.`modified_on` AS `modified_on`,`post`.`deleted_on` AS `deleted_on`,`post`.`is_del` AS `is_del` from ((select distinct `p_post_content`.`post_id` AS `post_id` from `p_post_content` where (`p_post_content`.`type` = 3 or `p_post_content`.`type` = 4 or `p_post_content`.`type` = 7 or `p_post_content`.`type` = 8) and `p_post_content`.`is_del` = 0) `media` join `p_post` `post` on(`media`.`post_id` = `post`.`id`)) where `post`.`is_del` = 0 */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `p_user_relation`
--

/*!50001 DROP VIEW IF EXISTS `p_user_relation`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_user_relation` AS select `p_contact`.`user_id` AS `user_id`,`p_contact`.`friend_id` AS `he_uid`,5 AS `style` from `p_contact` where `p_contact`.`status` = 2 and `p_contact`.`is_del` = 0 union select `p_following`.`user_id` AS `user_id`,`p_following`.`follow_id` AS `he_uid`,10 AS `style` from `p_following` where `p_following`.`is_del` = 0 */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2024-08-17 15:25:13
