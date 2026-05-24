# AI Project OS MVP 后端开发任务拆解

> 来源：`AI_PROJECT_OS_MVP_ARCHITECTURE.md`  
> 范围：仅拆解 MVP 后端开发任务，不包含前端页面实现。  
> 目标：基于 Go Fiber + GORM + PostgreSQL + Redis，提供稳定、可扩展、可鉴权的 REST API，支撑 AI Project OS MVP 主链路。

## 1. 后端总体目标

后端需要支撑 MVP 主链路：

1. 用户注册、登录、JWT 鉴权。
2. 用户创建和管理 AI Coding 项目。
3. 系统根据项目类型、技术栈、项目上下文生成 Prompt。
4. 系统根据项目上下文生成 Cursor Rules、`CLAUDE.md`、`AGENTS.md` 等规则文件。
5. 用户可以保存、编辑、复制、导出生成内容。
6. 用户可以维护 Project Brain 上下文文档。
7. 系统记录 Prompt / Rules 生成行为和 Prompt 历史版本。
8. 为后续真实 AI Provider、Context Manager、Agent Workflow 预留架构空间。

## 2. 后端技术栈任务

### 2.1 应用基础设施

- [ ] 初始化 `apps/api` Go 后端服务。
- [ ] 使用 Go 1.22+。
- [ ] 使用 Fiber 搭建 HTTP 服务。
- [ ] 使用 GORM 作为 ORM。
- [ ] 使用 PostgreSQL 作为主数据库。
- [ ] 使用 Redis 预留缓存 / 会话 / 限流能力。
- [ ] 使用 Viper 管理配置。
- [ ] 使用 Zap 或 Zerolog 管理日志。
- [ ] 配置 `.env.example`。
- [ ] 配置 Docker Compose 启动 PostgreSQL 和 Redis。
- [ ] 提供健康检查接口。
- [ ] 编写基础 README 后端启动说明。

### 2.2 后端目录结构

按照架构文档规划建立目录：

```text
apps/api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── ai/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── services/
│   ├── utils/
│   └── validators/
├── migrations/
├── go.mod
└── go.sum
```

开发任务：

- [ ] 创建 `cmd/server/main.go` 作为服务入口。
- [ ] 创建 `internal/config` 配置加载模块。
- [ ] 创建 `internal/database` 数据库初始化模块。
- [ ] 创建 `internal/routes` 路由注册模块。
- [ ] 创建 `internal/handlers` HTTP Handler 层。
- [ ] 创建 `internal/services` 业务服务层。
- [ ] 创建 `internal/repositories` 数据访问层。
- [ ] 创建 `internal/models` GORM 模型层。
- [ ] 创建 `internal/middleware` 中间件层。
- [ ] 创建 `internal/validators` 请求校验层。
- [ ] 创建 `internal/utils` 通用工具。
- [ ] 创建 `internal/ai` AI Provider 抽象层。
- [ ] 创建 `migrations` 数据库迁移目录。

## 3. 后端分层架构任务

后端采用分层架构：

```text
HTTP Request
  ↓
Routes
  ↓
Middleware
  ↓
Handler
  ↓
Service
  ↓
Repository
  ↓
GORM / PostgreSQL
```

### 3.1 Handler 层

- [ ] 解析 HTTP 请求参数。
- [ ] 获取当前登录用户。
- [ ] 调用 Service 层。
- [ ] 返回统一 JSON 响应。
- [ ] 不写业务逻辑。
- [ ] 不直接访问数据库。

### 3.2 Service 层

- [ ] 实现业务流程。
- [ ] 实现权限校验。
- [ ] 实现 Prompt / Rules 生成流程。
- [ ] 实现 Prompt 版本保存。
- [ ] 调用 AI Provider。
- [ ] 调用 Repository 完成数据读写。

### 3.3 Repository 层

- [ ] 负责数据库读写。
- [ ] 不处理业务规则。
- [ ] 不直接处理 HTTP 上下文。
- [ ] 所有涉及用户数据的查询必须带 `user_id`。
- [ ] 所有项目相关数据查询必须带 `project_id`。

### 3.4 Model 层

