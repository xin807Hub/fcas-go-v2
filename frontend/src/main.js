import './style/global-variables.scss'
import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import './style/element_visiable.scss'

import { createApp } from 'vue'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/boycent-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import { initDom } from './utils/positionToCode'

// 引入自定义样式
import "./style/boycent.scss"


initDom()

const app = createApp(App)
app.config.productionTip = false

app
    .use(run)
    .use(store)
    .use(auth)
    .use(router)
    .mount('#app')
export default app
