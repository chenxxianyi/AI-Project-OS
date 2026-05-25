<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AppShell from '@/components/layout/AppShell.vue'

const route = useRoute()
const isPublicLayout = computed(() =>
  ['landing', 'login'].includes(String(route.name)) || route.meta.publicLayout || route.path.startsWith('/admin')
)
</script>

<template>
  <AppShell v-if="!isPublicLayout">
    <router-view />
  </AppShell>
  <router-view v-else />
</template>
