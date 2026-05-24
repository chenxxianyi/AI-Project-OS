import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authApi, type User } from '@/api/auth'

export type { User }

export const useAuthStore = defineStore('auth', () => {
  const token = ref('')
  const user = ref<User | null>(null)

  const isLoggedIn = computed(() => !!token.value || !!user.value)
  const username = computed(() => user.value?.username || '')

  async function login(email: string, password: string) {
    const res = await authApi.login({ email, password })
    const { token: t, user: u } = res.data.data
    token.value = t
    user.value = u
    localStorage.removeItem('token')
  }

  async function register(email: string, username: string, password: string) {
    const res = await authApi.register({ email, username, password })
    const { token: t, user: u } = res.data.data
    token.value = t
    user.value = u
    localStorage.removeItem('token')
  }

  async function fetchUser() {
    try {
      const res = await authApi.me()
      user.value = res.data.data
    } catch {
      logout()
    }
  }

  function logout() {
    token.value = ''
    user.value = null
    localStorage.removeItem('token')
  }

  async function logoutRemote() {
    try {
      await authApi.logout()
    } finally {
      logout()
    }
  }

  return { token, user, isLoggedIn, username, login, register, fetchUser, logout, logoutRemote }
})
