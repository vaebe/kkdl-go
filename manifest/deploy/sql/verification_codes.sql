create table verification_codes
(
    id int auto_increment primary key,
    email varchar(100) not null comment '邮箱',
    type tinyint not null comment '验证码类型: 1注册 2登录 3重置密码',
    code varchar(10) not null comment '验证码',
    used tinyint(1) default 0 comment '是否已使用: 0未使用 1已使用',
    expired_at timestamp not null comment '过期时间',
    created_at timestamp default current_timestamp comment '创建时间',
    index idx_email_type_created (email, type, created_at),
    index idx_expired_at (expired_at)
) comment '邮箱验证码表';
