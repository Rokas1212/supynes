import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vite.dev/config/
export default defineConfig({
  plugins: [svelte()],
  server: {
    allowedHosts: ['karkasai.ciuplinskas.lt'],
    proxy: { '/api': 'http://app:8080' }
  }
})
