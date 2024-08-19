import { createApp } from 'vue'
import router from './router'
import store from './store'
import App from './App.vue'
import VueCookies from 'vue-cookies'
import { createI18n } from 'vue-i18n'
import enUS from './i18n/en-US.json'
import zhCN from './i18n/zh-CN.json'
import zhHK from './i18n/zh-HK.json'
import type { MessageApiInjection } from 'naive-ui/lib/message/src/MessageProvider'

// 通用字体
import 'vfonts/Lato.css'
// 等宽字体
import 'vfonts/FiraCode.css'

// Type-define 'en-US' as the master schema for the resource
type MessageSchema = typeof enUS

const i18n = createI18n<[MessageSchema], 'en-US' | 'zh-CN' | 'zh-HK'>({
    locale: 'en-US',
    messages: {
        'en-US': enUS,
        'zh-CN': zhCN,
        'zh-HK': zhHK
    }
})
createApp(App)
    .use(router)
    .use(store)
    .use(VueCookies)
    .use(i18n)
    .mount('#app')

declare global {
    interface Window {
        $message: MessageApiInjection,
        $store: any
    }
}
