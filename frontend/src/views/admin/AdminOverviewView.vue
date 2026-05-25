<script setup lang="ts">
import { onMounted } from 'vue'
import { useAdminStore } from '@/stores/admin'
import AppCard from '@/components/ui/AppCard.vue'

const store = useAdminStore()

const statusColors: Record<string, string> = {
  healthy: 'text-green',
  degraded: 'text-[#e5a000]',
  down: 'text-red',
  unknown: 'text-muted',
}

const statusDots: Record<string, string> = {
  healthy: 'bg-green',
  degraded: 'bg-[#e5a000]',
  down: 'bg-red',
  unknown: 'bg-muted',
}

onMounted(async () => {
  await Promise.all([
    store.fetchStats(),
    store.fetchHealth(),
  ])
})

function formatNumber(n: number) {
  if (n >= 1_000_000) return `${(n / 1_000_000).toFixed(1)}M`
  if (n >= 1_000) return `${(n / 1_000).toFixed(1)}K`
  return String(n)
}
</script>

<template>
  <div class="p-[32px_36px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">系统概览</h1>
      <p class="text-muted mt-1">AI Project OS 管理后台 — 实时系统状态与关键指标。</p>
    </section>

    <!-- Key Metrics -->
    <section class="stats grid grid-cols-4 gap-4 mb-6">
      <div class="stat flex items-center gap-3 p-5 rounded-[14px] glass glass-hover">
        <span class="icon-box relative z-10">♙</span>
        <div class="relative z-10">
          <strong class="block text-2xl">{{ formatNumber(store.stats.total_users) }}</strong>
          <span class="text-xs text-muted">总用户</span>
        </div>
      </div>
      <div class="stat flex items-center gap-3 p-5 rounded-[14px] glass glass-hover">
        <span class="icon-box green relative z-10">▣</span>
        <div class="relative z-10">
          <strong class="block text-2xl">{{ formatNumber(store.stats.total_projects) }}</strong>
          <span class="text-xs text-muted">总项目</span>
        </div>
      </div>
      <div class="stat flex items-center gap-3 p-5 rounded-[14px] glass glass-hover">
        <span class="icon-box purple relative z-10">✧</span>
        <div class="relative z-10">
          <strong class="block text-2xl">{{ formatNumber(store.stats.total_generations) }}</strong>
          <span class="text-xs text-muted">AI 生成</span>
        </div>
      </div>
      <div class="stat flex items-center gap-3 p-5 rounded-[14px] glass glass-hover">
        <span class="icon-box pink relative z-10">≋</span>
        <div class="relative z-10">
          <strong class="block text-2xl">{{ formatNumber(store.stats.total_tokens_used) }}</strong>
          <span class="text-xs text-muted">Token 用量</span>
        </div>
      </div>
    </section>

    <div class="grid grid-cols-[1fr_320px] gap-5">
      <div class="flex flex-col gap-5">
        <!-- Activity Chart -->
        <AppCard title="平台活跃度 ⓘ">
          <template #actions>
            <span class="text-xs text-muted">近 7 天</span>
          </template>
          <div class="chart h-[200px]">
            <svg viewBox="0 0 760 200" preserveAspectRatio="none" class="w-full h-full">
              <path d="M0 160 C60 140 80 100 130 110 S200 155 260 105 S340 150 400 75 S480 130 540 100 S620 145 680 95 S740 60 760 55" fill="none" stroke="#4167ff" stroke-width="3"/>
              <path d="M0 160 C60 140 80 100 130 110 S200 155 260 105 S340 150 400 75 S480 130 540 100 S620 145 680 95 S740 60 760 55 L760 200 L0 200Z" fill="rgba(65,103,255,.08)"/>
              <path d="M0 170 C60 155 80 135 130 140 S200 165 260 130 S340 160 400 110 S480 145 540 125 S620 155 680 130 S740 100 760 95" fill="none" stroke="#27c990" stroke-width="2" stroke-dasharray="6 4"/>
            </svg>
            <div class="flex gap-5 mt-3">
              <span class="flex items-center gap-1.5 text-xs text-muted"><span class="w-3 h-0.5 bg-primary rounded"></span>生成次数</span>
              <span class="flex items-center gap-1.5 text-xs text-muted"><span class="w-3 h-0.5 bg-green rounded" style="border-bottom: 2px dashed #27c990"></span>活跃用户</span>
            </div>
          </div>
        </AppCard>

        <!-- Secondary Stats -->
        <div class="grid grid-cols-3 gap-4">
          <div class="p-4 rounded-[14px] glass glass-hover">
            <div class="relative z-10 flex items-center justify-between mb-2">
              <span class="text-xs text-muted">提示词总数</span>
              <span class="icon-box">▤</span>
            </div>
            <strong class="relative z-10 block text-xl">{{ formatNumber(store.stats.total_prompts) }}</strong>
          </div>
          <div class="p-4 rounded-[14px] glass glass-hover">
            <div class="relative z-10 flex items-center justify-between mb-2">
              <span class="text-xs text-muted">活跃用户</span>
              <span class="icon-box green">◉</span>
            </div>
            <strong class="relative z-10 block text-xl">{{ formatNumber(store.stats.active_users) }}</strong>
          </div>
          <div class="p-4 rounded-[14px] glass glass-hover">
            <div class="relative z-10 flex items-center justify-between mb-2">
              <span class="text-xs text-muted">失败生成</span>
              <span class="icon-box" :class="store.stats.failed_generations > 0 ? 'pink' : ''">△</span>
            </div>
            <strong class="relative z-10 block text-xl" :class="store.stats.failed_generations > 0 ? 'text-red' : ''">{{ store.stats.failed_generations }}</strong>
          </div>
        </div>
      </div>

      <!-- System Health -->
      <aside class="flex flex-col gap-5">
        <div class="p-5 rounded-[14px] glass-panel">
          <h3 class="relative z-10 font-bold text-sm mb-4">系统健康</h3>
          <div class="relative z-10 flex flex-col gap-3">
            <div v-for="(value, key) in { database: store.health.database, redis: store.health.redis, ai_provider: store.health.ai_provider }" :key="key" class="flex items-center justify-between py-2 border-b border-line last:border-0">
              <span class="text-sm text-[#425071]">{{ { database: '数据库', redis: 'Redis', ai_provider: 'AI 服务' }[key] }}</span>
              <span class="flex items-center gap-1.5 text-sm font-semibold" :class="statusColors[value] || statusColors.unknown">
                <span class="w-2 h-2 rounded-full" :class="statusDots[value] || statusDots.unknown"></span>
                {{ { healthy: '正常', degraded: '降级', down: '故障', unknown: '未知' }[value] || '未知' }}
              </span>
            </div>
          </div>
        </div>

        <div class="p-5 rounded-[14px] glass-panel">
          <h3 class="relative z-10 font-bold text-sm mb-4">资源使用</h3>
          <div class="relative z-10 flex flex-col gap-4">
            <div>
              <div class="flex items-center justify-between mb-1.5">
                <span class="text-xs text-muted">内存</span>
                <span class="text-xs font-semibold">{{ store.health.memory_usage }}</span>
              </div>
              <div class="h-2 rounded-full bg-line overflow-hidden">
                <div class="h-full rounded-full bg-primary transition-all duration-500" :style="{ width: store.health.memory_usage }"></div>
              </div>
            </div>
            <div>
              <div class="flex items-center justify-between mb-1.5">
                <span class="text-xs text-muted">磁盘</span>
                <span class="text-xs font-semibold">{{ store.health.disk_usage }}</span>
              </div>
              <div class="h-2 rounded-full bg-line overflow-hidden">
                <div class="h-full rounded-full bg-green transition-all duration-500" :style="{ width: store.health.disk_usage }"></div>
              </div>
            </div>
          </div>
        </div>

        <div class="p-5 rounded-[14px] glass glass-hover">
          <h3 class="relative z-10 font-bold text-sm mb-3">系统信息</h3>
          <div class="relative z-10 flex flex-col gap-2">
            <div class="flex items-center justify-between text-sm">
              <span class="text-muted">运行时间</span>
              <span class="font-semibold">{{ store.stats.uptime || '—' }}</span>
            </div>
            <div class="flex items-center justify-between text-sm">
              <span class="text-muted">版本</span>
              <span class="font-semibold">{{ store.stats.version || '—' }}</span>
            </div>
          </div>
        </div>
      </aside>
    </div>
  </div>
</template>
