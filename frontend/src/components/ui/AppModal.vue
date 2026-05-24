<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { useProjectStore } from '@/stores/project'
import { useToast } from '@/composables/useToast'
import AppButton from '@/components/ui/AppButton.vue'

const store = useAppStore()
const projectStore = useProjectStore()
const toast = useToast()
const router = useRouter()

const step = ref(1)
const creating = ref(false)

const form = reactive({
  name: '',
  description: '',
  project_type: 'saas',
  frontend_stack: 'Next.js',
  backend_stack: 'Node.js',
  database_stack: 'MySQL',
  ai_stack: 'OpenAI',
  ui_style: 'Modern',
  target_user: '',
  product_goal: '',
})

const projectTypes = [
  { value: 'saas', label: 'SaaS 应用', icon: '⌁' },
  { value: 'ai_chat_app', label: 'AI 聊天', icon: '◌' },
  { value: 'ai_agent_app', label: 'AI 智能体', icon: '⊙' },
  { value: 'ecommerce', label: '电商平台', icon: '⌘' },
  { value: 'admin_dashboard', label: '管理后台', icon: '□' },
  { value: 'blog_cms', label: '博客 CMS', icon: '▤' },
  { value: 'landing_page', label: '落地页', icon: '✧' },
  { value: 'mobile_app', label: '移动应用', icon: '⬡' },
]

const frontendOptions = ['Next.js', 'React', 'Vue.js', 'Svelte', 'Nuxt.js', 'Angular']
const backendOptions = ['Node.js', 'Python', 'Go', 'Deno', 'Rust', 'Java']
const dbOptions = ['MySQL', 'PostgreSQL', 'MongoDB', 'Supabase', 'SQLite', 'Redis']
const aiOptions = ['OpenAI', 'Claude', 'Gemini', 'GLM', 'Ollama', 'None']

type StackField = 'frontend_stack' | 'backend_stack' | 'database_stack' | 'ai_stack'

function selectStack(field: StackField, value: string) {
  form[field] = value
}

