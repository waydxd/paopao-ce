import { createApp } from 'vue'
import router from './router'
import store from './store'
import App from './App.vue'
import VueCookies from 'vue-cookies'

import type { MessageApiInjection } from 'naive-ui/lib/message/src/MessageProvider'

// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'
//
import i18n from '@/i18n';
createApp(App)
    .use(i18n)
    .use(router)
    .use(store)
    .use(VueCookies)
    .mount('#app')

declare global {
    interface Window {
        $message: MessageApiInjection,
        $store: any
    }
}
