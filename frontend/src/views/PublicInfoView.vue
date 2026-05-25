<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AppButton from '@/components/ui/AppButton.vue'

const route = useRoute()

const title = computed(() => String(route.meta.title || 'AI Project OS'))
const description = computed(() => String(route.meta.description || '让 AI 更好地理解你的项目上下文。'))
</script>

<template>
  <div class="landing-body min-h-screen p-[22px_34px_32px] overflow-x-hidden">
    <header class="topbar h-[78px] rounded-[18px] flex items-center justify-between px-6 max-w-[1210px] mx-auto glass-lg">
      <router-link to="/" class="relative z-10 flex items-center gap-3 font-black text-lg text-[#0b173d]">
        <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
        AI Project OS
      </router-link>
      <nav class="topnav flex items-center gap-[54px] text-[#425071] font-semibold text-[13px]">
        <router-link to="/features">功能↴</router-link>
        <router-link to="/template-gallery">模板</router-link>
        <router-link to="/pricing">定价</router-link>
        <router-link to="/docs">文档</router-link>
        <router-link to="/changelog">更新日志</router-link>
      </nav>
      <div class="flex gap-3.5">
        <AppButton variant="ghost"><router-link to="/login">登录</router-link></AppButton>
        <AppButton variant="primary"><router-link to="/dashboard">开始使用</router-link></AppButton>
      </div>
    </header>

    <main class="max-w-[960px] mx-auto pt-20">
      <section class="min-h-[380px] flex flex-col justify-center">
        <h1 class="text-[64px] leading-none mb-6">{{ title }}</h1>
        <p class="text-xl leading-[1.8] text-[#62708d] max-w-[700px] mb-8">{{ description }}</p>
        <div class="flex gap-4">
          <AppButton variant="primary"><router-link to="/dashboard">进入工作台</router-link></AppButton>
          <AppButton><router-link to="/">返回首页</router-link></AppButton>
        </div>
      </section>
    </main>
  </div>
</template>
