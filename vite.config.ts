import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue'

import path from 'path'

export default defineConfig(({ command, mode, isSsrBuild, isPreview }) => {
    // const rootDir = process.cwd();

    return {
        plugins: [
            [vue({})]
        ],
        resolve: {
            alias: {
                '@': path.resolve(__dirname, './resources/js'),
                '~': path.resolve(__dirname, './'),
            },
        },
        build: {
            rollupOptions: {
                input: {
                    main: path.resolve(__dirname, './resources/js/index.html'), // Menetapkan input ke index.html
                }
            }
        }
    }

})