import api from './client'
import type { ApiResponse } from './types'

export interface Template {
  id: string
  name: string
  description: string
  project_type: string
  frontend_stack: string
  backend_stack: string
  database_stack: string
  ai_stack: string
  ui_style: string
  default_prompts: string
  default_rules: string
  created_at: string
  updated_at: string
}

export interface TemplateListResponse {
  templates: Template[]
  total: number
}

export const templateApi = {
  list(params?: { project_type?: string; page?: number; page_size?: number }) {
    return api.get<ApiResponse<TemplateListResponse>>('/templates', { params })
  },
  get(id: string) {
    return api.get<ApiResponse<Template>>(`/templates/${id}`)
  },
}
