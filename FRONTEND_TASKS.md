# AI Project OS MVP 前端开发任务拆解

> 来源：`AI_PROJECT_OS_MVP_ARCHITECTURE.md`  
> 范围：仅拆解 MVP 前端开发任务，不包含后端接口实现。  
> 目标：基于后端 REST API，完成可运行、可演示、可继续扩展的 AI Project OS Web 应用。

## 1. 前端总体目标

前端需要打通 MVP 主链路：

1. 用户注册 / 登录。
2. 用户进入 Dashboard 查看项目与最近生成内容。
3. 用户创建 AI Coding 项目并填写上下文。
4. 用户进入项目详情页管理 Prompt、Rules、Project Brain。
5. 用户在 Prompt Studio 中生成、编辑、保存、复制、导出 Prompt。
6. 用户在 Rules Builder 中生成、复制、导出 Cursor Rules、`CLAUDE.md`、`AGENTS.md` 等规则文件。
7. 用户在 Project Brain 中创建、编辑、删除项目上下文文档，并基于文档重新生成 Prompt / Rules。
8. 为后续 Context Manager、Agent Workflow、团队协作预留页面和状态结构。

## 2. 前端技术栈任务

### 2.1 应用基础设施

- [ ] 初始化 `apps/web` 前端应用。
- [ ] 使用 Next.js App Router 搭建页面路由。
- [ ] 配置 TypeScript。
- [ ] 配置 TailwindCSS。
- [ ] 配置 shadcn/ui。
- [ ] 配置全局样式、主题变量、字体和基础布局。
- [ ] 配置 `.env.local.example`，包含 API Base URL 等前端环境变量。
- [ ] 配置 ESLint、Prettier、TypeScript 检查脚本。
- [ ] 配置前端启动、构建、检查命令。

### 2.2 前端目录结构

按照架构文档规划建立目录：

```text
apps/web/
├── app/
├── components/
│   ├── ui/
│   ├── layout/
│   └── editor/
├── features/
│   ├── auth/
│   ├── projects/
│   ├── prompts/
│   ├── rules/
│   ├── brain/
│   └── templates/
├── hooks/
├── lib/
├── public/
├── store/
├── styles/
├── types/
├── package.json
└── next.config.ts
```

开发任务：

- [ ] 创建 `components/layout`：AppShell、Sidebar、Topbar、PageHeader、Breadcrumb。
- [ ] 创建 `components/ui`：Button、Input、Textarea、Select、Card、Tabs、Modal、Toast、Badge、Dropdown、EmptyState、Skeleton。
- [ ] 创建 `components/editor`：MarkdownEditor、MonacoEditorWrapper、ReadonlyMarkdownPreview。
- [ ] 创建 `features/*` 业务模块目录。
- [ ] 创建 `lib/api` 请求封装目录。
- [ ] 创建 `types` 前端共享类型定义。
- [ ] 创建 `store` Zustand 状态目录。

## 3. 请求与状态管理任务

### 3.1 API 请求封装

- [ ] 封装统一 API Client。
- [ ] 支持 `GET`、`POST`、`PUT`、`DELETE`。
- [ ] 自动携带 JWT Token。
- [ ] 统一处理后端响应格式：

```json
{
  "success": true,
  "data": {},
  "message": "ok",
  "error": null
}
```

- [ ] 统一处理错误响应和 Toast 提示。
- [ ] 处理 401 未登录状态并跳转登录页。
- [ ] 提供文件导出 / Markdown 下载工具函数。

### 3.2 TanStack Query 数据管理

需要用 TanStack Query 管理：

- [ ] 当前用户信息。
- [ ] 项目列表。
- [ ] 项目详情。
- [ ] Prompt 列表。
- [ ] Prompt 详情。
- [ ] Prompt 历史版本。
- [ ] Rules 生成结果。
- [ ] Project Brain 文档列表。
- [ ] Project Brain 文档详情。
- [ ] 模板列表。
- [ ] Dashboard 统计数据。

### 3.3 Zustand 本地 UI 状态

需要用 Zustand 管理：

