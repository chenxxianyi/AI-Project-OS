<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'

const route = useRoute()
const router = useRouter()
const store = useAppStore()

const navItems = [
  { icon: '⌂', label: '首页', name: 'dashboard' },
  { icon: '▣', label: '项目', name: 'project-detail' },
  { icon: '▤', label: '模板', name: 'dashboard' },
  { icon: '◇', label: '智能体', name: 'dashboard' },
  { sep: true },
  { icon: '↗', label: 'Prompt 工作室', name: 'prompt-studio' },
  { icon: '▥', label: '规则构建器', name: 'rules-builder' },
  { icon: '⬡', label: '项目大脑', name: 'project-brain' },
  { icon: '⚙', label: '设置', name: 'dashboard' },
]

function isActive(name: string) {
  return route.name === name
}

function navigate(name: string) {
  if (name === 'project-detail') {
    router.push({ name, params: { id: 'demo' } })
  } else {
    router.push({ name })
  }
}
</script>

<template>
  <aside
    class="sidebar shrink-0 w-[216px] min-h-screen flex flex-col p-5 glass !border-r !border-white/20 !rounded-none"
  >
    <router-link to="/" class="flex items-center gap-3 font-black text-lg text-[#0b173d] mb-6">
      <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
      AI Project OS
    </router-link>

    <nav class="flex flex-col gap-1 flex-1">
      <template v-for="(item, i) in navItems" :key="i">
        <div v-if="item.sep" class="nav-sep h-px bg-line my-2" />
        <a
          v-else
          class="nav-link flex items-center gap-3 px-3 py-2.5 rounded-[10px] text-sm font-semibold transition-colors cursor-pointer"
          :class="isActive(item.name!) ? 'bg-soft text-primary' : 'text-[#425071] hover:bg-[#f8fbff]'"
          @click="navigate(item.name!)"
        >
          <span>{{ item.icon }}</span>{{ item.label }}
        </a>
      </template>
    </nav>

    <div class="mt-auto flex flex-col gap-3 pt-4 border-t border-line">
      <button
        class="btn w-full bg-[#f8faff] border border-line rounded-[10px] h-[42px] font-bold text-sm flex items-center justify-center gap-2 hover:-translate-y-px transition"
        @click="store.openCreateModal()"
      >
        ◎ 升级计划
      </button>
      <div class="flex items-center gap-3 px-2">
        <div class="w-8 h-8 rounded-full bg-gradient-to-br from-[#222d58] to-[#f3a071]" />
        <div class="flex-1 min-w-0">
          <b class="block text-sm truncate">Alex Morgan</b>
          <span class="text-xs text-muted truncate block">alex@morgan.dev</span>
        </div>
        <span class="text-muted">⌄</span>
      </div>
    </div>
  </aside>
</template>
