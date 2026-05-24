<script setup lang="ts">
import { ref } from 'vue'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'

const toast = useToast()
const activeTab = ref(0)
const ruleTabs = ['Cursor 规则', 'CLAUDE.md', 'AGENTS.md']

const categories = [
  { icon: '▧', name: '通用', active: true },
  { icon: '⌘', name: '代码风格' },
  { icon: '☆', name: '最佳实践' },
  { icon: '▱', name: '架构' },
  { icon: '⚗', name: '测试' },
  { icon: '⬡', name: '安全' },
  { icon: '⌘', name: 'Git 工作流' },
  { icon: '⊕', name: '自定义' },
]

const rulesContent = `# 通用规则

## 核心原则
- 编写清晰、简洁、可维护的代码
- 优先考虑可读性和一致性
- 简洁优于取巧

## 代码质量
- 使用有意义的变量和函数名
- 保持函数小而专注
- 避免代码重复
- 优雅地处理错误

## 文档
- 为复杂逻辑添加注释
- 保持 README.md 更新
- 记录公共 API 和函数

## 协作
- 遵循已有的项目结构
- 尊重现有模式和约定
- 提交前审查变更
- 编写清晰的提交信息

## 性能
- 考虑性能影响
- 仅在必要时优化
- 优化前先做性能分析`

const previewSections = [
  { title: '核心原则', items: ['编写清晰、简洁、可维护的代码', '优先考虑可读性和一致性', '简洁优于取巧'] },
  { title: '代码质量', items: ['使用有意义的变量和函数名', '保持函数小而专注', '避免代码重复', '优雅地处理错误'] },
  { title: '文档', items: ['为复杂逻辑添加注释', '保持 README.md 更新', '记录公共 API 和函数'] },
  { title: '协作', items: ['遵循已有的项目结构', '尊重现有模式和约定', '提交前审查变更', '编写清晰的提交信息'] },
  { title: '性能', items: ['考虑性能影响', '仅在必要时优化', '优化前先做性能分析'] },
]
</script>

<template>
  <div class="p-[32px_36px]">
    <div class="workspace-top flex items-center justify-between mb-4">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[340px]">
        <span class="text-muted">⌕</span>
        <input class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索项目、提示词、文档..." />
        <span class="kbd text-[11px] text-muted bg-[#f1f5ff] px-1.5 py-0.5 rounded">⌘ K</span>
      </div>
      <div class="w-8 h-8 rounded-full bg-gradient-to-br from-[#222d58] to-[#f3a071]" />
    </div>

    <div class="workspace-top flex items-end justify-between mb-5">
      <section class="headline">
        <h1 class="text-2xl font-bold">规则构建器</h1>
        <p class="text-muted mt-1">为你的项目生成和管理规则。</p>
      </section>
      <div class="flex gap-3">
        <AppButton>⊙ 预览</AppButton>
        <AppButton>▣ 复制</AppButton>
        <AppButton>⇩ 下载</AppButton>
      </div>
    </div>

    <nav class="tabs flex gap-1 border-b border-line mb-5">
      <a v-for="(tab, i) in ruleTabs" :key="i" class="tab px-4 py-2.5 text-sm font-semibold cursor-pointer transition-colors border-b-2"
        :class="activeTab === i ? 'text-primary border-primary' : 'text-muted border-transparent hover:text-[#08163d]'"
        @click="activeTab = i">{{ tab }}</a>
    </nav>

    <section class="rules-layout grid grid-cols-[200px_1fr_280px] gap-5">
      <!-- Categories -->
      <AppCard title="分类" no-pad>
        <div class="flex flex-col gap-0">
          <div v-for="c in categories" :key="c.name" class="flex items-center gap-2.5 px-4 py-3 cursor-pointer transition-colors text-sm font-semibold"
            :class="c.active ? 'bg-[#f0f4ff] text-primary border-l-2 border-primary' : 'hover:bg-[#f8fbff] text-[#08163d]'">
            {{ c.icon }} {{ c.name }}
          </div>
        </div>
        <div class="px-3.5 py-6">
          <AppButton block>＋ 添加分类</AppButton>
        </div>
        <p class="text-xs text-muted px-4 pb-4">✦ 规则会自动应用到你的项目。</p>
      </AppCard>

      <!-- Editor -->
      <AppCard title="生成的规则" no-pad>
        <template #actions>
          <AppButton variant="ghost" class="!min-h-[34px]">Markdown⌄</AppButton>
        </template>
        <div class="editor-box p-5 font-mono text-sm leading-relaxed whitespace-pre-wrap text-[#53617f] min-h-[600px]">{{ rulesContent }}</div>
      </AppCard>

      <!-- Preview -->
      <AppCard title="预览" no-pad>
        <template #actions>
          <AppButton variant="ghost" class="!min-h-[32px] !px-2.5">↗</AppButton>
        </template>
        <div class="markdown-preview p-4 text-sm leading-relaxed">
          <h2 class="text-lg font-bold mb-4">通用规则</h2>
          <template v-for="s in previewSections" :key="s.title">
            <h3 class="font-bold mt-4 mb-2 text-[#08163d]">{{ s.title }}</h3>
            <ul class="list-disc pl-5 text-[#53617f]">
              <li v-for="item in s.items" :key="item" class="mb-1">{{ item }}</li>
            </ul>
          </template>
        </div>
      </AppCard>
    </section>

    <div class="savebar glass fixed bottom-0 left-[216px] right-0 flex items-center justify-between px-8 py-4 z-40">
      <span class="text-xs text-muted">⬡ 这些规则将包含在你的项目上下文中，并应用于所有 AI 交互。</span>
      <span class="text-xs text-muted">最后更新&nbsp;&nbsp;&nbsp; 2分钟前</span>
      <AppButton variant="primary" @click="toast.show('规则已保存')">保存规则</AppButton>
    </div>
  </div>
</template>