- [ ] 定义 GORM 模型。
- [ ] 管理数据表结构。
- [ ] 定义模型关联关系。
- [ ] 保持字段和架构文档一致。

### 3.5 AI 层

- [ ] 定义统一 AI Provider 接口。
- [ ] 实现 Mock Provider。
- [ ] 预留 OpenAI Provider。
- [ ] 预留 Claude Provider。
- [ ] 预留 Gemini Provider。
- [ ] 预留 GLM Provider。

## 4. 通用能力任务

## 4.1 统一响应格式

成功响应：

```json
{
  "success": true,
  "data": {},
  "message": "ok",
  "error": null
}
```

失败响应：

```json
{
  "success": false,
  "data": null,
  "message": "validation failed",
  "error": {
    "code": "INVALID_INPUT",
    "details": {}
  }
}
```

开发任务：

- [ ] 定义统一响应结构。
- [ ] 定义统一错误结构。
- [ ] 封装成功响应方法。
- [ ] 封装失败响应方法。
- [ ] 统一处理参数校验错误。
- [ ] 统一处理业务错误。
- [ ] 统一处理未授权错误。
- [ ] 统一处理资源不存在错误。

## 4.2 配置管理

- [ ] 使用 Viper 加载配置。
- [ ] 支持 `.env`。
- [ ] 支持环境变量覆盖。
- [ ] 配置服务端口。
- [ ] 配置 PostgreSQL 连接。
- [ ] 配置 Redis 连接。
- [ ] 配置 JWT Secret。
- [ ] 配置 AI Provider 默认类型和模型。
- [ ] 配置日志级别。

必须满足：

- [ ] JWT Secret 必须来自环境变量。
- [ ] 不允许硬编码敏感信息。

## 4.3 日志与错误处理

- [ ] 接入 Zap 或 Zerolog。
- [ ] 记录请求日志。
- [ ] 记录业务错误。
- [ ] 记录 AI 生成错误。
- [ ] 记录数据库错误。
- [ ] 避免日志泄露密码、Token、API Key。
- [ ] 提供全局 recover 中间件。

## 4.4 中间件

- [ ] 实现 Request ID 中间件。
- [ ] 实现 CORS 中间件。
- [ ] 实现 Logger 中间件。
- [ ] 实现 Recover 中间件。
- [ ] 实现 JWT Auth 中间件。
- [ ] 预留 Rate Limit 中间件。

## 5. 数据库与模型任务

## 5.1 数据库连接和迁移

- [ ] 配置 PostgreSQL 连接。
- [ ] 配置 Redis 连接。
- [ ] 实现数据库初始化。
- [ ] 实现自动迁移或迁移脚本。
- [ ] 创建数据库索引。
- [ ] 编写数据库文档。

核心表：

- `users`
- `projects`
- `prompts`
- `prompt_versions`
- `project_docs`
- `templates`
- `ai_generations`

设计要求：

- [ ] 所有用户数据必须包含 `user_id`。
- [ ] 所有项目相关数据必须包含 `project_id`。
- [ ] Prompt 编辑和生成必须写入 `prompt_versions`。
- [ ] AI 生成行为必须写入 `ai_generations`。

## 5.2 User 模型

字段：

- `id`
- `email`
- `username`
- `password_hash`
- `avatar`
- `role`
- `created_at`
- `updated_at`

开发任务：

- [ ] 定义 User GORM 模型。
- [ ] 添加邮箱唯一索引。
- [ ] 添加用户名字段。
- [ ] 添加密码 hash 字段。
- [ ] 添加角色字段。
- [ ] 实现 User Repository。

## 5.3 Project 模型

字段：

- `id`
- `user_id`
- `name`
- `description`
- `project_type`
- `frontend_stack`
- `backend_stack`
- `database_stack`
- `ai_stack`
- `ui_style`
- `target_user`
- `product_goal`
- `status`
- `created_at`
- `updated_at`

开发任务：

- [ ] 定义 Project GORM 模型。
- [ ] 建立 User 与 Project 关联。
- [ ] 添加 `user_id` 索引。
- [ ] 添加 `project_type` 索引。
- [ ] 实现 Project Repository。
- [ ] 实现项目类型枚举。
- [ ] 实现项目状态枚举。

