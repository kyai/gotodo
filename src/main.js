import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import request from './plugins/request'
import toast from './plugins/toast'
import cookies from 'vue-cookies'

import '@/plugins/permission'

import './styles/style.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(vuetify)
app.use(cookies)

app.config.globalProperties.$post = (url, data) => {
    return request({url, data, method: 'POST'})
}

app.config.globalProperties.$toast = toast

app.mount('#app')
