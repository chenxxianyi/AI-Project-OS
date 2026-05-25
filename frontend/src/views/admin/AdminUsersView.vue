<script setup lang="ts">
import { onMounted, ref, watch, computed } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppChip from '@/components/ui/AppChip.vue'

const store = useAdminStore()
const toast = useToast()

const searchQuery = ref('')
const roleFilter = ref('')
const editingUser = ref<string | null>(null)
const editRole = ref('')
const showDeleteConfirm = ref<string | null>(null)

const roleOptions = ['admin', 'user']
const statusOptions: Record<string, { label: string; color: string }> = {
  active: { label: '活跃', color: 'green' },
  suspended: { label: '停用', color: 'red' },
  pending: { label: '待激活', color: '[#e5a000]' },
}

onMounted(() => {
  store.fetchUsers({ page: 1, page_size: 20 })
})

watch([searchQuery, roleFilter], () => {
  store.fetchUsers({ page: 1, page_size: 20, search: searchQuery.value, role: roleFilter.value })
})

function handlePageChange(page: number) {
  store.fetchUsers({ page, page_size: 20, search: searchQuery.value, role: roleFilter.value })
}

function startEdit(userId: string, currentRole: string) {
  editingUser.value = userId
  editRole.value = currentRole
}

async function saveRole(userId: string) {
  await store.updateUser(userId, { role: editRole.value })
  editingUser.value = null
  toast.show('用户角色已更新')
}

async function toggleStatus(userId: string, currentStatus: string) {
  const newStatus = currentStatus === 'active' ? 'suspended' : 'active'
  await store.updateUser(userId, { status: newStatus })
  toast.show(`用户已${newStatus === 'active' ? '激活' : '停用'}`)
}

async function confirmDelete(userId: string) {
  await store.deleteUser(userId)
  showDeleteConfirm.value = null
  toast.show('用户已删除')
}

function timeAgo(dateStr: string) {
  const diff = Date.now() - new Date(dateStr).getTime()
  const mins = Math.floor(diff / 60000)
  if (mins < 60) return `${mins}分钟前`
  const hours = Math.floor(mins / 60)
  if (hours < 24) return `${hours}小时前`
  const days = Math.floor(hours / 24)
  return `${days}天前`
}

const totalPages = computed(() => Math.ceil(store.usersTotal / 20))
</script>

<template>
  <div class="p-[32px_36px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">用户管理</h1>
      <p class="text-muted mt-1">管理系统用户、角色和状态。</p>
    </section>

    <!-- Filters -->
    <div class="flex items-center gap-3 mb-5">
      <div class="search flex items-center gap-2 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn w-[300px]">
        <span class="text-muted">⌕</span>
        <input v-model="searchQuery" class="flex-1 bg-transparent outline-none text-sm" placeholder="搜索用户名或邮箱..." />
      </div>
      <div class="flex items-center gap-2">
        <button
          v-for="role in ['', 'admin', 'user']"
          :key="role"
          class="px-3 py-2 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
          :class="roleFilter === role ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
          @click="roleFilter = role"
        >
          {{ role || '全部' }}
        </button>
      </div>
      <span class="ml-auto text-sm text-muted">共 {{ store.usersTotal }} 位用户</span>
    </div>

    <!-- Users Table -->
    <div class="rounded-[14px] glass-panel overflow-hidden">
      <table class="w-full relative z-10">
        <thead>
          <tr class="border-b border-line text-left">
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">用户</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">角色</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">状态</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">项目</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">生成</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider">注册时间</th>
            <th class="px-5 py-3.5 text-xs font-bold text-muted uppercase tracking-wider text-right">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in store.users" :key="user.id" class="border-b border-line last:border-0 hover:bg-[#f8fbff] transition-colors">
            <td class="px-5 py-3.5">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 rounded-full bg-gradient-to-br from-[#222d58] to-[#f3a071] grid place-items-center text-white text-xs font-bold shrink-0">{{ (user.username || '?')[0].toUpperCase() }}</div>
                <div class="min-w-0">
                  <b class="block text-sm truncate">{{ user.username }}</b>
                  <span class="text-xs text-muted truncate block">{{ user.email }}</span>
                </div>
              </div>
            </td>
            <td class="px-5 py-3.5">
              <div v-if="editingUser === user.id" class="flex items-center gap-2">
                <select v-model="editRole" class="h-8 px-2 rounded-[8px] bg-white border border-line text-sm outline-none focus:border-primary">
                  <option v-for="r in roleOptions" :key="r" :value="r">{{ r }}</option>
                </select>
                <button class="text-primary text-xs font-bold cursor-pointer" @click="saveRole(user.id)">保存</button>
                <button class="text-muted text-xs cursor-pointer" @click="editingUser = null">取消</button>
              </div>
              <AppChip v-else :label="user.role" @click="startEdit(user.id, user.role)" class="cursor-pointer" />
            </td>
            <td class="px-5 py-3.5">
              <span
                class="inline-flex items-center gap-1.5 text-xs font-semibold cursor-pointer"
                :class="`text-${statusOptions[user.status]?.color || 'muted'}`"
                @click="toggleStatus(user.id, user.status)"
              >
                <span class="w-2 h-2 rounded-full" :class="`bg-${statusOptions[user.status]?.color || 'muted'}`"></span>
                {{ statusOptions[user.status]?.label || user.status }}
              </span>
            </td>
            <td class="px-5 py-3.5 text-sm">{{ user.project_count }}</td>
            <td class="px-5 py-3.5 text-sm">{{ user.generation_count }}</td>
            <td class="px-5 py-3.5 text-sm text-muted">{{ timeAgo(user.created_at) }}</td>
            <td class="px-5 py-3.5 text-right">
              <div v-if="showDeleteConfirm === user.id" class="flex items-center justify-end gap-2">
                <span class="text-xs text-red">确认删除?</span>
                <button class="text-red text-xs font-bold cursor-pointer" @click="confirmDelete(user.id)">确认</button>
                <button class="text-muted text-xs cursor-pointer" @click="showDeleteConfirm = null">取消</button>
              </div>
              <button v-else class="text-muted hover:text-red text-sm cursor-pointer transition-colors" @click="showDeleteConfirm = user.id">🗑</button>
            </td>
          </tr>
        </tbody>
      </table>

      <!-- Empty State -->
      <div v-if="store.users.length === 0 && !store.loading" class="py-12 text-center text-muted text-sm">
        暂无用户数据
      </div>

      <!-- Loading -->
      <div v-if="store.loading" class="py-8 text-center text-muted text-sm">
        加载中...
      </div>
    </div>

    <!-- Pagination -->
    <div v-if="totalPages > 1" class="flex items-center justify-center gap-2 mt-5">
      <button
        v-for="p in totalPages"
        :key="p"
        class="w-9 h-9 rounded-[10px] text-sm font-semibold cursor-pointer transition-all duration-ios-tab ease-ios-snappy"
        :class="store.usersPage === p ? 'glass-pill !rounded-[10px] text-primary' : 'text-muted hover:bg-white/20'"
        @click="handlePageChange(p)"
      >
        {{ p }}
      </button>
    </div>
  </div>
</template>
