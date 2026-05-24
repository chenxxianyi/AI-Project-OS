<script setup lang="ts">
withDefaults(defineProps<{
  title?: string
  noPad?: boolean
}>(), { title: '', noPad: false })
</script>

<template>
  <div class="panel rounded-[14px] glass">
    <div v-if="title" class="panel-head flex items-center justify-between px-5 py-4 border-b border-white/20">
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
