import { defineStore } from 'pinia'
import { ref } from 'vue'
import { adminApi, type SystemStats, type SystemHealth, type SystemSettings, type AdminUser, type AdminProject, type AdminGeneration } from '@/api/admin'

export const useAdminStore = defineStore('admin', () => {
  const stats = ref<SystemStats>({
    total_users: 0,
    active_users: 0,
    total_projects: 0,
    total_prompts: 0,
    total_generations: 0,
    total_tokens_used: 0,
    failed_generations: 0,
    uptime: '',
    version: '',
  })

  const health = ref<SystemHealth>({
    database: 'unknown',
    redis: 'unknown',
    ai_provider: 'unknown',
    disk_usage: '0%',
    memory_usage: '0%',
  })

  const settings = ref<SystemSettings>({
    allow_registration: true,
    default_user_role: 'user',
    max_projects_per_user: 10,
    ai_rate_limit: 100,
    maintenance_mode: false,
  })

  const users = ref<AdminUser[]>([])
  const usersTotal = ref(0)
  const usersPage = ref(1)
  const usersSearch = ref('')
  const usersRoleFilter = ref('')

  const projects = ref<AdminProject[]>([])
  const projectsTotal = ref(0)
  const projectsPage = ref(1)
  const projectsSearch = ref('')
  const projectsStatusFilter = ref('')

  const generations = ref<AdminGeneration[]>([])
  const generationsTotal = ref(0)
  const generationsPage = ref(1)
  const generationsStatusFilter = ref('')

  const loading = ref(false)
  const error = ref('')

  async function fetchStats() {
    try {
      loading.value = true
      error.value = ''
      const res = await adminApi.stats()
      stats.value = res.data.data
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '获取统计失败'
    } finally {
      loading.value = false
    }
  }

  async function fetchHealth() {
    try {
      const res = await adminApi.health()
      health.value = res.data.data
    } catch {
      // silent - health check is non-critical
    }
  }

  async function fetchSettings() {
    try {
      const res = await adminApi.getSettings()
      settings.value = res.data.data
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '获取设置失败'
    }
  }

  async function updateSettings(data: Partial<SystemSettings>) {
    try {
      const res = await adminApi.updateSettings(data)
      settings.value = res.data.data
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '更新设置失败'
    }
  }

  async function fetchUsers(params?: { page?: number; page_size?: number; search?: string; role?: string }) {
    try {
      loading.value = true
      error.value = ''
      const res = await adminApi.listUsers(params)
      users.value = res.data.data.users
      usersTotal.value = res.data.data.total
      if (params?.page) usersPage.value = params.page
      if (params?.search !== undefined) usersSearch.value = params.search
      if (params?.role !== undefined) usersRoleFilter.value = params.role
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '获取用户列表失败'
    } finally {
      loading.value = false
    }
  }

  async function updateUser(id: string, data: Partial<Pick<AdminUser, 'role' | 'status'>>) {
    try {
      await adminApi.updateUser(id, data)
      const idx = users.value.findIndex(u => u.id === id)
      if (idx !== -1) {
        const res = await adminApi.listUsers({ page: usersPage.value, page_size: 20, search: usersSearch.value, role: usersRoleFilter.value })
        users.value = res.data.data.users
      }
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '更新用户失败'
    }
  }

  async function deleteUser(id: string) {
    try {
      await adminApi.deleteUser(id)
      users.value = users.value.filter(u => u.id !== id)
      usersTotal.value--
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '删除用户失败'
    }
  }

  async function fetchProjects(params?: { page?: number; page_size?: number; search?: string; status?: string; project_type?: string }) {
    try {
      loading.value = true
      error.value = ''
      const res = await adminApi.listProjects(params)
      projects.value = res.data.data.projects
      projectsTotal.value = res.data.data.total
      if (params?.page) projectsPage.value = params.page
      if (params?.search !== undefined) projectsSearch.value = params.search
      if (params?.status !== undefined) projectsStatusFilter.value = params.status
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '获取项目列表失败'
    } finally {
      loading.value = false
    }
  }

  async function deleteProject(id: string) {
    try {
      await adminApi.deleteProject(id)
      projects.value = projects.value.filter(p => p.id !== id)
      projectsTotal.value--
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '删除项目失败'
    }
  }

  async function fetchGenerations(params?: { page?: number; page_size?: number; status?: string; generation_type?: string }) {
    try {
      loading.value = true
      error.value = ''
      const res = await adminApi.listGenerations(params)
      generations.value = res.data.data.generations
      generationsTotal.value = res.data.data.total
      if (params?.page) generationsPage.value = params.page
      if (params?.status !== undefined) generationsStatusFilter.value = params.status
    } catch (e: unknown) {
      error.value = e instanceof Error ? e.message : '获取生成记录失败'
    } finally {
      loading.value = false
    }
  }

  return {
    stats, health, settings,
    users, usersTotal, usersPage, usersSearch, usersRoleFilter,
    projects, projectsTotal, projectsPage, projectsSearch, projectsStatusFilter,
    generations, generationsTotal, generationsPage, generationsStatusFilter,
    loading, error,
    fetchStats, fetchHealth, fetchSettings, updateSettings,
    fetchUsers, updateUser, deleteUser,
    fetchProjects, deleteProject,
    fetchGenerations,
  }
})
