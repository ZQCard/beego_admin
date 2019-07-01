/*
Navicat MySQL Data Transfer

Source Server         : 47.97.215.189
Source Server Version : 50725
Source Host           : 47.97.215.189:3306
Source Database       : beego_admin_template

Target Server Type    : MYSQL
Target Server Version : 50725
File Encoding         : 65001

Date: 2019-07-01 22:06:30
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
  `deleted_at` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

-- ----------------------------
-- Records of administrator
-- ----------------------------
INSERT INTO `administrator` VALUES ('1', 'zhouqi', '96e79218965eb72c92a549dd5a330112', '周起', '445864742@qq.com', '2019-05-22 15:02:14', '2019-05-05 16:24:18', '0');
INSERT INTO `administrator` VALUES ('22', 'user1', '96e79218965eb72c92a549dd5a330112', '普通用户', '645144262@qq.com', '2019-06-12 16:06:05', '2019-06-12 16:06:05', '0');
INSERT INTO `administrator` VALUES ('23', 'zhouqi1', '96e79218965eb72c92a549dd5a330112', '111111', '1111@qq.com', '2019-06-12 16:09:47', '2019-06-12 16:09:47', '0');
INSERT INTO `administrator` VALUES ('24', '1', 'e10adc3949ba59abbe56e057f20f883e', '12', '111111@qq.com', '2019-06-19 14:44:01', '2019-06-19 14:43:51', '1560955853');

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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作表';

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

-- ----------------------------
-- Table structure for `auth_administrator_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_administrator_role`;
CREATE TABLE `auth_administrator_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `administrator_id` int(11) NOT NULL COMMENT '用户id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';

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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('14', '0', '设置', '1', '');
INSERT INTO `auth_menu` VALUES ('15', '14', '权限管理', '1', '');
INSERT INTO `auth_menu` VALUES ('16', '14', '我的设置', '2', '');
INSERT INTO `auth_menu` VALUES ('18', '15', '角色列表', '2', '/auth/role');
INSERT INTO `auth_menu` VALUES ('19', '15', '权限列表', '3', '/auth/permission');
INSERT INTO `auth_menu` VALUES ('20', '15', '行为列表', '4', '/auth/action');
INSERT INTO `auth_menu` VALUES ('21', '15', '菜单列表', '5', '/auth/menu');
INSERT INTO `auth_menu` VALUES ('22', '16', '基本信息', '1', '/administrator/adminInfo');
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
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_权限表';

-- ----------------------------
-- Records of auth_permission
-- ----------------------------
INSERT INTO `auth_permission` VALUES ('22', '权限管理');
INSERT INTO `auth_permission` VALUES ('23', '公共权限');
INSERT INTO `auth_permission` VALUES ('24', '游客权限');
INSERT INTO `auth_permission` VALUES ('26', '工具箱权限');

-- ----------------------------
-- Table structure for `auth_permission_action`
-- ----------------------------
DROP TABLE IF EXISTS `auth_permission_action`;
CREATE TABLE `auth_permission_action` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `permission_id` int(11) NOT NULL,
  `action_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=238 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_操作权限关联表';

-- ----------------------------
-- Records of auth_permission_action
-- ----------------------------
INSERT INTO `auth_permission_action` VALUES ('170', '24', '4');
INSERT INTO `auth_permission_action` VALUES ('171', '24', '5');
INSERT INTO `auth_permission_action` VALUES ('172', '23', '6');
INSERT INTO `auth_permission_action` VALUES ('173', '23', '8');
INSERT INTO `auth_permission_action` VALUES ('174', '23', '9');
INSERT INTO `auth_permission_action` VALUES ('211', '22', '14');
INSERT INTO `auth_permission_action` VALUES ('212', '22', '34');
INSERT INTO `auth_permission_action` VALUES ('213', '22', '35');
INSERT INTO `auth_permission_action` VALUES ('214', '22', '36');
INSERT INTO `auth_permission_action` VALUES ('215', '22', '37');
INSERT INTO `auth_permission_action` VALUES ('216', '22', '38');
INSERT INTO `auth_permission_action` VALUES ('217', '22', '39');
INSERT INTO `auth_permission_action` VALUES ('218', '22', '10');
INSERT INTO `auth_permission_action` VALUES ('219', '22', '11');
INSERT INTO `auth_permission_action` VALUES ('220', '22', '12');
INSERT INTO `auth_permission_action` VALUES ('221', '22', '13');
INSERT INTO `auth_permission_action` VALUES ('222', '22', '15');
INSERT INTO `auth_permission_action` VALUES ('223', '22', '16');
INSERT INTO `auth_permission_action` VALUES ('224', '22', '17');
INSERT INTO `auth_permission_action` VALUES ('225', '22', '18');
INSERT INTO `auth_permission_action` VALUES ('226', '22', '19');
INSERT INTO `auth_permission_action` VALUES ('227', '22', '20');
INSERT INTO `auth_permission_action` VALUES ('228', '22', '21');
INSERT INTO `auth_permission_action` VALUES ('229', '22', '22');
INSERT INTO `auth_permission_action` VALUES ('230', '22', '23');
INSERT INTO `auth_permission_action` VALUES ('231', '22', '24');
INSERT INTO `auth_permission_action` VALUES ('232', '22', '25');
INSERT INTO `auth_permission_action` VALUES ('233', '22', '26');
INSERT INTO `auth_permission_action` VALUES ('234', '22', '27');
INSERT INTO `auth_permission_action` VALUES ('235', '22', '28');
INSERT INTO `auth_permission_action` VALUES ('236', '22', '29');
INSERT INTO `auth_permission_action` VALUES ('237', '22', '30');

-- ----------------------------
-- Table structure for `auth_role`
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) NOT NULL COMMENT '角色名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色表';

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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COMMENT='权限管理_角色权限关联表';

-- ----------------------------
-- Records of auth_role_permission
-- ----------------------------
INSERT INTO `auth_role_permission` VALUES ('8', '87', '24');
INSERT INTO `auth_role_permission` VALUES ('9', '2', '23');
INSERT INTO `auth_role_permission` VALUES ('20', '1', '26');
INSERT INTO `auth_role_permission` VALUES ('21', '1', '22');
INSERT INTO `auth_role_permission` VALUES ('22', '1', '23');
INSERT INTO `auth_role_permission` VALUES ('23', '1', '24');
