# kkdl-go 项目指南

## 项目概述

kkdl-go 是一个基于 [GoFrame](https://goframe.org/) 框架构建的短链接生成服务。该项目提供完整的短链接管理功能，包括链接生成、统计分析、用户认证（支持邮箱、GitHub、验证码登录）、文件上传等功能。

### 主要技术栈

- **语言**: Go 1.24.0
- **框架**: GoFrame v2.10.0
- **数据库**: MySQL
- **认证**: JWT (github.com/gogf/gf-jwt/v2 v2.1.0)
- **对象存储**: 七牛云 (github.com/qiniu/go-sdk/v7 v7.25.6)
- **邮件**: jordan-wright/email v4.0.1
- **Excel 处理**: github.com/xuri/excelize/v2 v2.10.0
- **OAuth2**: golang.org/x/oauth2 v0.34.0
- **User Agent 解析**: github.com/mssola/useragent v1.0.0
- **WebSocket**: github.com/gorilla/websocket v1.5.3

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
   - WebSocket 实时通信

4. **用户管理**
   - 用户增删改查
   - 管理员权限控制

5. **通用功能**
   - 验证码生成（基于数据库存储，支持冷却期和原子性验证）
   - 文件上传（七牛云）
   - WebSocket 支持
   - 定时任务（短链码预生成）

## 构建和运行

### 前置要求

- Go 1.24.0 或更高版本
- MySQL 数据库
- GoFrame CLI 工具（可通过 `make up` 安装）

### 常用命令

#### 开发相关

首先需要安装 [GF CLI 工具](https://goframe.org/docs/cli/install)：

```bash
# 更新 GoFrame 及 CLI 到最新版本
make up

# 构建二进制文件
make build

# 生成 Controller（从 API 定义）
make ctrl

# 生成 DAO/DO/Entity（从数据库表结构）
make dao

# 生成 Service
make service

# 生成枚举
make enums

# 生成 protobuf 文件
make pb

# 生成数据库表的 protobuf 文件
make pbentity
```

#### GF CLI 命令说明

使用 `gf gen -help` 查看所有可用命令：

| 命令 | 说明 |
|------|------|
| `gf gen ctrl` | 解析 API 定义，生成 Controller 及 SDK 的 Go 文件 |
| `gf gen dao` | 自动生成 DAO / DO / Entity 相关的 Go 文件 |
| `gf gen enums` | 解析当前项目中的 Go 文件，生成枚举定义文件 |
| `gf gen pb` | 解析 `.proto` 文件，生成 Protobuf 的 Go 文件 |
| `gf gen pbentity` | 生成符合 Protobuf v3 格式的实体（Entity）消息文件 |
| `gf gen service` | 解析指定包中的结构体及其关联方法，生成 Service 层 Go 文件 |

#### 运行服务

```bash
# 指定配置文件运行
go run main.go --gf.gcfg.file=manifest/config/config.pro.yaml

# 使用部署脚本（需要先 make build）
./run.sh
```

#### Docker 相关

```bash
# 构建 Docker 镜像
make image

# 构建并推送 Docker 镜像
make image.push

# 部署到 Kubernetes
make deploy
```

### 配置文件

配置文件位于 `manifest/config/` 目录，当前示例配置文件为 `exampleConfig.yaml`。

主要配置项包括：

- **服务器配置**:
  - 端口（默认 6001）
  - OpenAPI 路径（/api.json）
  - Swagger 路径（/swagger）
  - 日志路径（log）
  - 访问日志（默认关闭）

- **数据库**: MySQL 连接配置（host、port、user、password、database）

- **七牛云**: access、secret、bucket、baseUrl

- **邮箱配置**: key（SMTP 授权码）、email（发件邮箱）

- **GitHub OAuth**: client_id、client_secret、redirect_uri

## 项目结构

```
kkdl-go/
├── api/                          # API 接口定义
│   ├── analytics/               # 分析相关接口
│   ├── common/                  # 通用接口（验证码、文件上传）
│   ├── login/                   # 登录相关接口
│   ├── short_url/               # 短链接管理接口
│   ├── short_url_code/          # 短链接码接口
│   ├── short_url_visits/        # 访问记录接口
│   └── user/                    # 用户管理接口
├── internal/                     # 内部实现
│   ├── cmd/                     # 命令行入口和路由注册
│   ├── consts/                  # 常量定义
│   ├── controller/              # HTTP 控制器
│   ├── dao/                     # 数据访问层
│   ├── logic/                   # 业务逻辑层
│   ├── middlewares/             # 中间件（认证、权限）
│   ├── model/                   # 数据模型
│   ├── packed/                  # 打包资源
│   └── service/                 # 服务接口定义
├── manifest/                     # 配置和部署文件
│   ├── config/                  # 配置文件
│   ├── deploy/                  # Kubernetes 部署配置
│   │   └── sql/                 # 数据库表结构 SQL 文件
│   └── docker/                  # Docker 相关文件
├── resource/                     # 静态资源
│   ├── public/                  # 公共资源
│   │   ├── html/                # HTML 文件
│   │   ├── plugin/              # 插件
│   │   └── resource/            # 资源文件
│   └── template/                # 模板文件
│       └── shortUrl/            # 短链接模板
├── utility/                      # 工具函数
│   ├── IntToBase62.go           # Base62 编码转换
│   └── utils.go                 # 通用工具函数
├── hack/                         # 构建脚本和工具
├── log/                          # 日志文件目录
├── main.go                       # 程序入口
├── Makefile                      # 构建配置
├── go.mod                        # Go 模块依赖
└── AGENTS.md                     # 项目指南（本文件）
```

## 开发约定

### 代码组织

- **API 定义**: 在 `api/` 目录下定义接口结构和输入输出参数
- **Controller**: 在 `internal/controller/` 下实现 HTTP 请求处理
- **Logic**: 在 `internal/logic/` 下实现核心业务逻辑
- **DAO**: 在 `internal/dao/` 下实现数据库操作
- **Service**: 在 `internal/service/` 下定义服务接口
- **Model**: 在 `internal/model/` 下定义数据模型（entity、do）

### 中间件

项目使用两个主要中间件：

1. **Auth 中间件** (`internal/middlewares/auth.go`): JWT 认证
2. **UserIsAdmin 中间件** (`internal/middlewares/admin.go`): 管理员权限验证

### 路由分组

所有路由都在 `/api` 前缀下，分为三个权限级别：

1. **公开接口**: 无需认证
   - 邮箱登录、注册、验证码登录
   - GitHub 登录
   - WebSocket 连接
   - 用户注册检查
   - 验证码生成
   - 统计分析接口（设备、区域、时间）

2. **认证接口**: 需要 JWT Token
   - 登出、Token 刷新
   - 短链接码批量创建
   - 文件上传
   - 短链接管理（创建、删除、获取列表、获取 URL、批量导入/导出、模板下载）
   - 访问记录查询

3. **管理员接口**: 需要管理员权限
   - 用户管理（创建、删除、获取列表、获取单个、更新）

### 数据库表结构

项目使用以下主要数据表（SQL 文件位于 `manifest/deploy/sql/`）：

- **user**: 用户信息表
- **short_url**: 短链接表
- **short_url_code**: 短链接码表
- **short_url_visits**: 访问记录表
- **verification_codes**: 验证码表

### 验证码机制

项目使用基于数据库的验证码机制，主要功能包括：

- **创建验证码**: 生成6位随机数字验证码，有效期10分钟
- **验证验证码**: 校验验证码是否正确且未过期
- **冷却期检查**: 同一邮箱1分钟内只能发送一次验证码
- **原子性验证和消费**: 使用 `VerifyAndConsumeVCode` 方法防止 TOCTOU 竞态条件
- **消费验证码**: 验证成功后标记为已使用
- **删除验证码**: 用于邮件发送失败时的回滚操作
- **清理过期验证码**: 定期清理已过期的验证码记录

相关实现位于 `internal/logic/common/verification-code.go`

### 定时任务

项目使用 GoFrame 的 `gcron` 实现定时任务：

- **短链码预生成任务**: 每天凌晨 3 点执行，检查未使用的短链码数量，如果少于 10000 个则生成 1000 个新码

相关实现位于 `internal/cmd/cron.go`

### 数据库优化

`verification_codes` 表已创建以下索引以优化查询性能：

```sql
KEY `idx_email_type_created` (`email`,`type`,`created_at`),
KEY `idx_expired_at` (`expired_at`)
```

## 开发注意事项

1. **验证码机制**: 验证码存储在数据库中，使用原子性操作 `VerifyAndConsumeVCode` 防止 TOCTOU 竞态条件
2. **登录逻辑**: 登录相关逻辑需要进一步完善，代码中已有 TODO 标记
3. **统计分析**: 区域和设备统计功能需要增加时间过滤条件
4. **配置文件**: 开发时请根据实际情况复制 `manifest/config/exampleConfig.yaml` 并修改为实际配置文件
5. **API 生成**: 修改 API 定义后，需要运行 `make ctrl` 重新生成 Controller
6. **数据库变更**: 修改数据库表结构后，需要运行 `make dao` 重新生成 DAO/DO/Entity
7. **Git 分支**: 当前项目主要在 `dev` 分支进行开发

## 近期变更

### 架构重构
- **移除 Redis 依赖**: 验证码机制从 Redis 迁移到数据库存储
- **用户模型优化**: 移除了 wxId 字段，简化用户数据结构
- **版本升级**: Go 版本升级到 1.24.0，GoFrame 升级到 v2.10.0
- **模块名更新**: 项目模块名改为 `compressURL`（保持一致性）

### 新增功能
- **验证码冷却期**: 防止验证码滥用，同一邮箱1分钟内只能发送一次
- **验证码自动清理**: 支持清理过期验证码记录
- **原子性验证码操作**: 实现 `VerifyAndConsumeVCode` 方法消除 TOCTOU 竞态条件
- **验证码回滚机制**: 邮件发送失败时支持删除已创建的验证码
- **改进的邮件模板**: 验证码邮件使用更美观的 HTML 模板
- **定时任务**: 添加短链码预生成定时任务

### 数据库优化
- 为 `verification_codes` 表添加复合索引 `idx_email_type_created` 优化冷却期查询
- 为 `verification_codes` 表添加 `idx_expired_at` 索引优化过期验证码清理

## 相关文档

- [GoFrame 官方文档](https://goframe.org/display/gf)
- [项目详细介绍](https://juejin.cn/post/7345310754470887458)
- [GF CLI 安装文档](https://goframe.org/docs/cli/install)