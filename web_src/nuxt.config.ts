// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from '@tailwindcss/vite'

export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@element-plus/nuxt'],
  css: ['~/assets/styles/main.css'],

  nitro: {
    devProxy: {
      '/api': {
        target: 'http://192.168.1.166:3000/api',
        changeOrigin: true,
      },
    },
  },

  vite: {
    plugins: [tailwindcss()],
    css: {
      preprocessorOptions: {
        scss: { api: 'modern-compiler' },
      },
    },
  },
})
