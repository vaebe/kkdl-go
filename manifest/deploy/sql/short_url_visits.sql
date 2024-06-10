-- auto-generated definition
create table short_url_visits
(
    id              bigint unsigned auto_increment
        primary key,
    user_id         varchar(40)                         not null comment '用户id',
    short_url       varchar(20)                         not null comment '短链,唯一，不能为空',
    raw_url         varchar(500)                        not null comment '原始 url 不能为空',
    user_agent      varchar(255)                        not null comment '用户代理字符串，存储提供的完整用户代理',
    browser_name    varchar(50)                         null comment '浏览器名称',
    browser_version varchar(50)                         null comment '浏览器版本',
    device_model    varchar(50)                         null comment '设备型号',
    engine_name     varchar(50)                         null comment '浏览器引擎名称',
    engine_version  varchar(50)                         null comment '浏览器引擎版本',
    os_name         varchar(50)                         null comment '操作系统名称',
    os_version      varchar(50)                         null comment '操作系统版本',
    ip              varchar(45)                         not null comment 'ip 不能为空',
    continent       varchar(50)                         null comment '大洲名称',
    continent_code  varchar(10)                         null comment '大洲代码',
    country         varchar(50)                         null comment '国家名称',
    country_code    varchar(10)                         null comment '国家代码',
    region          varchar(10)                         null comment '地区或州的短代码（FIPS或ISO）',
    region_name     varchar(50)                         null comment '地区或州名称',
    city            varchar(50)                         null comment '城市名称',
    district        varchar(50)                         null comment '位置的区（郡）',
    lat             double                              null comment '纬度',
    lon             double                              null comment '经度',
    created_at      timestamp default CURRENT_TIMESTAMP null comment '创建时间，默认为当前时间戳',
    updated_at      timestamp                           null comment '更新时间',
    deleted_at      timestamp                           null comment '删除时间',
    constraint id_unique
        unique (id)
);

create index idx_short_url_visits_on_created_at
    on short_url_visits (created_at);

create index idx_short_url_visits_on_short_url
    on short_url_visits (short_url);

