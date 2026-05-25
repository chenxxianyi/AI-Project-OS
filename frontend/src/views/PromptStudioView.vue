<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useToast } from '@/composables/useToast'
import { useProjectStore } from '@/stores/project'
import { promptApi, type Prompt } from '@/api/prompt'
import { generationApi } from '@/api/generation'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'

const toast = useToast()
const route = useRoute()
const projectStore = useProjectStore()
const running = ref(false)
const selectedPrompt = ref<Prompt | null>(null)
const outputContent = ref('')
const error = ref('')
const projectId = ref('')
const selectedModel = ref('deepseek-chat')

const modelOptions = [
  { value: 'deepseek-chat', label: 'DeepSeek V3 (deepseek-chat)' },
  { value: 'deepseek-reasoner', label: 'DeepSeek R1 (deepseek-reasoner)' },
]

const templates = [
  { icon: '▱', name: '通用', desc: '通用目的提示词', active: true },
  { icon: '⌬', name: '代码生成', desc: '基于上下文生成代码' },
  { icon: '↻', name: '审查与重构', desc: '审查和改进代码' },
  { icon: '◔', name: '数据分析', desc: '分析和可视化数据' },
  { icon: '✎', name: '内容写作', desc: '创建高质量内容' },
  { icon: '▤', name: '摘要', desc: '总结长内容' },
  { icon: '⚙', name: '自定义', desc: '从零开始' },
]

const promptContent = ref(`# 角色\n你是一位资深软件工程师。\n\n## 上下文\n你正在开发 {{project_name}} 项目\n使用 {{tech_stack}}。\n\n## 任务\n{{task_description}}\n\n## 要求\n- 遵循最佳实践\n- 编写整洁、可维护的代码\n- 在必要时添加注释\n\n## 输出格式\n- 提供代码\n- 解释实现思路`)

onMounted(async () => {
  try {
    error.value = ''
    projectId.value = route.params.id as string
    if (!projectId.value) {
      error.value = '缺少项目 ID'
      return
    }
    await projectStore.fetchProject(projectId.value)
    const res = await promptApi.list(projectId.value)
    const prompts = res.data.data
    if (prompts.length > 0) {
      selectedPrompt.value = prompts[0]
      promptContent.value = prompts[0].content
    }
  } catch (e: unknown) {
    const message = e instanceof Error ? e.message : '提示词加载失败'
    error.value = message
    toast.show(message)
  }
})

async function runPrompt() {
  if (!projectId.value) {
    toast.show('缺少项目 ID')
    return
  }
  running.value = true
  try {
    const res = await generationApi.generate({
      project_id: projectId.value,
      generation_type: 'prompt',
      input_payload: promptContent.value,
    })
    outputContent.value = res.data.data.output_content
    toast.show('提示词运行完成')
  } catch (e: unknown) {
    toast.show(e instanceof Error ? e.message : '运行失败')
  } finally {
    running.value = false
  }
}
</script>

<template>
  <div class="p-[32px_36px]">
    <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>
    <div class="workspace-top flex items-center justify-between mb-6">
      <section class="headline">
        <h1 class="text-2xl font-bold">Prompt 工作室</h1>
        <p class="text-muted mt-1">设计、测试和优化你的提示词。</p>
      </section>
      <div class="flex gap-3">
        <AppButton>▦ 模板</AppButton>
        <AppButton>◷ 历史</AppButton>
        <AppButton variant="primary">＋ 新建提示词</AppButton>
      </div>
    </div>

    <section class="studio-layout grid grid-cols-[240px_1fr_300px] gap-5">
      <!-- Template List -->
      <AppCard title="模板" no-pad>
        <div class="px-3.5 pb-3 flex gap-2">
          <div class="search flex items-center gap-2 h-9 px-3 rounded-[10px] bg-white border border-line w-full">
            <span class="text-muted text-sm">⌕</span>
            <input class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索模板..." />
          </div>
        </div>
        <div class="flex flex-col gap-0">
          <div v-for="t in templates" :key="t.name" class="flex items-center gap-3 px-4 py-3 cursor-pointer transition-colors"
            :class="t.active ? 'bg-[#f0f4ff] border-l-2 border-primary' : 'hover:bg-[#f8fbff]'">
            <span class="icon-box">{{ t.icon }}</span>
            <div><b class="block text-sm">{{ t.name }}</b><span class="text-xs text-muted">{{ t.desc }}</span></div>
          </div>
        </div>
        <div class="px-3.5 py-4">
          <AppButton block>＋ 新建模板</AppButton>
        </div>
      </AppCard>

      <!-- Editor -->
      <AppCard no-pad>
        <template #header>
          <div class="panel-head flex items-center justify-between px-5 py-4 border-b border-line">
            <h3 class="font-bold text-sm">{{ selectedPrompt?.title || '未命名提示词' }}</h3>
            <div class="flex items-center gap-2">
              <span v-if="selectedPrompt" class="text-xs text-[#21af7d]">✓ 已保存</span>
              <AppButton variant="ghost" class="!min-h-[32px] !px-2.5">⋯</AppButton>
            </div>
          </div>
        </template>
        <textarea v-model="promptContent" class="editor-box p-5 font-mono text-sm leading-relaxed text-[#53617f] min-h-[400px] w-full resize-none outline-none bg-transparent border-0" spellcheck="false" />
        <div class="toolbar-bottom flex items-center gap-4 px-5 py-3 border-t border-line text-sm text-muted">
          <span class="cursor-pointer hover:text-primary">＋ 变量</span>
          <span>Markdown⌄</span>
        </div>
      </AppCard>

      <!-- Test & Output -->
      <AppCard title="测试提示词" no-pad>
        <div class="p-4 flex flex-col gap-4">
          <div>
            <label class="block text-sm font-semibold mb-1.5">模型</label>
            <select v-model="selectedModel" class="w-full h-9 px-3 rounded-[10px] border border-line bg-white text-sm">
              <option v-for="m in modelOptions" :key="m.value" :value="m.value">{{ m.label }}</option>
            </select>
          </div>
          <div>
            <label class="block text-sm font-semibold mb-1.5">温度</label>
            <input type="range" min="0" max="1" step="0.1" value="0.7" class="w-full" />
          </div>
          <AppButton variant="primary" block @click="runPrompt">
            {{ running ? '⏳ 运行中...' : '▷ 运行提示词' }}
          </AppButton>
        </div>
        <div class="px-4 py-3 border-t border-line">
          <div class="flex items-center justify-between mb-2">
            <h3 class="font-bold text-sm">输出</h3>
            <p v-if="outputContent" class="text-xs text-muted">已生成</p>
          </div>
        </div>
        <div class="output px-4 pb-3">
          <pre v-if="outputContent" class="font-mono text-xs leading-relaxed text-[#53617f] bg-[#f8fbff] p-3 rounded-[10px] overflow-auto max-h-[280px]">{{ outputContent }}</pre>
          <p v-else class="text-xs text-muted">点击"运行提示词"查看输出</p>
        </div>
        <div class="flex gap-3 px-4 pb-4">
          <AppButton @click="toast.show('已复制到剪贴板')">▣ 复制</AppButton>
          <AppButton>⇩ 导出⌄</AppButton>
          <AppButton @click="outputContent = ''">⌫ 清空</AppButton>
        </div>
      </AppCard>
    </section>
  </div>
</template>
