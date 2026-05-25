import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'landing',
      component: () => import('@/views/LandingView.vue'),
    },
    {
      path: '/features',
      name: 'features',
      component: () => import('@/views/PublicInfoView.vue'),
      meta: {
        publicLayout: true,
        title: '功能',
        description: '围绕项目上下文、Prompt、规则和知识库组织 AI 原生工作流。',
      },
    },
    {
      path: '/template-gallery',
      name: 'template-gallery',
      component: () => import('@/views/PublicInfoView.vue'),
      meta: {
        publicLayout: true,
        title: '模板',
        description: '从常见产品形态和技术栈模板开始，减少从零搭建项目上下文的成本。',
      },
    },
    {
      path: '/pricing',
      name: 'pricing',
      component: () => import('@/views/PublicInfoView.vue'),
      meta: {
        publicLayout: true,
        title: '定价',
        description: '基础能力可用于个人项目，团队、额度和高级智能体能力会逐步开放。',
      },
    },
    {
      path: '/docs',
      name: 'docs',
      component: () => import('@/views/PublicInfoView.vue'),
      meta: {
        publicLayout: true,
        title: '文档',
        description: '查看产品使用方式、项目组织方法和 AI 工作流最佳实践。',
      },
    },
    {
      path: '/changelog',
      name: 'changelog',
      component: () => import('@/views/PublicInfoView.vue'),
      meta: {
        publicLayout: true,
        title: '更新日志',
        description: '跟踪 AI Project OS 的功能迭代、体验优化和平台能力更新。',
      },
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/AuthView.vue'),
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/DashboardView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/projects',
      name: 'projects',
      component: () => import('@/views/ProjectsView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/templates',
      name: 'templates',
      component: () => import('@/views/PlaceholderView.vue'),
      meta: {
        requiresAuth: true,
        title: '模板中心',
        description: '浏览和复用项目模板，后续会支持从模板一键创建项目。',
      },
    },
    {
      path: '/agents',
      name: 'agents',
      component: () => import('@/views/PlaceholderView.vue'),
      meta: {
        requiresAuth: true,
        title: '智能体',
        description: '这里将集中管理多智能体工作流、执行记录和能力配置。',
      },
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('@/views/PlaceholderView.vue'),
      meta: {
        requiresAuth: true,
        title: '设置',
        description: '账户、团队、偏好和集成配置会在这里统一管理。',
      },
    },
    {
      path: '/project/:id',
      name: 'project-detail',
      component: () => import('@/views/ProjectDetailView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/project/:id/prompt-studio',
      name: 'prompt-studio',
      component: () => import('@/views/PromptStudioView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/project/:id/rules-builder',
      name: 'rules-builder',
      component: () => import('@/views/RulesBuilderView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/project/:id/brain',
      name: 'project-brain',
      component: () => import('@/views/ProjectBrainView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/admin',
      component: () => import('@/components/layout/AdminLayout.vue'),
      meta: { requiresAuth: true, requiresAdmin: true },
      children: [
        {
          path: '',
          name: 'admin-overview',
          component: () => import('@/views/admin/AdminOverviewView.vue'),
        },
        {
          path: 'users',
          name: 'admin-users',
          component: () => import('@/views/admin/AdminUsersView.vue'),
        },
        {
          path: 'projects',
          name: 'admin-projects',
          component: () => import('@/views/admin/AdminProjectsView.vue'),
        },
        {
          path: 'generations',
          name: 'admin-generations',
          component: () => import('@/views/admin/AdminGenerationsView.vue'),
        },
        {
          path: 'settings',
          name: 'admin-settings',
          component: () => import('@/views/admin/AdminSettingsView.vue'),
        },
      ],
    },
  ],
})

router.beforeEach(async (to) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.user) {
    await auth.fetchUser()
  }
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    return { name: 'login', query: { redirect: to.fullPath } }
  }
  if (!to.meta.requiresAuth && auth.isLoggedIn && !auth.user) {
    await auth.fetchUser()
  }
  if (to.meta.requiresAdmin && auth.user?.role !== 'admin') {
    return { name: 'dashboard' }
  }
})

export default router
