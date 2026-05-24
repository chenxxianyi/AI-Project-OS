/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: { DEFAULT: '#4167ff', 2: '#7e95ff', 3: '#8f68ff' },
        cyan: '#23c7db',
        green: '#27c990',
        red: '#ff6a87',
        line: '#e3eafa',
        line2: '#d6e0f5',
        muted: '#6b7898',
        soft: '#f1f5ff',
        bg: '#f6f9ff',
        bg2: '#eef4ff',
        text: '#08163d',
      },
      fontFamily: {
        sans: ['Manrope', 'HarmonyOS Sans SC', 'PingFang SC', 'Microsoft YaHei', 'sans-serif'],
        mono: ['Cascadia Mono', 'Consolas', 'monospace'],
      },
      borderRadius: {
        card: '18px',
        panel: '14px',
        btn: '10px',
        chip: '9px',
      },
      boxShadow: {
        card: '0 12px 34px rgba(56,77,126,.09)',
        strong: '0 24px 70px rgba(58,86,170,.13)',
        btn: '0 8px 18px rgba(67,91,144,.05)',
        'btn-hover': '0 16px 28px rgba(67,91,144,.1)',
        primary: '0 14px 30px rgba(64,100,255,.28)',
      },
    },
  },
  plugins: [],
}
