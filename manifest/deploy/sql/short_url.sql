CREATE TABLE `short_url` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识，自增长整数',
    `shortUrl` varchar(20) NOT NULL COMMENT '短链,唯一，不能为空',
    `rawUrl` varchar(500) NOT NULL COMMENT '原始 url 不能为空',
    `expirationTime` timestamp NULL DEFAULT NULL COMMENT '过期时间',
    `userId` varchar(40) DEFAULT NULL COMMENT '用户id',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认为当前时间戳',
    `title` varchar(100) NOT NULL COMMENT '短链标题',
    `group_id` int(11) DEFAULT NULL COMMENT '短链分组id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `shortUrl` (`shortUrl`)
) COMMENT '短链列表';