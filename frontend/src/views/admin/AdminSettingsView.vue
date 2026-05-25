<script setup lang="ts">
import { onMounted, ref, reactive } from 'vue'
import { useAdminStore } from '@/stores/admin'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'
import AppCard from '@/components/ui/AppCard.vue'

const store = useAdminStore()
const toast = useToast()

const form = reactive({
  allow_registration: true,
  default_user_role: 'user',
  max_projects_per_user: 10,
  ai_rate_limit: 100,
  maintenance_mode: false,
})

const saving = ref(false)

onMounted(async () => {
  await store.fetchSettings()
  Object.assign(form, store.settings)
})

async function handleSave() {
  saving.value = true
  try {
    await store.updateSettings({ ...form })
    toast.show('设置已保存')
  } catch {
    toast.show('保存失败')
  } finally {
    saving.value = false
  }
}

function handleReset() {
  Object.assign(form, store.settings)
}
</script>

<template>
  <div class="p-[32px_36px] max-w-[800px]">
    <section class="headline mb-6">
      <h1 class="text-2xl font-bold">系统设置</h1>
      <p class="text-muted mt-1">管理平台全局配置与安全策略。</p>
    </section>

    <div class="flex flex-col gap-5">
      <!-- Registration -->
      <AppCard title="注册与访问">
        <div class="relative z-10 flex flex-col gap-5">
          <div class="flex items-center justify-between">
            <div>
              <b class="block text-sm">开放注册</b>
              <span class="text-xs text-muted">允许新用户自行注册账户</span>
            </div>
            <button
              class="w-12 h-7 rounded-full transition-all duration-ios-tab ease-ios-snappy relative"
              :class="form.allow_registration ? 'bg-primary' : 'bg-line'"
              @click="form.allow_registration = !form.allow_registration"
            >
              <span
                class="absolute top-0.5 w-6 h-6 rounded-full bg-white shadow-btn transition-all duration-ios-tab ease-ios-snappy"
                :class="form.allow_registration ? 'left-[22px]' : 'left-0.5'"
              ></span>
            </button>
          </div>

          <div class="flex items-center justify-between">
            <div>
              <b class="block text-sm">维护模式</b>
              <span class="text-xs text-muted">启用后普通用户无法访问系统</span>
            </div>
            <button
              class="w-12 h-7 rounded-full transition-all duration-ios-tab ease-ios-snappy relative"
              :class="form.maintenance_mode ? 'bg-red' : 'bg-line'"
              @click="form.maintenance_mode = !form.maintenance_mode"
            >
              <span
                class="absolute top-0.5 w-6 h-6 rounded-full bg-white shadow-btn transition-all duration-ios-tab ease-ios-snappy"
                :class="form.maintenance_mode ? 'left-[22px]' : 'left-0.5'"
              ></span>
            </button>
          </div>

          <div>
            <label class="block text-sm font-semibold mb-1.5">默认用户角色</label>
            <select v-model="form.default_user_role" class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors">
              <option value="user">普通用户</option>
              <option value="admin">管理员</option>
            </select>
          </div>
        </div>
      </AppCard>

      <!-- Limits -->
      <AppCard title="资源限制">
        <div class="relative z-10 flex flex-col gap-5">
          <div>
            <label class="block text-sm font-semibold mb-1.5">每用户最大项目数</label>
            <input v-model.number="form.max_projects_per_user" type="number" min="1" max="1000" class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" />
            <span class="text-xs text-muted mt-1 block">限制每个用户可创建的项目数量</span>
          </div>

          <div>
            <label class="block text-sm font-semibold mb-1.5">AI 请求频率限制</label>
            <div class="flex items-center gap-2">
              <input v-model.number="form.ai_rate_limit" type="number" min="1" max="10000" class="flex-1 h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" />
              <span class="text-sm text-muted whitespace-nowrap">次/小时</span>
            </div>
            <span class="text-xs text-muted mt-1 block">每个用户每小时的 AI 生成请求上限</span>
          </div>
        </div>
      </AppCard>

      <!-- Actions -->
      <div class="flex items-center justify-end gap-3 pt-2">
        <AppButton variant="ghost" @click="handleReset">重置</AppButton>
        <AppButton variant="primary" :disabled="saving" @click="handleSave">
          {{ saving ? '保存中...' : '保存设置' }}
        </AppButton>
      </div>
    </div>
  </div>
</template>
