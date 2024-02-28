create table short_url
(
    id             int auto_increment comment '唯一标识，自增长整数'
        primary key,
    shortUrl       varchar(20)                         not null comment '短链,唯一，不能为空',
    rawUrl         varchar(500)                        not null comment '原始 url 不能为空',
    expirationTime timestamp                           null comment '过期时间',
    userId         varchar(40)                         null comment '用户id',
    created_at     timestamp default CURRENT_TIMESTAMP null comment '创建时间，默认为当前时间戳',
    title          varchar(100)                        not null comment '短链标题',
    group_id       int                                 null comment '短链分组id',
    constraint shortUrl
        unique (shortUrl)
);