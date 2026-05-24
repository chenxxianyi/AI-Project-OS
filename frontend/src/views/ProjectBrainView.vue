<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useProjectStore } from '@/stores/project'
import { docApi, type ProjectDoc } from '@/api/doc'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'
import AppChip from '@/components/ui/AppChip.vue'

const route = useRoute()
const toast = useToast()
const projectStore = useProjectStore()
const projectId = computed(() => route.params.id as string)

const activeTab = ref(0)
const brainTabs = ['⌘ 知识图谱', '▧ 文档', '⌘ API 文档', '↱ 时间线']

const docs = ref<ProjectDoc[]>([])
const selectedDoc = ref<ProjectDoc | null>(null)
const error = ref('')

const nodes = [
  { id: 'n1', icon: '▣', label: '认证', cls: 'n1' },
  { id: 'n2', icon: '▤', label: '支付', cls: 'n2 purple' },
  { id: 'n3', icon: '▱', label: '组件', cls: 'n3' },
  { id: 'n4', icon: '♙', label: '用户管理', cls: 'n4 green' },
  { id: 'n5', icon: '♙', label: '用户管理', cls: 'n5 purple' },
  { id: 'n6', icon: '✹', label: 'AI 引擎', cls: 'n6 cyan' },
  { id: 'n7', icon: '▤', label: '数据库', cls: 'n7 green' },
]

const relatedNodes = [
  { icon: '♙', iconClass: 'green', name: '用户管理', tag: '依赖于' },
  { icon: '▤', iconClass: 'green', name: '数据库', tag: '使用' },
  { icon: '✹', iconClass: 'cyan', name: 'AI 引擎', tag: '集成于' },
  { icon: '▤', iconClass: 'purple', name: '支付', tag: '可选' },
]

const project = computed(() => projectStore.currentProject)

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
    if (projectId.value) {
      await projectStore.fetchProject(projectId.value)
      const res = await docApi.list(projectId.value)
      docs.value = res.data.data
      if (docs.value.length > 0) {
        selectedDoc.value = docs.value[0]
      }
    } else {
      error.value = '缺少项目 ID'
    }
  } catch (e: unknown) {
    const message = e instanceof Error ? e.message : '项目大脑加载失败'
    error.value = message
    toast.show(message)
  }
})
</script>

