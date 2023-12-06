create table short_url
(
    id             INT AUTO_INCREMENT PRIMARY KEY COMMENT '唯一标识，自增长整数',
    shortUrl       VARCHAR(20)  NOT NULL UNIQUE COMMENT '短链,唯一，不能为空',
    rawUrl         VARCHAR(500) NOT NULL COMMENT '原始 url 不能为空',
    expirationTime TIMESTAMP COMMENT '过期时间',
    userId         INT COMMENT '用户id',
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认为当前时间戳'
);