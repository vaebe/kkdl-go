-- auto-generated definition
create table short_url_visits
(
    id                 bigint unsigned auto_increment
        primary key,
    userId             varchar(40)                         not null comment '用户id',
    shortUrl           varchar(20)                         not null comment '短链,唯一，不能为空',
    rawUrl             varchar(500)                        not null comment '原始 url 不能为空',
    ip                 varchar(45)                         not null comment 'ip 不能为空',
    user_agent         text                                not null comment '客户端软件的类型、版本和其他相关信息',
    sec_ch_ua          varchar(255)                        not null comment '客户端使用的浏览器和版本',
    sec_ch_ua_mobile   varchar(10)                         not null comment '请求是否来自移动设备。?0 表示不是移动设备，?1 表示是移动设备',
    sec_ch_ua_platform varchar(255)                        not null comment '客户端所运行的平台',
    sec_fetch_user     varchar(10)                         null comment '请求是否是用户发起的,?1 表示是用户发起的请求',
    created_at         timestamp default CURRENT_TIMESTAMP null comment '创建时间，默认为当前时间戳',
    updated_at         timestamp                           null comment '更新时间',
    deleted_at         timestamp                           null comment '删除时间',
    constraint id
        unique (id)
);