项目类型：

- SaaS。
- AI Chat App。
- AI Agent App。
- Admin Dashboard。
- E-commerce。
- Blog CMS。
- Landing Page。
- Mobile App。
- Developer Tool。
- Internal System。

## 5.4 Prompt 模型

字段：

- `id`
- `project_id`
- `user_id`
- `type`
- `title`
- `content`
- `version`
- `created_at`
- `updated_at`

开发任务：

- [ ] 定义 Prompt GORM 模型。
- [ ] 建立 Project 与 Prompt 关联。
- [ ] 建立 User 与 Prompt 关联。
- [ ] 添加 `project_id`、`user_id`、`type` 索引。
- [ ] 实现 Prompt Repository。
- [ ] 实现 Prompt 类型枚举。

Prompt 类型：

- `backend_prompt`
- `frontend_prompt`
- `fullstack_prompt`
- `ui_design_prompt`
- `database_prompt`
- `api_design_prompt`
- `testing_prompt`
- `deployment_prompt`
- `cursor_rules`
- `claude_md`
- `agents_md`

## 5.5 PromptVersion 模型

字段：

- `id`
- `prompt_id`
- `version`
- `content`
- `change_note`
- `created_at`

开发任务：

- [ ] 定义 PromptVersion GORM 模型。
- [ ] 建立 Prompt 与 PromptVersion 关联。
- [ ] 添加 `prompt_id` 索引。
- [ ] 实现 PromptVersion Repository。
- [ ] 实现版本号递增逻辑。
- [ ] 预留版本回滚接口能力。
- [ ] MVP 可暂不实现复杂 diff。

## 5.6 ProjectDoc 模型

字段：

- `id`
- `project_id`
- `user_id`
- `doc_type`
- `title`
- `content`
- `created_at`
- `updated_at`

开发任务：

- [ ] 定义 ProjectDoc GORM 模型。
- [ ] 建立 Project 与 ProjectDoc 关联。
- [ ] 建立 User 与 ProjectDoc 关联。
- [ ] 添加 `project_id`、`user_id`、`doc_type` 索引。
- [ ] 实现 ProjectDoc Repository。
- [ ] 实现文档类型枚举。

文档类型：

- `architecture`
- `api_docs`
- `database_schema`
- `ui_guidelines`
- `coding_standards`
- `product_requirements`
- `changelog`
- `decisions`

## 5.7 Template 模型

字段：

- `id`
- `name`
- `description`
- `project_type`
- `frontend_stack`
- `backend_stack`
- `database_stack`
- `ai_stack`
- `ui_style`
- `default_prompts`
- `default_rules`

开发任务：

- [ ] 定义 Template GORM 模型。
- [ ] 实现 Template Repository。
- [ ] 添加内置模板种子数据。
- [ ] 支持根据模板创建项目。

内置模板：

- Go + Next.js SaaS。
- Python + Vue Admin。
- Go + Vue E-commerce。
- Next.js AI Chat。
- Go + React Dashboard。
- AI Agent Platform。
- Landing Page Builder。

## 5.8 AIGeneration 模型

字段：

- `id`
- `user_id`
- `project_id`
- `generation_type`
- `input_payload`
- `output_content`
- `model_provider`
- `model_name`
- `status`
- `error_message`
- `created_at`

开发任务：

- [ ] 定义 AIGeneration GORM 模型。
- [ ] 添加 `user_id`、`project_id`、`generation_type` 索引。
- [ ] 实现 AIGeneration Repository。
- [ ] Prompt 生成时记录生成行为。
- [ ] Rules 生成时记录生成行为。
- [ ] 失败时记录错误信息。

## 6. 核心业务模块任务

## 6.1 用户认证模块

功能：

- 注册。
- 登录。
- 退出。
- JWT 鉴权。
- 获取用户信息。
- 修改密码，MVP 可保留接口或后续实现。

接口：

```text
POST /api/auth/register
POST /api/auth/login
GET  /api/auth/me
```

开发任务：

