import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import * as fs from "fs";

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
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    server: {
        proxy: {
            '/api': {
                target: 'https://10.30.19.40:8080',
                changeOrigin: true,
            },
        },
        host: '10.30.19.40',
        port: 8081,
        https: true,
        key: fs.readFile('/home/jss40/assistant/agileserve.org.cn.key'),
        cert: fs.readFile('/home/jss40/assistant/agileserve.org.cn.pem'),
    }
})
