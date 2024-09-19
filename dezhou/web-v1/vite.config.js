import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        proxy: {
            '/local': {
                target: 'http://localhost:9000',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/local/, '')
            }
        }
    }
})
