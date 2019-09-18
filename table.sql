/*
Navicat MySQL Data Transfer

Source Server         : 47.97.215.189
Source Server Version : 80017
Source Host           : 47.97.215.189:3306
Source Database       : item_fmg

Target Server Type    : MYSQL
Target Server Version : 80017
File Encoding         : 65001

Date: 2019-09-18 11:09:59
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for administrator
-- ----------------------------
DROP TABLE IF EXISTS `administrator`;
CREATE TABLE `administrator` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` char(20) NOT NULL,
  `password` char(32) NOT NULL,
  `nickname` char(20) NOT NULL,
  `email` varchar(50) NOT NULL COMMENT '邮箱地址',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员表';

-- ----------------------------
-- Records of administrator
-- ----------------------------
INSERT INTO `administrator` VALUES ('1', 'zhouqi', 'e10adc3949ba59abbe56e057f20f883e', '周起', '445864742@qq.com', '2019-09-09 21:52:10', '2019-05-05 16:24:18', null);
INSERT INTO `administrator` VALUES ('22', 'user111', 'e10adc3949ba59abbe56e057f20f883e', '普通用户', '645144262@qq.com', '2019-06-12 16:06:05', '2019-06-12 16:06:05', '2019-08-29 22:07:31');
INSERT INTO `administrator` VALUES ('30', 'zhouqi1', 'e10adc3949ba59abbe56e057f20f883e', '22222', '1@qq.com', '2019-08-29 23:37:41', '2019-08-29 23:04:17', null);
INSERT INTO `administrator` VALUES ('31', 'zhouqi11', '111111', '1111', '1@qq.com', '2019-08-29 23:37:46', '2019-08-29 23:37:21', null);
INSERT INTO `administrator` VALUES ('32', 'zhouqi111', '123456', '111', '445864742@qq.com', '2019-09-09 22:11:29', '2019-09-09 22:11:10', null);

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '请求动作',
  `method` varchar(50) NOT NULL COMMENT '请求方法',
  `route` varchar(50) NOT NULL COMMENT '请求路由',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_操作表';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES ('4', '登录首页', 'GET', '/login');
INSERT INTO `auth_action` VALUES ('5', '登录操作', 'POST', '/login');
INSERT INTO `auth_action` VALUES ('6', '后台首页', 'GET', '/admin/');
INSERT INTO `auth_action` VALUES ('7', '测试接口', 'GET', '/admin/');
INSERT INTO `auth_action` VALUES ('8', '后台首页', 'GET', '/admin/index');
INSERT INTO `auth_action` VALUES ('9', '退出', 'GET', '/admin/logout');
INSERT INTO `auth_action` VALUES ('10', '管理员列表', 'GET', '/admin/auth/administrator');
INSERT INTO `auth_action` VALUES ('11', '管理员添加', 'POST', '/admin/auth/administrator');
INSERT INTO `auth_action` VALUES ('12', '管理员更新', 'PUT', '/admin/auth/administrator');
INSERT INTO `auth_action` VALUES ('13', '管理员删除', 'DELETE', '/admin/auth/administrator');
INSERT INTO `auth_action` VALUES ('14', '管理员恢复', 'PATCH', '/admin/auth/administrator');
INSERT INTO `auth_action` VALUES ('15', '角色列表', 'GET', '/admin/auth/role');
INSERT INTO `auth_action` VALUES ('16', '角色添加', 'POST', '/admin/auth/role');
INSERT INTO `auth_action` VALUES ('17', '角色更新', 'PUT', '/admin/auth/role');
INSERT INTO `auth_action` VALUES ('18', '角色删除', 'DELETE', '/admin/auth/role');
INSERT INTO `auth_action` VALUES ('19', '权限列表', 'GET', '/admin/auth/permission');
INSERT INTO `auth_action` VALUES ('20', '权限添加', 'POST', '/admin/auth/permission');
INSERT INTO `auth_action` VALUES ('21', '权限更新', 'PUT', '/admin/auth/permission');
INSERT INTO `auth_action` VALUES ('22', '权限删除', 'DELETE', '/admin/auth/permission');
INSERT INTO `auth_action` VALUES ('23', '行为列表', 'GET', '/admin/auth/action');
INSERT INTO `auth_action` VALUES ('24', '行为添加', 'POST', '/admin/auth/action');
INSERT INTO `auth_action` VALUES ('25', '行为更新', 'PUT', '/admin/auth/action');
INSERT INTO `auth_action` VALUES ('26', '行为删除', 'DELETE', '/admin/auth/action');
INSERT INTO `auth_action` VALUES ('27', '菜单列表', 'GET', '/admin/auth/menu');
INSERT INTO `auth_action` VALUES ('28', '菜单添加', 'POST', '/admin/auth/menu');
INSERT INTO `auth_action` VALUES ('29', '菜单更新', 'PUT', '/admin/auth/menu');
INSERT INTO `auth_action` VALUES ('30', '菜单删除', 'DELETE', '/admin/auth/menu');
INSERT INTO `auth_action` VALUES ('31', '基本信息', 'GET', '/admin/administrator/adminInfo');
INSERT INTO `auth_action` VALUES ('33', '上传文件', 'POST', '/admin/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('34', '查看管理员角色', 'GET', '/admin/auth/administrator/roles');
INSERT INTO `auth_action` VALUES ('35', '授予管理员角色', 'PUT', '/admin/auth/administrator/roles');
INSERT INTO `auth_action` VALUES ('36', '查看角色权限', 'GET', '/admin/auth/role/permissions');
INSERT INTO `auth_action` VALUES ('37', '授予权限角色', 'PUT', '/admin/auth/role/permissions');
INSERT INTO `auth_action` VALUES ('38', '查看权限行为', 'GET', '/admin/auth/permission/actions');
INSERT INTO `auth_action` VALUES ('39', '授予权限行为', 'PUT', '/admin/auth/permission/actions');
INSERT INTO `auth_action` VALUES ('45', '文件上传页面', 'GET', '/admin/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('46', '文件上传', 'POST', '/admin/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('47', '刷新权限', 'GET', '/admin/auth/administrator/refreshAuth');
INSERT INTO `auth_action` VALUES ('48', '首页模块控制页面', 'GET', '/admin/homepage/setting');
INSERT INTO `auth_action` VALUES ('49', '首页模块控制设置', 'PATCH', '/admin/homepage/setting');
INSERT INTO `auth_action` VALUES ('50', '视频分类列表', 'GET', '/admin/video/category');
INSERT INTO `auth_action` VALUES ('51', '视频分类创建', 'POST', '/admin/video/category');
INSERT INTO `auth_action` VALUES ('52', '视频分类更新', 'PUT', '/admin/video/category');
INSERT INTO `auth_action` VALUES ('53', '视频分类删除', 'DELETE', '/admin/video/category');
INSERT INTO `auth_action` VALUES ('54', '视频分类恢复', 'PATCH', '/admin/video/category');
INSERT INTO `auth_action` VALUES ('55', '资料分类列表', 'GET', '/admin/documentation/category');
INSERT INTO `auth_action` VALUES ('56', '资料分类创建', 'POST', '/admin/documentation/category');
INSERT INTO `auth_action` VALUES ('57', '资料分类更新', 'PUT', '/admin/documentation/category');
INSERT INTO `auth_action` VALUES ('58', '资料分类删除', 'DELETE', '/admin/documentation/category');
INSERT INTO `auth_action` VALUES ('59', '资料分类恢复', 'PATCH', '/admin/documentation/category');
INSERT INTO `auth_action` VALUES ('60', '视频列表', 'GET', '/admin/video');
INSERT INTO `auth_action` VALUES ('61', '视频详情页面', 'GET', '/admin/video/info');
INSERT INTO `auth_action` VALUES ('62', '视频创建', 'POST', '/admin/video');
INSERT INTO `auth_action` VALUES ('63', '视频编辑', 'PUT', '/admin/video');
INSERT INTO `auth_action` VALUES ('64', '视频删除', 'DELETE', '/admin/video');
INSERT INTO `auth_action` VALUES ('65', '视频恢复', 'PATCH', '/admin/video');
INSERT INTO `auth_action` VALUES ('66', '资料列表', 'GET', '/admin/documentation');
INSERT INTO `auth_action` VALUES ('67', '资料详情页面', 'GET', '/admin/documentation/info');
INSERT INTO `auth_action` VALUES ('68', '资料创建', 'POST', '/admin/documentation');
INSERT INTO `auth_action` VALUES ('69', '资料编辑', 'PUT', '/admin/documentation');
INSERT INTO `auth_action` VALUES ('70', '资料删除', 'DELETE', '/admin/documentation');
INSERT INTO `auth_action` VALUES ('71', '资料恢复', 'PATCH', '/admin/documentation');
INSERT INTO `auth_action` VALUES ('72', '轮播图列表', 'GET', '/admin/banner');
INSERT INTO `auth_action` VALUES ('73', '轮播图创建', 'POST', '/admin/banner');
INSERT INTO `auth_action` VALUES ('74', '轮播图编辑', 'PUT', '/admin/banner');
INSERT INTO `auth_action` VALUES ('75', '轮播图删除', 'DELETE', '/admin/banner');
INSERT INTO `auth_action` VALUES ('76', '轮播图恢复', 'PATCH', '/admin/banner');
INSERT INTO `auth_action` VALUES ('77', '留言板查看', 'GET', '/admin/message');
INSERT INTO `auth_action` VALUES ('78', '用户查看', 'GET', '/admin/user');
INSERT INTO `auth_action` VALUES ('79', '用户删除', 'DELETE', '/admin/user');
INSERT INTO `auth_action` VALUES ('80', '用户恢复', 'PATCH', '/admin/user');
INSERT INTO `auth_action` VALUES ('82', '管理员信息查看', 'GET', '/admin//administrator/info');
INSERT INTO `auth_action` VALUES ('83', '导航栏列表', 'GET', '/admin/navigator');
INSERT INTO `auth_action` VALUES ('84', '导航栏更新', 'PUT', '/admin/navigator');
INSERT INTO `auth_action` VALUES ('85', '导航栏删除', 'DELETE', '/admin/navigator');
INSERT INTO `auth_action` VALUES ('86', '导航栏创建', 'POST', '/admin/navigator');
INSERT INTO `auth_action` VALUES ('87', '导航栏恢复', 'PATCH', '/admin/navigator');
INSERT INTO `auth_action` VALUES ('88', '错误提示页面', 'GET', '/admin/error');

-- ----------------------------
-- Table structure for auth_administrator_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_administrator_role`;
CREATE TABLE `auth_administrator_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色关联表';

-- ----------------------------
-- Records of auth_administrator_role
-- ----------------------------
INSERT INTO `auth_administrator_role` VALUES ('28', '30', '94');
INSERT INTO `auth_administrator_role` VALUES ('32', '22', '94');
INSERT INTO `auth_administrator_role` VALUES ('33', '1', '1');
INSERT INTO `auth_administrator_role` VALUES ('34', '1', '94');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '父级菜单id',
  `name` char(20) NOT NULL COMMENT '菜单名单',
  `sort` tinyint(4) DEFAULT '0' COMMENT '排序值',
  `route` varchar(60) DEFAULT '' COMMENT '跳转链接',
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('14', '0', '设置', '98', '');
INSERT INTO `auth_menu` VALUES ('15', '14', '权限管理', '1', '');
INSERT INTO `auth_menu` VALUES ('18', '15', '角色列表', '2', '/admin/auth/role');
INSERT INTO `auth_menu` VALUES ('19', '15', '权限列表', '3', '/admin/auth/permission');
INSERT INTO `auth_menu` VALUES ('20', '15', '行为列表', '4', '/admin/auth/action');
INSERT INTO `auth_menu` VALUES ('21', '15', '菜单列表', '5', '/admin/auth/menu');
INSERT INTO `auth_menu` VALUES ('23', '15', '管理员列表', '8', '/admin/auth/administrator');
INSERT INTO `auth_menu` VALUES ('31', '0', '企业站管理', '1', '');
INSERT INTO `auth_menu` VALUES ('32', '31', '首页管理', '1', '');
INSERT INTO `auth_menu` VALUES ('33', '32', '首页模块控制', '1', '/admin/homepage/setting');
INSERT INTO `auth_menu` VALUES ('34', '31', '视频管理', '2', '');
INSERT INTO `auth_menu` VALUES ('35', '34', '视频分类管理', '1', '/admin/video/category');
INSERT INTO `auth_menu` VALUES ('36', '34', '视频详情管理', '2', '/admin/video');
INSERT INTO `auth_menu` VALUES ('37', '31', '资料管理', '3', '');
INSERT INTO `auth_menu` VALUES ('38', '37', '资料分类管理', '1', '/admin/documentation/category');
INSERT INTO `auth_menu` VALUES ('39', '37', '资料详情管理', '2', '/admin/documentation');
INSERT INTO `auth_menu` VALUES ('40', '31', '轮播图管理', '4', '');
INSERT INTO `auth_menu` VALUES ('41', '40', '轮播图', '1', '/admin/banner');
INSERT INTO `auth_menu` VALUES ('43', '31', '留言板管理', '5', '');
INSERT INTO `auth_menu` VALUES ('44', '43', '留言查看', '1', '/admin/message');
INSERT INTO `auth_menu` VALUES ('46', '31', '用户管理', '6', '');
INSERT INTO `auth_menu` VALUES ('47', '46', '用户查看', '1', '/admin/user');
INSERT INTO `auth_menu` VALUES ('50', '31', '导航栏管理', '7', '');
INSERT INTO `auth_menu` VALUES ('51', '50', '导航栏查看', '1', '/admin/navigator');

-- ----------------------------
-- Table structure for auth_permission
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission`;
CREATE TABLE `auth_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_权限表';

-- ----------------------------
-- Records of auth_permission
-- ----------------------------
INSERT INTO `auth_permission` VALUES ('22', '权限管理');
INSERT INTO `auth_permission` VALUES ('23', '公共权限');
INSERT INTO `auth_permission` VALUES ('24', '游客权限');
INSERT INTO `auth_permission` VALUES ('35', '开发权限');
INSERT INTO `auth_permission` VALUES ('37', '首页控制权限');
INSERT INTO `auth_permission` VALUES ('38', '视频管理权限');
INSERT INTO `auth_permission` VALUES ('39', '资料管理权限');
INSERT INTO `auth_permission` VALUES ('40', '轮播图管理');
INSERT INTO `auth_permission` VALUES ('41', '留言板管理');
INSERT INTO `auth_permission` VALUES ('42', '用户管理');
INSERT INTO `auth_permission` VALUES ('44', '导航栏管理');

-- ----------------------------
-- Table structure for auth_permission_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission_action`;
CREATE TABLE `auth_permission_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=515 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_操作权限关联表';

-- ----------------------------
-- Records of auth_permission_action
-- ----------------------------
INSERT INTO `auth_permission_action` VALUES ('239', '29', '5');
INSERT INTO `auth_permission_action` VALUES ('240', '29', '4');
INSERT INTO `auth_permission_action` VALUES ('241', '26', '4');
INSERT INTO `auth_permission_action` VALUES ('242', '26', '5');
INSERT INTO `auth_permission_action` VALUES ('243', '30', '4');
INSERT INTO `auth_permission_action` VALUES ('244', '30', '5');
INSERT INTO `auth_permission_action` VALUES ('245', '31', '4');
INSERT INTO `auth_permission_action` VALUES ('246', '31', '5');
INSERT INTO `auth_permission_action` VALUES ('247', '31', '6');
INSERT INTO `auth_permission_action` VALUES ('248', '31', '7');
INSERT INTO `auth_permission_action` VALUES ('249', '31', '8');
INSERT INTO `auth_permission_action` VALUES ('250', '31', '9');
INSERT INTO `auth_permission_action` VALUES ('251', '31', '10');
INSERT INTO `auth_permission_action` VALUES ('252', '31', '11');
INSERT INTO `auth_permission_action` VALUES ('253', '31', '12');
INSERT INTO `auth_permission_action` VALUES ('254', '31', '13');
INSERT INTO `auth_permission_action` VALUES ('291', '22', '31');
INSERT INTO `auth_permission_action` VALUES ('292', '22', '10');
INSERT INTO `auth_permission_action` VALUES ('293', '22', '11');
INSERT INTO `auth_permission_action` VALUES ('294', '22', '12');
INSERT INTO `auth_permission_action` VALUES ('295', '22', '13');
INSERT INTO `auth_permission_action` VALUES ('296', '22', '14');
INSERT INTO `auth_permission_action` VALUES ('297', '22', '15');
INSERT INTO `auth_permission_action` VALUES ('298', '22', '16');
INSERT INTO `auth_permission_action` VALUES ('299', '22', '17');
INSERT INTO `auth_permission_action` VALUES ('300', '22', '18');
INSERT INTO `auth_permission_action` VALUES ('301', '22', '19');
INSERT INTO `auth_permission_action` VALUES ('302', '22', '20');
INSERT INTO `auth_permission_action` VALUES ('303', '22', '21');
INSERT INTO `auth_permission_action` VALUES ('304', '22', '22');
INSERT INTO `auth_permission_action` VALUES ('305', '22', '23');
INSERT INTO `auth_permission_action` VALUES ('306', '22', '24');
INSERT INTO `auth_permission_action` VALUES ('307', '22', '25');
INSERT INTO `auth_permission_action` VALUES ('308', '22', '26');
INSERT INTO `auth_permission_action` VALUES ('309', '22', '27');
INSERT INTO `auth_permission_action` VALUES ('310', '22', '28');
INSERT INTO `auth_permission_action` VALUES ('311', '22', '29');
INSERT INTO `auth_permission_action` VALUES ('312', '22', '30');
INSERT INTO `auth_permission_action` VALUES ('313', '22', '34');
INSERT INTO `auth_permission_action` VALUES ('314', '22', '35');
INSERT INTO `auth_permission_action` VALUES ('315', '22', '36');
INSERT INTO `auth_permission_action` VALUES ('316', '22', '37');
INSERT INTO `auth_permission_action` VALUES ('317', '22', '38');
INSERT INTO `auth_permission_action` VALUES ('318', '22', '39');
INSERT INTO `auth_permission_action` VALUES ('409', '37', '48');
INSERT INTO `auth_permission_action` VALUES ('410', '37', '49');
INSERT INTO `auth_permission_action` VALUES ('461', '38', '60');
INSERT INTO `auth_permission_action` VALUES ('462', '38', '61');
INSERT INTO `auth_permission_action` VALUES ('463', '38', '62');
INSERT INTO `auth_permission_action` VALUES ('464', '38', '63');
INSERT INTO `auth_permission_action` VALUES ('465', '38', '64');
INSERT INTO `auth_permission_action` VALUES ('466', '38', '65');
INSERT INTO `auth_permission_action` VALUES ('467', '38', '50');
INSERT INTO `auth_permission_action` VALUES ('468', '38', '51');
INSERT INTO `auth_permission_action` VALUES ('469', '38', '52');
INSERT INTO `auth_permission_action` VALUES ('470', '38', '53');
INSERT INTO `auth_permission_action` VALUES ('471', '38', '54');
INSERT INTO `auth_permission_action` VALUES ('472', '39', '66');
INSERT INTO `auth_permission_action` VALUES ('473', '39', '67');
INSERT INTO `auth_permission_action` VALUES ('474', '39', '68');
INSERT INTO `auth_permission_action` VALUES ('475', '39', '69');
INSERT INTO `auth_permission_action` VALUES ('476', '39', '70');
INSERT INTO `auth_permission_action` VALUES ('477', '39', '71');
INSERT INTO `auth_permission_action` VALUES ('478', '39', '55');
INSERT INTO `auth_permission_action` VALUES ('479', '39', '56');
INSERT INTO `auth_permission_action` VALUES ('480', '39', '57');
INSERT INTO `auth_permission_action` VALUES ('481', '39', '58');
INSERT INTO `auth_permission_action` VALUES ('482', '39', '59');
INSERT INTO `auth_permission_action` VALUES ('483', '40', '72');
INSERT INTO `auth_permission_action` VALUES ('484', '40', '73');
INSERT INTO `auth_permission_action` VALUES ('485', '40', '74');
INSERT INTO `auth_permission_action` VALUES ('486', '40', '75');
INSERT INTO `auth_permission_action` VALUES ('487', '40', '76');
INSERT INTO `auth_permission_action` VALUES ('488', '41', '77');
INSERT INTO `auth_permission_action` VALUES ('494', '23', '82');
INSERT INTO `auth_permission_action` VALUES ('495', '23', '6');
INSERT INTO `auth_permission_action` VALUES ('496', '23', '8');
INSERT INTO `auth_permission_action` VALUES ('497', '23', '9');
INSERT INTO `auth_permission_action` VALUES ('498', '23', '46');
INSERT INTO `auth_permission_action` VALUES ('499', '23', '47');
INSERT INTO `auth_permission_action` VALUES ('500', '44', '83');
INSERT INTO `auth_permission_action` VALUES ('501', '44', '84');
INSERT INTO `auth_permission_action` VALUES ('502', '44', '85');
INSERT INTO `auth_permission_action` VALUES ('503', '44', '86');
INSERT INTO `auth_permission_action` VALUES ('504', '44', '87');
INSERT INTO `auth_permission_action` VALUES ('509', '42', '78');
INSERT INTO `auth_permission_action` VALUES ('510', '42', '79');
INSERT INTO `auth_permission_action` VALUES ('511', '42', '80');
INSERT INTO `auth_permission_action` VALUES ('512', '24', '4');
INSERT INTO `auth_permission_action` VALUES ('513', '24', '5');
INSERT INTO `auth_permission_action` VALUES ('514', '24', '88');

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_角色表';

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES ('1', '超级管理员');
INSERT INTO `auth_role` VALUES ('87', '游客');
INSERT INTO `auth_role` VALUES ('94', '企业站管理员');

-- ----------------------------
-- Table structure for auth_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_permission`;
CREATE TABLE `auth_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `permission_id` int(11) NOT NULL COMMENT '权限id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限管理_角色权限关联表';

-- ----------------------------
-- Records of auth_role_permission
-- ----------------------------
INSERT INTO `auth_role_permission` VALUES ('149', '1', '40');
INSERT INTO `auth_role_permission` VALUES ('150', '1', '41');
INSERT INTO `auth_role_permission` VALUES ('151', '1', '42');
INSERT INTO `auth_role_permission` VALUES ('152', '1', '22');
INSERT INTO `auth_role_permission` VALUES ('153', '1', '23');
INSERT INTO `auth_role_permission` VALUES ('154', '1', '24');
INSERT INTO `auth_role_permission` VALUES ('155', '1', '35');
INSERT INTO `auth_role_permission` VALUES ('156', '1', '37');
INSERT INTO `auth_role_permission` VALUES ('157', '1', '38');
INSERT INTO `auth_role_permission` VALUES ('158', '1', '39');
INSERT INTO `auth_role_permission` VALUES ('159', '94', '44');
INSERT INTO `auth_role_permission` VALUES ('160', '94', '37');
INSERT INTO `auth_role_permission` VALUES ('161', '94', '38');
INSERT INTO `auth_role_permission` VALUES ('162', '94', '39');
INSERT INTO `auth_role_permission` VALUES ('163', '94', '40');
INSERT INTO `auth_role_permission` VALUES ('164', '94', '41');
INSERT INTO `auth_role_permission` VALUES ('165', '94', '42');
INSERT INTO `auth_role_permission` VALUES ('169', '87', '23');
INSERT INTO `auth_role_permission` VALUES ('170', '87', '24');

-- ----------------------------
-- Table structure for company_banner
-- ----------------------------
DROP TABLE IF EXISTS `company_banner`;
CREATE TABLE `company_banner` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '轮播图名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '跳转地址 为空不跳转',
  `is_show` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否展示 0否 1是',
  `sort` tinyint(4) NOT NULL DEFAULT '1' COMMENT '排序值 1-100 从小到大',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='首页轮播图标';

-- ----------------------------
-- Records of company_banner
-- ----------------------------
INSERT INTO `company_banner` VALUES ('4', '121', 'http://localhost:8080/static/uploadFile/banner/2019-09-09/ZPXVYFXCXJ_zhouqi1.png', '0', '11', '2019-09-09 23:01:06', '2019-09-09 23:00:40', null);
INSERT INTO `company_banner` VALUES ('5', '测试', 'http://localhost:8080/static/uploadFile/banner/2019-09-09/AIXEIBWOZF_微信图片_20190905212641.jpg', '1', '2', '2019-09-09 23:42:37', '2019-09-09 23:42:37', null);

-- ----------------------------
-- Table structure for company_documentation
-- ----------------------------
DROP TABLE IF EXISTS `company_documentation`;
CREATE TABLE `company_documentation` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `company_documentation_category_id` int(11) NOT NULL COMMENT '分类id',
  `name` varchar(255) NOT NULL COMMENT '资料名称',
  `url` varchar(255) NOT NULL COMMENT '资料url地址',
  `description` text COMMENT '描述',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='资料表';

-- ----------------------------
-- Records of company_documentation
-- ----------------------------
INSERT INTO `company_documentation` VALUES ('6', '2', '121', 'http://localhost:8080/static/uploadFile/documentation/2019-09-09/JCJMOYLXCJ_83cd31e65aded69b6408308f9beda33.jpg', '121', '2019-09-09 22:58:35', '2019-09-09 22:58:35', null);
INSERT INTO `company_documentation` VALUES ('7', '3', '1211111', 'http://localhost:8080/static/uploadFile/documentation/2019-09-09/HCBQSRSOYH_83cd31e65aded69b6408308f9beda33.jpg', '121', '2019-09-09 23:00:08', '2019-09-09 23:00:08', null);

-- ----------------------------
-- Table structure for company_documentation_category
-- ----------------------------
DROP TABLE IF EXISTS `company_documentation_category`;
CREATE TABLE `company_documentation_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '分类名称',
  `sort` tinyint(4) NOT NULL DEFAULT '1' COMMENT '排序值 1-100 从小到大',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='资料分类表';

-- ----------------------------
-- Records of company_documentation_category
-- ----------------------------
INSERT INTO `company_documentation_category` VALUES ('1', '外语资料', '1', '2019-09-05 22:31:08', '2019-09-05 22:18:45', null);
INSERT INTO `company_documentation_category` VALUES ('2', '高中资料', '2', '2019-09-05 22:18:59', '2019-09-05 22:19:02', null);
INSERT INTO `company_documentation_category` VALUES ('3', '初中资料', '3', '2019-09-05 22:19:27', '2019-09-05 22:19:32', null);
INSERT INTO `company_documentation_category` VALUES ('6', '1', '12', '2019-09-09 22:58:14', '2019-09-09 22:57:52', null);

-- ----------------------------
-- Table structure for company_homepage
-- ----------------------------
DROP TABLE IF EXISTS `company_homepage`;
CREATE TABLE `company_homepage` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL,
  `is_show` tinyint(4) DEFAULT '1' COMMENT '是否展示 1展示 0关闭',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='首页模块显示控制表';

-- ----------------------------
-- Records of company_homepage
-- ----------------------------
INSERT INTO `company_homepage` VALUES ('1', '轮播图', '1');
INSERT INTO `company_homepage` VALUES ('2', '视频列表', '0');
INSERT INTO `company_homepage` VALUES ('3', '优秀视频', '1');
INSERT INTO `company_homepage` VALUES ('4', '亮点介绍', '1');

-- ----------------------------
-- Table structure for company_message
-- ----------------------------
DROP TABLE IF EXISTS `company_message`;
CREATE TABLE `company_message` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `mobile` varchar(20) NOT NULL DEFAULT '',
  `content` text,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of company_message
-- ----------------------------
INSERT INTO `company_message` VALUES ('1', '1', '1', '1', '2019-09-09 20:18:50', '2019-09-09 20:18:53', null);

-- ----------------------------
-- Table structure for company_navigator
-- ----------------------------
DROP TABLE IF EXISTS `company_navigator`;
CREATE TABLE `company_navigator` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级导航条id',
  `name` varchar(20) NOT NULL COMMENT '名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '跳转地址',
  `sort` tinyint(4) NOT NULL DEFAULT '1',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='导航条表';

-- ----------------------------
-- Records of company_navigator
-- ----------------------------
INSERT INTO `company_navigator` VALUES ('1', '0', '导航条1', 'http://www.baidu.com', '1', '2019-09-10 23:39:53', '2019-09-10 23:39:55', null);
INSERT INTO `company_navigator` VALUES ('2', '0', '导航条2', '', '2', '2019-09-10 23:40:16', '2019-09-10 23:40:19', null);
INSERT INTO `company_navigator` VALUES ('3', '0', '导航条3', '', '3', '2019-09-10 23:40:38', '2019-09-10 23:40:41', null);
INSERT INTO `company_navigator` VALUES ('4', '1', '子导航1', '', '1', '2019-09-10 23:41:12', '2019-09-10 23:41:14', null);
INSERT INTO `company_navigator` VALUES ('5', '1', '子导航11', '', '1', '2019-09-10 23:41:30', '2019-09-10 23:41:35', null);
INSERT INTO `company_navigator` VALUES ('8', '2', '子导航2', '', '1', '2019-09-10 23:42:53', '2019-09-10 23:42:56', null);

-- ----------------------------
-- Table structure for company_user
-- ----------------------------
DROP TABLE IF EXISTS `company_user`;
CREATE TABLE `company_user` (
  `id` int(11) NOT NULL,
  `union_id` varchar(255) NOT NULL DEFAULT '',
  `weixin_open_id` varchar(255) NOT NULL DEFAULT '',
  `username` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(60) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `sex` tinyint(4) NOT NULL DEFAULT '1' COMMENT '性别 1男 0女',
  `province` varchar(255) NOT NULL DEFAULT '' COMMENT '省份',
  `city` varchar(255) NOT NULL DEFAULT '' COMMENT '城市',
  `headimgurl` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `mobile` varchar(60) NOT NULL DEFAULT '' COMMENT '电话号码',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `login_times` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `weixin_open_id_index` (`weixin_open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='企业站用户表';

-- ----------------------------
-- Records of company_user
-- ----------------------------
INSERT INTO `company_user` VALUES ('1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '2019-09-09 23:42:18', '2019-09-09 23:41:53', null);

-- ----------------------------
-- Table structure for company_video
-- ----------------------------
DROP TABLE IF EXISTS `company_video`;
CREATE TABLE `company_video` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `company_video_category_id` int(11) NOT NULL COMMENT '视频所属分类',
  `title` varchar(255) NOT NULL COMMENT '视频标题',
  `url` varchar(255) NOT NULL COMMENT '视频播放地址',
  `view_times` int(11) NOT NULL DEFAULT '0' COMMENT '视频观看次数',
  `is_show_home_page` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否展示在首页推荐 0否 1是',
  `is_on_sale` tinyint(4) NOT NULL DEFAULT '0' COMMENT '视频是否上架 0否 1是',
  `description` text COMMENT '视频描述',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='企业站视频表';

-- ----------------------------
-- Records of company_video
-- ----------------------------
INSERT INTO `company_video` VALUES ('1', '1', '外语课程视频', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '1', '视频描述12312', '2019-09-06 22:29:08', '2019-09-05 22:44:24', null);
INSERT INTO `company_video` VALUES ('2', '3', '外语课程1', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '0', '视频描述', '2019-09-05 22:44:44', '2019-09-05 22:44:46', null);
INSERT INTO `company_video` VALUES ('3', '2', '高中课程', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '1', '1', '视频描述', '2019-09-05 22:45:11', '2019-09-05 22:45:13', null);
INSERT INTO `company_video` VALUES ('7', '1', '123', 'http://localhost:8080/static/uploadFile/video/2019-09-09/KPYSOZUCZM_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '0', '121', '2019-09-09 22:21:08', '2019-09-09 22:21:08', null);
INSERT INTO `company_video` VALUES ('8', '1', '231', 'http://localhost:8080/static/uploadFile/video/2019-09-09/AFFBBWGWHL_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '0', '121', '2019-09-09 22:53:28', '2019-09-09 22:22:48', null);
INSERT INTO `company_video` VALUES ('9', '1', '12312321', 'http://localhost:8080/static/uploadFile/video/2019-09-09/TJZMNYBYAC_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '0', '1232', '2019-09-09 22:53:47', '2019-09-09 22:53:47', null);

-- ----------------------------
-- Table structure for company_video_category
-- ----------------------------
DROP TABLE IF EXISTS `company_video_category`;
CREATE TABLE `company_video_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '课程分类名称',
  `sort` tinyint(4) NOT NULL DEFAULT '1' COMMENT '排序值 1-100 从小到大',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='视频分类名称';

-- ----------------------------
-- Records of company_video_category
-- ----------------------------
INSERT INTO `company_video_category` VALUES ('1', '外语课程', '12', '2019-09-05 22:05:01', '2019-09-05 22:05:05', null);
INSERT INTO `company_video_category` VALUES ('2', '高中课程', '2', '2019-09-05 22:05:20', '2019-09-05 22:05:22', null);
INSERT INTO `company_video_category` VALUES ('3', '其他课程', '3', '2019-09-05 22:05:34', '2019-09-05 22:05:36', null);

-- ----------------------------
-- Table structure for demo_excel
-- ----------------------------
DROP TABLE IF EXISTS `demo_excel`;
CREATE TABLE `demo_excel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '姓名',
  `age` tinyint(4) NOT NULL COMMENT '年龄',
  `birthday` int(11) NOT NULL COMMENT '出生日期(时间戳)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='导入导出excel数据表';

-- ----------------------------
-- Records of demo_excel
-- ----------------------------
