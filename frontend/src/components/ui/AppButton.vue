<script setup lang="ts">
defineProps<{
  variant?: 'default' | 'primary' | 'ghost'
  block?: boolean
}>()
</script>

<template>
  <button
    class="btn inline-flex items-center justify-center gap-2 min-h-[42px] px-5 rounded-[12px] font-bold text-sm transition-all duration-ios-tab ease-ios-snappy"
    :class="[
      variant === 'primary'
        ? 'bg-gradient-to-br from-[#6f8dff] to-[#355cff] text-white border border-[#6b82ff] shadow-primary hover:-translate-y-px hover:shadow-[0_18px_36px_rgba(64,100,255,.36)] active:scale-[0.97]'
        : variant === 'ghost'
          ? 'glass-pill !rounded-[12px] text-[#425071] hover:-translate-y-px active:scale-[0.97]'
          : 'glass-pill !rounded-[12px] text-[#08163d] hover:-translate-y-px active:scale-[0.97]',
      block ? 'w-full' : '',
    ]"
  >
    <slot />
  </button>
</template>
