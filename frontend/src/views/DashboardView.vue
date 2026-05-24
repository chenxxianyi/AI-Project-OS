<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAppStore } from '@/stores/app'
import { useAuthStore } from '@/stores/auth'
import { useProjectStore } from '@/stores/project'
import { dashboardApi } from '@/api/dashboard'
import { generationApi, type AIGeneration } from '@/api/generation'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'
import AppModal from '@/components/ui/AppModal.vue'

const store = useAppStore()
const auth = useAuthStore()
const projectStore = useProjectStore()
const toast = useToast()

const stats = ref({ project_count: 0, prompt_count: 0, generation_count: 0 })
const recentGenerations = ref<AIGeneration[]>([])
const error = ref('')

const quickActions = [
  '▣ 创建新项目',
  '▤ 打开 Prompt 工作室',
  '⊙ 生成规则',
  '⌕ 搜索模板',
  '⬡ 添加知识',
]

const projectTypeIcons: Record<string, { icon: string; iconClass?: string }> = {
  saas: { icon: '⌁' },
  ai_chat_app: { icon: '◌', iconClass: 'purple' },
  ai_agent_app: { icon: '⊙' },
  ecommerce: { icon: '⌘' },
  admin_dashboard: { icon: '□', iconClass: 'purple' },
  blog_cms: { icon: '▤' },
  landing_page: { icon: '✧' },
  mobile_app: { icon: '⬡' },
  developer_tool: { icon: '▹' },
  internal_system: { icon: '⌁' },
}

function getProjectIcon(type: string) {
  return projectTypeIcons[type] || { icon: '▣' }
}

function timeAgo(dateStr: string) {
  const diff = Date.now() - new Date(dateStr).getTime()
  const mins = Math.floor(diff / 60000)
  if (mins < 60) return `${mins}分钟前`
  const hours = Math.floor(mins / 60)
  if (hours < 24) return `${hours}小时前`
  const days = Math.floor(hours / 24)
  return `${days}天前`
}

onMounted(async () => {
  try {
    error.value = ''
    const [statsRes, , genRes] = await Promise.all([
      dashboardApi.stats(),
      projectStore.fetchProjects({ page: 1, page_size: 5 }),
      generationApi.list({ page: 1, page_size: 5 }),
    ])
    stats.value = statsRes.data.data
    recentGenerations.value = genRes.data.data
  } catch (e: unknown) {
    const message = e instanceof Error ? e.message : '工作台加载失败'
    error.value = message
    toast.show(message)
  }
})
</script>

