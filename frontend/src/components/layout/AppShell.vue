<script setup lang="ts">
import Sidebar from './Sidebar.vue'
import AppToast from '@/components/ui/AppToast.vue'
</script>

<template>
  <div class="flex min-h-screen">
    <Sidebar />
    <main class="flex-1 min-w-0 overflow-auto">
      <slot />
    </main>
  </div>
  <AppToast />
</template>
