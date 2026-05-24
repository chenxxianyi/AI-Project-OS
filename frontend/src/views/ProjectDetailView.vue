<script setup lang="ts">
import { ref } from 'vue'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'

const activeTab = ref(0)
const tabs = ['概览', '提示词', '规则', '项目大脑', '设置']

const techStack = [
  { logo: 'N', name: 'Next.js', color: '' },
  { logo: '≈', name: 'Tailwind CSS', color: 'text-[#21b8d5]' },
  { logo: '⚡', name: 'Supabase', color: 'text-green', iconClass: 'green' },
  { logo: '//', name: 'Shadcn', color: '' },
  { logo: '◎', name: 'OpenAI', color: '' },
  { logo: 'TS', name: 'TypeScript', color: 'text-white', bg: 'bg-[#3178c6]' },
]

const recentPrompts = [
  { name: '用户认证流程', time: '2分钟前' },
  { name: '订阅计划组件', time: '5小时前' },
  { name: '仪表盘分析组件', time: '1天前' },
]

const activity = [
  { icon: '☷', title: '已生成提示词', desc: '用户认证流程', time: '2分钟前' },
  { icon: '⛓', title: '规则已更新', desc: '营销智能体', time: '15分钟前' },
  { icon: '▧', title: '上下文已刷新', desc: 'AI 文档', time: '1小时前' },
  { icon: '⌬', iconClass: 'green', title: '智能体已执行', desc: '数据库迁移', time: '2小时前' },
]
</script>

<template>
  <div class="p-[32px_36px]">
    <div class="workspace-top flex items-center justify-between mb-6">
      <div class="breadcrumb flex items-center gap-2 text-sm">
        <span>▣</span><span>›</span><span class="text-muted">项目</span><span>›</span><b>AI SaaS Starter</b>
      </div>
      <div class="flex gap-3">
        <AppButton>↗ 在编辑器中打开</AppButton>
        <AppButton>⋯</AppButton>
      </div>
    </div>

    <section class="project-hero flex items-start gap-5 mb-6 p-6 rounded-[14px] bg-white/78 border border-line shadow-card">
      <div class="project-logo w-14 h-14 rounded-2xl bg-gradient-to-br from-[#eef3ff] to-white border border-line grid place-items-center text-2xl text-primary shadow-card">▱</div>
      <div>
        <h1 class="text-xl font-bold mb-2">AI SaaS Starter</h1>
        <p class="text-muted text-sm mb-3">基于 Next.js 和现代工具构建的 AI SaaS 启动套件，用于快速开发应用。</p>
        <div class="flex gap-2.5">
          <AppChip label="Next.js" /><AppChip label="Tailwind" /><AppChip label="Supabase" /><AppChip label="OpenAI" />
        </div>
      </div>
    </section>

    <nav class="tabs flex gap-1 border-b border-line mb-5">
      <a v-for="(tab, i) in tabs" :key="i" class="tab px-4 py-2.5 text-sm font-semibold cursor-pointer transition-colors border-b-2"
        :class="activeTab === i ? 'text-primary border-primary' : 'text-muted border-transparent hover:text-[#08163d]'"
        @click="activeTab = i">{{ tab }}</a>
    </nav>

    <section v-if="activeTab === 0" class="project-grid grid grid-cols-2 gap-5">
      <AppCard title="项目概览">
        <template #icon><span class="icon-box">▧</span></template>
        <div class="text-[#53617f] leading-[1.8] text-sm mb-4">
          AI SaaS Starter 是一个生产就绪的启动套件，用于构建带认证、支付和现代技术栈的 AI 驱动 SaaS 应用。
        </div>
        <div class="flex flex-col gap-3">
          <div class="flex justify-between text-sm"><span class="text-muted">◷ 创建时间</span><b>2024年5月12日</b></div>
          <div class="flex justify-between text-sm"><span class="text-muted">↧ 最后更新</span><b>2小时前</b></div>
          <div class="flex justify-between text-sm"><span class="text-muted">⊙ 可见性</span><b>私有</b></div>
          <div class="flex justify-between text-sm items-center"><span class="text-muted">♙ 成员</span><div class="member-stack"><i></i><i></i><i></i><i></i><AppChip label="+3" /></div></div>
        </div>
      </AppCard>

      <AppCard title="技术栈">
        <template #icon><span class="icon-box">▱</span></template>
        <div class="grid grid-cols-3 gap-3">
          <div v-for="t in techStack" :key="t.name" class="flex items-center gap-2.5 p-3 rounded-[10px] border border-line bg-[#f8fbff]">
            <span class="tech-logo w-8 h-8 rounded-lg grid place-items-center text-sm font-bold border border-line bg-white" :class="[t.color, t.bg]">{{ t.logo }}</span>
            <span class="text-sm font-semibold">{{ t.name }}</span>
          </div>
        </div>
      </AppCard>

      <AppCard title="最近提示词">
        <template #icon><span class="icon-box">☷</span></template>
        <template #actions>
          <router-link to="/prompt-studio" class="text-primary text-xs font-semibold">查看全部提示词 →</router-link>
        </template>
        <div class="flex flex-col gap-0">
          <div v-for="p in recentPrompts" :key="p.name" class="flex items-center gap-3 py-3 border-b border-line last:border-0">
            <span class="icon-box">▣</span>
            <div class="flex-1"><b class="text-sm">{{ p.name }}</b></div>
            <span class="text-xs text-muted">{{ p.time }}</span>
          </div>
        </div>
      </AppCard>

      <AppCard title="项目动态">
        <template #icon><span class="icon-box">⌁</span></template>
        <template #actions>
          <router-link to="/dashboard" class="text-primary text-xs font-semibold">查看全部动态 →</router-link>
        </template>
        <div class="flex flex-col gap-0">
          <div v-for="a in activity" :key="a.title + a.time" class="flex items-center gap-3 py-3 border-b border-line last:border-0">
            <span class="icon-box" :class="a.iconClass">{{ a.icon }}</span>
            <div class="flex-1 min-w-0"><b class="text-sm">{{ a.title }}</b><span class="text-xs text-muted ml-2">{{ a.desc }}</span></div>
            <span class="text-xs text-muted whitespace-nowrap">{{ a.time }}</span>
          </div>
        </div>
      </AppCard>
    </section>

    <div v-else class="text-center py-20 text-muted">
      <p class="text-lg font-semibold mb-2">{{ tabs[activeTab] }}</p>
      <p class="text-sm">此标签页内容将在后续迭代中实现</p>
    </div>
  </div>
</template>
