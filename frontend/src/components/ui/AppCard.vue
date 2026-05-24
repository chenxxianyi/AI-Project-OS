<script setup lang="ts">
withDefaults(defineProps<{
  title?: string
  noPad?: boolean
}>(), { title: '', noPad: false })
</script>

<template>
  <div class="panel rounded-[14px] bg-white/78 border border-line shadow-card backdrop-blur-[22px]">
    <div v-if="title" class="panel-head flex items-center justify-between px-5 py-4 border-b border-line">
      <h3 class="panel-title font-bold text-sm flex items-center gap-2">
        <slot name="icon" />
        {{ title }}
      </h3>
      <slot name="actions" />
    </div>
    <div :class="noPad ? '' : 'p-5'">
      <slot />
    </div>
  </div>
</template>
