<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  tabs: string[]
}>()

const active = ref(0)
</script>

<template>
  <nav class="tabs flex gap-1 border-b border-line mb-5">
    <a
      v-for="(tab, i) in tabs"
      :key="i"
      class="tab px-4 py-2.5 text-sm font-semibold cursor-pointer transition-colors border-b-2"
      :class="active === i ? 'text-primary border-primary' : 'text-muted border-transparent hover:text-[#08163d]'"
      @click="active = i"
    >
      {{ tab }}
    </a>
  </nav>
  <slot :active="active" />
</template>