async function createProject() {
  if (!form.name.trim()) {
    toast.show('请输入项目名称')
    return
  }
  creating.value = true
  try {
    const p = await projectStore.createProject(form)
    store.closeCreateModal()
    toast.show('项目创建成功')
    router.push(`/project/${p.id}`)
  } catch (e: unknown) {
    toast.show(e instanceof Error ? e.message : '创建失败')
  } finally {
    creating.value = false
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="store.createModalOpen"
        class="modal-backdrop fixed inset-0 z-50 bg-black/30 backdrop-blur-sm flex items-center justify-center"
        @click.self="store.closeCreateModal()"
      >
        <section class="project-modal w-[820px] max-h-[85vh] grid grid-cols-[200px_1fr] bg-white rounded-[22px] shadow-strong overflow-hidden">
          <aside class="modal-side bg-[#f8faff] border-r border-line p-6 flex flex-col">
            <div class="modal-title flex items-center gap-3 mb-8">
              <span class="brand-mark w-[34px] h-[34px] rounded-[11px] bg-gradient-to-br from-[#eef3ff] to-white grid place-items-center text-primary shadow-[inset_0_0_0_1px_#dfe7ff]">✦</span>
              <span class="font-bold text-sm">创建新项目</span>
            </div>
            <div class="steps flex flex-col gap-4 flex-1">
              <div v-for="s in 5" :key="s" class="step flex items-center gap-3 text-sm"
                :class="step >= s ? 'font-bold text-primary' : 'text-muted'">
                <span class="step-num w-6 h-6 rounded-full grid place-items-center text-xs font-bold"
                  :class="step >= s ? 'bg-primary text-white' : 'bg-line text-muted'">{{ s }}</span>
                {{ ['基本信息', '技术栈', 'AI 与数据库', '风格与偏好', '预览'][s - 1] }}
              </div>
            </div>
            <AppButton variant="ghost" block @click="store.closeCreateModal()">取消</AppButton>
          </aside>
          <main class="modal-main p-8 overflow-auto relative">
            <button class="absolute top-4 right-4 w-8 h-8 rounded-full hover:bg-[#f1f5ff] grid place-items-center text-lg text-muted" @click="store.closeCreateModal()">×</button>

            <!-- Step 1: Basic Info -->
            <template v-if="step === 1">
              <h2 class="text-xl font-bold mb-1">基本信息</h2>
              <p class="text-muted text-sm mb-6">输入项目名称和类型</p>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">项目名称</label>
                <input v-model="form.name" class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" placeholder="My Awesome Project" />
              </div>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">项目描述</label>
                <textarea v-model="form.description" rows="3" class="w-full px-4 py-3 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors resize-none" placeholder="描述你的项目..." />
              </div>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">项目类型</label>
                <div class="grid grid-cols-4 gap-3">
                  <button v-for="pt in projectTypes" :key="pt.value" class="stack-card px-3 py-2.5 rounded-[10px] font-bold text-sm flex items-center gap-2 transition-colors"
                    :class="form.project_type === pt.value ? 'border-2 border-primary bg-[#f0f4ff]' : 'border border-line bg-white hover:border-primary/40'"
                    @click="form.project_type = pt.value">
                    <span>{{ pt.icon }}</span>{{ pt.label }}
                  </button>
                </div>
              </div>
            </template>

            <!-- Step 2: Tech Stack -->
            <template v-if="step === 2">
              <h2 class="text-xl font-bold mb-1">技术栈</h2>
              <p class="text-muted text-sm mb-6">选择你偏好的技术栈</p>
              <div class="mb-6">
                <h4 class="font-bold text-sm mb-3">前端</h4>
                <div class="grid grid-cols-4 gap-3">
                  <button v-for="opt in frontendOptions" :key="opt" class="stack-card px-4 py-3 rounded-[10px] font-bold text-sm transition-colors"
                    :class="form.frontend_stack === opt ? 'border-2 border-primary bg-[#f0f4ff]' : 'border border-line bg-white hover:border-primary/40'"
                    @click="selectStack('frontend_stack', opt)">{{ opt }}</button>
                </div>
              </div>
              <div class="mb-6">
                <h4 class="font-bold text-sm mb-3">后端</h4>
                <div class="grid grid-cols-4 gap-3">
                  <button v-for="opt in backendOptions" :key="opt" class="stack-card px-4 py-3 rounded-[10px] font-bold text-sm transition-colors"
                    :class="form.backend_stack === opt ? 'border-2 border-primary bg-[#f0f4ff]' : 'border border-line bg-white hover:border-primary/40'"
                    @click="selectStack('backend_stack', opt)">{{ opt }}</button>
                </div>
              </div>
            </template>

            <!-- Step 3: AI & Database -->
            <template v-if="step === 3">
              <h2 class="text-xl font-bold mb-1">AI 与数据库</h2>
              <p class="text-muted text-sm mb-6">选择 AI 提供商和数据库</p>
              <div class="mb-6">
                <h4 class="font-bold text-sm mb-3">数据库</h4>
                <div class="grid grid-cols-4 gap-3">
                  <button v-for="opt in dbOptions" :key="opt" class="stack-card px-4 py-3 rounded-[10px] font-bold text-sm transition-colors"
                    :class="form.database_stack === opt ? 'border-2 border-primary bg-[#f0f4ff]' : 'border border-line bg-white hover:border-primary/40'"
                    @click="selectStack('database_stack', opt)">{{ opt }}</button>
                </div>
              </div>
              <div class="mb-6">
                <h4 class="font-bold text-sm mb-3">AI 提供商</h4>
                <div class="grid grid-cols-4 gap-3">
                  <button v-for="opt in aiOptions" :key="opt" class="stack-card px-4 py-3 rounded-[10px] font-bold text-sm transition-colors"
                    :class="form.ai_stack === opt ? 'border-2 border-primary bg-[#f0f4ff]' : 'border border-line bg-white hover:border-primary/40'"
                    @click="selectStack('ai_stack', opt)">{{ opt }}</button>
                </div>
              </div>
            </template>

            <!-- Step 4: Style & Preferences -->
            <template v-if="step === 4">
              <h2 class="text-xl font-bold mb-1">风格与偏好</h2>
              <p class="text-muted text-sm mb-6">定义 UI 风格和目标用户</p>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">UI 风格</label>
                <input v-model="form.ui_style" class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" placeholder="Modern, Minimal, Material..." />
              </div>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">目标用户</label>
                <input v-model="form.target_user" class="w-full h-[42px] px-4 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors" placeholder="开发者、设计师、企业用户..." />
              </div>
              <div class="mb-5">
                <label class="block text-sm font-semibold mb-1.5">产品目标</label>
                <textarea v-model="form.product_goal" rows="3" class="w-full px-4 py-3 rounded-[10px] bg-white border border-line shadow-btn text-sm outline-none focus:border-primary transition-colors resize-none" placeholder="描述产品目标..." />
              </div>
            </template>

            <!-- Step 5: Preview -->
            <template v-if="step === 5">
              <h2 class="text-xl font-bold mb-1">预览</h2>
              <p class="text-muted text-sm mb-6">确认项目信息</p>
              <div class="flex flex-col gap-3">
                <div class="flex justify-between text-sm"><span class="text-muted">项目名称</span><b>{{ form.name }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">项目类型</span><b>{{ form.project_type }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">描述</span><b class="max-w-[300px] truncate">{{ form.description || '-' }}</b></div>
                <div class="h-px bg-line my-2" />
                <div class="flex justify-between text-sm"><span class="text-muted">前端</span><b>{{ form.frontend_stack }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">后端</span><b>{{ form.backend_stack }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">数据库</span><b>{{ form.database_stack }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">AI</span><b>{{ form.ai_stack }}</b></div>
                <div class="h-px bg-line my-2" />
                <div class="flex justify-between text-sm"><span class="text-muted">UI 风格</span><b>{{ form.ui_style || '-' }}</b></div>
                <div class="flex justify-between text-sm"><span class="text-muted">目标用户</span><b>{{ form.target_user || '-' }}</b></div>
              </div>
            </template>

            <div class="flex gap-3 justify-end mt-8">
              <AppButton v-if="step > 1" @click="step--">上一步</AppButton>
              <AppButton v-if="step < 5" variant="primary" @click="step++">下一步</AppButton>
              <AppButton v-if="step === 5" variant="primary" :disabled="creating" @click="createProject">
                {{ creating ? '创建中...' : '✦ 创建项目' }}
              </AppButton>
            </div>
          </main>
        </section>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.2s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
