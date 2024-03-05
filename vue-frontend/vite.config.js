import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import fs from 'fs'
// import mkcert from 'vite-plugin-mkcert'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        AutoImport({
            resolvers: [
                ElementPlusResolver(),
            ],
        }),
        Components({
            resolvers: [
                ElementPlusResolver(),
            ],
        }),
        // mkcert(),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    server: {
        proxy: {
            '/api': {
                // target: 'https://10.30.19.40:8080',
                changeOrigin: true,
                // target: 'https://callme.agileserve.org.cn:30941'
                target: 'http://127.0.0.1:8080'
            },
        },
        host: '127.0.0.1',
        port: 8081,
        // https: true,
        // key: fs.readFileSync('/home/jss40/assistant/10.30.19.40-key.pem'),
        // cert: fs.readFileSync('/home/jss40/assistant/10.30.19.40.pem'),
    }
})
