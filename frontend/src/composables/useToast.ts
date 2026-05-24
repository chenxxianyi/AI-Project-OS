import { ref } from 'vue'

const visible = ref(false)
const message = ref('')
let timer: ReturnType<typeof setTimeout>

export function useToast() {
  function show(msg: string, duration = 1800) {
    message.value = msg
    visible.value = true
    clearTimeout(timer)
    timer = setTimeout(() => { visible.value = false }, duration)
  }

  return { visible, message, show }
}
