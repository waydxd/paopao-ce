
import { createI18n } from 'vue-i18n'
import enUS from "@/i18n/en-US.json";
import zhCN from "@/i18n/zh-CN.json";
import zhHK from "@/i18n/zh-HK.json";

// Type-define 'en-US' as the master schema for the resource
type MessageSchema = typeof enUS | typeof zhCN | typeof zhHK;

const storedLanguage = localStorage.getItem('language') || 'en-US';

const i18n = createI18n<[MessageSchema], 'en-US' | 'zh-CN' | 'zh-HK'>({
    locale: storedLanguage,
    messages: {
        'en-US': enUS,
        'zh-CN': zhCN,
        'zh-HK': zhHK
    }
})

export default i18n;