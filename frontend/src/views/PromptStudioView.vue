<script setup lang="ts">
import { ref } from 'vue'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'

const toast = useToast()
const running = ref(false)

const templates = [
  { icon: '▱', name: '通用', desc: '通用目的提示词', active: true },
  { icon: '⌬', name: '代码生成', desc: '基于上下文生成代码' },
  { icon: '↻', name: '审查与重构', desc: '审查和改进代码' },
  { icon: '◔', name: '数据分析', desc: '分析和可视化数据' },
  { icon: '✎', name: '内容写作', desc: '创建高质量内容' },
  { icon: '▤', name: '摘要', desc: '总结长内容' },
  { icon: '⚙', name: '自定义', desc: '从零开始' },
]

const promptContent = `# 角色
你是一位资深软件工程师。

## 上下文
你正在开发 {{project_name}} 项目
使用 {{tech_stack}}。

## 任务
{{task_description}}

## 要求
- 遵循最佳实践
- 编写整洁、可维护的代码
- 在必要时添加注释

## 输出格式
- 提供代码
- 解释实现思路`

const outputCode = `1  // authentication.js
2  import jwt from 'jsonwebtoken';
3  import { User } from './models/User.js';
4
5  export const authenticate = async (req, res, next) => {
6    try {
7      const token = req.headers.authorization
8        ?.split(' ')[1];
9
10     if (!token) {
11       return res.status(401).json({
12         success: false,
13         message: 'No token provided'
14       });
15     }`

function runPrompt() {
  running.value = true
  setTimeout(() => {
    running.value = false
    toast.show('提示词运行完成')
  }, 1500)
}
</script>

<template>
  <div class="p-[32px_36px]">
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
            <h3 class="font-bold text-sm">未命名提示词</h3>
            <div class="flex items-center gap-2">
              <span class="text-xs text-[#21af7d]">✓ 已保存</span>
              <AppButton variant="ghost" class="!min-h-[32px] !px-2.5">⋯</AppButton>
            </div>
          </div>
        </template>
        <div class="editor-box p-5 font-mono text-sm leading-relaxed whitespace-pre-wrap text-[#53617f] min-h-[400px]">{{ promptContent }}</div>
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
            <select class="w-full h-9 px-3 rounded-[10px] border border-line bg-white text-sm">
              <option>GPT-4o</option>
              <option>Claude 3.5 Sonnet</option>
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
            <p class="text-xs text-muted">已生成 2,048 Token · 12秒 · GPT-4o · 今天 10:15</p>
          </div>
        </div>
        <div class="output px-4 pb-3">
          <pre class="font-mono text-xs leading-relaxed text-[#53617f] bg-[#f8fbff] p-3 rounded-[10px] overflow-auto max-h-[280px]">{{ outputCode }}</pre>
        </div>
        <div class="flex gap-3 px-4 pb-4">
          <AppButton @click="toast.show('已复制到剪贴板')">▣ 复制</AppButton>
          <AppButton>⇩ 导出⌄</AppButton>
          <AppButton>⌫ 清空</AppButton>
        </div>
      </AppCard>
    </section>
  </div>
</template>
