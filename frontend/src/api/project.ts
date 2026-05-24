import api from './client'
import type { ApiResponse } from './types'

export interface Project {
  id: string
  user_id: string
  name: string
  description: string
  project_type: string
  frontend_stack: string
  backend_stack: string
  database_stack: string
  ai_stack: string
  ui_style: string
  target_user: string
  product_goal: string
  status: string
  created_at: string
  updated_at: string
}

export interface ProjectListResponse {
  projects: Project[]
  total: number
}

export const projectApi = {
  create(data: Partial<Project> & { name: string; project_type: string }) {
    return api.post<ApiResponse<Project>>('/projects', data)
  },
  list(params?: { project_type?: string; search?: string; page?: number; page_size?: number }) {
    return api.get<ApiResponse<ProjectListResponse>>('/projects', { params })
  },
  get(id: string) {
    return api.get<ApiResponse<Project>>(`/projects/${id}`)
  },
  update(id: string, data: Partial<Project>) {
    return api.put<ApiResponse<Project>>(`/projects/${id}`, data)
  },
  delete(id: string) {
    return api.delete<ApiResponse<null>>(`/projects/${id}`)
  },
}
