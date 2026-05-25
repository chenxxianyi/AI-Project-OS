import api from './client'
import type { ApiResponse } from './types'

export interface AdminUser {
  id: string
  email: string
  username: string
  avatar: string
  role: string
  status: string
  project_count: number
  generation_count: number
  created_at: string
  updated_at: string
}

export interface AdminUserListResponse {
  users: AdminUser[]
  total: number
}

export interface AdminProject {
  id: string
  user_id: string
  name: string
  description: string
  project_type: string
  frontend_stack: string
  backend_stack: string
  status: string
  owner_name: string
  created_at: string
  updated_at: string
}

export interface AdminProjectListResponse {
  projects: AdminProject[]
  total: number
}

export interface AdminGeneration {
  id: string
  user_id: string
  project_id: string
  generation_type: string
  model_provider: string
  model_name: string
  status: string
  error_message: string
  user_email: string
  project_name: string
  created_at: string
}

export interface AdminGenerationListResponse {
  generations: AdminGeneration[]
  total: number
}

export interface SystemStats {
  total_users: number
  active_users: number
  total_projects: number
  total_prompts: number
  total_generations: number
  total_tokens_used: number
  failed_generations: number
  uptime: string
  version: string
}

export interface SystemHealth {
  database: string
  redis: string
  ai_provider: string
  disk_usage: string
  memory_usage: string
}

export interface SystemSettings {
  allow_registration: boolean
  default_user_role: string
  max_projects_per_user: number
  ai_rate_limit: number
  maintenance_mode: boolean
}

export const adminApi = {
  // System
  stats() {
    return api.get<ApiResponse<SystemStats>>('/admin/stats')
  },
  health() {
    return api.get<ApiResponse<SystemHealth>>('/admin/health')
  },
  getSettings() {
    return api.get<ApiResponse<SystemSettings>>('/admin/settings')
  },
  updateSettings(data: Partial<SystemSettings>) {
    return api.put<ApiResponse<SystemSettings>>('/admin/settings', data)
  },

  // Users
  listUsers(params?: { page?: number; page_size?: number; search?: string; role?: string }) {
    return api.get<ApiResponse<AdminUserListResponse>>('/admin/users', { params })
  },
  getUser(id: string) {
    return api.get<ApiResponse<AdminUser>>(`/admin/users/${id}`)
  },
  updateUser(id: string, data: Partial<Pick<AdminUser, 'role' | 'status'>>) {
    return api.put<ApiResponse<AdminUser>>(`/admin/users/${id}`, data)
  },
  deleteUser(id: string) {
    return api.delete<ApiResponse<null>>(`/admin/users/${id}`)
  },

  // Projects
  listProjects(params?: { page?: number; page_size?: number; search?: string; status?: string; project_type?: string }) {
    return api.get<ApiResponse<AdminProjectListResponse>>('/admin/projects', { params })
  },
  deleteProject(id: string) {
    return api.delete<ApiResponse<null>>(`/admin/projects/${id}`)
  },

  // Generations
  listGenerations(params?: { page?: number; page_size?: number; status?: string; generation_type?: string }) {
    return api.get<ApiResponse<AdminGenerationListResponse>>('/admin/generations', { params })
  },
}
