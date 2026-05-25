<div align="center">

# AI Project OS

**面向 AI 编程时代的开发者工作台**

管理 AI Coding 项目的 Prompt · Rules · 项目上下文 · 知识文档 · AI 生成记录

[快速开始](#-快速开始) · [核心功能](#-核心功能) · [API 概览](#-api-概览) · [架构文档](#-相关文档)

</div>

---

## 项目简介

AI Project OS 是一个面向 AI 原生开发流程的项目工作台，目标是把项目上下文、Prompt 工程、规则管理、知识文档和 AI 生成记录集中到同一个系统中，帮助开发者沉淀可复用的研发资产。

它不是普通 Prompt 平台，而是一个逐步演进的 **AI Developer Platform**。

### 产品演进方向

| 阶段 | 能力 | 状态 |
|------|------|------|
| V1 | Rules Builder + Prompt Studio | ✅ 已完成 MVP |
| V2 | Project Brain + Context Manager | 🔧 基础框架已搭建 |
| V3 | 代码仓库扫描 | 📋 规划中 |
| V4 | Multi-Agent Workflow | 📋 规划中 |
| V5 | 团队协作 | 📋 规划中 |
| V6 | 浏览器插件 / VSCode 插件 / Cursor 插件 | 📋 规划中 |
| V7 | AI Developer OS | 📋 规划中 |

### MVP 核心价值

> 帮助开发者为每个 AI Coding 项目沉淀一套可复用、可维护、可持续演进的项目上下文、Prompt 和 AI 编程规则。

当前仓库包含：

- **后端 API** — Go Fiber + GORM + MySQL，提供认证、项目、Prompt、文档、模板、AI 生成和 Dashboard 统计接口
- **前端应用** — Vue 3 + Vite + TypeScript + Pinia + Vue Router，提供 Landing、登录注册、工作台、项目详情、Prompt Studio、Rules Builder 和 Project Brain 页面
- **H5 原型** — 包含核心页面交互原型
- **规划文档** — 包含 MVP 架构说明、前后端任务拆解

---

## 📊 当前完成度

**整体状态：MVP 核心闭环已基本打通，处于可本地运行和继续迭代阶段**

| 模块 | 完成度 | 状态说明 |
|------|:------:|----------|
| 后端基础设施 | 🟢 高 | 配置加载、日志、全局中间件、数据库连接、GORM AutoMigrate、统一响应格式 |
| 用户认证 | 🟢 高 | 注册、登录、当前用户、登出；JWT HttpOnly Cookie + Bearer Token 兼容 |
| 项目管理 | 🟢 高 | 项目 CRUD，按用户归属做资源隔离，支持搜索和类型筛选 |
| Prompt 管理 | 🟢 高 | 项目下 Prompt CRUD、版本历史、用户/项目归属校验 |
| 项目文档 / Project Brain | 🟡 中高 | 项目文档 CRUD、页面基础展示和项目级数据加载 |
| AI 生成 | 🟡 中 | 生成接口、生成记录、Mock Provider；真实大模型 Provider 尚待接入 |
| 模板中心 | 🟡 中 | 模板查询接口和前端类型对接；数据维护/后台管理仍待完善 |
| Dashboard | 🟡 中高 | 统计、最近项目、AI 动态、错误提示和项目级跳转 |
| 前端 UI | 🟡 中高 | 主要页面和基础设计系统已完成；部分图谱/图表仍为静态展示 |
| 安全修复 | 🟢 高 | IDOR 风险已修复；认证 Cookie 化已完成 |
| 测试体系 | 🔴 低 | 暂无后端 `_test.go` 或前端业务测试文件 |
| 生产化能力 | 🟠 中低 | 本地开发链路完整；生产部署、CI/CD、监控告警仍待建设 |

---

## 🛠 技术栈

### 后端

| 技术 | 说明 |
|------|------|
| Go 1.22 | 后端语言 |
| Fiber v2 | HTTP 框架 |
| GORM | ORM |
| MySQL 8.0 | 主数据库 |
| Redis 7 | 缓存（基础设施已预留） |
| JWT HS256 | 认证方案，HttpOnly Cookie |
| Zap | 结构化日志 |
| godotenv | 配置管理 |

### 前端

| 技术 | 说明 |
|------|------|
| Vue 3 | 前端框架 |
| Vite 6 | 构建工具 |
| TypeScript | 类型安全 |
| Pinia | 状态管理 |
| Vue Router 4 | 路由 |
| Axios | HTTP 客户端 |
| TailwindCSS | 样式系统 + 自定义设计 Token |

---

## 📁 目录结构

```text
AI Project OS/
├── apps/
│   └── api/                              # Go 后端 API 服务
│       ├── cmd/server/main.go            # 后端入口
│       ├── internal/
│       │   ├── ai/                       # AI Provider 抽象与 Mock 实现
│       │   ├── config/                   # 环境变量与配置加载
│       │   ├── database/                 # MySQL 初始化与 AutoMigrate
│       │   ├── handlers/                 # HTTP Handler（9 个）
│       │   ├── logger/                   # Zap 日志初始化
│       │   ├── middleware/               # 全局中间件与 JWT 认证
│       │   ├── models/                   # GORM 数据模型（7 个）
│       │   ├── repositories/             # 数据访问层（6 个）
│       │   ├── routes/                   # 路由注册
│       │   ├── services/                 # 业务服务层（7 个）
│       │   └── utils/                    # 统一响应工具
│       ├── docker-compose.yml            # MySQL / Redis 本地依赖
│       ├── .env.example                  # 后端环境变量模板
│       └── README.md
├── frontend/                             # Vue 前端应用
│   ├── src/
│   │   ├── api/                          # API Client 与接口封装（8 个模块）
│   │   ├── assets/                       # 静态资源
│   │   ├── components/                   # Layout 与 UI 组件
│   │   │   ├── layout/                   # 布局组件
│   │   │   └── ui/                       # 基础 UI 组件
│   │   ├── composables/                  # 组合式函数
│   │   ├── router/                       # 前端路由
│   │   ├── stores/                       # Pinia 状态管理（3 个）
│   │   └── views/                        # 页面视图（7 个）
│   ├── package.json
│   ├── vite.config.ts
│   └── README.md
├── h5-prototype/                         # H5 交互原型
│   ├── index.html                        # Landing 原型
│   ├── dashboard.html                    # Dashboard 原型
│   ├── project-detail.html               # 项目详情原型
│   ├── prompt-studio.html                # Prompt Studio 原型
│   ├── rules-builder.html                # Rules Builder 原型
│   ├── project-brain.html                # Project Brain 原型
│   ├── app.js                            # 原型交互逻辑
│   └── styles.css                        # 原型样式
├── AI_PROJECT_OS_MVP_ARCHITECTURE.md     # MVP 架构设计文档
├── BACKEND_TASKS.md                      # 后端任务拆解
├── FRONTEND_TASKS.md                     # 前端任务拆解
└── README.md
```

---

## ✨ 核心功能

### 1. 用户认证

- 注册与登录
- 获取当前用户信息
- 登出清理 Cookie
- 认证态通过 `access_token` HttpOnly Cookie 传递
- 前端 `withCredentials` + Bearer Token 兼容
- 未登录访问受保护页面自动跳转登录页，保留 `redirect` 参数

### 2. 项目管理

- 创建项目（含技术栈、项目类型、目标用户等完整上下文）
- 项目列表（支持搜索和类型筛选）
- 项目详情、更新、删除
- 按用户归属做资源隔离，避免越权访问

### 3. Prompt Studio

- 项目级 Prompt 列表加载
- Prompt 运行调用 AI 生成接口
- 生成结果展示与编辑
- Prompt 版本历史管理
- 路由：`/project/:id/prompt-studio`

### 4. Rules Builder

- 项目级规则 Prompt 加载
- Markdown 风格规则内容编辑与预览
- 支持 Cursor Rules / CLAUDE.md / AGENTS.md 等规则类型
- 路由：`/project/:id/rules-builder`

### 5. Project Brain

- 项目级文档加载与 CRUD
- 文档选择与详情展示
- 知识图谱基础 UI
- 路由：`/project/:id/brain`

### 6. AI Generation

- 支持 `prompt` 和 `rules` 类型生成
- 当前 Provider 为 `mock`，用于本地闭环开发
- 生成记录保存项目 ID、用户 ID、模型信息、状态和输出
- 后端校验生成请求中的项目归属

### 7. Dashboard

- 展示项目数、Prompt 数、AI 生成数等统计
- 展示最近项目和 AI 动态
- 按生成记录跳转到对应项目页面
- 错误提示通过 Toast 展示

---

## 🚀 快速开始

### 前置要求

- **Go** 1.22+
- **Node.js** 18+ 或 20+
- **npm**
- **Docker Desktop**（用于本地 MySQL / Redis）

### 1. 启动后端依赖

```bash
cd apps/api
docker compose up -d
```

将启动：

- MySQL → `localhost:3306`
- Redis → `localhost:6379`

### 2. 配置后端环境变量

```bash
cd apps/api
cp .env.example .env
```

默认 `.env.example` 已适配本地 Docker Compose：

```env
SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=aiproject_dev
DB_NAME=ai_project_os
REDIS_HOST=localhost
REDIS_PORT=6379
JWT_SECRET=change-me-in-production
JWT_EXPIRATION_HOURS=72
AI_PROVIDER=mock
AI_MODEL=mock-v1
LOG_LEVEL=info
```

> ⚠️ **生产环境必须修改 `JWT_SECRET`，不要提交真实密钥。**

### 3. 启动后端 API

```bash
cd apps/api
go mod tidy
go run cmd/server/main.go
```

后端默认运行在 `http://localhost:8080`

```bash
# 验证健康检查
curl http://localhost:8080/health
```

### 4. 启动前端应用

```bash
cd frontend
npm install
npm run dev
```

前端默认运行在 `http://localhost:3000`

Vite 已配置 `/api` 代理到 `http://localhost:8080`，前端 API baseURL 为 `/api/v1`。

---

## 📋 常用命令

### 后端

```bash
cd apps/api

# 启动本地依赖
docker compose up -d

# 启动 API 服务
go run cmd/server/main.go

# 编译检查
go build ./...
```

### 前端

```bash
cd frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 类型检查 + 生产构建
npm run build

# 预览生产构建
npm run preview
```

---

## 📡 API 概览

后端 API 前缀：`/api/v1`

### 认证

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|:----:|
| POST | `/api/v1/auth/register` | 注册 | 否 |
| POST | `/api/v1/auth/login` | 登录 | 否 |
| POST | `/api/v1/auth/logout` | 登出 | 否 |
| GET | `/api/v1/auth/me` | 当前用户 | 是 |

### 项目

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|:----:|
| GET | `/api/v1/projects` | 项目列表 | 是 |
| POST | `/api/v1/projects` | 创建项目 | 是 |
| GET | `/api/v1/projects/:id` | 项目详情 | 是 |
| PUT | `/api/v1/projects/:id` | 更新项目 | 是 |
| DELETE | `/api/v1/projects/:id` | 删除项目 | 是 |

### Prompt

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|:----:|
| GET | `/api/v1/projects/:projectId/prompts` | Prompt 列表 | 是 |
| POST | `/api/v1/projects/:projectId/prompts` | 创建 Prompt | 是 |
| GET | `/api/v1/projects/:projectId/prompts/:id` | Prompt 详情 | 是 |
| PUT | `/api/v1/projects/:projectId/prompts/:id` | 更新 Prompt | 是 |
| DELETE | `/api/v1/projects/:projectId/prompts/:id` | 删除 Prompt | 是 |
| GET | `/api/v1/projects/:projectId/prompts/:id/versions` | 版本历史 | 是 |

### 项目文档

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|:----:|
| GET | `/api/v1/projects/:projectId/docs` | 文档列表 | 是 |
| POST | `/api/v1/projects/:projectId/docs` | 创建文档 | 是 |
| GET | `/api/v1/projects/:projectId/docs/:docId` | 文档详情 | 是 |
| PUT | `/api/v1/projects/:projectId/docs/:docId` | 更新文档 | 是 |
| DELETE | `/api/v1/projects/:projectId/docs/:docId` | 删除文档 | 是 |

### 其他

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|:----:|
| GET | `/health` | 健康检查 | 否 |
| GET | `/api/v1/templates` | 模板列表 | 否 |
| GET | `/api/v1/templates/:id` | 模板详情 | 否 |
| POST | `/api/v1/generations` | AI 生成 | 是 |
| GET | `/api/v1/generations` | 生成记录 | 是 |
| GET | `/api/v1/dashboard/stats` | Dashboard 统计 | 是 |

### 统一响应格式

**成功：**

```json
{
  "success": true,
  "data": {},
  "message": "ok"
}
```

**失败：**

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

---

## 🖥 前端路由

| 路径 | 页面 | 说明 |
|------|------|------|
| `/` | Landing | 产品首页 |
| `/login` | Auth | 登录 / 注册 |
| `/dashboard` | Dashboard | 工作台 |
| `/project/:id` | Project Detail | 项目详情 |
| `/project/:id/prompt-studio` | Prompt Studio | 项目级 Prompt 工作室 |
| `/project/:id/rules-builder` | Rules Builder | 项目级规则构建器 |
| `/project/:id/brain` | Project Brain | 项目级知识与文档页面 |

---

## 🔒 安全与权限

### 已完成

- 受保护 API 统一经过 JWT 认证
- JWT HttpOnly Cookie，降低前端 token 泄露风险
- 项目、Prompt、文档、AI 生成记录均按用户做归属校验
- 前端不再依赖 localStorage 持久化 token
- 登录页不渲染业务侧边栏

### 建议补充

- CSRF 防护（SameSite 策略 / CSRF Token / 双提交 Cookie）
- 生产环境 Cookie `Secure` 配置
- API Rate Limit
- 更细粒度角色权限（管理员后台、模板管理）
- 审计日志

---

## 🗺 Roadmap

### 近期优先

- [ ] **补充测试体系** — 后端单元测试、前端组件测试、API 集成测试
- [ ] **接入真实 AI Provider** — OpenAI / Claude / Gemini / GLM，API Key 环境变量注入
- [ ] **增强 Project Brain** — 文档上传/解析、关系抽取、知识检索

### 中期目标

- [ ] **完善模板中心** — 种子数据、管理员 CRUD、按项目类型推荐
- [ ] **生产化部署** — Dockerfile 编排、CI/CD、监控告警
- [ ] **Redis 业务缓存** — 热点数据缓存、会话管理、限流计数

### 远期愿景

- [ ] 代码仓库扫描
- [ ] Multi-Agent Workflow
- [ ] 团队协作
- [ ] 浏览器 / VSCode / Cursor 插件
- [ ] AI Developer OS

---

## 🤝 参与贡献

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/your-feature`
3. 提交更改：`git commit -m 'Add some feature'`
4. 推送分支：`git push origin feature/your-feature`
5. 提交 Pull Request

### 开发规范

- 后端采用分层架构：Handler → Service → Repository，Handler 不写业务逻辑
- 所有 API 返回统一 JSON 格式
- 所有输入需要校验，密码必须 hash，JWT Secret 从环境变量读取
- 前端请求统一封装，表单必须有校验，重要操作需要 Toast 提示
- 不硬编码敏感信息，不提交 `.env` 文件

---

## 📄 相关文档

| 文档 | 说明 |
|------|------|
| [`AI_PROJECT_OS_MVP_ARCHITECTURE.md`](./AI_PROJECT_OS_MVP_ARCHITECTURE.md) | MVP 架构设计 |
| [`BACKEND_TASKS.md`](./BACKEND_TASKS.md) | 后端任务拆解 |
| [`FRONTEND_TASKS.md`](./FRONTEND_TASKS.md) | 前端任务拆解 |
| [`apps/api/README.md`](./apps/api/README.md) | 后端 API 说明 |
| [`frontend/README.md`](./frontend/README.md) | 前端应用说明 |

---

## 📜 License

本项目仅供学习和开发参考使用，暂未指定开源协议。
