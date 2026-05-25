<script setup lang="ts">
import { useToast } from '@/composables/useToast'

const { visible, message } = useToast()
</script>

<template>
  <Transition name="toast">
    <div
      v-if="visible"
      class="toast glass-panel !rounded-[20px] fixed bottom-6 left-1/2 -translate-x-1/2 z-[999] px-5 py-3 text-[#08163d] text-sm font-semibold !shadow-strong"
    >
      {{ message }}
    </div>
  </Transition>
</template>

<style scoped>
.toast-enter-active {
  transition: all 0.4s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}
.toast-leave-active {
  transition: all 0.2s ease-in;
}
.toast-enter-from {
  opacity: 0;
  transform: translate(-50%, 16px) scale(0.95);
}
.toast-leave-to {
  opacity: 0;
  transform: translate(-50%, 8px);
}
</style>
