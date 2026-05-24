import api from './client'
import type { ApiResponse } from './types'

export interface ProjectDoc {
  id: string
  project_id: string
  user_id: string
  doc_type: string
  title: string
  content: string
  created_at: string
  updated_at: string
}

export const docApi = {
  create(projectId: string, data: { doc_type: string; title: string; content: string }) {
    return api.post<ApiResponse<ProjectDoc>>(`/projects/${projectId}/docs`, data)
  },
  list(projectId: string, params?: { doc_type?: string }) {
    return api.get<ApiResponse<ProjectDoc[]>>(`/projects/${projectId}/docs`, { params })
  },
  get(projectId: string, docId: string) {
    return api.get<ApiResponse<ProjectDoc>>(`/projects/${projectId}/docs/${docId}`)
  },
  update(projectId: string, docId: string, data: Partial<Pick<ProjectDoc, 'doc_type' | 'title' | 'content'>>) {
    return api.put<ApiResponse<ProjectDoc>>(`/projects/${projectId}/docs/${docId}`, data)
  },
  delete(projectId: string, docId: string) {
    return api.delete<ApiResponse<null>>(`/projects/${projectId}/docs/${docId}`)
  },
}
