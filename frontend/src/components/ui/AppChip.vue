<script setup lang="ts">
defineProps<{
  label: string
}>()
</script>

<template>
  <span class="chip inline-flex items-center gap-2 border border-line bg-[#f8fbff] rounded-[9px] px-3 py-[7px] text-xs font-bold text-[#40506f]">
    {{ label }}
  </span>
</template>
