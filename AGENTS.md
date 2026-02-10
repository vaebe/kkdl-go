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

## 构建和运行

### 前置要求

- Go 1.24 或更高版本
- MySQL 数据库
- GoFrame CLI 工具（可通过 `make up` 安装）

### 常用命令

#### 开发相关

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

配置文件位于 `manifest/config/` 目录，主要配置项包括：

- **服务器配置**: 端口（默认 6001）、日志路径、访问日志等
- **数据库**: MySQL 连接配置
- **七牛云**: access key、secret、bucket、base URL
- **邮箱配置**: SMTP key 和 email
- **GitHub OAuth**: client_id 和 client_secret

## 项目结构

```
compressURL/
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
│   └── docker/                  # Docker 相关文件
├── resource/                     # 静态资源
│   ├── public/                  # 公共资源
│   └── template/                # 模板文件
├── utility/                      # 工具函数
├── hack/                        # 构建脚本和工具
├── main.go                      # 程序入口
├── Makefile                     # 构建配置
└── go.mod                       # Go 模块依赖
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

路由分为三个权限级别：

1. **公开接口**: 无需认证
   - 登录、注册、验证码登录
   - GitHub 登录
   - 验证码生成
   - 统计分析接口

2. **认证接口**: 需要 JWT Token
   - 登出、Token 刷新
   - 短链接管理
   - 文件上传
   - 访问记录查询

3. **管理员接口**: 需要管理员权限
   - 用户管理

### 数据库表结构

项目使用以下主要数据表：

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
- **消费验证码**: 验证成功后标记为已使用
- **清理过期验证码**: 定期清理已过期的验证码记录

相关实现位于 `internal/logic/common/verification-code.go`

### 数据库优化

建议为以下表创建索引以优化查询性能：

```sql
CREATE INDEX idx_short_url_visits_on_created_at ON short_url_visits(created_at);
CREATE INDEX idx_short_url_visits_on_short_url ON short_url_visits(short_url);
CREATE INDEX idx_verification_codes_on_email ON verification_codes(email);
CREATE INDEX idx_verification_codes_on_expired_at ON verification_codes(expired_at);
```

## 开发注意事项

1. **验证码机制**: 验证码现在存储在数据库中，不再依赖 Redis
2. **登录逻辑**: 目前登录相关逻辑需要进一步完善，代码中已有 TODO 标记
3. **统计分析**: 区域和设备统计功能需要增加时间过滤条件
4. **配置文件**: 开发时请根据实际情况修改 `manifest/config/` 下的配置文件
5. **API 生成**: 修改 API 定义后，需要运行 `make ctrl` 重新生成 Controller
6. **数据库变更**: 修改数据库表结构后，需要运行 `make dao` 重新生成 DAO/DO/Entity

## 近期变更

### 架构重构
- **移除 Redis 依赖**: 验证码机制从 Redis 迁移到数据库存储
- **用户模型优化**: 移除了 wxId 字段，简化用户数据结构
- **版本升级**: Go 版本升级到 1.24.0，GoFrame 升级到 v2.10.0

### 新增功能
- **验证码冷却期**: 防止验证码滥用，同一邮箱1分钟内只能发送一次
- **验证码自动清理**: 支持清理过期验证码记录
- **改进的邮件模板**: 验证码邮件使用更美观的 HTML 模板

## 相关文档

- [GoFrame 官方文档](https://goframe.org/display/gf)
- [项目详细介绍](https://juejin.cn/post/7345310754470887458)