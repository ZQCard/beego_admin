/*
SQLyog Ultimate v12.5.0 (64 bit)
MySQL - 8.0.17 : Database - item_taige
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`item_taige` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `item_taige`;

/*Table structure for table `administrators` */

DROP TABLE IF EXISTS `administrators`;

CREATE TABLE `administrators` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(20) NOT NULL,
  `password` char(32) NOT NULL,
  `nickname` char(20) NOT NULL,
  `email` varchar(50) NOT NULL COMMENT '邮箱地址',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员表';

/*Data for the table `administrators` */

insert  into `administrators`(`id`,`username`,`password`,`nickname`,`email`,`updated_at`,`created_at`,`deleted_at`) values 
(1,'taige_admin','36168aa9dc70eb53dee09cde123bc5e2','超级管理员','445864742@qq.com','2019-11-02 17:07:27','2019-05-05 16:24:18',NULL),
(33,'taige','e10adc3949ba59abbe56e057f20f883e','管理员','4458647421@qq.com','2019-11-02 17:07:59','2019-10-24 22:08:47',NULL);

/*Table structure for table `auth_action` */

DROP TABLE IF EXISTS `auth_action`;

CREATE TABLE `auth_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '请求动作',
  `method` varchar(50) NOT NULL COMMENT '请求方法',
  `route` varchar(50) NOT NULL COMMENT '请求路由',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_操作表';

/*Data for the table `auth_action` */

insert  into `auth_action`(`id`,`name`,`method`,`route`) values 
(4,'登录首页','GET','/login'),
(5,'登录操作','POST','/login'),
(6,'后台首页','GET','/admin/'),
(8,'后台首页','GET','/admin/index'),
(9,'退出','GET','/admin/logout'),
(10,'管理员列表','GET','/admin/auth/administrator'),
(11,'管理员添加','POST','/admin/auth/administrator'),
(12,'管理员更新','PUT','/admin/auth/administrator'),
(13,'管理员删除','DELETE','/admin/auth/administrator'),
(14,'管理员恢复','PATCH','/admin/auth/administrator'),
(15,'角色列表','GET','/admin/auth/role'),
(16,'角色添加','POST','/admin/auth/role'),
(17,'角色更新','PUT','/admin/auth/role'),
(18,'角色删除','DELETE','/admin/auth/role'),
(19,'权限列表','GET','/admin/auth/permission'),
(20,'权限添加','POST','/admin/auth/permission'),
(21,'权限更新','PUT','/admin/auth/permission'),
(22,'权限删除','DELETE','/admin/auth/permission'),
(23,'行为列表','GET','/admin/auth/action'),
(24,'行为添加','POST','/admin/auth/action'),
(25,'行为更新','PUT','/admin/auth/action'),
(26,'行为删除','DELETE','/admin/auth/action'),
(27,'菜单列表','GET','/admin/auth/menu'),
(28,'菜单添加','POST','/admin/auth/menu'),
(29,'菜单更新','PUT','/admin/auth/menu'),
(30,'菜单删除','DELETE','/admin/auth/menu'),
(31,'基本信息','GET','/admin/administrator/adminInfo'),
(34,'查看管理员角色','GET','/admin/auth/administrator/roles'),
(35,'授予管理员角色','PUT','/admin/auth/administrator/roles'),
(36,'查看角色权限','GET','/admin/auth/role/permissions'),
(37,'授予权限角色','PUT','/admin/auth/role/permissions'),
(38,'查看权限行为','GET','/admin/auth/permission/actions'),
(39,'授予权限行为','PUT','/admin/auth/permission/actions'),
(47,'刷新权限','GET','/admin/auth/administrator/refreshAuth'),
(60,'视频列表','GET','/admin/video'),
(61,'视频详情页面','GET','/admin/video/info'),
(62,'视频创建','POST','/admin/video'),
(63,'视频编辑','PUT','/admin/video'),
(64,'视频删除','DELETE','/admin/video'),
(65,'视频恢复','PATCH','/admin/video'),
(82,'管理员信息查看','GET','/admin/administrator/info'),
(88,'错误提示页面','GET','/admin/error'),
(90,'管理员信息编辑','PUT','/admin/administrator/info');

/*Table structure for table `auth_administrator_role` */

DROP TABLE IF EXISTS `auth_administrator_role`;

CREATE TABLE `auth_administrator_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色关联表';

/*Data for the table `auth_administrator_role` */

insert  into `auth_administrator_role`(`id`,`administrator_id`,`role_id`) values 
(28,30,94),
(32,22,94),
(33,1,1),
(34,1,94),
(39,33,94);

/*Table structure for table `auth_menu` */

DROP TABLE IF EXISTS `auth_menu`;

CREATE TABLE `auth_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '父级菜单id',
  `name` char(20) NOT NULL COMMENT '菜单名单',
  `sort` tinyint(4) DEFAULT '0' COMMENT '排序值',
  `route` varchar(60) DEFAULT '' COMMENT '跳转链接',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_菜单表';

/*Data for the table `auth_menu` */

insert  into `auth_menu`(`id`,`pid`,`name`,`sort`,`route`) values 
(14,0,'设置',98,''),
(15,14,'权限管理',1,''),
(18,15,'角色列表',2,'/admin/auth/role'),
(19,15,'权限列表',3,'/admin/auth/permission'),
(20,15,'行为列表',4,'/admin/auth/action'),
(21,15,'菜单列表',5,'/admin/auth/menu'),
(23,15,'管理员列表',8,'/admin/auth/administrator'),
(31,0,'数据管理',1,''),
(34,31,'视频管理',2,''),
(36,34,'视频详情管理',2,'/admin/video'),
(51,50,'导航栏查看',1,'/admin/navigator'),
(53,52,'测试菜单2',1,'');

/*Table structure for table `auth_permission` */

DROP TABLE IF EXISTS `auth_permission`;

CREATE TABLE `auth_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_权限表';

