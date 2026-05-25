<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'

const route = useRoute()
const router = useRouter()

const title = computed(() => String(route.meta.title || '功能建设中'))
const description = computed(() => String(route.meta.description || '这个模块正在规划中，后续会逐步开放。'))
</script>

<template>
  <div class="p-[32px_36px] max-w-[860px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">{{ title }}</h1>
      <p class="text-muted mt-1">{{ description }}</p>
    </section>

    <AppCard title="即将上线">
      <div class="relative z-10 py-10 text-center">
        <div class="mx-auto mb-4 w-14 h-14 rounded-[18px] bg-[#f0f4ff] grid place-items-center text-primary text-2xl">✦</div>
        <h2 class="text-lg font-bold mb-2">{{ title }}正在开发中</h2>
        <p class="text-muted mb-6">你可以先返回首页创建项目，或进入已有项目使用 Prompt、规则和项目大脑功能。</p>
        <AppButton variant="primary" @click="router.push({ name: 'dashboard' })">返回首页</AppButton>
      </div>
    </AppCard>
  </div>
</template>