<template>
  <div class="p-[32px_36px]">
    <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>
    <div class="breadcrumb flex items-center gap-2 text-sm mb-4">
      <span>＋</span><span>›</span><span class="text-muted">项目</span><span>›</span><span>＋</span><b>{{ project?.name || '...' }}</b>
    </div>

    <div class="workspace-top flex items-end justify-between mb-5">
      <section class="headline">
        <h1 class="text-2xl font-bold">项目大脑</h1>
        <p class="text-muted mt-1">你的项目知识和上下文中心。</p>
      </section>
      <AppButton>＋ 添加知识</AppButton>
    </div>

    <section class="brain-card rounded-[14px] bg-white/78 border border-line shadow-card overflow-hidden">
      <nav class="tabs flex gap-1 border-b border-line px-5">
        <a v-for="(tab, i) in brainTabs" :key="i" class="tab px-4 py-2.5 text-sm font-semibold cursor-pointer transition-colors border-b-2"
          :class="activeTab === i ? 'text-primary border-primary' : 'text-muted border-transparent hover:text-[#08163d]'"
          @click="activeTab = i">{{ tab }}</a>
      </nav>

      <div v-if="activeTab === 0" class="brain-area grid grid-cols-[1fr_300px] gap-0">
        <!-- Knowledge Graph -->
        <div class="graph relative h-[620px] overflow-hidden">
          <svg viewBox="0 0 760 620" preserveAspectRatio="none" class="w-full h-full">
            <g stroke="#9cafed" stroke-width="1.4" opacity=".8">
              <line x1="380" y1="315" x2="390" y2="120"/>
              <line x1="380" y1="315" x2="180" y2="200"/>
              <line x1="380" y1="315" x2="92" y2="365"/>
              <line x1="380" y1="315" x2="360" y2="500"/>
              <line x1="380" y1="315" x2="565" y2="460"/>
              <line x1="380" y1="315" x2="625" y2="335"/>
              <line x1="380" y1="315" x2="585" y2="205"/>
            </g>
            <g fill="#edf2ff" opacity=".9">
              <circle cx="265" cy="180" r="6"/>
              <circle cx="135" cy="285" r="4"/>
              <circle cx="520" cy="260" r="6"/>
              <circle cx="655" cy="430" r="5"/>
              <circle cx="320" cy="390" r="6"/>
              <circle cx="230" cy="420" r="4"/>
              <circle cx="460" cy="170" r="5"/>
            </g>
          </svg>

          <!-- Center Node -->
          <div class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 flex items-center gap-2 px-4 py-3 rounded-2xl bg-white border-2 border-primary shadow-primary cursor-pointer">
            <span class="icon-box">⌘</span><b class="text-sm font-bold">{{ project?.name || '项目' }}</b>
          </div>

          <!-- Satellite Nodes -->
          <div v-for="n in nodes" :key="n.id" class="node absolute flex items-center gap-2 px-3 py-2 rounded-xl bg-white/90 border border-line shadow-card cursor-pointer hover:border-primary/40 transition-colors"
            :class="n.cls">
            <span class="icon-box" :class="n.cls.split(' ')[1] || ''">{{ n.icon }}</span>
            <span class="text-xs font-semibold">{{ n.label }}</span>
          </div>

          <!-- Zoom Controls -->
          <div class="zoom absolute bottom-4 left-4 flex gap-2">
            <button v-for="z in ['−','＋','⛶','▣']" :key="z" class="w-8 h-8 rounded-lg bg-white border border-line shadow-btn grid place-items-center text-sm hover:bg-[#f8fbff] transition">{{ z }}</button>
          </div>
        </div>

        <!-- Detail Panel -->
        <aside class="detail border-l border-line p-5 overflow-auto">
          <h3 class="font-bold text-sm mb-4">节点详情</h3>
          <div v-if="selectedDoc" class="flex items-center gap-3 mb-3">
            <span class="icon-box big">▧</span>
            <div><h3 class="font-bold text-base">{{ selectedDoc.title }}</h3><AppChip :label="selectedDoc.doc_type" /></div>
          </div>
          <p v-if="selectedDoc" class="text-muted text-sm leading-relaxed">{{ selectedDoc.content?.slice(0, 200) }}{{ selectedDoc.content?.length > 200 ? '...' : '' }}</p>

          <div class="h-px bg-line my-5" />
          <h4 class="font-bold text-sm mb-3">关联节点</h4>
          <div class="flex flex-col gap-2">
            <div v-for="r in relatedNodes" :key="r.name" class="flex items-center gap-2.5 py-2">
              <span class="icon-box" :class="r.iconClass">{{ r.icon }}</span>
              <b class="text-sm flex-1">{{ r.name }}</b>
              <span class="tag text-[11px] text-muted bg-[#f1f5ff] px-2 py-0.5 rounded">{{ r.tag }}</span>
            </div>
          </div>

          <div class="h-px bg-line my-5" />
          <h4 class="font-bold text-sm mb-3">知识文档</h4>
          <div class="flex flex-col gap-0">
            <div v-for="d in docs" :key="d.id" class="flex items-center gap-2.5 py-2.5 border-b border-line last:border-0 cursor-pointer hover:bg-[#f8fbff]"
              @click="selectedDoc = d">
              <span class="icon-box">▧</span>
              <b class="text-sm flex-1">{{ d.title }}</b>
              <span class="text-xs text-muted">{{ timeAgo(d.updated_at) }}</span>
            </div>
          </div>

          <AppButton variant="primary" block class="mt-5">在上下文中打开 ↗</AppButton>
        </aside>
      </div>

      <div v-else class="p-10 text-center text-muted">
        <p class="text-lg font-semibold mb-2">{{ brainTabs[activeTab] }}</p>
        <p class="text-sm">此标签页内容将在后续迭代中实现</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.node.n1 { top: 15%; left: 48%; }
.node.n2 { top: 25%; left: 18%; }
.node.n3 { top: 15%; right: 18%; }
.node.n4 { top: 55%; left: 8%; }
.node.n5 { top: 75%; left: 42%; }
.node.n6 { top: 70%; right: 18%; }
.node.n7 { top: 50%; right: 8%; }
</style>