- [ ] Sidebar 展开 / 收起状态。
- [ ] 当前主题状态。
- [ ] 当前编辑器偏好。
- [ ] 当前选中的 Prompt 类型。
- [ ] 当前选中的 Rules 类型。
- [ ] Command Menu 打开状态。
- [ ] 创建项目向导临时状态。

## 4. 页面开发任务

## 4.1 Landing Page

目标：展示产品定位、核心价值和 CTA。

开发任务：

- [ ] 实现首页 Hero 区。
- [ ] 展示产品定位：AI Project OS / AI Developer Platform。
- [ ] 展示核心功能：Prompt Studio、Rules Builder、Project Brain、Context Manager、Agent Workflow。
- [ ] 展示 MVP 使用流程。
- [ ] 展示适合人群。
- [ ] 提供注册 / 登录 / 创建项目 CTA。
- [ ] 视觉风格保持 AI Native、高级开发者工具风。
- [ ] 支持响应式布局。

验收标准：

- [ ] 未登录用户可访问 Landing Page。
- [ ] CTA 能跳转到登录 / 注册页。
- [ ] 页面风格与产品定位一致。

## 4.2 登录 / 注册页

目标：完成前端认证入口。

开发任务：

- [ ] 实现登录页。
- [ ] 实现注册页。
- [ ] 使用 React Hook Form 管理表单。
- [ ] 使用 Zod 校验邮箱、密码、用户名。
- [ ] 登录成功后保存 Token。
- [ ] 注册成功后可自动登录或跳转登录页。
- [ ] 显示后端错误提示。
- [ ] 实现退出登录。
- [ ] 实现登录态路由保护。
- [ ] 实现 `/api/auth/me` 当前用户拉取。