- [ ] 实现注册请求 DTO 和校验。
- [ ] 实现登录请求 DTO 和校验。
- [ ] 实现 User Service。
- [ ] 实现邮箱唯一性检查。
- [ ] 实现密码安全 hash。
- [ ] 实现密码校验。
- [ ] 实现 JWT 签发。
- [ ] 实现 JWT 校验。
- [ ] 实现 Auth Middleware。
- [ ] 实现获取当前用户接口。
- [ ] 预留修改密码接口。

鉴权要求：

- [ ] 除注册和登录外，所有接口默认需要 JWT 鉴权。
- [ ] JWT Secret 必须来自环境变量。
- [ ] 密码必须使用安全 hash。
- [ ] 用户不能访问其他用户的数据。

验收标准：

- [ ] 用户可以注册。
- [ ] 用户可以登录。
- [ ] 登录后可以访问 `/api/auth/me`。
- [ ] 未登录不能访问受保护接口。

## 6.2 项目模块

接口：

```text
GET    /api/projects
POST   /api/projects
GET    /api/projects/:id
PUT    /api/projects/:id
DELETE /api/projects/:id
```

开发任务：

- [ ] 实现项目创建请求 DTO 和校验。
- [ ] 实现项目更新请求 DTO 和校验。
- [ ] 实现 Project Service。
- [ ] 实现项目创建。
- [ ] 实现项目列表。
- [ ] 实现项目详情。
- [ ] 实现项目更新。
- [ ] 实现项目删除。
- [ ] 实现项目搜索。
- [ ] 实现按项目类型筛选。
- [ ] 实现项目权限校验。
- [ ] 确保所有项目查询都校验 `user_id`。

验收标准：

- [ ] 用户可以创建、查看、编辑、删除自己的项目。
- [ ] 用户不能访问其他用户项目。
- [ ] 项目类型枚举可用。

## 6.3 Prompt 生成模块

接口：

```text
GET    /api/projects/:projectId/prompts
POST   /api/projects/:projectId/prompts/generate
GET    /api/prompts/:id
PUT    /api/prompts/:id
DELETE /api/prompts/:id
POST   /api/prompts/:id/duplicate
GET    /api/prompts/:id/versions
```

开发任务：

- [ ] 实现 Prompt Service。
- [ ] 实现 Prompt 生成请求 DTO 和校验。
- [ ] 实现 Prompt 更新请求 DTO 和校验。
- [ ] 实现项目上下文读取。
- [ ] 实现 Project Docs 读取。
- [ ] 实现 Context Assembler。
- [ ] 实现 Template Renderer。
- [ ] 调用 AI Provider 生成内容。
- [ ] 保存生成结果到 Prompt。
- [ ] 保存生成版本到 PromptVersion。
- [ ] 记录生成行为到 AIGeneration。
- [ ] 实现 Prompt 列表。
- [ ] 实现 Prompt 详情。
- [ ] 实现 Prompt 编辑。
- [ ] 实现 Prompt 删除。
- [ ] 实现 Prompt 复制。
- [ ] 实现 Prompt 历史版本列表。
- [ ] 实现 Prompt 权限校验。

生成内容必须包含：

- 角色设定。
- 项目背景。
- 技术栈。
- 功能模块。
- 数据库设计要求。
- API 设计要求。
- 前端页面要求。
- 代码规范。
- 安全要求。
- 测试要求。
- 输出要求。
- 禁止事项。

验收标准：

- [ ] 可以基于项目生成不同类型 Prompt。
- [ ] 生成结果是结构化 Markdown。
- [ ] 编辑后会保存新版本。
- [ ] 可以查看历史版本。

## 6.4 Rules Builder 模块

接口：

```text
POST /api/projects/:projectId/rules/generate
GET  /api/projects/:projectId/rules
```

开发任务：

