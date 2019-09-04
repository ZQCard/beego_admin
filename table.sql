/*
Navicat MySQL Data Transfer

Source Server         : 47.97.215.189
Source Server Version : 50727
Source Host           : 47.97.215.189:3306
Source Database       : edu_study

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2019-09-04 21:12:49
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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作表';

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

-- ----------------------------
-- Table structure for `auth_administrator_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_administrator_role`;
CREATE TABLE `auth_administrator_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

-- ----------------------------
-- Records of auth_administrator_role
-- ----------------------------
INSERT INTO `auth_administrator_role` VALUES ('15', '1', '1');
INSERT INTO `auth_administrator_role` VALUES ('16', '22', '2');

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
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('14', '0', '设置', '1', '');
INSERT INTO `auth_menu` VALUES ('15', '14', '权限管理', '1', '');
INSERT INTO `auth_menu` VALUES ('18', '15', '角色列表', '2', '/auth/role');
INSERT INTO `auth_menu` VALUES ('19', '15', '权限列表', '3', '/auth/permission');
INSERT INTO `auth_menu` VALUES ('20', '15', '行为列表', '4', '/auth/action');
INSERT INTO `auth_menu` VALUES ('21', '15', '菜单列表', '5', '/auth/menu');
INSERT INTO `auth_menu` VALUES ('23', '15', '管理员列表', '8', '/auth/administrator');
INSERT INTO `auth_menu` VALUES ('26', '0', '工具箱', '0', '');
INSERT INTO `auth_menu` VALUES ('27', '26', '常规功能', '0', '');
INSERT INTO `auth_menu` VALUES ('28', '27', '文件上传', '0', '/tools/uploadFile');

-- ----------------------------
-- Table structure for `auth_permission`
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission`;
CREATE TABLE `auth_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_权限表';

-- ----------------------------
-- Records of auth_permission
-- ----------------------------
INSERT INTO `auth_permission` VALUES ('22', '权限管理');
INSERT INTO `auth_permission` VALUES ('23', '公共权限');
INSERT INTO `auth_permission` VALUES ('24', '游客权限');
INSERT INTO `auth_permission` VALUES ('35', '开发权限');

-- ----------------------------
-- Table structure for `auth_permission_action`
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission_action`;
CREATE TABLE `auth_permission_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=409 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作权限关联表';

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
INSERT INTO `auth_permission_action` VALUES ('368', '35', '45');
INSERT INTO `auth_permission_action` VALUES ('369', '35', '4');
INSERT INTO `auth_permission_action` VALUES ('370', '35', '5');
INSERT INTO `auth_permission_action` VALUES ('371', '35', '6');
INSERT INTO `auth_permission_action` VALUES ('372', '35', '7');
INSERT INTO `auth_permission_action` VALUES ('373', '35', '8');
INSERT INTO `auth_permission_action` VALUES ('374', '35', '9');
INSERT INTO `auth_permission_action` VALUES ('375', '35', '10');
INSERT INTO `auth_permission_action` VALUES ('376', '35', '11');
INSERT INTO `auth_permission_action` VALUES ('377', '35', '12');
INSERT INTO `auth_permission_action` VALUES ('378', '35', '13');
INSERT INTO `auth_permission_action` VALUES ('379', '35', '14');
INSERT INTO `auth_permission_action` VALUES ('380', '35', '15');
INSERT INTO `auth_permission_action` VALUES ('381', '35', '16');
INSERT INTO `auth_permission_action` VALUES ('382', '35', '17');
INSERT INTO `auth_permission_action` VALUES ('383', '35', '18');
INSERT INTO `auth_permission_action` VALUES ('384', '35', '19');
INSERT INTO `auth_permission_action` VALUES ('385', '35', '20');
INSERT INTO `auth_permission_action` VALUES ('386', '35', '21');
INSERT INTO `auth_permission_action` VALUES ('387', '35', '22');
INSERT INTO `auth_permission_action` VALUES ('388', '35', '23');
INSERT INTO `auth_permission_action` VALUES ('389', '35', '24');
INSERT INTO `auth_permission_action` VALUES ('390', '35', '25');
INSERT INTO `auth_permission_action` VALUES ('391', '35', '26');
INSERT INTO `auth_permission_action` VALUES ('392', '35', '27');
INSERT INTO `auth_permission_action` VALUES ('393', '35', '28');
INSERT INTO `auth_permission_action` VALUES ('394', '35', '29');
INSERT INTO `auth_permission_action` VALUES ('395', '35', '30');
INSERT INTO `auth_permission_action` VALUES ('396', '35', '31');
INSERT INTO `auth_permission_action` VALUES ('397', '35', '33');
INSERT INTO `auth_permission_action` VALUES ('398', '35', '34');
INSERT INTO `auth_permission_action` VALUES ('399', '35', '35');
INSERT INTO `auth_permission_action` VALUES ('400', '35', '36');
INSERT INTO `auth_permission_action` VALUES ('401', '35', '37');
INSERT INTO `auth_permission_action` VALUES ('402', '35', '38');
INSERT INTO `auth_permission_action` VALUES ('403', '35', '39');
INSERT INTO `auth_permission_action` VALUES ('404', '23', '47');
INSERT INTO `auth_permission_action` VALUES ('405', '23', '6');
INSERT INTO `auth_permission_action` VALUES ('406', '23', '8');
INSERT INTO `auth_permission_action` VALUES ('407', '23', '9');
INSERT INTO `auth_permission_action` VALUES ('408', '23', '46');

-- ----------------------------
-- Table structure for `auth_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色表';

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES ('1', '超级管理员');
INSERT INTO `auth_role` VALUES ('2', '普通用户');
INSERT INTO `auth_role` VALUES ('87', '游客');

-- ----------------------------
-- Table structure for `auth_role_permission`
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_permission`;
CREATE TABLE `auth_role_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `permission_id` int(11) NOT NULL COMMENT '权限id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色权限关联表';

-- ----------------------------
-- Records of auth_role_permission
-- ----------------------------
INSERT INTO `auth_role_permission` VALUES ('8', '87', '24');
INSERT INTO `auth_role_permission` VALUES ('9', '2', '23');
INSERT INTO `auth_role_permission` VALUES ('40', '1', '35');
INSERT INTO `auth_role_permission` VALUES ('41', '1', '22');
INSERT INTO `auth_role_permission` VALUES ('42', '1', '23');
INSERT INTO `auth_role_permission` VALUES ('43', '1', '24');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资料分类表';

-- ----------------------------
-- Records of company_documentation_category
-- ----------------------------

-- ----------------------------
-- Table structure for `company_documention`
-- ----------------------------
DROP TABLE IF EXISTS `company_documention`;
CREATE TABLE `company_documention` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '资料名称',
  `url` varchar(255) NOT NULL COMMENT '资料url地址',
  `description` text COMMENT '描述',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='资料表';

-- ----------------------------
-- Records of company_documention
-- ----------------------------

-- ----------------------------
-- Table structure for `company_home_page_module_controller`
-- ----------------------------
DROP TABLE IF EXISTS `company_home_page_module_controller`;
CREATE TABLE `company_home_page_module_controller` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `banner_is_show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '轮播图模块是否展示 0否1是',
  `slogan_is_show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '标语模块是否展示 0否1是',
  `video_is_show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '视频模块是否展示 0否1是',
  `advantage_is_show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '亮点模块是否展示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='首页模块显示控制表';

-- ----------------------------
-- Records of company_home_page_module_controller
-- ----------------------------

-- ----------------------------
-- Table structure for `company_navigator`
-- ----------------------------
DROP TABLE IF EXISTS `company_navigator`;
CREATE TABLE `company_navigator` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级导航条id',
  `name` varchar(20) NOT NULL COMMENT '名称',
  `url` varchar(255) NOT NULL COMMENT '跳转地址',
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
  `open_id` int(11) DEFAULT NULL,
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='企业站视频表';

-- ----------------------------
-- Records of company_video
-- ----------------------------

-- ----------------------------
-- Table structure for `company_video_category`
-- ----------------------------
DROP TABLE IF EXISTS `company_video_category`;
CREATE TABLE `company_video_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` int(11) NOT NULL COMMENT '课程分类名称',
  `sort` tinyint(4) NOT NULL DEFAULT '1' COMMENT '排序值 1-100 从小到大',
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='视频分类名称';

-- ----------------------------
-- Records of company_video_category
-- ----------------------------

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
