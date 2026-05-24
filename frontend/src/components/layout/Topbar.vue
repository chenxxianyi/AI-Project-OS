<script setup lang="ts">
defineProps<{
  title?: string
}>()
</script>

<template>
  <header class="topbar h-[78px] rounded-[18px] flex items-center justify-between px-6 max-w-[1210px] mx-auto glass">
    <slot />
  </header>
</template>
