import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAppStore = defineStore('app', () => {
  const sidebarCollapsed = ref(false)
  const createModalOpen = ref(false)

  function toggleSidebar() {
    sidebarCollapsed.value = !sidebarCollapsed.value
  }

  function openCreateModal() {
    createModalOpen.value = true
  }

  function closeCreateModal() {
    createModalOpen.value = false
  }

  return { sidebarCollapsed, createModalOpen, toggleSidebar, openCreateModal, closeCreateModal }
})
