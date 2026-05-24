import api from './client'
import type { ApiResponse } from './types'

export interface User {
  id: string
  email: string
  username: string
  avatar: string
  role: string
}

export interface AuthData {
  token: string
  user: User
}

export const authApi = {
  register(data: { email: string; username: string; password: string }) {
    return api.post<ApiResponse<AuthData>>('/auth/register', data)
  },
  login(data: { email: string; password: string }) {
    return api.post<ApiResponse<AuthData>>('/auth/login', data)
  },
  me() {
    return api.get<ApiResponse<User>>('/auth/me')
  },
  logout() {
    return api.post<ApiResponse<null>>('/auth/logout')
  },
}
