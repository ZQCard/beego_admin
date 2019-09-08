/*
Navicat MySQL Data Transfer

Source Server         : 47.97.215.189
Source Server Version : 50727
Source Host           : 47.97.215.189:3306
Source Database       : edu_study

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2019-09-08 22:42:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `administrator`
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
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of administrator
-- ----------------------------
INSERT INTO `administrator` VALUES ('1', 'zhouqi', 'e10adc3949ba59abbe56e057f20f883e', '周起', '445864742@qq.com', '2019-09-03 00:25:37', '2019-05-05 16:24:18', null);
INSERT INTO `administrator` VALUES ('22', 'user111', '96e79218965eb72c92a549dd5a330112', '普通用户', '645144262@qq.com', '2019-06-12 16:06:05', '2019-06-12 16:06:05', '2019-08-29 22:07:31');
INSERT INTO `administrator` VALUES ('30', 'zhouqi1', 'e10adc3949ba59abbe56e057f20f883e', '22222', '1@qq.com', '2019-08-29 23:37:41', '2019-08-29 23:04:17', null);
INSERT INTO `administrator` VALUES ('31', 'zhouqi11', '111111', '1111', '1@qq.com', '2019-08-29 23:37:46', '2019-08-29 23:37:21', null);

-- ----------------------------
-- Table structure for `auth_action`
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '请求动作',
  `method` varchar(50) NOT NULL COMMENT '请求方法',
  `route` varchar(50) NOT NULL COMMENT '请求路由',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作表';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES ('4', '登录首页', 'GET', '/login');
INSERT INTO `auth_action` VALUES ('5', '登录操作', 'POST', '/login');
INSERT INTO `auth_action` VALUES ('6', '后台首页', 'GET', '/');
INSERT INTO `auth_action` VALUES ('7', '测试接口', 'GET', '/');
INSERT INTO `auth_action` VALUES ('8', '后台首页', 'GET', 'index');
INSERT INTO `auth_action` VALUES ('9', '退出', 'GET', '/logout');
INSERT INTO `auth_action` VALUES ('10', '管理员列表', 'GET', '/auth/administrator');
INSERT INTO `auth_action` VALUES ('11', '管理员添加', 'POST', '/auth/administrator');
INSERT INTO `auth_action` VALUES ('12', '管理员更新', 'PUT', '/auth/administrator');
INSERT INTO `auth_action` VALUES ('13', '管理员删除', 'DELETE', '/auth/administrator');
INSERT INTO `auth_action` VALUES ('14', '管理员恢复', 'PATCH', '/auth/administrator');
INSERT INTO `auth_action` VALUES ('15', '角色列表', 'GET', '/auth/role');
INSERT INTO `auth_action` VALUES ('16', '角色添加', 'POST', '/auth/role');
INSERT INTO `auth_action` VALUES ('17', '角色更新', 'PUT', '/auth/role');
INSERT INTO `auth_action` VALUES ('18', '角色删除', 'DELETE', '/auth/role');
INSERT INTO `auth_action` VALUES ('19', '权限列表', 'GET', '/auth/permission');
INSERT INTO `auth_action` VALUES ('20', '权限添加', 'POST', '/auth/permission');
INSERT INTO `auth_action` VALUES ('21', '权限更新', 'PUT', '/auth/permission');
INSERT INTO `auth_action` VALUES ('22', '权限删除', 'DELETE', '/auth/permission');
INSERT INTO `auth_action` VALUES ('23', '行为列表', 'GET', '/auth/action');
INSERT INTO `auth_action` VALUES ('24', '行为添加', 'POST', '/auth/action');
INSERT INTO `auth_action` VALUES ('25', '行为更新', 'PUT', '/auth/action');
INSERT INTO `auth_action` VALUES ('26', '行为删除', 'DELETE', '/auth/action');
INSERT INTO `auth_action` VALUES ('27', '菜单列表', 'GET', '/auth/menu');
INSERT INTO `auth_action` VALUES ('28', '菜单添加', 'POST', '/auth/menu');
INSERT INTO `auth_action` VALUES ('29', '菜单更新', 'PUT', '/auth/menu');
INSERT INTO `auth_action` VALUES ('30', '菜单删除', 'DELETE', '/auth/menu');
INSERT INTO `auth_action` VALUES ('31', '基本信息', 'GET', '/administrator/adminInfo');
INSERT INTO `auth_action` VALUES ('33', '上传文件', 'POST', '/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('34', '查看管理员角色', 'GET', '/auth/administrator/roles');
INSERT INTO `auth_action` VALUES ('35', '授予管理员角色', 'PUT', '/auth/administrator/roles');
INSERT INTO `auth_action` VALUES ('36', '查看角色权限', 'GET', '/auth/role/permissions');
INSERT INTO `auth_action` VALUES ('37', '授予权限角色', 'PUT', '/auth/role/permissions');
INSERT INTO `auth_action` VALUES ('38', '查看权限行为', 'GET', '/auth/permission/actions');
INSERT INTO `auth_action` VALUES ('39', '授予权限行为', 'PUT', '/auth/permission/actions');
INSERT INTO `auth_action` VALUES ('45', '文件上传页面', 'GET', '/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('46', '文件上传', 'POST', '/tools/uploadFile');
INSERT INTO `auth_action` VALUES ('47', '刷新权限', 'GET', '/auth/administrator/refreshAuth');
INSERT INTO `auth_action` VALUES ('48', '首页模块控制页面', 'GET', '/homepage/setting');
INSERT INTO `auth_action` VALUES ('49', '首页模块控制设置', 'PATCH', '/homepage/setting');
INSERT INTO `auth_action` VALUES ('50', '视频分类列表', 'GET', '/video/category');
INSERT INTO `auth_action` VALUES ('51', '视频分类创建', 'POST', '/video/category');
INSERT INTO `auth_action` VALUES ('52', '视频分类更新', 'PUT', '/video/category');
INSERT INTO `auth_action` VALUES ('53', '视频分类删除', 'DELETE', '/video/category');
INSERT INTO `auth_action` VALUES ('54', '视频分类恢复', 'PATCH', '/video/category');
INSERT INTO `auth_action` VALUES ('55', '资料分类列表', 'GET', '/documentation/category');
INSERT INTO `auth_action` VALUES ('56', '资料分类创建', 'POST', '/documentation/category');
INSERT INTO `auth_action` VALUES ('57', '资料分类更新', 'PUT', '/documentation/category');
INSERT INTO `auth_action` VALUES ('58', '资料分类删除', 'DELETE', '/documentation/category');
INSERT INTO `auth_action` VALUES ('59', '资料分类恢复', 'PATCH', '/documentation/category');
INSERT INTO `auth_action` VALUES ('60', '视频列表', 'GET', '/video');
INSERT INTO `auth_action` VALUES ('61', '视频详情页面', 'GET', '/video/info');
INSERT INTO `auth_action` VALUES ('62', '视频创建', 'POST', '/video');
INSERT INTO `auth_action` VALUES ('63', '视频编辑', 'PUT', '/video');
INSERT INTO `auth_action` VALUES ('64', '视频删除', 'DELETE', '/video');
INSERT INTO `auth_action` VALUES ('65', '视频恢复', 'PATCH', '/video');
INSERT INTO `auth_action` VALUES ('66', '资料列表', 'GET', '/documentation');
INSERT INTO `auth_action` VALUES ('67', '资料详情页面', 'GET', '/documentation/info');
INSERT INTO `auth_action` VALUES ('68', '资料创建', 'POST', '/documentation');
INSERT INTO `auth_action` VALUES ('69', '资料编辑', 'PUT', '/documentation');
INSERT INTO `auth_action` VALUES ('70', '编辑删除', 'DELETE', '/documentation');
INSERT INTO `auth_action` VALUES ('71', '资料恢复', 'PATCH', '/documentation');

-- ----------------------------
-- Table structure for `auth_administrator_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_administrator_role`;
CREATE TABLE `auth_administrator_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- ----------------------------
-- Records of auth_administrator_role
-- ----------------------------
INSERT INTO `auth_administrator_role` VALUES ('25', '31', '1');
INSERT INTO `auth_administrator_role` VALUES ('26', '31', '94');
INSERT INTO `auth_administrator_role` VALUES ('28', '30', '94');
INSERT INTO `auth_administrator_role` VALUES ('29', '1', '94');
INSERT INTO `auth_administrator_role` VALUES ('30', '1', '1');

-- ----------------------------
-- Table structure for `auth_menu`
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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('14', '0', '设置', '98', '');
INSERT INTO `auth_menu` VALUES ('15', '14', '权限管理', '1', '');
INSERT INTO `auth_menu` VALUES ('18', '15', '角色列表', '2', '/auth/role');
INSERT INTO `auth_menu` VALUES ('19', '15', '权限列表', '3', '/auth/permission');
INSERT INTO `auth_menu` VALUES ('20', '15', '行为列表', '4', '/auth/action');
INSERT INTO `auth_menu` VALUES ('21', '15', '菜单列表', '5', '/auth/menu');
INSERT INTO `auth_menu` VALUES ('23', '15', '管理员列表', '8', '/auth/administrator');
INSERT INTO `auth_menu` VALUES ('26', '0', '工具箱', '99', '');
INSERT INTO `auth_menu` VALUES ('27', '26', '常规功能', '0', '');
INSERT INTO `auth_menu` VALUES ('28', '27', '文件上传', '0', '/tools/uploadFile');
INSERT INTO `auth_menu` VALUES ('31', '0', '企业站管理', '1', '');
INSERT INTO `auth_menu` VALUES ('32', '31', '首页管理', '1', '');
INSERT INTO `auth_menu` VALUES ('33', '32', '首页模块控制', '1', '/homepage/setting');
INSERT INTO `auth_menu` VALUES ('34', '31', '视频管理', '2', '');
INSERT INTO `auth_menu` VALUES ('35', '34', '视频分类管理', '1', '/video/category');
INSERT INTO `auth_menu` VALUES ('36', '34', '视频详情管理', '2', '/video');
INSERT INTO `auth_menu` VALUES ('37', '31', '资料管理', '3', '');
INSERT INTO `auth_menu` VALUES ('38', '37', '资料分类管理', '1', '/documentation/category');
INSERT INTO `auth_menu` VALUES ('39', '37', '资料详情管理', '2', '/documentation');

-- ----------------------------
-- Table structure for `auth_permission`
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission`;
CREATE TABLE `auth_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_权限表';

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

-- ----------------------------
-- Table structure for `auth_permission_action`
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission_action`;
CREATE TABLE `auth_permission_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=483 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作权限关联表';

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
INSERT INTO `auth_permission_action` VALUES ('269', '24', '4');
INSERT INTO `auth_permission_action` VALUES ('270', '24', '5');
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
INSERT INTO `auth_permission_action` VALUES ('404', '23', '47');
INSERT INTO `auth_permission_action` VALUES ('405', '23', '6');
INSERT INTO `auth_permission_action` VALUES ('406', '23', '8');
INSERT INTO `auth_permission_action` VALUES ('407', '23', '9');
INSERT INTO `auth_permission_action` VALUES ('408', '23', '46');
INSERT INTO `auth_permission_action` VALUES ('409', '37', '48');
INSERT INTO `auth_permission_action` VALUES ('410', '37', '49');
INSERT INTO `auth_permission_action` VALUES ('411', '35', '46');
INSERT INTO `auth_permission_action` VALUES ('412', '35', '47');
INSERT INTO `auth_permission_action` VALUES ('413', '35', '48');
INSERT INTO `auth_permission_action` VALUES ('414', '35', '49');
INSERT INTO `auth_permission_action` VALUES ('415', '35', '4');
INSERT INTO `auth_permission_action` VALUES ('416', '35', '5');
INSERT INTO `auth_permission_action` VALUES ('417', '35', '6');
INSERT INTO `auth_permission_action` VALUES ('418', '35', '7');
INSERT INTO `auth_permission_action` VALUES ('419', '35', '8');
INSERT INTO `auth_permission_action` VALUES ('420', '35', '9');
INSERT INTO `auth_permission_action` VALUES ('421', '35', '10');
INSERT INTO `auth_permission_action` VALUES ('422', '35', '11');
INSERT INTO `auth_permission_action` VALUES ('423', '35', '12');
INSERT INTO `auth_permission_action` VALUES ('424', '35', '13');
INSERT INTO `auth_permission_action` VALUES ('425', '35', '14');
INSERT INTO `auth_permission_action` VALUES ('426', '35', '15');
INSERT INTO `auth_permission_action` VALUES ('427', '35', '16');
INSERT INTO `auth_permission_action` VALUES ('428', '35', '17');
INSERT INTO `auth_permission_action` VALUES ('429', '35', '18');
INSERT INTO `auth_permission_action` VALUES ('430', '35', '19');
INSERT INTO `auth_permission_action` VALUES ('431', '35', '20');
INSERT INTO `auth_permission_action` VALUES ('432', '35', '21');
INSERT INTO `auth_permission_action` VALUES ('433', '35', '22');
INSERT INTO `auth_permission_action` VALUES ('434', '35', '23');
INSERT INTO `auth_permission_action` VALUES ('435', '35', '24');
INSERT INTO `auth_permission_action` VALUES ('436', '35', '25');
INSERT INTO `auth_permission_action` VALUES ('437', '35', '26');
INSERT INTO `auth_permission_action` VALUES ('438', '35', '27');
INSERT INTO `auth_permission_action` VALUES ('439', '35', '28');
INSERT INTO `auth_permission_action` VALUES ('440', '35', '29');
INSERT INTO `auth_permission_action` VALUES ('441', '35', '30');
INSERT INTO `auth_permission_action` VALUES ('442', '35', '31');
INSERT INTO `auth_permission_action` VALUES ('443', '35', '33');
INSERT INTO `auth_permission_action` VALUES ('444', '35', '34');
INSERT INTO `auth_permission_action` VALUES ('445', '35', '35');
INSERT INTO `auth_permission_action` VALUES ('446', '35', '36');
INSERT INTO `auth_permission_action` VALUES ('447', '35', '37');
INSERT INTO `auth_permission_action` VALUES ('448', '35', '38');
INSERT INTO `auth_permission_action` VALUES ('449', '35', '39');
INSERT INTO `auth_permission_action` VALUES ('450', '35', '45');
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

-- ----------------------------
-- Table structure for `auth_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色表';

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES ('1', '超级管理员');
INSERT INTO `auth_role` VALUES ('87', '游客');
INSERT INTO `auth_role` VALUES ('94', '企业站管理员');

-- ----------------------------
-- Table structure for `auth_role_permission`
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_permission`;
CREATE TABLE `auth_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `permission_id` int(11) NOT NULL COMMENT '权限id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=124 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色权限关联表';

-- ----------------------------
-- Records of auth_role_permission
-- ----------------------------
INSERT INTO `auth_role_permission` VALUES ('46', '87', '23');
INSERT INTO `auth_role_permission` VALUES ('47', '87', '24');
INSERT INTO `auth_role_permission` VALUES ('114', '94', '39');
INSERT INTO `auth_role_permission` VALUES ('115', '94', '37');
INSERT INTO `auth_role_permission` VALUES ('116', '94', '38');
INSERT INTO `auth_role_permission` VALUES ('117', '1', '39');
INSERT INTO `auth_role_permission` VALUES ('118', '1', '22');
INSERT INTO `auth_role_permission` VALUES ('119', '1', '23');
INSERT INTO `auth_role_permission` VALUES ('120', '1', '24');
INSERT INTO `auth_role_permission` VALUES ('121', '1', '35');
INSERT INTO `auth_role_permission` VALUES ('122', '1', '37');
INSERT INTO `auth_role_permission` VALUES ('123', '1', '38');

-- ----------------------------
-- Table structure for `company_banner`
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='首页轮播图标';

-- ----------------------------
-- Records of company_banner
-- ----------------------------

-- ----------------------------
-- Table structure for `company_documentation`
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='资料表';

-- ----------------------------
-- Records of company_documentation
-- ----------------------------

-- ----------------------------
-- Table structure for `company_documentation_category`
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='资料分类表';

-- ----------------------------
-- Records of company_documentation_category
-- ----------------------------
INSERT INTO `company_documentation_category` VALUES ('1', '外语资料', '1', '2019-09-05 22:31:08', '2019-09-05 22:18:45', null);
INSERT INTO `company_documentation_category` VALUES ('2', '高中资料', '2', '2019-09-05 22:18:59', '2019-09-05 22:19:02', null);
INSERT INTO `company_documentation_category` VALUES ('3', '初中资料', '3', '2019-09-05 22:19:27', '2019-09-05 22:19:32', null);

-- ----------------------------
-- Table structure for `company_homepage`
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
INSERT INTO `company_homepage` VALUES ('1', '轮播图', '0');
INSERT INTO `company_homepage` VALUES ('2', '视频列表', '1');
INSERT INTO `company_homepage` VALUES ('3', '优秀视频', '1');
INSERT INTO `company_homepage` VALUES ('4', '亮点介绍', '1');

-- ----------------------------
-- Table structure for `company_navigator`
-- ----------------------------
DROP TABLE IF EXISTS `company_navigator`;
CREATE TABLE `company_navigator` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级导航条id',
  `name` varchar(20) NOT NULL COMMENT '名称',
  `url` varchar(255) NOT NULL COMMENT '跳转地址',
  `sort` tinyint(4) NOT NULL DEFAULT '1',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='导航条表';

-- ----------------------------
-- Records of company_navigator
-- ----------------------------

-- ----------------------------
-- Table structure for `company_user`
-- ----------------------------
DROP TABLE IF EXISTS `company_user`;
CREATE TABLE `company_user` (
  `id` int(11) NOT NULL,
  `open_id` varchar(255) DEFAULT '',
  `username` varchar(60) NOT NULL COMMENT '用户名',
  `password` varchar(60) NOT NULL,
  `mobile` varchar(60) NOT NULL DEFAULT '' COMMENT '电话号码',
  `email` varchar(60) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `login_times` int(11) NOT NULL DEFAULT '0' COMMENT '登录次数',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='企业站用户表';

-- ----------------------------
-- Records of company_user
-- ----------------------------

-- ----------------------------
-- Table structure for `company_video`
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='企业站视频表';

-- ----------------------------
-- Records of company_video
-- ----------------------------
INSERT INTO `company_video` VALUES ('1', '1', '外语课程视频', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '1', '视频描述12312', '2019-09-06 22:29:08', '2019-09-05 22:44:24', null);
INSERT INTO `company_video` VALUES ('2', '3', '外语课程1', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '0', '0', '视频描述', '2019-09-05 22:44:44', '2019-09-05 22:44:46', null);
INSERT INTO `company_video` VALUES ('3', '2', '高中课程', 'http://localhost:8080/static/uploadFile/video/2019-09-08/HMYECDQVOX_83cd31e65aded69b6408308f9beda33.jpg', '0', '1', '1', '视频描述', '2019-09-05 22:45:11', '2019-09-05 22:45:13', null);

-- ----------------------------
-- Table structure for `company_video_category`
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='视频分类名称';

-- ----------------------------
-- Records of company_video_category
-- ----------------------------
INSERT INTO `company_video_category` VALUES ('1', '外语课程', '1', '2019-09-05 22:05:01', '2019-09-05 22:05:05', null);
INSERT INTO `company_video_category` VALUES ('2', '高中课程', '2', '2019-09-05 22:05:20', '2019-09-05 22:05:22', null);
INSERT INTO `company_video_category` VALUES ('3', '其他课程', '3', '2019-09-05 22:05:34', '2019-09-05 22:05:36', null);

-- ----------------------------
-- Table structure for `demo_excel`
-- ----------------------------
DROP TABLE IF EXISTS `demo_excel`;
CREATE TABLE `demo_excel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '姓名',
  `age` tinyint(4) NOT NULL COMMENT '年龄',
  `birthday` int(11) NOT NULL COMMENT '出生日期(时间戳)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='导入导出excel数据表';

-- ----------------------------
-- Records of demo_excel
-- ----------------------------
