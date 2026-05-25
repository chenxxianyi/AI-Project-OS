<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useAppStore } from '@/stores/app'
import { useProjectStore } from '@/stores/project'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'
import AppModal from '@/components/ui/AppModal.vue'

const appStore = useAppStore()
const projectStore = useProjectStore()
const toast = useToast()

const search = ref('')
const error = ref('')

const filteredProjects = computed(() => {
  const keyword = search.value.trim().toLowerCase()
  if (!keyword) return projectStore.projects

  return projectStore.projects.filter((project) => {
    const searchable = [
      project.name,
      project.description,
      project.project_type,
      project.frontend_stack,
      project.backend_stack,
      project.database_stack,
      project.ai_stack,
    ]

    return searchable.some((value) => value?.toLowerCase().includes(keyword))
  })
})

function timeAgo(dateStr: string) {
  const date = new Date(dateStr)
  if (Number.isNaN(date.getTime())) return '-'

  const diff = Date.now() - date.getTime()
  const mins = Math.floor(diff / 60000)
  if (mins < 1) return '刚刚'
  if (mins < 60) return `${mins}分钟前`

  const hours = Math.floor(mins / 60)
  if (hours < 24) return `${hours}小时前`

  const days = Math.floor(hours / 24)
  return `${days}天前`
}

async function loadProjects() {
  try {
    error.value = ''
    await projectStore.fetchProjects({ page: 1, page_size: 50 })
  } catch (e: unknown) {
    const message = e instanceof Error ? e.message : '项目加载失败'
    error.value = message
    toast.show(message)
  }
}

onMounted(loadProjects)
</script>

<template>
  <div class="p-[32px_36px]">
    <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>

    <div class="workspace-top flex items-center justify-between mb-6">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[340px]">
        <span class="text-muted">⌕</span>
        <input v-model="search" class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索项目、技术栈、描述..." />
      </div>
      <AppButton variant="primary" @click="appStore.openCreateModal()">＋ 新建项目</AppButton>
    </div>

    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">项目</h1>
      <p class="text-muted mt-1">管理所有项目，并进入项目工作区。</p>
    </section>

    <AppCard title="全部项目" no-pad>
      <template #actions>
        <span class="text-xs font-semibold text-muted">{{ filteredProjects.length }} / {{ projectStore.total }}</span>
      </template>

      <div v-if="projectStore.loading" class="p-8 text-center text-muted text-sm">项目加载中...</div>

      <div v-else-if="filteredProjects.length === 0" class="p-10 text-center">
        <h2 class="text-lg font-bold mb-2">暂无匹配项目</h2>
        <p class="text-muted mb-6">创建一个项目后，就可以进入 Prompt、规则和项目大脑工作流。</p>
        <AppButton variant="primary" @click="appStore.openCreateModal()">创建项目</AppButton>
      </div>

      <div v-else class="divide-y divide-line">
        <router-link
          v-for="project in filteredProjects"
          :key="project.id"
          :to="{ name: 'project-detail', params: { id: project.id } }"
          class="flex items-center gap-4 px-5 py-4 hover:bg-[#f8fbff] transition-colors"
        >
          <span class="icon-box shrink-0">▣</span>
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-1">
              <b class="text-sm truncate">{{ project.name }}</b>
              <AppChip v-if="project.status" :label="project.status === 'active' ? '活跃' : project.status" />
            </div>
            <p class="text-xs text-muted truncate">{{ project.description || project.product_goal || '暂无项目描述' }}</p>
            <div class="flex flex-wrap gap-2 mt-3">
              <AppChip v-if="project.project_type" :label="project.project_type" />
              <AppChip v-if="project.frontend_stack" :label="project.frontend_stack" />
              <AppChip v-if="project.backend_stack" :label="project.backend_stack" />
              <AppChip v-if="project.ai_stack" :label="project.ai_stack" />
            </div>
          </div>
          <span class="text-xs text-muted whitespace-nowrap">{{ timeAgo(project.updated_at) }}</span>
        </router-link>
      </div>
    </AppCard>
  </div>

  <AppModal />
</template>
