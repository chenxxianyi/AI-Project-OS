<script setup lang="ts">
import { useAppStore } from '@/stores/app'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'
import AppModal from '@/components/ui/AppModal.vue'

const store = useAppStore()
const toast = useToast()

const recentProjects = [
  { icon: '⌁', name: 'AI SaaS Starter', desc: 'Next.js, Tailwind, Supabase', time: '2小时前更新' },
  { icon: '◌', iconClass: 'purple', name: 'Marketing Agent', desc: 'LangChain, OpenAI, Pinecone', time: '5小时前更新' },
  { icon: '⊙', name: 'E-commerce Platform', desc: 'Next.js, Stripe, Postgres', time: '1天前更新' },
  { icon: '⌘', name: 'Internal Dashboard', desc: 'React, UI Kit, Chart.js', time: '2天前更新' },
  { icon: '□', iconClass: 'purple', name: 'Docs Assistant', desc: 'GPT-4o, LangChain, Chroma', time: '3天前更新' },
]

const aiActivity = [
  { icon: '✧', title: '已生成提示词', desc: 'AI SaaS Starter', time: '2分钟前' },
  { icon: '⛓', title: '规则已更新', desc: '营销智能体', time: '15分钟前' },
  { icon: '▧', title: '已添加上下文', desc: '电商平台', time: '1小时前' },
  { icon: '▷', title: '智能体已执行', desc: '文档助手', time: '2小时前' },
  { icon: '✣', title: '已生成提示词', desc: '内部仪表盘', time: '3小时前' },
]

const quickActions = [
  '▣ 创建新项目',
  '▤ 打开 Prompt 工作室',
  '⊙ 生成规则',
  '⌕ 搜索模板',
  '⬡ 添加知识',
]
</script>

<template>
  <div class="p-[32px_36px]">
    <div class="workspace-top flex items-center justify-between mb-6">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[340px]">
        <span class="text-muted">⌕</span>
        <input class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索项目、提示词、文档..." />
        <span class="kbd text-[11px] text-muted bg-[#f1f5ff] px-1.5 py-0.5 rounded">⌘ K</span>
      </div>
      <AppButton variant="primary" @click="store.openCreateModal()">＋ 新建项目</AppButton>
    </div>

    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">早上好，Alex 👋</h1>
      <p class="text-muted mt-1">你的工作台动态一览。</p>
    </section>

    <div class="dashboard-layout grid grid-cols-[1fr_280px] gap-5">
      <div>
        <section class="stats grid grid-cols-5 gap-4 mb-5">
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box">▣</span>
            <div><strong class="block text-lg">12</strong><span class="text-xs text-muted">项目</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box green">›_</span>
            <div><strong class="block text-lg">248</strong><span class="text-xs text-muted">提示词</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box">▤</span>
            <div><strong class="block text-lg">8</strong><span class="text-xs text-muted">智能体</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box purple">♙</span>
            <div><strong class="block text-lg">4</strong><span class="text-xs text-muted">用户</span></div>
          </div>
          <div class="stat flex items-center gap-3 p-4 rounded-[14px] glass glass-hover">
            <span class="icon-box pink">≋</span>
            <div><strong class="block text-lg">4.2k</strong><span class="text-xs text-muted">Token 用量</span></div>
          </div>
        </section>

        <section class="grid grid-cols-2 gap-5 mb-5">
          <AppCard title="最近项目">
            <template #actions>
              <router-link to="/project/demo" class="text-primary text-xs font-semibold">查看全部</router-link>
            </template>
            <div class="flex flex-col gap-0">
              <router-link v-for="p in recentProjects" :key="p.name" to="/project/demo" class="flex items-center gap-3 py-3 border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors">
                <span class="icon-box" :class="p.iconClass">{{ p.icon }}</span>
                <div class="flex-1 min-w-0"><b class="block text-sm truncate">{{ p.name }}</b><span class="text-xs text-muted">{{ p.desc }}</span></div>
                <span class="text-xs text-muted whitespace-nowrap">{{ p.time }}</span>
              </router-link>
            </div>
          </AppCard>

          <AppCard title="AI 动态">
            <template #actions>
              <router-link to="/prompt-studio" class="text-primary text-xs font-semibold">查看全部动态</router-link>
            </template>
            <div class="flex flex-col gap-0">
              <div v-for="a in aiActivity" :key="a.title + a.time" class="flex items-center gap-3 py-3 border-b border-line last:border-0">
                <span class="icon-box">{{ a.icon }}</span>
                <div class="flex-1 min-w-0"><b class="block text-sm">{{ a.title }}</b><span class="text-xs text-muted">{{ a.desc }}</span></div>
                <span class="text-xs text-muted whitespace-nowrap">{{ a.time }}</span>
              </div>
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
