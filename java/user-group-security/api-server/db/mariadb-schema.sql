
DROP DATABASE IF EXISTS example;
CREATE DATABASE example;
GRANT ALL PRIVILEGES ON example.* TO 'dbuser'@'localhost' IDENTIFIED BY 'dbpass';
GRANT ALL PRIVILEGES ON example.* TO 'dbuser'@'%' IDENTIFIED BY 'dbpass';

USE example;

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '*用户名',
  `password` varchar(50) NOT NULL COMMENT '*密码',
  `email` varchar(50) DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(11) DEFAULT '' COMMENT '手机',
  `status` tinyint(8) unsigned DEFAULT '0' COMMENT '*用户状态, 0: 开通, 0x80: 不可用, 0xFF: 销户',
  PRIMARY KEY (`id`),
  INDEX (`username`),
  CONSTRAINT UNIQUE (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Table structure for `profile`
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `id` bigint(20) NOT NULL COMMENT '*用户id',
  `fullname` varchar(100) DEFAULT NULL COMMENT '*姓名',
  `email_with_labels` varchar(100) DEFAULT NULL COMMENT '电子邮件',
  `phone_with_labels` varchar(100) DEFAULT NULL,
  `avatar` varchar(100) DEFAULT NULL,
  `cv` text,
  `theme` varchar(100) DEFAULT NULL,
  `first_created` datetime DEFAULT NULL COMMENT '创建时间',
  `last_updated` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '销户时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户home面貌表';

-- ----------------------------
-- Table structure for `t_group`
-- ----------------------------
DROP TABLE IF EXISTS `t_group`;
CREATE TABLE `t_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `owner_id` bigint(20) DEFAULT NULL COMMENT '*创建者',
  `name` varchar(200) DEFAULT NULL,
  `classifier` tinyint(8) unsigned DEFAULT '0' COMMENT '分类, 0: 群, 1: 个人圈子, 0x80: 机构, 0xC0: 分支机构, 0xE0: 部门, 0xF0: 团队',
  `created_time` date DEFAULT NULL,
  `last_updated` date DEFAULT NULL,
  `deleted_time` date DEFAULT NULL,
  `logo` varchar(100) DEFAULT NULL,
  `kind` varchar(100) DEFAULT NULL,
  `status` tinyint(8) unsigned DEFAULT '0' COMMENT '状态位, 0: 有效, 0xFF: 无效',
  PRIMARY KEY (`id`),
  UNIQUE(`name`),
  INDEX (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='群';

  # FOREIGN KEY (`owner_id`) REFERENCES user (`id`) ON DELETE SET NULL

ALTER TABLE `t_group`
  ADD FOREIGN KEY (`owner_id`)
  REFERENCES user (`id`)
  ON DELETE SET NULL;

-- ----------------------------
-- Table structure for `organization`
-- ----------------------------
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization` (
  `id` bigint(20) NOT NULL,
  `classifier` tinyint(8) unsigned NOT NULL COMMENT '分类, 0: 群, 1: 个人圈子, 0x80: 机构, 0xC0: 分支机构, 0xE0: 部门, 0xF0: 团队',
  `address` varchar(200) DEFAULT NULL,
  `postcode` varchar(20) DEFAULT NULL,
  `fax` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `industry` tinyint(8) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8;




-- ----------------------------
-- Table structure for `t_member`
-- ----------------------------
DROP TABLE IF EXISTS `t_member`;
CREATE TABLE `t_member` (
  `group_id` bigint(20) NOT NULL COMMENT '群id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `participate_time` datetime DEFAULT NULL COMMENT '进群时间',
  `departure_time` datetime DEFAULT NULL COMMENT '退群时间',
  `invite_time` datetime DEFAULT NULL COMMENT '邀请',
  `apply_time` datetime DEFAULT NULL COMMENT '申请',
  `status` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`group_id`, `user_id`),
  INDEX (`group_id`),
  INDEX (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='群关系表';

  # FOREIGN KEY (`group_id`) REFERENCES t_group (`id`) ON DELETE CASCADE,
  # FOREIGN KEY (`user_id`) REFERENCES user (`id`) ON DELETE CASCADE

ALTER TABLE `t_member`
  ADD FOREIGN KEY (`group_id`)
  REFERENCES t_group (`id`)
  ON DELETE CASCADE,
  ADD FOREIGN KEY (`user_id`)
  REFERENCES user (`id`)
  ON DELETE CASCADE;


-- ----------------------------
-- Table structure for `role`
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `classifier` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'various class',
  `created_time` datetime DEFAULT NULL COMMENT 'first created',
  `last_updated` datetime DEFAULT NULL COMMENT 'last updated',
  `deleted_time` datetime DEFAULT NULL COMMENT 'deleted',
  `status` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  INDEX (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色表';




-- ----------------------------
-- Table structure for `t_resource`
-- ----------------------------
DROP TABLE IF EXISTS `t_resource`;
CREATE TABLE `t_resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `kind` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'various type',
  `description` varchar(100) NULL,
  `created_time` datetime DEFAULT NULL COMMENT 'first created',
  `last_updated` datetime DEFAULT NULL COMMENT 'last updated',
  `deleted_time` datetime DEFAULT NULL COMMENT 'deleted',
  `status` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  INDEX (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Table structure for `rule`
-- ----------------------------
DROP TABLE IF EXISTS `rule`;
CREATE TABLE `rule` (
  `resource_id` bigint(20) NOT NULL COMMENT '群id',
  `role_id` bigint(20) NOT NULL COMMENT '用户id',
  `crud_permission` varchar(100) NOT NULL,
  `kind` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'various type',
  `description` varchar(100) NULL,
  `created_time` datetime DEFAULT NULL COMMENT 'first created',
  `last_updated` datetime DEFAULT NULL COMMENT 'last updated',
  `deleted_time` datetime DEFAULT NULL COMMENT 'deleted',
  `status` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`resource_id`, `role_id`),
  FOREIGN KEY (`resource_id`) REFERENCES t_resource (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`role_id`) REFERENCES role (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Table structure for `role_binding`
-- ----------------------------
DROP TABLE IF EXISTS `role_binding`;
CREATE TABLE `role_binding` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL COMMENT '群id',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  `group_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `crud_permission` varchar(100) NOT NULL,
  `kind` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'various type',
  `description` varchar(100) NULL,
  `created_time` datetime DEFAULT NULL COMMENT 'first created',
  `last_updated` datetime DEFAULT NULL COMMENT 'last updated',
  `deleted_time` datetime DEFAULT NULL COMMENT 'deleted',
  `status` tinyint(8) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`role_id`) REFERENCES role (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES user (`id`) ON DELETE SET NULL,
  FOREIGN KEY (`group_id`) REFERENCES t_group (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8 COMMENT='角色表';
