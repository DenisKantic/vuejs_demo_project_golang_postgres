/** @type {import('tailwindcss').Config} */
import daisyui from 'daisyui/src'
export default {
  content: ['./index.html', './src/**/*.{vue,js,jsx,ts,tsx}'],
  theme: {
    extend: {}
  },
  plugins: [daisyui]
}
