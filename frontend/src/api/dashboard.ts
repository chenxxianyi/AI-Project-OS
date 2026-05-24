import api from './client'
import type { ApiResponse } from './types'

export interface DashboardStats {
  project_count: number
  prompt_count: number
  generation_count: number
}

export const dashboardApi = {
  stats() {
    return api.get<ApiResponse<DashboardStats>>('/dashboard/stats')
  },
}
