/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : 127.0.0.1:3306
 Source Schema         : go_admin

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 02/08/2021 10:54:58
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_log_data
-- ----------------------------
DROP TABLE IF EXISTS `admin_log_data`;
CREATE TABLE `admin_log_data`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `admin_log_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '日志ID',
    `data`         text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '日志内容',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台操作日志数据';

-- ----------------------------
-- Records of admin_log_data
-- ----------------------------
BEGIN;
INSERT INTO `admin_log_data`
VALUES (1, 1,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
INSERT INTO `admin_log_data`
VALUES (2, 2,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
INSERT INTO `admin_log_data`
VALUES (3, 3,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
INSERT INTO `admin_log_data`
VALUES (4, 4,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
INSERT INTO `admin_log_data`
VALUES (5, 5,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
INSERT INTO `admin_log_data`
VALUES (6, 6,
        'zxb2G3/wVHbjphQrt06Ogkf5mbRN+35mZ2t9iBrkmaiDNWMKR93lN6Slr501iuTGnUuTarX0tiaC1lcaf6biRSXXLs7VqM0ITdTu3eePcyAClBzho5v8KyzECPGq1G49qb4KwQ9S8MJPhvCoGW2RuOQAkRzEZE25qygLRI3wmt4=');
COMMIT;

-- ----------------------------
-- Table structure for admin_logs
-- ----------------------------
DROP TABLE IF EXISTS `admin_logs`;
CREATE TABLE `admin_logs`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `admin_user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `name`          varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '操作',
    `url`           varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'URL',
    `log_method`    varchar(8) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '不记录' COMMENT '记录日志方法',
    `log_ip`        varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '操作IP',
    `created_at`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    `updated_at`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台操作日志';

