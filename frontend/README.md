# AI Project OS — 前端应用

> AI 原生工作台，用于 Prompt 工程、项目上下文、规则管理和多智能体开发。

## 技术栈

| 技术 | 版本 | 用途 |
|------|------|------|
| Vue 3 | ^3.5 | 核心框架，Composition API |
| Vite | ^6.1 | 构建工具与开发服务器 |
| TypeScript | ~5.7 | 类型安全 |
| TailwindCSS | ^3.4 | 原子化样式 |
| Vue Router | ^4.5 | 路由管理 |
| Pinia | ^2.3 | 全局状态管理 |

## 快速开始

```bash
# 安装依赖
npm install

# 启动开发服务器（默认 http://localhost:3000）
npm run dev

# 构建生产版本
npm run build

# 预览生产构建
npm run preview
```

## 项目结构

```
frontend/
├── index.html                    # Vite 入口 HTML
├── package.json
├── vite.config.ts                # Vite 配置（含 @ 路径别名）
├── tsconfig.json
├── tailwind.config.js            # Tailwind 主题扩展
├── postcss.config.js
├── env.d.ts
└── src/
    ├── main.ts                   # 应用入口
    ├── App.vue                   # 根组件（Landing 独立布局 / 其他页面 AppShell 包裹）
    ├── assets/styles/main.css    # 全局样式 + Tailwind 指令 + 自定义组件类
    ├── router/index.ts           # 路由定义
    ├── stores/app.ts             # Pinia Store（侧栏状态、弹窗状态）
    ├── composables/useToast.ts   # Toast 提示组合式函数
    ├── components/
    │   ├── layout/
    │   │   ├── AppShell.vue      # 应用外壳（Sidebar + 主内容区 + Toast）
    │   │   ├── Sidebar.vue       # 侧边栏导航
    │   │   └── Topbar.vue        # 顶部栏
    │   └── ui/
    │       ├── AppButton.vue     # 按钮（default / primary / ghost）
    │       ├── AppCard.vue       # 卡片面板
    │       ├── AppTabs.vue       # 标签页切换
    │       ├── AppModal.vue      # 创建项目弹窗（5步向导）
    │       ├── AppToast.vue      # Toast 全局提示
    │       └── AppChip.vue       # 标签/徽章
    └── views/
        ├── LandingView.vue       # 首页（Hero + 功能条 + 信任区）
        ├── DashboardView.vue     # 工作台（统计 + 项目列表 + AI 动态 + 图表 + 右侧栏）
        ├── ProjectDetailView.vue # 项目详情（概览 + 技术栈 + 提示词 + 动态）
        ├── PromptStudioView.vue  # Prompt 工作室（模板 + 编辑器 + 测试/输出）
        ├── RulesBuilderView.vue  # 规则构建器（分类 + Markdown 编辑器 + 预览）
        └── ProjectBrainView.vue  # 项目大脑（知识图谱 + 节点详情）
```

## 页面路由

| 路径 | 视图 | 说明 |
|------|------|------|
| `/` | LandingView | 产品首页，独立布局（无侧栏） |
| `/dashboard` | DashboardView | 工作台首页 |
| `/project/:id` | ProjectDetailView | 项目详情 |
| `/prompt-studio` | PromptStudioView | Prompt 工作室 |
| `/rules-builder` | RulesBuilderView | 规则构建器 |
| `/project-brain` | ProjectBrainView | 项目大脑 |

## 设计系统

### 颜色变量

| 变量 | 值 | 用途 |
|------|------|------|
| `--primary` | `#4167ff` | 主色 |
| `--primary-2` | `#7e95ff` | 主色浅 |
| `--primary-3` | `#8f68ff` | 紫色 |
| `--cyan` | `#23c7db` | 青色 |
| `--green` | `#27c990` | 绿色/成功 |
| `--red` | `#ff6a87` | 红色/错误 |
| `--muted` | `#6b7898` | 辅助文字 |
| `--line` | `#e3eafa` | 边框/分割线 |
| `--bg` | `#f6f9ff` | 背景色 |

### Tailwind 扩展

项目通过 `tailwind.config.js` 扩展了以下设计 Token：

- **colors** — `primary`, `cyan`, `green`, `red`, `line`, `muted`, `bg` 等
- **borderRadius** — `card: 18px`, `panel: 14px`, `btn: 10px`, `chip: 9px`
- **boxShadow** — `card`, `strong`, `btn`, `btn-hover`, `primary`
- **fontFamily** — `sans`（Manrope + 中文字体栈）, `mono`

## 状态管理

### Pinia Store (`stores/app.ts`)

| 状态 | 类型 | 说明 |
|------|------|------|
| `sidebarCollapsed` | `boolean` | 侧栏折叠状态 |
| `createModalOpen` | `boolean` | 创建项目弹窗开关 |

| 方法 | 说明 |
|------|------|
| `toggleSidebar()` | 切换侧栏折叠 |
| `openCreateModal()` | 打开创建项目弹窗 |
| `closeCreateModal()` | 关闭创建项目弹窗 |

### Composable (`composables/useToast.ts`)

全局单例 Toast，提供 `show(msg, duration?)` 方法，在 `AppToast.vue` 中自动渲染。

## MVP 开发阶段规划

> 详见 `FRONTEND_TASKS.md`

| 阶段 | 内容 | 状态 |
|------|------|------|
| 一 | 基础设施和 App Shell | ✅ 已完成 |
| 二 | 用户认证（登录/注册） | 🔲 待开发 |
| 三 | 项目模块（列表/创建/详情） | 🔲 待开发 |
| 四 | Prompt Studio | 🔲 待开发 |
| 五 | Rules Builder | 🔲 待开发 |
| 六 | Project Brain | 🔲 待开发 |
| 七 | 模板与 Dashboard 完善 | 🔲 待开发 |
| 八 | 完善与交付 | 🔲 待开发 |

## 待对接后端接口

MVP 阶段需要对接的核心 API：

- `POST /api/auth/register` — 用户注册
- `POST /api/auth/login` — 用户登录
- `GET /api/auth/me` — 获取当前用户
- `GET /api/projects` — 项目列表
- `POST /api/projects` — 创建项目
- `GET /api/projects/:id` — 项目详情
- `PUT /api/projects/:id` — 更新项目
- `DELETE /api/projects/:id` — 删除项目
- `GET /api/projects/:projectId/prompts` — Prompt 列表
- `POST /api/projects/:projectId/prompts/generate` — 生成 Prompt
- `GET /api/projects/:projectId/rules` — 获取规则
- `POST /api/projects/:projectId/rules/generate` — 生成规则
- `GET /api/projects/:projectId/docs` — 文档列表
- `POST /api/projects/:projectId/docs` — 创建文档
- `GET /api/templates` — 模板列表

## 环境变量

在项目根目录创建 `.env.local`：

```env
VITE_API_BASE_URL=http://localhost:8000/api
```

## 约定

- 所有用户可见文案默认使用**简体中文**
- 组件命名采用 `App` 前缀（`AppButton`, `AppCard`）
- 视图命名采用 `View` 后缀（`DashboardView`）
- 路径别名 `@` 指向 `src/`
- 使用 `<script setup lang="ts">` 编写组件