- [ ] 实现 Rules Service。
- [ ] 实现 Rules 生成请求 DTO 和校验。
- [ ] 实现 Cursor Rules 模板。
- [ ] 实现 `CLAUDE.md` 模板。
- [ ] 实现 `AGENTS.md` 模板。
- [ ] 实现 `frontend-rules.md` 模板。
- [ ] 实现 `backend-rules.md` 模板。
- [ ] 实现 `ui-rules.md` 模板。
- [ ] 实现 `testing-rules.md` 模板。
- [ ] 支持从项目基础信息组装规则上下文。
- [ ] 支持从 Project Docs 组装规则上下文。
- [ ] 保存 Rules 结果，可复用 Prompt 表或单独内部逻辑。
- [ ] 记录生成行为到 AIGeneration。
- [ ] 实现 Rules 查询接口。
- [ ] 实现项目权限校验。

生成内容必须包括：

- 项目目标。
- 技术栈。
- 目录结构规范。
- 命名规范。
- API 规范。
- 数据库规范。
- 错误处理规范。
- 安全规范。
- UI 规范。
- 测试规范。
- 禁止事项。
- Agent 开发流程。

验收标准：

- [ ] 可以生成 Cursor、Claude、Agents 三类核心规则。
- [ ] 规则内容包含项目上下文、技术栈、目录规范、命名规范、Agent 工作方式和禁止事项。
- [ ] 生成结果可以被前端复制和导出。

## 6.5 Project Brain 模块

接口：

```text
GET    /api/projects/:projectId/docs
POST   /api/projects/:projectId/docs
GET    /api/docs/:id
PUT    /api/docs/:id
DELETE /api/docs/:id
```

开发任务：

- [ ] 实现 ProjectDoc Service。
- [ ] 实现文档创建请求 DTO 和校验。
- [ ] 实现文档更新请求 DTO 和校验。
- [ ] 实现文档列表。
- [ ] 实现文档详情。
- [ ] 实现文档创建。
- [ ] 实现文档编辑。
- [ ] 实现文档删除。
- [ ] 实现按文档类型筛选。
- [ ] 实现文档权限校验。
- [ ] 支持 Prompt 生成读取项目上下文文档。
- [ ] 支持 Rules 生成读取项目上下文文档。

验收标准：

- [ ] 用户可以创建架构文档、API 文档、数据库文档、产品需求等。
- [ ] 文档可以编辑和删除。
- [ ] Prompt / Rules 生成时可以读取项目上下文文档。

## 6.6 模板模块

接口：

```text
GET /api/templates
GET /api/templates/:id
```

开发任务：

- [ ] 实现 Template Service。
- [ ] 实现模板列表接口。
- [ ] 实现模板详情接口。
- [ ] 实现内置模板种子数据。
- [ ] 支持项目创建时引用模板。
- [ ] 模板中包含默认 Prompts 和默认 Rules。

验收标准：

- [ ] 用户可以获取模板列表。
- [ ] 用户可以查看模板详情。
- [ ] 用户可以从模板创建项目。

## 6.7 Dashboard 统计模块

架构文档未明确列出独立 Dashboard API，但第七阶段要求实现 Dashboard 统计接口。

建议接口：

```text
GET /api/dashboard/summary
```

建议返回：

- 项目数量。
- Prompt 数量。
- 最近项目。
- 最近 Prompt。
- 最近 AI 生成记录。

开发任务：

- [ ] 实现 Dashboard Service。
- [ ] 实现项目数量统计。
- [ ] 实现 Prompt 数量统计。
- [ ] 实现最近项目查询。
- [ ] 实现最近 Prompt 查询。
- [ ] 实现最近 AI 生成记录查询。
- [ ] 所有统计仅限当前用户数据。

验收标准：

- [ ] Dashboard 可以展示项目数量、最近项目、最近 Prompt。

## 7. AI Provider 与生成流程任务

## 7.1 AI Provider 接口

接口设计：

```go
type AIProvider interface {
    Generate(ctx context.Context, req GenerateRequest) (*GenerateResponse, error)
}
```

GenerateRequest：

- `Provider`
- `Model`
- `SystemPrompt`
- `UserPrompt`
- `Temperature`
- `MaxTokens`

GenerateResponse：

- `Content`
- `Usage`
- `Raw`

开发任务：

