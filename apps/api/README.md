# AI Project OS API

AI Project OS 后端服务，基于 Go Fiber + GORM + MySQL + Redis。

## 技术栈

- **Web 框架**: Go Fiber v2
- **ORM**: GORM
- **数据库**: MySQL 8.0
- **缓存**: Redis 7
- **认证**: JWT (HS256)
- **日志**: Zap
- **配置**: godotenv + 环境变量

## 快速开始

### 1. 启动 MySQL 和 Redis

```bash
docker compose up -d
```

### 2. 配置环境变量

```bash
cp .env.example .env
# 按需修改 .env 中的配置
```

### 3. 安装依赖并启动

```bash
cd apps/api
go mod tidy
go run cmd/server/main.go
```

服务默认运行在 `http://localhost:8080`。

### 4. 健康检查

```bash
curl http://localhost:8080/health
```

## 项目结构

```
apps/api/
├── cmd/server/main.go          # 入口
├── internal/
│   ├── ai/                     # AI Provider 抽象层
│   │   ├── provider.go         # 接口定义
│   │   └── mock_provider.go    # Mock 实现
│   ├── config/                 # 配置加载
│   ├── database/               # 数据库初始化 + AutoMigrate
│   ├── handlers/               # HTTP Handler（Controller 层）
│   ├── logger/                 # Zap 日志初始化
│   ├── middleware/             # JWT 认证 + 全局中间件
│   ├── models/                 # GORM 模型
│   ├── repositories/           # 数据访问层
│   ├── routes/                 # 路由注册
│   ├── services/               # 业务逻辑层
│   └── utils/                  # 统一响应格式
├── docker-compose.yml
├── .env.example
├── go.mod
└── README.md
```

## API 端点

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| GET | `/health` | 健康检查 | ❌ |
| POST | `/api/v1/auth/register` | 用户注册 | ❌ |
| POST | `/api/v1/auth/login` | 用户登录 | ❌ |
| GET | `/api/v1/auth/me` | 当前用户 | ✅ |
| POST | `/api/v1/projects` | 创建项目 | ✅ |
| GET | `/api/v1/projects` | 项目列表 | ✅ |
| GET | `/api/v1/projects/:id` | 项目详情 | ✅ |
| PUT | `/api/v1/projects/:id` | 更新项目 | ✅ |
| DELETE | `/api/v1/projects/:id` | 删除项目 | ✅ |
| POST | `/api/v1/projects/:projectId/prompts` | 创建 Prompt | ✅ |
| GET | `/api/v1/projects/:projectId/prompts` | Prompt 列表 | ✅ |
| GET | `/api/v1/projects/:projectId/prompts/:id` | Prompt 详情 | ✅ |
| PUT | `/api/v1/projects/:projectId/prompts/:id` | 更新 Prompt | ✅ |
| DELETE | `/api/v1/projects/:projectId/prompts/:id` | 删除 Prompt | ✅ |
| GET | `/api/v1/projects/:projectId/prompts/:id/versions` | Prompt 版本历史 | ✅ |
| POST | `/api/v1/projects/:projectId/docs` | 创建文档 | ✅ |
| GET | `/api/v1/projects/:projectId/docs` | 文档列表 | ✅ |
| GET | `/api/v1/projects/:projectId/docs/:docId` | 文档详情 | ✅ |
| PUT | `/api/v1/projects/:projectId/docs/:docId` | 更新文档 | ✅ |
| DELETE | `/api/v1/projects/:projectId/docs/:docId` | 删除文档 | ✅ |
| GET | `/api/v1/templates` | 模板列表 | ❌ |
| GET | `/api/v1/templates/:id` | 模板详情 | ❌ |
| POST | `/api/v1/generations` | AI 生成 | ✅ |
| GET | `/api/v1/generations` | 生成记录 | ✅ |
| GET | `/api/v1/dashboard/stats` | 仪表盘统计 | ✅ |

## 统一响应格式

```json
{
  "success": true,
  "data": {},
  "message": "ok"
}
```

错误响应：

```json
{
  "success": false,
  "data": null,
  "message": "ERROR_CODE",
  "error": {
    "code": "ERROR_CODE",
    "details": null
  }
}
```
