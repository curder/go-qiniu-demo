-- 系统用户
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_username` (`username`),
  UNIQUE KEY `uniq_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 七牛账户
CREATE TABLE IF NOT EXISTS `qiniu_accounts` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT '账户邮箱',
    `access_key` varchar(255) NOT NULL DEFAULT '' COMMENT '七牛AK',
    `secret_key` varchar(255) NOT NULL DEFAULT '' COMMENT '七牛SK',
    `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '账户描述',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='七牛账户';

-- 七牛存储桶
CREATE TABLE IF NOT EXISTS `qiniu_buckets` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) unsigned NOT NULL COMMENT '所属账户',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '存储桶名称',
    `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储桶描述',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='七牛存储桶';

-- 七牛域名
CREATE TABLE IF NOT EXISTS `qiniu_domains` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `bucket_id` bigint(20) unsigned NOT NULL COMMENT '所属存储桶',
  `protocol` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '服务协议，包含http/https，默认https',
  `hostname` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '域名主机',
  `description` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '域名描述',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='七牛云域名';

-- 七牛资源
CREATE TABLE IF NOT EXISTS `qiniu_files` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `bucket_id` bigint(20) unsigned NOT NULL COMMENT '所属存储桶',
    `file_name` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件名',
    `file_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件类型',
    `storage_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储类型',
    `file_size` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件大小',
    `last_modified` timestamp NULL DEFAULT NULL COMMENT '最后更新',
    `e_tag` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件etag信息',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='七牛云文件资源';
