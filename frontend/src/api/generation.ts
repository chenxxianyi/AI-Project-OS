import api from './client'
import type { ApiResponse } from './types'

export interface AIGeneration {
  id: string
  user_id: string
  project_id: string
  generation_type: string
  input_payload: string
  output_content: string
  model_provider: string
  model_name: string
  status: string
  error_message: string
  created_at: number
}

export const generationApi = {
  generate(data: { project_id: string; generation_type: string; input_payload: string }) {
    return api.post<ApiResponse<AIGeneration>>('/generations', data)
  },
  list(params?: { page?: number; page_size?: number }) {
    return api.get<ApiResponse<AIGeneration[]>>('/generations', { params })
  },
}
