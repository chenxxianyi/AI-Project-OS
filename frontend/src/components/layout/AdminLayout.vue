<script setup lang="ts">
import { useRoute, useRouter, RouterView } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import AppToast from '@/components/ui/AppToast.vue'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()

const navItems = [
  { icon: '◉', label: '系统概览', name: 'admin-overview' },
  { icon: '♙', label: '用户管理', name: 'admin-users' },
  { icon: '▣', label: '项目管理', name: 'admin-projects' },
  { icon: '✧', label: '生成记录', name: 'admin-generations' },
  { icon: '⚙', label: '系统设置', name: 'admin-settings' },
]

function isActive(name: string) {
  return route.name === name
}

function handleBackToApp() {
  router.push({ name: 'dashboard' })
}

function handleLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <div class="admin-layout flex min-h-screen">
    <aside
      class="admin-sidebar shrink-0 w-[216px] min-h-screen flex flex-col p-5 glass-lg !border-r !border-white/20 !rounded-none"
    >
      <div class="relative z-10 flex items-center gap-3 font-black text-lg text-[#0b173d] mb-2">
        <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
        AI Project OS
      </div>
      <div class="relative z-10 mb-4 flex items-center gap-1.5 px-1">
        <span class="text-[11px] font-bold text-white bg-primary/90 rounded-[6px] px-2 py-0.5">ADMIN</span>
        <span class="text-[11px] text-muted">管理后台</span>
      </div>

      <nav class="relative z-10 flex flex-col gap-1 flex-1">
        <a
          v-for="item in navItems"
          :key="item.name"
          class="nav-link flex items-center gap-3 px-3 py-2.5 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="isActive(item.name!) ? 'glass-pill !rounded-[10px] text-primary' : 'text-[#425071] hover:bg-white/20'"
          @click="router.push({ name: item.name })"
        >
          <span>{{ item.icon }}</span>{{ item.label }}
        </a>
      </nav>

      <div class="relative z-10 mt-auto flex flex-col gap-3 pt-4 border-t border-line">
        <button
          class="btn glass-pill !rounded-[10px] h-[38px] font-bold text-sm flex items-center justify-center gap-2 transition-all duration-ios-tab ease-ios-snappy"
          @click="handleBackToApp"
        >
          ↩ 返回应用
        </button>
        <div class="flex items-center gap-3 px-2">
          <div class="w-8 h-8 rounded-full bg-gradient-to-br from-[#222d58] to-[#f3a071] grid place-items-center text-white text-xs font-bold">{{ (auth.username || '?')[0].toUpperCase() }}</div>
          <div class="flex-1 min-w-0">
            <b class="block text-sm truncate">{{ auth.username || '未登录' }}</b>
            <span class="text-xs text-muted truncate block">{{ auth.user?.email || '' }}</span>
          </div>
          <span class="text-muted cursor-pointer hover:text-primary transition-colors duration-ios-tab" @click="handleLogout" title="退出登录">⏻</span>
        </div>
      </div>
    </aside>

    <main class="flex-1 min-w-0 bg-bg overflow-auto">
      <RouterView />
    </main>
  </div>
  <AppToast />
</template>
