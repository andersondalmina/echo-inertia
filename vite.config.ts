import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  publicDir: false,
  build: {
    manifest: "mix-manifest.json",
    outDir: "public/build",
    rollupOptions: {
      input: ['frontend/js/app.tsx'],
    },
  },
})
