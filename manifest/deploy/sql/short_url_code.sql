create table short_url_code
(
    id         INT AUTO_INCREMENT PRIMARY KEY COMMENT '唯一标识，自增长整数',
    code       VARCHAR(20) NOT NULL UNIQUE COMMENT '短链,唯一，不能为空',
    status     BOOL        NOT NULL COMMENT '是否使用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认为当前时间戳'
);