-- ----------------------------
-- Records of admin_logs
-- ----------------------------
BEGIN;
INSERT INTO `admin_logs`
VALUES (1, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:20:40', '2021-08-02 10:20:40');
INSERT INTO `admin_logs`
VALUES (2, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:31:46', '2021-08-02 10:31:46');
INSERT INTO `admin_logs`
VALUES (3, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:32:25', '2021-08-02 10:32:25');
INSERT INTO `admin_logs`
VALUES (4, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:33:55', '2021-08-02 10:33:55');
INSERT INTO `admin_logs`
VALUES (5, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:34:24', '2021-08-02 10:34:24');
INSERT INTO `admin_logs`
VALUES (6, 1, '登录', 'admin/auth/login', 'POST', '127.0.0.1', '2021-08-02 10:35:21', '2021-08-02 10:35:21');
COMMIT;

-- ----------------------------
-- Table structure for admin_menus
-- ----------------------------
DROP TABLE IF EXISTS `admin_menus`;
CREATE TABLE `admin_menus`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `parent_id`  int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单',
    `name`       varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '名称',
    `url`        varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'url',
    `route_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由名或者路由标识',
    `icon`       varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT 'fa-list' COMMENT '图标',
    `is_show`    tinyint(4) NOT NULL DEFAULT '1' COMMENT '等级',
    `sort_id`    int(10) unsigned NOT NULL DEFAULT '1000' COMMENT '排序',
    `log_method` varchar(8) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '不记录' COMMENT '记录日志方法',
    PRIMARY KEY (`id`) USING BTREE,
    KEY          `index_url` (`url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台菜单';

-- ----------------------------
-- Records of admin_menus
-- ----------------------------
BEGIN;
INSERT INTO `admin_menus`
VALUES (1, 0, '后台首页', 'admin/index/index', 'admin.index.index', 'fa-home', 1, 99, '不记录');
INSERT INTO `admin_menus`
VALUES (2, 0, '系统管理', 'admin/sys', 'admin.sys', 'fa-desktop', 1, 1099, '不记录');
INSERT INTO `admin_menus`
VALUES (3, 2, '用户管理', 'admin/admin_user/index', 'admin.admin_user.index', 'fa-user', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (4, 3, '添加用户界面', 'admin/admin_user/add', 'admin.admin_user.add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (5, 3, '修改用户界面', 'admin/admin_user/edit', 'admin.admin_user.edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (6, 3, '删除用户', 'admin/admin_user/del', 'admin.admin_user.del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (7, 2, '角色管理', 'admin/admin_role/index', 'admin.admin_role.index', 'fa-group', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (8, 7, '添加角色界面', 'admin/admin_role/add', 'admin.admin_role.add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (9, 7, '修改角色界面', 'admin/admin_role/edit', 'admin.admin_role.edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (10, 7, '删除角色', 'admin/admin_role/del', 'admin.admin_role.del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (11, 7, '角色授权界面', 'admin/admin_role/access', 'admin.admin_role.access', 'fa-key', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (12, 2, '菜单管理', 'admin/admin_menu/index', 'admin.admin_menu.index', 'fa-align-justify', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (13, 12, '添加菜单界面', 'admin/admin_menu/add', 'admin.admin_menu.add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (14, 12, '修改菜单界面', 'admin/admin_menu/edit', 'admin.admin_menu.edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (15, 12, '删除菜单', 'admin/admin_menu/del', 'admin.admin_menu.del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (16, 2, '操作日志', 'admin/admin_log/index', 'admin.admin_log.index', 'fa-keyboard-o', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (17, 16, '查看操作日志详情', 'admin/admin_log/view', 'admin.admin_log.view', 'fa-search-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (18, 2, '个人资料', 'admin/admin_user/profile', 'admin.admin_user.profile', 'fa-smile-o', 1, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (19, 2, '开发管理', 'admin/develop/manager', 'admin.develop.manager', 'fa-code', 1, 1005, '不记录');
INSERT INTO `admin_menus`
VALUES (20, 19, '设置配置', 'admin/develop/setting', 'admin.develop.setting', 'fa-cogs', 1, 995, '不记录');
INSERT INTO `admin_menus`
VALUES (21, 20, '设置管理', 'admin/setting/index', 'admin.setting.index', 'fa-cog', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (22, 21, '添加设置界面', 'admin/setting/add', 'admin.setting.add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (23, 21, '修改设置界面', 'admin/setting/edit', 'admin.setting.edit', 'fa-pencil', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (24, 21, '删除设置', 'admin/setting/del', 'admin.setting.del', 'fa-trash', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (25, 20, '设置分组管理', 'admin/setting_group/index', 'admin.setting_group.index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (26, 25, '添加设置分组界面', 'admin/setting_group/add', 'admin.setting_group.add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (27, 25, '修改设置分组界面', 'admin/setting_group/edit', 'admin.setting_group.edit', 'fa-pencil', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (28, 25, '删除设置分组', 'admin/setting_group/del', 'admin.setting_group.del', 'fa-trash', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (29, 0, '设置中心', 'admin/setting/center', 'admin.setting.center', 'fa-cogs', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (30, 29, '所有配置', 'admin/setting/all', 'admin.setting.all', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (31, 29, '后台设置', 'admin/setting/admin', 'admin.setting.admin', 'fa-adjust', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (32, 29, '更新设置', 'admin/setting/update', 'admin.setting.update', 'fa-pencil', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (33, 19, '数据维护', 'admin/database/table', 'admin.database.table', 'fa-database', 1, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (34, 33, '查看表详情', 'admin/database/view', 'admin.database.view', 'fa-eye', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (35, 33, '优化表', 'admin/database/optimize', 'admin.database.optimize', 'fa-refresh', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (36, 33, '修复表', 'admin/database/repair', 'admin.database.repair', 'fa-circle-o-notch', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (37, 3, '修改头像', 'admin/admin_user/update_avatar', 'admin.admin_user.update_avatar', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (38, 3, '添加用户', 'admin/admin_user/create', 'admin.admin_user.create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (39, 3, '修改用户', 'admin/admin_user/update', 'admin.admin_user.update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (40, 3, '用户启用', 'admin/admin_user/enable', 'admin.admin_user.enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (41, 3, '用户禁用', 'admin/admin_user/disable', 'admin.admin_user.disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (42, 3, '修改昵称', 'admin/admin_user/update_nickname', 'admin.admin_user.update_nickname', 'fa-list', 0, 1000,
        'POST');
INSERT INTO `admin_menus`
VALUES (43, 3, '修改密码', 'admin/admin_user/update_password', 'admin.admin_user.update_password', 'fa-list', 0, 1000,
        'POST');
INSERT INTO `admin_menus`
VALUES (44, 7, '创建角色', 'admin/admin_role/create', 'admin.admin_role.create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (45, 7, '修改角色', 'admin/admin_role/update', 'admin.admin_role.update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (46, 7, '启用角色', 'admin/admin_role/enable', 'admin.admin_role.enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (47, 7, '禁用角色', 'admin/admin_role/disable', 'admin.admin_role.disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (48, 7, '角色授权', 'admin/admin_role/access_operate', 'admin.admin_role.access_operate', 'fa-list', 0, 1000,
        'POST');
INSERT INTO `admin_menus`
VALUES (49, 12, '创建菜单', 'admin/admin_menu/create', 'admin.admin_menu.create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (50, 12, '修改菜单', 'admin/admin_menu/update', 'admin.admin_menu.update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (51, 21, '创建设置', 'admin/setting/create', 'admin.setting.create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (52, 21, '修改设置', 'admin/setting/do_update', 'admin.setting.do_update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (53, 31, '修改设置', 'admin/setting/info', 'admin.setting.info', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menus`
VALUES (54, 25, '创建设置分组', 'admin/setting_group/create', 'admin.setting_group.create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (55, 25, '修改设置', 'admin/setting_group/update', 'admin.setting_group.update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (56, 25, '生成菜单', 'admin/setting_group/menu', 'admin.setting_group.menu', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menus`
VALUES (57, 25, '生成文件', 'admin/setting_group/file', 'admin.setting_group.file', 'fa-list', 0, 1000, 'POST');
COMMIT;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`        varchar(50) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '' COMMENT '名称',
    `description` varchar(100) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '简介',
    `url`         varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限',
    `status`      tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台角色';

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
BEGIN;
INSERT INTO `admin_roles`
VALUES (1, '超级管理员', '后台管理员角色',
        '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,34,35,36,37,38,39,40,41,42,43,44,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79',
        1);
INSERT INTO `admin_roles`
VALUES (3, '管理员', '管理员', '1,2,3,16,17,18,42,43,58,59,60,61,62,63,64', 1);
INSERT INTO `admin_roles`
VALUES (4, '客服', '客服', '1,2,18,62', 1);
COMMIT;

-- ----------------------------
-- Table structure for admin_users
-- ----------------------------
DROP TABLE IF EXISTS `admin_users`;
CREATE TABLE `admin_users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`   varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'JDJ5JDEwJGUvaXZQcUMvbHBFcHUvVm16RWRrbU9ROFROYlMvMktwZnZqaGhWQ29uUi5GTGQ5Sng3SzlD' COMMENT '密码',
    `nickname`   varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`     varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '/static/admin/images/avatar.png' COMMENT '头像',
    `role`       varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色',
    `status`     tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 0：否 1：是',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='后台用户';

-- ----------------------------
-- Records of admin_users
-- ----------------------------
BEGIN;
INSERT INTO `admin_users`
VALUES (1, 'super_admin', 'JDJ5JDEwJGUzUzVkeHZYSmtBN2Y3SWVORHN4N3VjS29sSDY5UXA4MlJTbjFCYlhCU1J5Y2Rza3JZTkVD', '超级管理员',
        '/storage/attachment/N7je0ClI6pNpjslLJnkwM1T7j0EvpBZ6b6sQXpuH.jpg', '1', 1, NULL);
COMMIT;

-- ----------------------------
-- Table structure for attachments
-- ----------------------------
DROP TABLE IF EXISTS `attachments`;
CREATE TABLE `attachments`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT,
    `admin_user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '后台用户ID',
    `user_id`       int(10) unsigned NOT NULL DEFAULT '0' COMMENT '前台用户ID',
    `original_name` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '原文件名',
    `save_name`     varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '保存文件名',
    `save_path`     varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '系统完整路径',
    `url`           varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '网站url路径',
    `extension`     varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '后缀',
    `mime`          varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类型',
    `size`          bigint(20) NOT NULL DEFAULT '0' COMMENT '大小',
    `md5`           varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT 'MD5',
    `sha1`          varchar(40) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT 'SHA1',
    `created_at`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`    timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='附件表';

-- ----------------------------
-- Records of attachments
-- ----------------------------
BEGIN;
INSERT INTO `attachments`
VALUES (1, 0, 0, '450_450.jpg', 'QGcEEV2N4Vh5iaT2UreX4lVvxtd1nAqERSBZAPoo.jpg',
        '/mnt/www/website/redeemcode_server/storage/app/public/attachment/QGcEEV2N4Vh5iaT2UreX4lVvxtd1nAqERSBZAPoo.jpg',
        '/storage/attachment/QGcEEV2N4Vh5iaT2UreX4lVvxtd1nAqERSBZAPoo.jpg', 'jpg', 'image/jpeg', 31410,
        '9ed9076c6f01ef114c05ae7c671f2825', 'f5dc69383c5c6b607e009a7c9c69911ee82d7251', '2021-07-22 09:08:52',
        '2021-07-22 09:08:52', NULL);
INSERT INTO `attachments`
VALUES (2, 0, 0, '450_450.jpg', 'N7je0ClI6pNpjslLJnkwM1T7j0EvpBZ6b6sQXpuH.jpg',
        '/mnt/www/website/redeemcode_server/storage/app/public/attachment/N7je0ClI6pNpjslLJnkwM1T7j0EvpBZ6b6sQXpuH.jpg',
        '/storage/attachment/N7je0ClI6pNpjslLJnkwM1T7j0EvpBZ6b6sQXpuH.jpg', 'jpg', 'image/jpeg', 31410,
        '0a4cbf06818b497bab0f3375a3ae68fc', '976559775b311be31dff2f8b80e275b16e708ba9', '2021-07-22 09:09:10',
        '2021-07-22 09:09:10', NULL);
COMMIT;

-- ----------------------------
-- Table structure for setting
-- ----------------------------
DROP TABLE IF EXISTS `setting`;
CREATE TABLE `setting`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `setting_group_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属分组',
    `name`             varchar(50) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '名称',
    `description`      varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
    `code`             varchar(50) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '代码',
    `content`          text COLLATE utf8mb4_unicode_ci         NOT NULL COMMENT '设置配置及内容',
    `sort_number`      int(10) unsigned NOT NULL DEFAULT '1000' COMMENT '排序',
    `created_at`       timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`       timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='设置';

-- ----------------------------
-- Records of setting
-- ----------------------------
BEGIN;
INSERT INTO `setting`
VALUES (1, 1, '基本设置', '后台的基本信息设置', 'base',
        '[{\"Name\":\"后台名称\",\"Field\":\"name\",\"Type\":\"text\",\"Content\":\"beego通用后台系统\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"后台简称\",\"Field\":\"short_name\",\"Type\":\"text\",\"Content\":\"beego通用后台系统\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"后台作者\",\"Field\":\"author\",\"Type\":\"text\",\"Content\":\"bigpang\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"后台版本\",\"Field\":\"version\",\"Type\":\"text\",\"Content\":\"1.0\",\"Option\":\"\",\"Form\":\"\"}]',
        1000, '2020-04-26 13:44:31', '2020-07-29 15:36:47', NULL);
INSERT INTO `setting`
VALUES (2, 1, '登录设置', '后台登录相关设置', 'login',
        '[{\"Name\":\"登录token验证\",\"Field\":\"token\",\"Type\":\"switch\",\"Content\":\"1\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"验证码\",\"Field\":\"captcha\",\"Type\":\"select\",\"Content\":\"0\",\"Option\":\"0||不开启\\r\\n1||图形验证码\",\"Form\":\"\"},{\"Name\":\"登录背景\",\"Field\":\"background\",\"Type\":\"image\",\"Content\":\"/static/uploads/attachment/ea08c391-0eb4-4c6a-8e4f-9846c51d61cc.jpg\",\"Option\":\"\",\"Form\":\"\"}]',
        1, '2020-04-26 13:44:31', '2020-07-24 10:11:43', NULL);
INSERT INTO `setting`
VALUES (3, 1, '首页设置', '后台首页参数设置', 'index',
        '[{\"Name\":\"默认密码警告\",\"Field\":\"password_warning\",\"Type\":\"switch\",\"Content\":\"1\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"是否显示提示信息\",\"Field\":\"show_notice\",\"Type\":\"switch\",\"Content\":\"1\",\"Option\":\"\",\"Form\":\"\"},{\"Name\":\"提示信息内容\",\"Field\":\"notice_content\",\"Type\":\"text\",\"Content\":\"欢迎来到使用本系统，左侧为菜单区域，右侧为功能区。\",\"Option\":\"\",\"Form\":\"\"}]',
        1, '2020-04-26 13:44:31', '2020-07-29 15:36:59', NULL);
COMMIT;

-- ----------------------------
-- Table structure for setting_group
-- ----------------------------
DROP TABLE IF EXISTS `setting_group`;
CREATE TABLE `setting_group`
(
    `id`               int(10) unsigned NOT NULL AUTO_INCREMENT,
    `module`           varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '作用模块',
    `name`             varchar(50) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '名称',
    `description`      varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
    `code`             varchar(50) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '代码',
    `sort_number`      int(10) unsigned NOT NULL DEFAULT '1000' COMMENT '排序',
    `auto_create_menu` tinyint(4) NOT NULL DEFAULT '0' COMMENT '自动生成菜单',
    `auto_create_file` tinyint(4) NOT NULL DEFAULT '0' COMMENT '自动生成配置文件',
    `icon`             varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT 'fa-list' COMMENT '图标',
    `created_at`       timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`       timestamp NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='设置分组';

-- ----------------------------
-- Records of setting_group
-- ----------------------------
BEGIN;
INSERT INTO `setting_group`
VALUES (1, 'admin', '后台设置', '后台管理方面的设置', 'admin', 1000, 1, 1, 'fa-adjust', '2021-06-30 11:28:49', '2021-06-30 11:28:49',
        NULL);
COMMIT;

SET
FOREIGN_KEY_CHECKS = 1;
