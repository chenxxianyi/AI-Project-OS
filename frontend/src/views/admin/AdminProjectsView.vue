<script setup lang="ts">
import { onMounted, ref, watch, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAdminStore } from '@/stores/admin'
import { useToast } from '@/composables/useToast'
import AppChip from '@/components/ui/AppChip.vue'

const store = useAdminStore()
const toast = useToast()
const router = useRouter()

const searchQuery = ref('')
const statusFilter = ref('')
const showDeleteConfirm = ref<string | null>(null)

const projectStatusMap: Record<string, { label: string; color: string }> = {
  active: { label: '活跃', color: 'green' },
  archived: { label: '归档', color: 'muted' },
  draft: { label: '草稿', color: '[#e5a000]' },
}

onMounted(() => {
  store.fetchProjects({ page: 1, page_size: 20 })
})

watch([searchQuery, statusFilter], () => {
  store.fetchProjects({ page: 1, page_size: 20, search: searchQuery.value, status: statusFilter.value })
})

function handlePageChange(page: number) {
  store.fetchProjects({ page, page_size: 20, search: searchQuery.value, status: statusFilter.value })
}

async function confirmDelete(projectId: string) {
  await store.deleteProject(projectId)
  showDeleteConfirm.value = null
  toast.show('项目已删除')
}

function goToProject(projectId: string) {
  router.push(`/project/${projectId}`)
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

const totalPages = computed(() => Math.ceil(store.projectsTotal / 20))
</script>

<template>
  <div class="p-[32px_36px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">项目管理</h1>
      <p class="text-muted mt-1">查看和管理所有平台项目。</p>
    </section>

    <!-- Filters -->
    <div class="flex items-center gap-3 mb-5">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[300px]">
        <span class="text-muted">⌕</span>
        <input v-model="searchQuery" class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索项目名称..." />
      </div>
      <div class="flex items-center gap-2">
        <button
          v-for="s in ['', 'active', 'archived', 'draft']"
          :key="s"
          class="px-3 py-2 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="statusFilter === s ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
          @click="statusFilter = s"
        >
          {{ s ? projectStatusMap[s]?.label || s : '全部' }}
        </button>
      </div>
      <span class="ml-auto text-sm text-muted">共 {{ store.projectsTotal }} 个项目</span>
    </div>

    <!-- Projects Table -->
    <div class="rounded-[14px] glass-panel overflow-hidden">
      <table class="w-full relative z-10">
        <thead>
          <tr class="border-b border-line text-left">
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">项目</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">类型</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">技术栈</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">状态</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">所有者</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">更新时间</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider text-right">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="project in store.projects" :key="project.id" class="border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors cursor-pointer" @click="goToProject(project.id)">
            <td class="px-5 py-3.5">
              <div class="min-w-0">
                <b class="block text-sm truncate">{{ project.name }}</b>
                <span class="text-xs text-muted truncate block">{{ project.description || '无描述' }}</span>
              </div>
            </td>
            <td class="px-5 py-3.5">
              <AppChip :label="project.project_type" />
            </td>
            <td class="px-5 py-3.5">
              <div class="flex flex-wrap gap-1.5">
                <AppChip v-if="project.frontend_stack" :label="project.frontend_stack" />
                <AppChip v-if="project.backend_stack" :label="project.backend_stack" />
              </div>
            </td>
            <td class="px-5 py-3.5">
              <span class="inline-flex items-center gap-1.5 text-xs font-semibold" :class="`text-${projectStatusMap[project.status]?.color || 'muted'}`">
                <span class="w-2 h-2 rounded-full" :class="`bg-${projectStatusMap[project.status]?.color || 'muted'}`"></span>
                {{ projectStatusMap[project.status]?.label || project.status }}
              </span>
            </td>
            <td class="px-5 py-3.5 text-sm text-[#425071]">{{ project.owner_name }}</td>
            <td class="px-5 py-3.5 text-sm text-muted">{{ timeAgo(project.updated_at) }}</td>
            <td class="px-5 py-3.5 text-right" @click.stop>
              <div v-if="showDeleteConfirm === project.id" class="flex items-center justify-end gap-2">
                <span class="text-xs text-red">确认删除?</span>
                <button class="text-red text-xs font-bold cursor-pointer" @click="confirmDelete(project.id)">确认</button>
                <button class="text-muted text-xs cursor-pointer" @click="showDeleteConfirm = null">取消</button>
              </div>
              <button v-else class="text-muted hover:text-red text-sm cursor-pointer transition-colors" @click="showDeleteConfirm = project.id">🗑</button>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-if="store.projects.length === 0 && !store.loading" class="py-12 text-center text-muted text-sm">
        暂无项目数据
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
        :class="store.projectsPage === p ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
        @click="handlePageChange(p)"
      >
        {{ p }}
      </button>
    </div>
  </div>
</template>
