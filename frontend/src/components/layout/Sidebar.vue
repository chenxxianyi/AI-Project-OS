<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { useAuthStore } from '@/stores/auth'
import { useToast } from '@/composables/useToast'

const route = useRoute()
const router = useRouter()
const store = useAppStore()
const auth = useAuthStore()
const toast = useToast()

const navItems = [
  { icon: '⌂', label: '首页', name: 'dashboard', activeName: 'dashboard' },
  { icon: '▣', label: '项目', name: 'projects', activeName: 'projects', activeNames: ['projects', 'project-detail', 'prompt-studio', 'rules-builder', 'project-brain'] },
  { icon: '▤', label: '模板', name: 'templates', activeName: 'templates' },
  { icon: '◇', label: '智能体', name: 'agents', activeName: 'agents' },
  { sep: true },
  { icon: '↗', label: 'Prompt 工作室', name: 'prompt-studio', activeName: 'prompt-studio' },
  { icon: '▥', label: '规则构建器', name: 'rules-builder', activeName: 'rules-builder' },
  { icon: '⬡', label: '项目大脑', name: 'project-brain', activeName: 'project-brain' },
  { icon: '⚙', label: '设置', name: 'settings', activeName: 'settings' },
]

const isAdmin = computed(() => auth.user?.role === 'admin')

function isActive(activeName: string, activeNames?: string[]) {
  const name = String(route.name)
  if (activeNames) return activeNames.includes(name)
  return name === activeName
}

function navigate(name: string) {
  if (['prompt-studio', 'rules-builder', 'project-brain'].includes(name)) {
    const projectId = route.params.id
    if (projectId) {
      router.push({ name, params: { id: projectId } }).catch(() => {})
    } else {
      toast.show('请先进入一个项目，再使用此功能')
      router.push({ name: 'projects' }).catch(() => {})
    }
  } else {
    router.push({ name }).catch(() => {})
  }
}

function handleLogout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <aside
    class="sidebar shrink-0 w-[216px] min-h-screen flex flex-col p-5 glass-lg !border-r !border-white/20 !rounded-none"
  >
    <router-link to="/" class="relative z-10 flex items-center gap-3 font-black text-lg text-[#0b173d] mb-6">
      <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
      AI Project OS
    </router-link>

    <nav class="relative z-10 flex flex-col gap-1 flex-1">
      <template v-for="(item, i) in navItems" :key="i">
        <div v-if="item.sep" class="nav-sep h-px bg-line my-2" />
        <a
          v-else
          class="nav-link flex items-center gap-3 px-3 py-2.5 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="isActive(item.activeName!, item.activeNames) ? 'glass-pill !rounded-[10px] text-primary' : 'text-[#425071] hover:bg-white/20'"
          @click="navigate(item.name!)"
        >
          <span>{{ item.icon }}</span>{{ item.label }}
        </a>
      </template>
      <template v-if="isAdmin">
        <div class="nav-sep h-px bg-line my-2" />
        <a
          class="nav-link flex items-center gap-3 px-3 py-2.5 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="route.path.startsWith('/admin') ? 'glass-pill !rounded-[10px] text-primary' : 'text-[#425071] hover:bg-white/20'"
          @click="router.push('/admin').catch(() => {})"
        >
          <span>◉</span>管理后台
        </a>
      </template>
    </nav>

    <div class="relative z-10 mt-auto flex flex-col gap-3 pt-4 border-t border-line">
      <button
        class="btn glass-pill !rounded-[10px] h-[42px] font-bold text-sm flex items-center justify-center gap-2 transition-all duration-ios-tab ease-ios-snappy"
        @click="store.openCreateModal()"
      >
        ◎ 升级计划
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
</template>
