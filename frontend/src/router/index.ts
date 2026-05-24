import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'landing',
      component: () => import('@/views/LandingView.vue'),
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/views/DashboardView.vue'),
    },
    {
      path: '/project/:id',
      name: 'project-detail',
      component: () => import('@/views/ProjectDetailView.vue'),
    },
    {
      path: '/prompt-studio',
      name: 'prompt-studio',
      component: () => import('@/views/PromptStudioView.vue'),
    },
    {
      path: '/rules-builder',
      name: 'rules-builder',
      component: () => import('@/views/RulesBuilderView.vue'),
    },
    {
      path: '/project-brain',
      name: 'project-brain',
      component: () => import('@/views/ProjectBrainView.vue'),
    },
  ],
})

export default router