- [ ] 定义 `AIProvider` 接口。
- [ ] 定义 `GenerateRequest`。
- [ ] 定义 `GenerateResponse`。
- [ ] 实现 Provider Factory。
- [ ] 实现 MockProvider。
- [ ] 预留 OpenAIProvider。
- [ ] 预留 ClaudeProvider。
- [ ] 预留 GeminiProvider。
- [ ] 预留 GLMProvider。

MVP 要求：

- [ ] MVP 使用 Mock Provider。
- [ ] Mock Provider 返回结构化 Markdown。
- [ ] 不要求真实模型多供应商完整接入。

## 7.2 Prompt / Rules 生成流程

生成流程：

```text
Project + Docs
  ↓
Context Assembler
  ↓
Template Renderer
  ↓
AI Provider Factory
  ↓
Mock Provider
  ↓
Generation Recorder
  ↓
Prompt / Rules / Version
```

开发任务：

- [ ] 实现 Context Assembler。
- [ ] 聚合项目基础信息。
- [ ] 聚合项目技术栈。
- [ ] 聚合 Project Brain 文档。
- [ ] 渲染 Prompt 模板。
- [ ] 渲染 Rules 模板。
- [ ] 调用 AI Provider。
- [ ] 记录生成输入和输出。
- [ ] 生成成功后保存 Prompt / Rules。
- [ ] 生成失败后保存失败记录。

## 8. REST API 总清单

## 8.1 认证

```text
POST /api/auth/register
POST /api/auth/login
GET  /api/auth/me
```

## 8.2 项目

```text
GET    /api/projects
POST   /api/projects
GET    /api/projects/:id
PUT    /api/projects/:id
DELETE /api/projects/:id
```

## 8.3 Prompt

```text
GET    /api/projects/:projectId/prompts
POST   /api/projects/:projectId/prompts/generate
GET    /api/prompts/:id
PUT    /api/prompts/:id
DELETE /api/prompts/:id
POST   /api/prompts/:id/duplicate
GET    /api/prompts/:id/versions
```

## 8.4 Rules

```text
POST /api/projects/:projectId/rules/generate
GET  /api/projects/:projectId/rules
```

## 8.5 Project Brain

```text
GET    /api/projects/:projectId/docs
POST   /api/projects/:projectId/docs
GET    /api/docs/:id
PUT    /api/docs/:id
DELETE /api/docs/:id
```

## 8.6 Templates

```text
GET /api/templates
GET /api/templates/:id
```

## 8.7 建议补充接口

```text
GET /api/health
GET /api/dashboard/summary
```

## 9. 后端开发阶段拆解

## 阶段一：基础设施

- [ ] 初始化 Go Fiber API。
- [ ] 配置 GORM。
- [ ] 配置 PostgreSQL。
- [ ] 配置 Redis。
- [ ] 配置 Viper。
- [ ] 配置日志。
- [ ] 配置 Docker Compose。
- [ ] 配置 `.env.example`。
- [ ] 实现健康检查接口。
- [ ] 编写基础 README。

验收标准：

- [ ] 后端可以启动。
- [ ] PostgreSQL 和 Redis 可以通过 Docker 启动。
- [ ] 健康检查接口可访问。

## 阶段二：用户认证

- [ ] 实现 User 模型。
- [ ] 实现注册接口。
- [ ] 实现登录接口。
- [ ] 实现获取当前用户接口。
- [ ] 实现密码 hash。
- [ ] 实现 JWT 签发与校验。
- [ ] 实现 Auth Middleware。

验收标准：

- [ ] 用户可以注册。
- [ ] 用户可以登录。
- [ ] 登录后可以访问 `/api/auth/me`。
- [ ] 未登录不能访问受保护接口。

## 阶段三：项目模块

- [ ] 实现 Project 模型。
- [ ] 实现项目 CRUD API。
- [ ] 实现项目权限校验。
- [ ] 实现项目类型枚举。
- [ ] 实现搜索和筛选。

验收标准：

- [ ] 用户可以创建、查看、编辑、删除自己的项目。
- [ ] 用户不能访问其他用户项目。

## 阶段四：Prompt 生成模块

