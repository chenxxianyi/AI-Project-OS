<script setup lang="ts">
defineProps<{
  label: string
}>()
</script>

<template>
  <span class="chip glass-sm inline-flex items-center gap-2 !rounded-[9px] px-3 py-[7px] text-xs font-bold text-[#40506f]">
    {{ label }}
  </span>
</template>
