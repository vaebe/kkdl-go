CREATE TABLE verification_codes (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `email` varchar(100) NOT NULL COMMENT '邮箱',
    `type` tinyint(4) DEFAULT '1' COMMENT '验证码类型: 目前只有一种用途',
    `CODE` varchar(10) NOT NULL COMMENT '验证码',
    `used` tinyint(1) DEFAULT '0' COMMENT '是否已使用: 0未使用 1已使用',
    `expired_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '过期时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_email_type_created` (`email`,`type`,`created_at`),
    KEY `idx_expired_at` (`expired_at`)
) COMMENT '邮箱验证码表';