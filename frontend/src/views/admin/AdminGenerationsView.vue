<script setup lang="ts">
import { onMounted, ref, watch, computed } from 'vue'
import { useAdminStore } from '@/stores/admin'
import AppChip from '@/components/ui/AppChip.vue'

const store = useAdminStore()

const statusFilter = ref('')
const typeFilter = ref('')

const genStatusMap: Record<string, { label: string; color: string }> = {
  completed: { label: '完成', color: 'green' },
  pending: { label: '处理中', color: '[#e5a000]' },
  failed: { label: '失败', color: 'red' },
}

const genTypeMap: Record<string, string> = {
  prompt: '提示词',
  rule: '规则',
  code: '代码',
  architecture: '架构',
}

onMounted(() => {
  store.fetchGenerations({ page: 1, page_size: 20 })
})

watch([statusFilter, typeFilter], () => {
  store.fetchGenerations({ page: 1, page_size: 20, status: statusFilter.value, generation_type: typeFilter.value })
})

function handlePageChange(page: number) {
  store.fetchGenerations({ page, page_size: 20, status: statusFilter.value, generation_type: typeFilter.value })
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

const totalPages = computed(() => Math.ceil(store.generationsTotal / 20))
</script>

<template>
  <div class="p-[32px_36px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">生成记录</h1>
      <p class="text-muted mt-1">AI 生成任务日志与状态追踪。</p>
    </section>

    <!-- Filters -->
    <div class="flex items-center gap-3 mb-5">
      <div class="flex items-center gap-2">
        <span class="text-xs text-muted font-semibold">状态</span>
        <button
          v-for="s in ['', 'completed', 'pending', 'failed']"
          :key="s"
          class="px-3 py-2 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="statusFilter === s ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
          @click="statusFilter = s"
        >
          {{ s ? genStatusMap[s]?.label || s : '全部' }}
        </button>
      </div>
      <div class="flex items-center gap-2 ml-4">
        <span class="text-xs text-muted font-semibold">类型</span>
        <button
          v-for="t in ['', 'prompt', 'rule', 'code', 'architecture']"
          :key="t"
          class="px-3 py-2 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="typeFilter === t ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
          @click="typeFilter = t"
        >
          {{ t ? genTypeMap[t] || t : '全部' }}
        </button>
      </div>
      <span class="ml-auto text-sm text-muted">共 {{ store.generationsTotal }} 条记录</span>
    </div>

    <!-- Generations Table -->
    <div class="rounded-[14px] glass-panel overflow-hidden">
      <table class="w-full relative z-10">
        <thead>
          <tr class="border-b border-line text-left">
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">ID</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">类型</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">模型</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">状态</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">用户</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">项目</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">时间</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">错误</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="gen in store.generations" :key="gen.id" class="border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors">
            <td class="px-5 py-3.5 text-xs text-muted font-mono">{{ gen.id.slice(0, 8) }}</td>
            <td class="px-5 py-3.5">
              <AppChip :label="genTypeMap[gen.generation_type] || gen.generation_type" />
            </td>
            <td class="px-5 py-3.5">
              <div>
                <span class="text-sm font-semibold">{{ gen.model_provider }}</span>
                <span class="text-xs text-muted ml-1.5">{{ gen.model_name }}</span>
              </div>
            </td>
            <td class="px-5 py-3.5">
              <span class="inline-flex items-center gap-1.5 text-xs font-semibold" :class="`text-${genStatusMap[gen.status]?.color || 'muted'}`">
                <span class="w-2 h-2 rounded-full" :class="`bg-${genStatusMap[gen.status]?.color || 'muted'}`"></span>
                {{ genStatusMap[gen.status]?.label || gen.status }}
              </span>
            </td>
            <td class="px-5 py-3.5 text-sm text-[#425071]">{{ gen.user_email }}</td>
            <td class="px-5 py-3.5 text-sm text-[#425071]">{{ gen.project_name }}</td>
            <td class="px-5 py-3.5 text-sm text-muted">{{ timeAgo(gen.created_at) }}</td>
            <td class="px-5 py-3.5 text-xs text-red max-w-[200px] truncate" :title="gen.error_message">{{ gen.error_message || '—' }}</td>
          </tr>
        </tbody>
      </table>

      <div v-if="store.generations.length === 0 && !store.loading" class="py-12 text-center text-muted text-sm">
        暂无生成记录
      </div>

      <div v-if="store.loading" class="py-8 text-center text-muted text-sm">
        加载中...
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="flex items-center justify-center gap-2 mt-5">
      <button
        v-for="p in totalPages"
        :key="p"
        class="w-9 h-9 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
        :class="store.generationsPage === p ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
        @click="handlePageChange(p)"
      >
        {{ p }}
      </button>
    </div>
  </div>
</template>
