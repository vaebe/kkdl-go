CREATE TABLE `user`(
    `id` varchar(32) NOT NULL COMMENT '唯一标识',
    `email` varchar(50) DEFAULT NULL COMMENT '邮箱,唯一',
    `password` varchar(200) DEFAULT NULL COMMENT '密码, 小程序登录无密码',
    `nickName` varchar(20) NOT NULL COMMENT '昵称, 创建默认生成',
    `accountType` varchar(4) NOT NULL DEFAULT '01' COMMENT '账号类型: 01 邮箱 02 小程序',
    `role` varchar(4) NOT NULL DEFAULT '01' COMMENT '角色: 00 admin 01 普通用户 02 vip',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
    `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `salt` varchar(12) DEFAULT NULL COMMENT '用户盐值',
    `avatar` varchar(300) DEFAULT NULL COMMENT '用户头像',
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`)
) COMMENT '用户表';