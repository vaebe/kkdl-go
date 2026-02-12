# compressURL 项目指南

## 项目概述

compressURL 是一个基于 [GoFrame](https://goframe.org/) 框架构建的短链接生成服务。该项目提供完整的短链接管理功能，包括链接生成、统计分析、用户认证（支持邮箱、GitHub、验证码登录）、文件上传等功能。

### 主要技术栈

- **语言**: Go 1.24+
- **框架**: GoFrame v2.10.0
- **数据库**: MySQL
- **认证**: JWT (github.com/gogf/gf-jwt/v2)
- **对象存储**: 七牛云 (github.com/qiniu/go-sdk/v7)
- **邮件**: jordan-wright/email
- **Excel 处理**: excelize/v2
- **OAuth2**: golang.org/x/oauth2
- **User Agent 解析**: github.com/mssola/useragent

### 核心功能模块

1. **短链接管理**
    - 创建、删除、查询短链接
    - 批量导入/导出
    - 短链接码批量生成

2. **统计分析**
    - 点击设备统计
    - 点击区域统计
    - 点击时间统计
    - 访问记录查询

3. **用户认证**
    - 邮箱登录
    - GitHub OAuth 登录
    - 验证码登录
    - JWT Token 刷新

4. **用户管理**
    - 用户增删改查
    - 管理员权限控制

5. **通用功能**
    - 验证码生成（基于数据库存储）
    - 文件上传（七牛云）
    - WebSocket 支持

## 掘金文章介绍

[猛击查看](https://juejin.cn/post/7345310754470887458)

![image.png](https://github.com/user-attachments/assets/74a99c39-1000-4c17-a0e8-09319985238d)

## 常用命令

### 开发

首先需要安装 [gf cli 工具](https://goframe.org/docs/cli/install) 非常好用！

使用 `gf run main.go` 启动项目

最常用的命令是 `gf gen xxx`

使用 `gf gen -help` 可以查看有哪些命令，如 `gf gen service`

| 命令         | 说明                                    |
|------------|---------------------------------------|
| `ctrl`     | 解析 API 定义，生成 Controller 及 SDK 的 Go 文件 |
| `dao`      | 自动生成 DAO / DO / Entity 相关的 Go 文件      |
| `enums`    | 解析当前项目中的 Go 文件，生成枚举定义文件               |
| `pb`       | 解析 `.proto` 文件，生成 Protobuf 的 Go 文件    |
| `pbentity` | 生成符合 Protobuf v3 格式的实体（Entity）消息文件    |
| `service`  | 解析指定包中的结构体及其关联方法，生成 Service 层 Go 文件   |


## 配置

### 指定配置文件启动
`go run main.go --gf.gcfg.file=manifest/config/config.pro.yaml`

## todo

### 统计分析
1. 根据区域、设备统计 增加 时间过滤条件

### login

1. 梳理所有登录相关的逻辑,目前写的不太严谨