/*Data for the table `auth_permission` */

insert  into `auth_permission`(`id`,`name`) values 
(22,'权限管理'),
(23,'公共权限'),
(24,'游客权限'),
(38,'数据管理权限');

/*Table structure for table `auth_permission_action` */

DROP TABLE IF EXISTS `auth_permission_action`;

CREATE TABLE `auth_permission_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=523 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_操作权限关联表';

/*Data for the table `auth_permission_action` */

insert  into `auth_permission_action`(`id`,`permission_id`,`action_id`) values 
(239,29,5),
(240,29,4),
(241,26,4),
(242,26,5),
(243,30,4),
(244,30,5),
(245,31,4),
(246,31,5),
(247,31,6),
(248,31,7),
(249,31,8),
(250,31,9),
(251,31,10),
(252,31,11),
(253,31,12),
(254,31,13),
(291,22,31),
(292,22,10),
(293,22,11),
(294,22,12),
(295,22,13),
(296,22,14),
(297,22,15),
(298,22,16),
(299,22,17),
(300,22,18),
(301,22,19),
(302,22,20),
(303,22,21),
(304,22,22),
(305,22,23),
(306,22,24),
(307,22,25),
(308,22,26),
(309,22,27),
(310,22,28),
(311,22,29),
(312,22,30),
(313,22,34),
(314,22,35),
(315,22,36),
(316,22,37),
(317,22,38),
(318,22,39),
(461,38,60),
(462,38,61),
(463,38,62),
(464,38,63),
(465,38,64),
(466,38,65),
(467,38,50),
(468,38,51),
(469,38,52),
(470,38,53),
(471,38,54),
(512,24,4),
(513,24,5),
(514,24,88),
(517,23,6),
(518,23,8),
(519,23,9),
(520,23,47),
(521,23,82),
(522,23,90);

/*Table structure for table `auth_role` */

DROP TABLE IF EXISTS `auth_role`;

CREATE TABLE `auth_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_角色表';

/*Data for the table `auth_role` */

insert  into `auth_role`(`id`,`name`) values 
(1,'超级管理员'),
(87,'游客'),
(94,'数据管理员');

/*Table structure for table `auth_role_permission` */

DROP TABLE IF EXISTS `auth_role_permission`;

CREATE TABLE `auth_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `permission_id` int(11) NOT NULL COMMENT '权限id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=175 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_角色权限关联表';

/*Data for the table `auth_role_permission` */

insert  into `auth_role_permission`(`id`,`role_id`,`permission_id`) values 
(149,1,40),
(150,1,41),
(151,1,42),
(152,1,22),
(153,1,23),
(154,1,24),
(155,1,35),
(156,1,37),
(157,1,38),
(158,1,39),
(159,94,44),
(160,94,37),
(161,94,38),
(162,94,39),
(163,94,40),
(164,94,41),
(165,94,42),
(171,87,23),
(172,87,24);

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(60) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '账号',
  `password` char(32) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '密码',
  `nickname` varchar(60) COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `wechat_open_id` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '微信openid',
  `head_img_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '头像',
  `sex` tinyint(1) DEFAULT '0' COMMENT '性别 0未知 1男 2女',
  `province` varchar(60) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '省份',
  `city` varchar(60) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '城市',
  `login_ip` varchar(20) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '登录ip',
  `login_times` int(11) DEFAULT '0' COMMENT '登录次数',
  `salt` char(8) COLLATE utf8mb4_general_ci DEFAULT '""' COMMENT '8位salt密码盐',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `users` */

insert  into `users`(`id`,`account`,`password`,`nickname`,`wechat_open_id`,`head_img_url`,`sex`,`province`,`city`,`login_ip`,`login_times`,`salt`,`created_at`,`updated_at`,`deleted_at`) values 
(2,'','','后来呢','oDdyz0kGN4nIDEn2WOw39_GYfoG0','http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL66S8DVK8QqakM2czRKQXViawt3lEQBk1Tqykfts83l3rsPVvC0BZzM4WeicQjgI3Tv3M954rnEyHw/132',1,'上海','浦东新区','61.171.209.12',0,'NNUUDPEJ','2019-10-29 00:17:01','2019-10-29 00:17:01',NULL);

/*Table structure for table `videos` */

DROP TABLE IF EXISTS `videos`;

CREATE TABLE `videos` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '视频标题',
  `url` varchar(255) NOT NULL COMMENT '视频播放地址',
  `poster` varchar(255) NOT NULL COMMENT '视频封面图地址',
  `view_times` int(11) NOT NULL DEFAULT '0' COMMENT '视频观看次数',
  `description` text COMMENT '视频描述',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='企业站视频表';

/*Data for the table `videos` */

insert  into `videos`(`id`,`title`,`url`,`poster`,`view_times`,`description`,`updated_at`,`created_at`,`deleted_at`) values 
(35,'WKAE','http://personal-item-taige.oss-cn-shanghai.aliyuncs.com/test.mp4','http://taige.niu12.com/static/uploadFile/poster/2019-11-02/MOBXJKWIBK_poster.jpg',4,'<p>拍摄 | 人员1</p><p>剪辑 | 人员2</p><p>导演 | 人员3</p>','2019-11-02 16:26:12','2019-11-02 16:26:12',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
