<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useProjectStore } from '@/stores/project'
import { promptApi, type Prompt } from '@/api/prompt'
import { docApi, type ProjectDoc } from '@/api/doc'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'

const route = useRoute()
const projectStore = useProjectStore()
const projectId = computed(() => route.params.id as string)

const activeTab = ref(0)
const tabs = ['概览', '提示词', '规则', '项目大脑', '设置']

const prompts = ref<Prompt[]>([])
const docs = ref<ProjectDoc[]>([])
const error = ref('')

onMounted(async () => {
  try {
    error.value = ''
    await projectStore.fetchProject(projectId.value)
    const [promptRes, docRes] = await Promise.all([
      promptApi.list(projectId.value),
      docApi.list(projectId.value),
    ])
    prompts.value = promptRes.data.data
    docs.value = docRes.data.data
  } catch (e: unknown) {
    error.value = e instanceof Error ? e.message : '项目加载失败'
  }
})

const project = computed(() => projectStore.currentProject)

const techStack = computed(() => {
  if (!project.value) return []
  const stack: { logo: string; name: string; color: string; iconClass?: string }[] = []
  if (project.value.frontend_stack) stack.push({ logo: 'N', name: project.value.frontend_stack, color: '' })
  if (project.value.backend_stack) stack.push({ logo: '//', name: project.value.backend_stack, color: '' })
  if (project.value.database_stack) stack.push({ logo: '⚡', name: project.value.database_stack, color: 'text-green', iconClass: 'green' })
  if (project.value.ai_stack) stack.push({ logo: '◎', name: project.value.ai_stack, color: '' })
  return stack
})

function timeAgo(dateStr: string) {
  const diff = Date.now() - new Date(dateStr).getTime()
  const mins = Math.floor(diff / 60000)
  if (mins < 60) return `${mins}分钟前`
  const hours = Math.floor(mins / 60)
  if (hours < 24) return `${hours}小时前`
  const days = Math.floor(hours / 24)
  return `${days}天前`
}
</script>

<template>
  <div class="p-[32px_36px]">
    <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>
    <div class="workspace-top flex items-center justify-between mb-6">
      <div class="breadcrumb flex items-center gap-2 text-sm">
        <span>▣</span><span>›</span><span class="text-muted">项目</span><span>›</span><b>{{ project?.name || '加载中...' }}</b>
      </div>
      <div class="flex gap-3">
        <AppButton>↗ 在编辑器中打开</AppButton>
        <AppButton>⋯</AppButton>
      </div>
    </div>

    <section class="project-hero flex items-start gap-5 mb-6 p-6 rounded-[14px] bg-white/78 border border-line shadow-card">
      <div class="project-logo w-14 h-14 rounded-2xl bg-gradient-to-br from-[#eef3ff] to-white border border-line grid place-items-center text-2xl text-primary shadow-card">▱</div>
      <div>
        <h1 class="text-xl font-bold mb-2">{{ project?.name || '...' }}</h1>
        <p class="text-muted text-sm mb-3">{{ project?.description || '暂无描述' }}</p>
        <div class="flex gap-2.5">
          <AppChip v-if="project?.frontend_stack" :label="project.frontend_stack" />
          <AppChip v-if="project?.backend_stack" :label="project.backend_stack" />
          <AppChip v-if="project?.ai_stack" :label="project.ai_stack" />
          <AppChip :label="project?.project_type || ''" />
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
          {{ project?.product_goal || project?.description || '暂无项目描述' }}
        </div>
        <div class="flex flex-col gap-3">
          <div class="flex justify-between text-sm"><span class="text-muted">◷ 创建时间</span><b>{{ project?.created_at ? new Date(project.created_at).toLocaleDateString('zh-CN') : '-' }}</b></div>
          <div class="flex justify-between text-sm"><span class="text-muted">↧ 最后更新</span><b>{{ project?.updated_at ? timeAgo(project.updated_at) : '-' }}</b></div>
          <div class="flex justify-between text-sm"><span class="text-muted">⊙ 状态</span><b>{{ project?.status === 'active' ? '活跃' : '已归档' }}</b></div>
          <div class="flex justify-between text-sm items-center"><span class="text-muted">♙ 成员</span><div class="member-stack"><i></i><i></i><i></i><i></i><AppChip label="+3" /></div></div>
        </div>
      </AppCard>

      <AppCard title="技术栈">
        <template #icon><span class="icon-box">▱</span></template>
        <div class="grid grid-cols-3 gap-3">
          <div v-for="t in techStack" :key="t.name" class="flex items-center gap-2.5 p-3 rounded-[10px] border border-line bg-[#f8fbff]">
            <span class="tech-logo w-8 h-8 rounded-lg grid place-items-center text-sm font-bold border border-line bg-white" :class="t.color">{{ t.logo }}</span>
            <span class="text-sm font-semibold">{{ t.name }}</span>
          </div>
        </div>
      </AppCard>

      <AppCard title="最近提示词">
        <template #icon><span class="icon-box">☷</span></template>
        <template #actions>
          <router-link :to="`/project/${projectId}/prompt-studio`" class="text-primary text-xs font-semibold">查看全部提示词 →</router-link>
        </template>
        <div class="flex flex-col gap-0">
          <div v-for="p in prompts" :key="p.id" class="flex items-center gap-3 py-3 border-b border-line last:border-0">
            <span class="icon-box">▣</span>
            <div class="flex-1"><b class="text-sm">{{ p.title }}</b></div>
            <span class="text-xs text-muted">v{{ p.version }}</span>
          </div>
        </div>
      </AppCard>

      <AppCard title="项目动态">
        <template #icon><span class="icon-box">⌁</span></template>
        <template #actions>
          <router-link to="/dashboard" class="text-primary text-xs font-semibold">查看全部动态 →</router-link>
        </template>
        <div class="flex flex-col gap-0">
          <div v-for="d in docs" :key="d.id" class="flex items-center gap-3 py-3 border-b border-line last:border-0">
            <span class="icon-box">▧</span>
            <div class="flex-1 min-w-0"><b class="text-sm">{{ d.title }}</b><span class="text-xs text-muted ml-2">{{ d.doc_type }}</span></div>
            <span class="text-xs text-muted whitespace-nowrap">{{ timeAgo(d.updated_at) }}</span>
          </div>
        </div>
      </AppCard>
    </section>

    <div v-else-if="activeTab === 1" class="text-center py-20 text-muted">
      <p class="text-lg font-semibold mb-2">提示词</p>
      <router-link :to="`/project/${projectId}/prompt-studio`" class="text-primary text-sm font-semibold">打开 Prompt 工作室 →</router-link>
    </div>

    <div v-else-if="activeTab === 2" class="text-center py-20 text-muted">
      <p class="text-lg font-semibold mb-2">规则</p>
      <router-link :to="`/project/${projectId}/rules-builder`" class="text-primary text-sm font-semibold">打开规则构建器 →</router-link>
    </div>

    <div v-else-if="activeTab === 3" class="text-center py-20 text-muted">
      <p class="text-lg font-semibold mb-2">项目大脑</p>
      <router-link :to="`/project/${projectId}/brain`" class="text-primary text-sm font-semibold">打开项目大脑 →</router-link>
    </div>

    <div v-else class="text-center py-20 text-muted">
      <p class="text-lg font-semibold mb-2">{{ tabs[activeTab] }}</p>
      <p class="text-sm">此标签页内容将在后续迭代中实现</p>
    </div>
  </div>
</template>
