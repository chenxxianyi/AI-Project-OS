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
})

export default router
