/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 03/04/2020 16:39:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_article
-- ----------------------------
DROP TABLE IF EXISTS `admin_article`;
CREATE TABLE `admin_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
  `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '封面',
  `content` text NOT NULL COMMENT '内容',
  `is_scroll` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否轮播，0不轮播1轮播',
  `is_recommend` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '是否推荐，0不推荐1推荐',
  `allow_comment` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许评论，0不允许1允许',
  `category_id` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '栏目id',
  `favor_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `manager_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态，0已撤回1已发布',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- ----------------------------
-- Table structure for admin_category
-- ----------------------------
DROP TABLE IF EXISTS `admin_category`;
CREATE TABLE `admin_category` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `short_name` varchar(20) NOT NULL DEFAULT '' COMMENT '简称',
  `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='栏目表';

-- ----------------------------
-- Table structure for admin_log
-- ----------------------------
DROP TABLE IF EXISTS `admin_log`;
CREATE TABLE `admin_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `ip` varchar(20) NOT NULL DEFAULT '' COMMENT 'ip',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '访问的url',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '请求方式',
  `query` varchar(2048) NOT NULL DEFAULT '' COMMENT '地址栏参数',
  `headers` text NOT NULL COMMENT '请求头',
  `body` text NOT NULL COMMENT '请求体',
  `response` text NOT NULL COMMENT '响应',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '操作内容',
  `reason` text NOT NULL COMMENT '原因，操作失败时记录',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台日志表';

-- ----------------------------
-- Table structure for admin_manager
-- ----------------------------
DROP TABLE IF EXISTS `admin_manager`;
CREATE TABLE `admin_manager` (
  `id` tinyint(4) NOT NULL,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Table structure for admin_pemission
-- ----------------------------
DROP TABLE IF EXISTS `admin_pemission`;
CREATE TABLE `admin_pemission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Table structure for admin_role_manager
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_manager`;
CREATE TABLE `admin_role_manager` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员角色表';

-- ----------------------------
-- Table structure for admin_tag
-- ----------------------------
DROP TABLE IF EXISTS `admin_tag`;
CREATE TABLE `admin_tag` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `manager_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '管理员id',
  `article_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章数',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='标签表';

-- ----------------------------
-- Table structure for home_comment
-- ----------------------------
DROP TABLE IF EXISTS `home_comment`;
CREATE TABLE `home_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `article_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级评论id',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- ----------------------------
-- Table structure for home_favor_record
-- ----------------------------
DROP TABLE IF EXISTS `home_favor_record`;
CREATE TABLE `home_favor_record` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `article_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `ip` varchar(30) NOT NULL DEFAULT '' COMMENT '用户ip',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='点赞记录表';

-- ----------------------------
-- Table structure for home_log
-- ----------------------------
DROP TABLE IF EXISTS `home_log`;
CREATE TABLE `home_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `ip` varchar(20) NOT NULL DEFAULT '' COMMENT 'ip',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '访问的url',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '请求方式',
  `query` varchar(2048) NOT NULL DEFAULT '' COMMENT '地址栏参数',
  `headers` text NOT NULL COMMENT '请求头',
  `body` text NOT NULL COMMENT '请求体',
  `response` text NOT NULL COMMENT '响应',
  `content` varchar(255) NOT NULL DEFAULT '' COMMENT '操作内容',
  `reason` text NOT NULL COMMENT '原因，操作失败时记录',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='前台日志表';

-- ----------------------------
-- Table structure for home_user
-- ----------------------------
DROP TABLE IF EXISTS `home_user`;
CREATE TABLE `home_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `allow_comment` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许评论，0不允许1允许',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

SET FOREIGN_KEY_CHECKS = 1;
