CREATE TABLE short_url_code (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '唯一标识，自增长整数',
    `code` varchar(20) NOT NULL COMMENT '短链,唯一，不能为空',
    `status` tinyint(1) NOT NULL COMMENT '是否使用',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认为当前时间戳',
    PRIMARY KEY (`id`),
    UNIQUE KEY `code` (`code`)
) COMMENT '预生成的短链 code';