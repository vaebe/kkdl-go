create table short_url_code
(
    id         int auto_increment comment '唯一标识，自增长整数'
        primary key,
    code       varchar(20) not null comment '短链,唯一，不能为空',
    status     tinyint(1) not null comment '是否使用',
    created_at timestamp default CURRENT_TIMESTAMP null comment '创建时间，默认为当前时间戳',
    constraint code
        unique (code)
);

