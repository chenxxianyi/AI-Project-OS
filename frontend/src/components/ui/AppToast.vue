<script setup lang="ts">
import { useToast } from '@/composables/useToast'

const { visible, message } = useToast()
</script>

<template>
  <Transition name="toast">
    <div
      v-if="visible"
      class="toast fixed bottom-6 left-1/2 -translate-x-1/2 z-[999] px-5 py-3 rounded-[12px] bg-[#08163d] text-white text-sm font-semibold shadow-strong"
    >
      {{ message }}
    </div>
  </Transition>
</template>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.25s ease;
}
.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translate(-50%, 12px);
}
</style>
