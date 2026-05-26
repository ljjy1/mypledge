# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

sponge 生成的 Go HTTP 单体服务 (module: `pledge-be`, binary: `pledge_be`)，采用分层架构。

## 常用命令

```bash
make run                    # 编译并运行服务 (默认端口 8080)
make run Config=configs/dev.yml  # 指定配置文件运行
make test                   # 运行所有单测 (禁用缓存)
make docs                   # 重新生成 swagger 文档 (API 变更后执行)
make build                  # 编译 linux amd64 二进制
make clean                  # 清理编译产物
make update-config          # 从 yaml 配置同步 Go 结构体代码
make ci-lint                # 代码规范检查 (gofmt + golangci-lint)
```

单测运行: `go test -count=1 -short ./internal/handler/...`

macOS debug 构建需添加 `-buildmode=pie` 避免 dyld LC_UUID 缺失错误。

## 技术栈

| 组件 | 库 |
|------|-----|
| Web 框架 | Gin |
| ORM | GORM (MySQL) |
| 缓存 | go-redis / memory |
| 配置 | Viper |
| 日志 | Zap |
| 监控 | Prometheus |
| 链路追踪 | Jaeger (OpenTelemetry) |
| 认证 | sponge JWT (HS256) |

## JWT 认证

签名密钥硬编码在 `internal/auth/jwt.go`，`auth.Middleware()` 中间件保护 `/api/v1` 组下的所有路由。
登录接口 `POST /api/v1/login` 不需要 JWT 认证，登录成功返回 token。

调用受保护的 API 时在请求头添加: `Authorization: Bearer <token>`

## 架构

**初始化入口**: `cmd/pledge_be/main.go` → `cmd/pledge_be/initial/` (配置、数据库、缓存、日志、链路追踪)

**分层调用链路**:

```
cmd/pledge_be/main.go
  → internal/server/http.go       (HTTP 服务启动/关闭)
    → internal/routers/routers.go  (路由注册 + 中间件链)
      → internal/handler/          (API 处理层：参数校验、响应)
        → internal/dao/            (数据访问层：数据库 + 缓存 + singleflight 防击穿)
          → internal/model/         (数据模型)
```

各层职责：
- **initial**: 初始化配置、数据库、缓存、日志、链路追踪、资源统计
- **routers**: 全局中间件 (CORS、RequestID、日志、限流、熔断、链路追踪、指标采集)
- **handler**: 接口定义 + 参数校验 + swagger 注解
- **dao**: 数据操作 + 缓存读写 + `singleflight` 防缓存击穿 + 事务支持
- **model**: GORM 模型定义 + 查询字段白名单防 SQL 注入
- **cache**: 缓存封装 (Redis/Memory 双实现)
- **ecode**: 业务错误码 (sponge 的 `errcode.HCode` 定义)
- **types**: 请求/响应结构体 (含 validator 标签)
- **database**: MySQL/Redis 客户端初始化

## API 路由

API 统一前缀 `/api/v1`，Swagger 文档访问 `/swagger/index.html` (仅非 prod 环境)。

当前用户模块路由 (见 `internal/routers/user.go`):
- `POST /api/v1/login` — 用户登录 (无需 JWT)
- `POST /api/v1/user` — 创建用户
- `DELETE /api/v1/user/:id` — 删除用户
- `PUT /api/v1/user/:id` — 更新用户
- `GET /api/v1/user/:id` — 查询用户
- `POST /api/v1/user/list` — 用户列表 (分页)

## 项目目录结构

```
cmd/pledge_be/       # 应用入口
internal/
  auth/              # JWT 认证 (生成令牌、中间件)
  cache/             # 缓存层
  config/            # 配置结构体
  dao/               # 数据访问层
  database/          # MySQL/Redis 客户端
  ecode/             # 业务错误码
  handler/           # API 处理层
  model/             # 数据模型
  routers/           # 路由注册
  server/            # HTTP 服务
  types/             # 请求/响应结构体
scripts/             # 构建/运行/部署脚本
deployments/         # 部署配置 (binary/docker/k8s)
```
