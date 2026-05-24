<script setup lang="ts">
import { useAppStore } from '@/stores/app'

const store = useAppStore()
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="store.createModalOpen"
        class="modal-backdrop fixed inset-0 z-50 bg-black/30 backdrop-blur-sm flex items-center justify-center"
        @click.self="store.closeCreateModal()"
      >
        <section class="project-modal w-[820px] max-h-[85vh] grid grid-cols-[200px_1fr] bg-white rounded-[22px] shadow-strong overflow-hidden">
          <aside class="modal-side bg-[#f8faff] border-r border-line p-6 flex flex-col">
            <div class="modal-title flex items-center gap-3 mb-8">
              <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
              <span class="font-bold text-sm">创建新项目</span>
            </div>
            <div class="steps flex flex-col gap-4 flex-1">
              <div class="step flex items-center gap-3 text-sm"><span class="step-num w-6 h-6 rounded-full bg-primary text-white grid place-items-center text-xs font-bold">1</span>基本信息</div>
              <div class="step flex items-center gap-3 text-sm font-bold text-primary"><span class="step-num w-6 h-6 rounded-full bg-primary text-white grid place-items-center text-xs font-bold">2</span>技术栈</div>
              <div class="step flex items-center gap-3 text-sm text-muted"><span class="step-num w-6 h-6 rounded-full bg-line text-muted grid place-items-center text-xs font-bold">3</span>AI 与数据库</div>
              <div class="step flex items-center gap-3 text-sm text-muted"><span class="step-num w-6 h-6 rounded-full bg-line text-muted grid place-items-center text-xs font-bold">4</span>风格与偏好</div>
              <div class="step flex items-center gap-3 text-sm text-muted"><span class="step-num w-6 h-6 rounded-full bg-line text-muted grid place-items-center text-xs font-bold">5</span>预览</div>
            </div>
            <AppButton variant="ghost" block>返回</AppButton>
          </aside>
          <main class="modal-main p-8 overflow-auto relative">
            <button class="absolute top-4 right-4 w-8 h-8 rounded-full hover:bg-[#f1f5ff] grid place-items-center text-lg text-muted" @click="store.closeCreateModal()">×</button>
            <h2 class="text-xl font-bold mb-1">技术栈</h2>
            <p class="text-muted text-sm mb-6">选择你偏好的技术栈</p>
            <div class="mb-6">
              <h4 class="font-bold text-sm mb-3">选择前端技术栈</h4>
              <div class="grid grid-cols-4 gap-3">
                <button class="stack-card px-4 py-3 rounded-[10px] border-2 border-primary bg-[#f0f4ff] font-bold text-sm">● Next.js</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">⚛ React</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">▼ Vue.js</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">S Svelte</button>
              </div>
            </div>
            <div class="mb-6">
              <h4 class="font-bold text-sm mb-3">选择后端技术栈</h4>
              <div class="grid grid-cols-4 gap-3">
                <button class="stack-card px-4 py-3 rounded-[10px] border-2 border-primary bg-[#f0f4ff] font-bold text-sm">⬢ Node.js</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">🐍 Python</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">GO Go</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">◒ Deno</button>
              </div>
            </div>
            <div class="mb-6">
              <h4 class="font-bold text-sm mb-3">选择数据库</h4>
              <div class="grid grid-cols-4 gap-3">
                <button class="stack-card px-4 py-3 rounded-[10px] border-2 border-primary bg-[#f0f4ff] font-bold text-sm">🐘 PostgreSQL</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">🐬 MySQL</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">🍃 MongoDB</button>
                <button class="stack-card px-4 py-3 rounded-[10px] border border-line bg-white font-bold text-sm hover:border-primary/40">⚡ Supabase</button>
              </div>
            </div>
            <div class="flex gap-3 justify-end mt-8">
              <AppButton>返回</AppButton>
              <AppButton variant="primary">下一步</AppButton>
            </div>
          </main>
        </section>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