依赖接口：

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`

验收标准：

- [ ] 用户可以注册。
- [ ] 用户可以登录。
- [ ] 登录后进入 Dashboard。
- [ ] 未登录用户不能访问受保护页面。
- [ ] 表单错误提示清晰。

## 4.3 Dashboard 首页

目标：让用户快速了解当前工作区状态。

开发任务：

- [ ] 实现 Dashboard 页面布局。
- [ ] 展示项目数量。
- [ ] 展示最近项目。
- [ ] 展示最近生成的 Prompt。
- [ ] 展示快捷入口：创建项目、Prompt Studio、Rules Builder、Project Brain。
- [ ] 展示模板入口。
- [ ] 展示用户信息和工作区信息。
- [ ] 加载状态使用 Skeleton。
- [ ] 空状态引导用户创建第一个项目。

依赖接口：

- `GET /api/projects`
- `GET /api/projects/:projectId/prompts`
- 后续可扩展 Dashboard 聚合统计接口。

验收标准：

- [ ] 用户登录后可以看到自己的项目摘要。
- [ ] 最近项目可以跳转项目详情。
- [ ] 快捷入口可跳转对应功能页。

## 4.4 项目列表页

目标：用户可以查看和管理自己的 AI Coding 项目。

开发任务：

- [ ] 实现项目列表页。
- [ ] 支持项目卡片 / 表格视图。
- [ ] 支持搜索项目。
- [ ] 支持按项目类型筛选。
- [ ] 支持创建项目入口。
- [ ] 支持编辑项目入口。
- [ ] 支持删除项目并二次确认。
- [ ] 支持空状态。
- [ ] 支持加载状态和错误状态。

依赖接口：

- `GET /api/projects`
- `DELETE /api/projects/:id`

验收标准：

- [ ] 用户只能看到自己的项目。
- [ ] 搜索和筛选交互可用。
- [ ] 删除项目有确认提示。

## 4.5 项目创建页 / 创建项目向导

目标：收集项目基础信息和技术栈，创建项目。

表单字段：

- 项目名称。
- 项目描述。
- 项目类型。
- 前端技术栈。
- 后端技术栈。
- 数据库。
- AI 模型。
- UI 风格。
- 目标用户。
- 产品目标。

开发任务：

- [ ] 实现项目创建页或创建项目弹窗向导。
- [ ] 使用 React Hook Form + Zod 校验。
- [ ] 项目类型使用架构文档枚举。
- [ ] 技术栈字段支持下拉选择或卡片选择。
- [ ] 支持从模板创建项目。
- [ ] 创建成功后跳转项目详情页。
- [ ] 创建失败时显示错误提示。

依赖接口：

- `POST /api/projects`
- `GET /api/templates`
- `GET /api/templates/:id`

验收标准：

- [ ] 用户可以创建项目。
- [ ] 必填字段校验可用。
- [ ] 创建成功后项目出现在列表中。

## 4.6 项目详情页

目标：作为项目级工作台入口。

Tabs：

- Overview。
- Prompts。
- Rules。
- Project Brain。
- Settings。

开发任务：

- [ ] 实现项目详情基础布局。
- [ ] 展示项目基础信息。
- [ ] 展示技术栈标签。
- [ ] 展示最近 Prompt。
- [ ] 展示最近 Rules。
- [ ] 展示最近 Project Brain 文档。
- [ ] 实现 Tabs 切换。
- [ ] Prompts Tab 展示项目 Prompt 列表。
- [ ] Rules Tab 展示当前项目 Rules 结果。
- [ ] Project Brain Tab 展示文档列表。
- [ ] Settings Tab 支持编辑项目基础信息。
- [ ] 支持删除项目。

依赖接口：

- `GET /api/projects/:id`
- `PUT /api/projects/:id`
- `DELETE /api/projects/:id`
- `GET /api/projects/:projectId/prompts`
- `GET /api/projects/:projectId/rules`
- `GET /api/projects/:projectId/docs`

验收标准：

- [ ] 项目详情信息正确展示。
- [ ] Tabs 内容可正常切换。
- [ ] 用户不能访问其他用户项目。

## 4.7 Prompt Studio 页面

目标：完成 Prompt 生成、编辑、保存、复制、导出、版本查看。

开发任务：

- [ ] 实现 Prompt Studio 页面。
- [ ] 支持选择项目。
- [ ] 支持选择 Prompt 类型：
  - `backend_prompt`
  - `frontend_prompt`
  - `fullstack_prompt`
  - `ui_design_prompt`
  - `database_prompt`
  - `api_design_prompt`
  - `testing_prompt`
  - `deployment_prompt`
- [ ] 支持输入补充需求。
- [ ] 调用接口生成 Prompt。
- [ ] 使用 Monaco Editor 或 Markdown Editor 展示生成内容。
- [ ] 支持编辑 Prompt 内容。
- [ ] 支持保存 Prompt。
- [ ] 支持复制 Prompt。
- [ ] 支持导出 Markdown。
- [ ] 支持查看 Prompt 历史版本。
- [ ] 支持复制 / Duplicate Prompt。
- [ ] 支持加载中、生成中、保存中状态。
- [ ] 支持生成失败重试。

依赖接口：

- `GET /api/projects/:projectId/prompts`
- `POST /api/projects/:projectId/prompts/generate`
- `GET /api/prompts/:id`
- `PUT /api/prompts/:id`
- `DELETE /api/prompts/:id`
- `POST /api/prompts/:id/duplicate`
- `GET /api/prompts/:id/versions`

验收标准：

- [ ] 可以基于项目生成不同类型 Prompt。
- [ ] 生成内容为结构化 Markdown。
- [ ] 编辑后可以保存。
- [ ] 保存后可以查看版本。
- [ ] 可以复制和导出 Markdown。

## 4.8 Rules Builder 页面

目标：完成 AI 编程规则文件生成、预览、复制、导出。

开发任务：

- [ ] 实现 Rules Builder 页面。
- [ ] 支持选择项目。
- [ ] 支持生成 Cursor Rules。
- [ ] 支持生成 `CLAUDE.md`。
- [ ] 支持生成 `AGENTS.md`。
- [ ] 支持生成 `frontend-rules.md`。
- [ ] 支持生成 `backend-rules.md`。
- [ ] 支持生成 `ui-rules.md`。
- [ ] 支持生成 `testing-rules.md`。
- [ ] 支持 Markdown 预览。
- [ ] 支持复制规则内容。
- [ ] 支持导出规则文件。
- [ ] 支持生成中状态和失败重试。

依赖接口：

- `POST /api/projects/:projectId/rules/generate`
- `GET /api/projects/:projectId/rules`

验收标准：

- [ ] 可以生成 Cursor、Claude、Agents 三类核心规则。
- [ ] 规则内容包含项目上下文、技术栈、目录规范、命名规范、Agent 工作方式和禁止事项。
- [ ] 可以复制和导出生成结果。

## 4.9 Project Brain 页面

目标：完成项目上下文文档管理。

文档类型：

- `architecture`
- `api_docs`
- `database_schema`
- `ui_guidelines`
- `coding_standards`
- `product_requirements`
- `changelog`
- `decisions`

开发任务：

- [ ] 实现 Project Brain 页面。
- [ ] 展示文档列表。
- [ ] 支持按文档类型筛选。
- [ ] 支持创建上下文文档。
- [ ] 支持编辑上下文文档。
- [ ] 支持删除上下文文档并二次确认。
- [ ] 支持 Markdown 编辑器。
- [ ] 支持 Markdown 预览。
- [ ] 支持从文档重新生成 Prompt 的入口。
- [ ] 支持从文档重新生成 Rules 的入口。
- [ ] 预留知识图谱 / Context Manager UI 扩展区域。

依赖接口：

- `GET /api/projects/:projectId/docs`
- `POST /api/projects/:projectId/docs`
- `GET /api/docs/:id`
- `PUT /api/docs/:id`
- `DELETE /api/docs/:id`
- `POST /api/projects/:projectId/prompts/generate`
- `POST /api/projects/:projectId/rules/generate`

验收标准：

- [ ] 用户可以创建架构文档、API 文档、数据库文档、产品需求等。
- [ ] 文档可以编辑和删除。
- [ ] Prompt / Rules 生成入口可以读取项目上下文文档。

## 4.10 模板页面 / 模板入口

目标：提升项目初始化效率。

开发任务：

- [ ] 实现模板列表 UI。
- [ ] 展示内置项目模板：
  - Go + Next.js SaaS。
  - Python + Vue Admin。
  - Go + Vue E-commerce。
  - Next.js AI Chat。
  - Go + React Dashboard。
  - AI Agent Platform。
  - Landing Page Builder。
- [ ] 支持查看模板详情。
- [ ] 支持从模板创建项目。
- [ ] 在 Dashboard 和创建项目页提供模板入口。

依赖接口：

- `GET /api/templates`
- `GET /api/templates/:id`
- `POST /api/projects`

验收标准：

- [ ] 用户可以浏览模板。
- [ ] 用户可以基于模板创建项目。

## 4.11 设置页

目标：提供基础个人设置和模型偏好设置入口。

开发任务：

- [ ] 实现个人信息设置 UI。
- [ ] 实现 API Key 设置 UI，MVP 只做界面。
- [ ] 实现模型偏好设置 UI。
- [ ] 预留修改密码入口。
- [ ] 预留主题切换入口。

依赖接口：

- `GET /api/auth/me`
- 修改密码接口 MVP 可后续实现。

验收标准：

- [ ] 页面 UI 可访问。
- [ ] MVP 范围内不要求真实 API Key 生效。

## 5. 前端开发阶段拆解

## 阶段一：基础设施

- [ ] 初始化 Next.js App Router。
- [ ] 配置 TailwindCSS。
- [ ] 配置 shadcn/ui。
- [ ] 配置全局布局和主题。
- [ ] 配置 API Client。
- [ ] 配置 TanStack Query。
- [ ] 配置 Zustand。
- [ ] 创建基础 UI 组件。
- [ ] 创建应用 Shell：Sidebar、Topbar、内容布局。

验收标准：

- [ ] 前端可以启动。
- [ ] 基础布局可访问。
- [ ] 可以请求后端健康检查接口。

## 阶段二：用户认证

- [ ] 实现登录页。
- [ ] 实现注册页。
- [ ] 实现 Token 存储。
- [ ] 实现请求拦截和 401 处理。
- [ ] 实现登录态路由保护。
- [ ] 实现获取当前用户信息。
- [ ] 实现退出登录。

验收标准：

- [ ] 用户可以注册和登录。
- [ ] 登录后可以访问 Dashboard。
- [ ] 未登录不能访问受保护页面。

## 阶段三：项目模块

- [ ] 实现项目列表页。
- [ ] 实现项目创建页 / 创建项目向导。
- [ ] 实现项目详情页基础框架。
- [ ] 实现项目编辑。
- [ ] 实现项目删除。
- [ ] 实现项目搜索和筛选。

验收标准：

- [ ] 用户可以创建、查看、编辑、删除自己的项目。
- [ ] 前端项目列表和详情页可用。

## 阶段四：Prompt Studio

- [ ] 实现 Prompt Studio 页面。
- [ ] 实现 Prompt 类型选择。
- [ ] 实现 Prompt 生成交互。
- [ ] 实现 Markdown / Monaco 编辑器。
- [ ] 实现 Prompt 保存。
- [ ] 实现 Prompt 复制。
- [ ] 实现 Markdown 导出。
- [ ] 实现版本列表。

验收标准：

- [ ] 可以基于项目生成不同类型 Prompt。
- [ ] 编辑后可以保存新版本。
- [ ] 可以查看历史版本。

## 阶段五：Rules Builder

- [ ] 实现 Rules Builder 页面。
- [ ] 实现规则类型切换。
- [ ] 实现 Rules 生成交互。
- [ ] 实现 Markdown 预览。
- [ ] 实现复制和导出。

验收标准：

- [ ] 可以生成 Cursor、Claude、Agents 三类核心规则。
- [ ] 生成结果可以复制和导出。

## 阶段六：Project Brain

- [ ] 实现 Project Brain 页面。
- [ ] 实现文档列表。
- [ ] 实现文档创建。
- [ ] 实现文档编辑。
- [ ] 实现文档删除。
- [ ] 实现文档类型筛选。
- [ ] 实现基于文档重新生成 Prompt / Rules 的入口。

验收标准：

- [ ] 用户可以维护项目上下文文档。
- [ ] Prompt / Rules 生成可以使用文档上下文入口。

## 阶段七：模板与 Dashboard

- [ ] 实现模板列表。
- [ ] 实现模板详情。
- [ ] 实现从模板创建项目。
- [ ] 完善 Dashboard 首页 UI。
- [ ] 展示项目数量、最近项目、最近 Prompt。

验收标准：

- [ ] 用户可以从模板创建项目。
- [ ] Dashboard 能展示核心工作区信息。

## 阶段八：完善与交付

- [ ] 完善全局错误处理。
- [ ] 完善 Toast 提示。
- [ ] 完善加载状态。
- [ ] 完善空状态。
- [ ] 完善响应式。
- [ ] 优化深色主题。
- [ ] 增加基础组件测试或页面冒烟测试。
- [ ] 完善 README 前端启动说明。

验收标准：

- [ ] 新开发者可以启动前端项目。
- [ ] 核心链路可以完整跑通。
- [ ] UI 细节达到可演示状态。

## 6. 前端质量要求

- [ ] 请求必须统一封装。
- [ ] 表单必须有校验。
- [ ] 重要操作必须有 Toast 提示。
- [ ] 删除类操作必须二次确认。
- [ ] 组件复用清晰。
- [ ] 页面支持响应式。
- [ ] 深色主题优先。
- [ ] 信息密度适中。
- [ ] 代码必须可运行。
- [ ] 业务模块边界清晰。
- [ ] 不在页面组件中堆积复杂请求逻辑。
- [ ] API 类型、表单类型、业务类型需要统一维护。

## 7. MVP 前端暂不实现

- [ ] 真实多 Agent 执行 UI 的完整流程。
- [ ] 复杂团队协作 UI。
- [ ] 支付页面和套餐购买流程。
- [ ] 浏览器插件相关页面。
- [ ] IDE 插件相关页面。
- [ ] 复杂权限管理页面。
- [ ] 真实模型供应商 Key 的完整配置和调用闭环。

## 8. 推荐开发顺序

1. 前端基础设施和 App Shell。
2. 登录 / 注册和路由保护。
3. Dashboard 基础版。
4. 项目列表、创建、详情。
5. Prompt Studio。
6. Rules Builder。
7. Project Brain。
8. 模板入口和 Dashboard 完善。
9. 设置页和交付优化。