<template>
  <div class="p-[32px_36px]">
    <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>
    <div class="workspace-top flex items-center justify-between mb-6">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[340px]">
        <span class="text-muted">⌕</span>
        <input class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索项目、提示词、文档..." />
        <span class="kbd text-[11px] text-muted bg-[#f1f5ff] px-1.5 py-0.5 rounded">⌘ K</span>
      </div>
      <AppButton variant="primary" @click="store.openCreateModal()">＋ 新建项目</AppButton>
    </div>

    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">早上好，{{ auth.username }} 👋</h1>
      <p class="text-muted mt-1">你的工作台动态一览。</p>
    </section>

    <div class="dashboard-layout grid grid-cols-[1fr_280px] gap-5">
      <div>
        <section class="stats grid grid-cols-5 gap-4 mb-5">
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box">▣</span>
            <div><strong class="block text-lg">{{ stats.project_count }}</strong><span class="text-xs text-muted">项目</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box green">›_</span>
            <div><strong class="block text-lg">{{ stats.prompt_count }}</strong><span class="text-xs text-muted">提示词</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box">▤</span>
            <div><strong class="block text-lg">{{ stats.generation_count }}</strong><span class="text-xs text-muted">AI 生成</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box purple">♙</span>
            <div><strong class="block text-lg">1</strong><span class="text-xs text-muted">用户</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box pink">≋</span>
            <div><strong class="block text-lg">0</strong><span class="text-xs text-muted">Token 用量</span></div>
          </div>
        </section>

        <section class="grid grid-cols-2 gap-5 mb-5">
          <AppCard title="最近项目">
            <template #actions>
              <router-link to="/dashboard" class="text-primary text-xs font-semibold">查看全部</router-link>
            </template>
            <div class="flex flex-col gap-0">
              <router-link v-for="p in projectStore.projects" :key="p.id" :to="`/project/${p.id}`" class="flex items-center gap-3 py-3 border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors">
                <span class="icon-box" :class="getProjectIcon(p.project_type).iconClass">{{ getProjectIcon(p.project_type).icon }}</span>
                <div class="flex-1 min-w-0"><b class="block text-sm truncate">{{ p.name }}</b><span class="text-xs text-muted">{{ p.frontend_stack || p.project_type }}</span></div>
                <span class="text-xs text-muted whitespace-nowrap">{{ timeAgo(p.updated_at) }}</span>
              </router-link>
            </div>
          </AppCard>

          <AppCard title="AI 动态">
            <template #actions>
              <router-link v-if="recentGenerations[0]" :to="`/project/${recentGenerations[0].project_id}/prompt-studio`" class="text-primary text-xs font-semibold">查看全部动态</router-link>
            </template>
            <div class="flex flex-col gap-0">
              <router-link v-for="g in recentGenerations" :key="g.id" :to="`/project/${g.project_id}/prompt-studio`" class="flex items-center gap-3 py-3 border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors">
                <span class="icon-box">✧</span>
                <div class="flex-1 min-w-0"><b class="block text-sm">{{ g.generation_type === 'prompt' ? '已生成提示词' : '已生成规则' }}</b><span class="text-xs text-muted">{{ g.model_provider }}</span></div>
                <span class="text-xs text-muted whitespace-nowrap">{{ g.status }}</span>
              </router-link>
            </div>
          </AppCard>
        </section>

        <AppCard title="提示词用量 ⓘ" class="mt-[18px]">
          <template #actions>
            <AppButton variant="ghost">本周⌄</AppButton>
          </template>
          <div class="chart h-[180px]">
            <svg viewBox="0 0 760 180" preserveAspectRatio="none" class="w-full h-full">
              <path d="M0 130 C60 112 75 80 115 86 S185 130 235 92 S320 130 360 65 S450 118 515 92 S615 136 660 92 S720 45 760 50" fill="none" stroke="#4167ff" stroke-width="3"/>
              <path d="M0 130 C60 112 75 80 115 86 S185 130 235 92 S320 130 360 65 S450 118 515 92 S615 136 660 92 S720 45 760 50 L760 180 L0 180Z" fill="rgba(65,103,255,.08)"/>
            </svg>
          </div>
        </AppCard>
      </div>

      <aside class="flex flex-col gap-5">
        <div class="p-5 rounded-[14px] glass glass-hover">
          <h3 class="font-bold text-sm mb-4">工作台概览</h3>
          <p class="text-xs text-muted mt-5">计划</p><b class="text-sm">专业版</b>
          <p class="text-xs text-muted mt-[22px]">用量</p><b class="text-sm">4.2M / 10M Token <span class="text-muted float-right">42%</span></b>
          <div class="usage-bar mt-2.5"><i></i></div>
          <p class="text-xs text-muted mt-[26px]">团队成员</p>
          <div class="member-stack mt-3 mb-[22px]"><i></i><i></i><i></i><i></i><AppChip label="+3" /></div>
          <AppButton block @click="toast.show('邀请链接已复制')">♙ 邀请</AppButton>
        </div>

        <div class="p-5 rounded-[14px] glass glass-hover">
          <h3 class="font-bold text-sm mb-3">快捷操作</h3>
          <div class="flex flex-col gap-0">
            <div v-for="q in quickActions" :key="q" class="flex items-center py-2.5 border-b border-line last:border-0 text-sm cursor-pointer hover:bg-[#f8fbff] transition-colors">
              {{ q }} <span class="ml-auto">›</span>
            </div>
          </div>
        </div>

        <div class="p-5 rounded-[14px] glass glass-hover !bg-gradient-to-br !from-[#eef3ff]/50 !to-white/30">
          <h3 class="font-bold text-sm">AI Project OS 专业版</h3>
          <p class="text-xs text-muted my-2.5">解锁无限 Token、高级智能体等更多功能。</p>
          <AppButton variant="primary" block>立即升级</AppButton>
        </div>
      </aside>
    </div>
  </div>
  <AppModal />
</template>
