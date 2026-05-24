<script setup lang="ts">
defineProps<{
  variant?: 'default' | 'primary' | 'ghost'
  block?: boolean
}>()
</script>

<template>
  <button
    class="btn inline-flex items-center justify-center gap-2 min-h-[42px] px-5 rounded-[10px] font-bold text-sm transition-all duration-200"
    :class="[
      variant === 'primary'
        ? 'bg-gradient-to-br from-[#6f8dff] to-[#355cff] text-white border border-[#6b82ff] shadow-primary hover:-translate-y-px hover:shadow-[0_18px_36px_rgba(64,100,255,.36)]'
        : variant === 'ghost'
          ? 'bg-[#f8faff] border border-line shadow-btn hover:-translate-y-px hover:shadow-btn-hover'
          : 'bg-white border border-line shadow-btn hover:-translate-y-px hover:shadow-btn-hover',
      block ? 'w-full' : '',
    ]"
  >
    <slot />
  </button>
</template>
