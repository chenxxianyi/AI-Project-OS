import { defineStore } from 'pinia'
import { ref } from 'vue'
import { projectApi, type Project, type ProjectListResponse } from '@/api/project'

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>([])
  const currentProject = ref<Project | null>(null)
  const total = ref(0)
  const loading = ref(false)

  async function fetchProjects(params?: { project_type?: string; search?: string; page?: number; page_size?: number }) {
    loading.value = true
    try {
      const res = await projectApi.list(params)
      const data = res.data.data
      projects.value = data.projects
      total.value = data.total
    } finally {
      loading.value = false
    }
  }

  async function fetchProject(id: string) {
    loading.value = true
    try {
      const res = await projectApi.get(id)
      currentProject.value = res.data.data
    } finally {
      loading.value = false
    }
  }

  async function createProject(data: Partial<Project> & { name: string; project_type: string }) {
    const res = await projectApi.create(data)
    const p = res.data.data
    projects.value.unshift(p)
    return p
  }

  async function updateProject(id: string, data: Partial<Project>) {
    const res = await projectApi.update(id, data)
    const updated = res.data.data
    if (currentProject.value?.id === id) {
      currentProject.value = updated
    }
    const idx = projects.value.findIndex((p) => p.id === id)
    if (idx !== -1) projects.value[idx] = updated
    return updated
  }

  async function deleteProject(id: string) {
    await projectApi.delete(id)
    projects.value = projects.value.filter((p) => p.id !== id)
    if (currentProject.value?.id === id) currentProject.value = null
  }

  return { projects, currentProject, total, loading, fetchProjects, fetchProject, createProject, updateProject, deleteProject }
})
