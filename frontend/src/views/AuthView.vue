<script setup lang="ts">
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import AppButton from '@/components/ui/AppButton.vue'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()

const isLogin = ref(true)
const email = ref('')
const username = ref('')
const password = ref('')
const error = ref('')
const loading = ref(false)

async function submit() {
  error.value = ''
  loading.value = true
  try {
    if (isLogin.value) {
      await auth.login(email.value, password.value)
    } else {
      if (!username.value) {
        error.value = '请输入用户名'
        loading.value = false
        return
      }
      await auth.register(email.value, username.value, password.value)
    }
    router.push((route.query.redirect as string) || '/dashboard')
  } catch (e: unknown) {
    error.value = e instanceof Error ? e.message : '操作失败'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center p-6">
    <div class="w-full max-w-[420px]">
      <div class="text-center mb-8">
        <router-link to="/" class="inline-flex items-center gap-3 font-black text-lg text-[#0b173d]">
          <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
          AI Project OS
        </router-link>
        <h1 class="text-2xl font-bold mt-6">{{ isLogin ? '欢迎回来' : '创建账户' }}</h1>
        <p class="text-muted text-sm mt-2">{{ isLogin ? '登录以继续使用' : '注册以开始使用' }}</p>
      </div>

      <form class="p-6 rounded-[18px] glass-panel" @submit.prevent="submit">
        <div v-if="error" class="mb-4 p-3 rounded-[10px] bg-red-50 border border-red-200 text-red-600 text-sm">{{ error }}</div>

        <div class="mb-4">
          <label class="block text-sm font-semibold mb-1.5">邮箱</label>
          <input v-model="email" type="email" required placeholder="you@example.com"
            class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" />
        </div>

        <div v-if="!isLogin" class="mb-4">
          <label class="block text-sm font-semibold mb-1.5">用户名</label>
          <input v-model="username" type="text" required placeholder="your_name"
            class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" />
        </div>

        <div class="mb-5">
          <label class="block text-sm font-semibold mb-1.5">密码</label>
          <input v-model="password" type="password" required minlength="6" placeholder="至少6位"
            class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" />
        </div>

        <AppButton variant="primary" block :disabled="loading" type="submit">
          {{ loading ? '处理中...' : (isLogin ? '登录' : '注册') }}
        </AppButton>

        <p class="text-center text-sm text-muted mt-4">
          {{ isLogin ? '还没有账户？' : '已有账户？' }}
          <a class="text-primary font-semibold cursor-pointer" @click="isLogin = !isLogin; error = ''">
            {{ isLogin ? '注册' : '登录' }}
          </a>
        </p>
      </form>
    </div>
  </div>
</template>
