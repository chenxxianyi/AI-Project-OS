import axios from 'axios'
import { useAuthStore } from '@/stores/auth'

const api = axios.create({
  baseURL: '/api/v1',
  timeout: 15000,
  withCredentials: true,
  headers: { 'Content-Type': 'application/json' },
})

api.interceptors.request.use((config) => {
  const auth = useAuthStore()
  if (auth.token) {
    config.headers.Authorization = `Bearer ${auth.token}`
  }
  return config
})

api.interceptors.response.use(
  (res) => {
    const body = res.data
    if (body.success === false) {
      const msg = body.error?.code || body.message || '请求失败'
      return Promise.reject(new Error(msg))
    }
    return res
  },
  (err) => {
    if (err.response?.status === 401) {
      const auth = useAuthStore()
      auth.logout()
      window.location.href = '/login'
    }
    const msg = err.response?.data?.message || err.response?.data?.error?.code || err.message
    return Promise.reject(new Error(msg))
  },
)

export default api
