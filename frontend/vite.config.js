import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true,       // Esto expone la app a la red (0.0.0.0)
    port: 5173,       // Fija el puerto
    watch: {
      usePolling: true // IMPORTANTE: Docker en Linux a veces no detecta cambios sin esto
    }
  }
})