- [ ] 实现 Prompt 模型。
- [ ] 实现 PromptVersion 模型。
- [ ] 实现 AIGeneration 模型。
- [ ] 实现 AIProvider 接口。
- [ ] 实现 MockProvider。
- [ ] 实现 Prompt 生成服务。
- [ ] 实现 Prompt CRUD。
- [ ] 实现 Prompt 版本保存。

验收标准：

- [ ] 可以基于项目生成不同类型 Prompt。
- [ ] 生成结果是结构化 Markdown。
- [ ] 编辑后会保存新版本。
- [ ] 可以查看历史版本。

## 阶段五：Rules Builder

- [ ] 实现 Rules 生成服务。
- [ ] 实现 Cursor Rules 模板。
- [ ] 实现 `CLAUDE.md` 模板。
- [ ] 实现 `AGENTS.md` 模板。
- [ ] 实现 frontend/backend/ui/testing rules 模板。

验收标准：

- [ ] 可以生成 Cursor、Claude、Agents 三类核心规则。
- [ ] 规则内容包含项目上下文、技术栈、目录规范、命名规范、Agent 工作方式和禁止事项。

## 阶段六：Project Brain

- [ ] 实现 ProjectDoc 模型。
- [ ] 实现文档 CRUD API。
- [ ] 实现文档类型枚举。
- [ ] 实现基于文档重新生成 Prompt / Rules 的后端能力。

验收标准：

- [ ] 用户可以创建架构文档、API 文档、数据库文档、产品需求等。
- [ ] Prompt 生成时可以读取项目上下文文档。
- [ ] 文档可以编辑和删除。

## 阶段七：模板与 Dashboard

- [ ] 实现 Template 模型。
- [ ] 实现内置模板种子数据。
- [ ] 实现模板列表接口。
- [ ] 实现从模板创建项目。
- [ ] 实现 Dashboard 统计接口。

验收标准：

- [ ] 用户可以从模板创建项目。
- [ ] Dashboard 能展示项目数量、最近项目、最近 Prompt。

## 阶段八：完善与交付

- [ ] 完善错误处理。
- [ ] 完善 API 文档。
- [ ] 完善数据库文档。
- [ ] 完善架构文档。
- [ ] 完善开发文档。
- [ ] 完善 README 启动说明。
- [ ] 增加基础测试。
- [ ] 验证 Docker 启动流程。

验收标准：

- [ ] 新开发者可以根据 README 启动项目。
- [ ] 核心流程可以完整跑通。
- [ ] MVP 边界清晰。
- [ ] 后续路线清晰。

## 10. 后端质量要求

- [ ] 采用分层架构。
- [ ] Handler 不写业务逻辑。
- [ ] Service 负责业务。
- [ ] Repository 负责数据库。
- [ ] 所有 API 返回统一 JSON 格式。
- [ ] 所有错误统一处理。
- [ ] 除登录注册外，所有接口需要鉴权。
- [ ] 所有输入需要校验。
- [ ] 密码必须 hash。
- [ ] JWT Secret 从环境变量读取。
- [ ] 不硬编码敏感信息。
- [ ] 查询项目、Prompt、文档时必须校验 `user_id`。
- [ ] 用户不能访问其他用户的数据。
- [ ] Prompt 编辑和生成必须写入 `prompt_versions`。
- [ ] AI 生成行为必须写入 `ai_generations`。

## 11. MVP 后端暂不实现

- [ ] 真实多 Agent 执行。
- [ ] 复杂向量数据库。
- [ ] 团队协作权限系统。
- [ ] 支付系统。
- [ ] 浏览器插件接口。
- [ ] IDE 插件接口。
- [ ] 真实模型多供应商完整接入。
- [ ] 复杂 Prompt 版本 diff。

## 12. 推荐开发顺序

1. 后端基础设施和数据库连接。
2. 用户注册登录和 JWT 鉴权。
3. 项目 CRUD 和权限校验。
4. AI Provider 抽象层和 Mock Provider。
5. Prompt 生成、保存、版本管理。
6. Rules 生成。
7. Project Brain 文档管理。
8. 模板和 Dashboard 统计。
9. 错误处理、测试、文档和 Docker 启动验证。
