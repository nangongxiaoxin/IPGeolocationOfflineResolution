import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  define:{
    __VUE_OPTIONS_API__: 'true',      // 启用选项API
    __VUE_PROD_DEVTOOLS__: 'true',   // 生产环境禁用devtools
    __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: 'true' // 禁用hydration不匹配细节
  }
})
