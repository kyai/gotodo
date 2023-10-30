import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import request from './plugins/request'

import './styles/style.scss'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(vuetify)

app.config.globalProperties.$post = (url, data) => {
    return request({url, data, method: 'POST'})
}

app.mount('#app')
