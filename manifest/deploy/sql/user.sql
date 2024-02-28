create table user
(
    id          varchar(32)                          not null comment '唯一标识'
        primary key,
    email       varchar(50)                          null comment '邮箱,唯一',
    wxId        varchar(50)                          null comment '小程序id,唯一',
    password    varchar(200)                         null comment '密码, 小程序登录无密码',
    nickName    varchar(20)                          not null comment '昵称, 创建默认生成',
    accountType varchar(4) default '01'              not null comment '账号类型: 01 邮箱 02 小程序',
    role        varchar(4) default '01'              not null comment '角色: 00 admin 01 普通用户 02 vip',
    deleted_at  timestamp                            null comment '删除时间',
    updated_at  timestamp                            null comment '更新时间',
    created_at  timestamp  default CURRENT_TIMESTAMP null comment '创建时间',
    salt        varchar(12)                          not null comment '用户盐值',
    avatar      varchar(300)                         null comment '用户头像',
    constraint email
        unique (email),
    constraint wxId
        unique (wxId)
);