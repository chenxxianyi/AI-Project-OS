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
        sans: ['"SF Pro"', 'Manrope', '"HarmonyOS Sans SC"', '"PingFang SC"', '-apple-system', 'system-ui', 'sans-serif'],
        mono: ['Cascadia Mono', 'Consolas', 'monospace'],
      },
      borderRadius: {
        card: '18px',
        panel: '14px',
        btn: '10px',
        chip: '9px',
        /* iOS 26 semantic radius */
        'ios-card': '12px',
        'ios-button': '12px',
        'ios-alert': '14px',
        'ios-menu': '14px',
        'ios-notification': '20px',
        'ios-pill': '9999px',
        'ios-glass-md': '20px',
        'ios-glass-lg': '24px',
      },
      boxShadow: {
        card: '0 12px 34px rgba(56,77,126,.09)',
        strong: '0 24px 70px rgba(58,86,170,.13)',
        btn: '0 8px 18px rgba(67,91,144,.05)',
        'btn-hover': '0 16px 28px rgba(67,91,144,.1)',
        primary: '0 14px 30px rgba(64,100,255,.28)',
        /* iOS 26 Liquid Glass shadow layers */
        'glass-sm': 'inset 0 0.5px 0.5px rgba(255,255,255,0.55), 0 4px 20px -8px rgba(56,77,126,0.08)',
        'glass-md': 'inset 0 0.5px 0.5px rgba(255,255,255,0.55), inset 0 0 0 0.5px rgba(255,255,255,0.15), 0 8px 40px -12px rgba(56,77,126,0.08), 0 2px 12px -4px rgba(56,77,126,0.08)',
        'glass-lg': 'inset 0 1px 1px rgba(255,255,255,0.55), inset 0 0 0 0.5px rgba(255,255,255,0.2), 0 16px 80px -20px rgba(56,77,126,0.14), 0 4px 20px -4px rgba(56,77,126,0.08)',
      },
      transitionDuration: {
        'ios-tab': '300ms',
        'ios-sheet': '500ms',
        'ios-alert': '200ms',
        'ios-menu': '250ms',
        'ios-indicator': '450ms',
      },
      transitionTimingFunction: {
        'ios-snappy': 'cubic-bezier(0.34, 1.56, 0.64, 1.0)',
        'ios-bouncy': 'cubic-bezier(0.68, -0.55, 0.265, 1.55)',
        'ios-gentle': 'cubic-bezier(0.25, 0.46, 0.45, 0.94)',
        'ios-stiff': 'cubic-bezier(0.25, 0.1, 0.25, 1.0)',
      },
    },
  },
  plugins: [],
}
