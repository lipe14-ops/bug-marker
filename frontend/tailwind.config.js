/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class', // ativa dark mode via classe
  theme: {
    extend: {
      colors: {
        'background-dark': '#121212',
        'background-light': '#ffffff',
        'text-dark': '#ffffff',
        'text-light': '#000000',
        'subtle-text-dark': '#94a3b8',
        'subtle-text-light': '#64748b',
        'border-dark': '#2a2a2a',
        'border-light': '#e2e8f0',
        'primary': '#06f906',
        'accent': '#06f906',
      },
      fontFamily: {
        'sans': ['Inter', 'system-ui', 'sans-serif'],
      },
    },
  },
  plugins: [
    require('tailwind-scrollbar'),
  ],
  variants: {
    scrollbar: ['rounded', 'dark'],
  },
}
