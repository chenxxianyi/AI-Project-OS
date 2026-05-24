import api from './client'
import type { ApiResponse } from './types'

export interface Prompt {
  id: string
  project_id: string
  user_id: string
  type: string
  title: string
  content: string
  version: number
  created_at: string
  updated_at: string
}

export interface PromptVersion {
  id: string
  prompt_id: string
  version: number
  content: string
  change_note: string
  created_at: number
}

export const promptApi = {
  create(projectId: string, data: { type: string; title: string; content: string }) {
    return api.post<ApiResponse<Prompt>>(`/projects/${projectId}/prompts`, data)
  },
  list(projectId: string, params?: { type?: string }) {
    return api.get<ApiResponse<Prompt[]>>(`/projects/${projectId}/prompts`, { params })
  },
  get(projectId: string, id: string) {
    return api.get<ApiResponse<Prompt>>(`/projects/${projectId}/prompts/${id}`)
  },
  update(projectId: string, id: string, data: Partial<Pick<Prompt, 'type' | 'title' | 'content'>>) {
    return api.put<ApiResponse<Prompt>>(`/projects/${projectId}/prompts/${id}`, data)
  },
  delete(projectId: string, id: string) {
    return api.delete<ApiResponse<null>>(`/projects/${projectId}/prompts/${id}`)
  },
  versions(projectId: string, promptId: string) {
    return api.get<ApiResponse<PromptVersion[]>>(`/projects/${projectId}/prompts/${promptId}/versions`)
  },
}
