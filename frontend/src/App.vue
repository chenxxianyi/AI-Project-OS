<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AppShell from '@/components/layout/AppShell.vue'

const route = useRoute()
const isLanding = computed(() => route.name === 'landing')
</script>

<template>
  <AppShell v-if="!isLanding">
    <router-view />
  </AppShell>
  <router-view v-else />
</template>
