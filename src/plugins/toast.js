import Vue3Toasity, { toast } from 'vue3-toastify'
import 'vue3-toastify/dist/index.css'

const options = {
    autoClose: 2000,
    position: toast.POSITION.TOP_CENTER,
    closeButton: false,
    containerClassName: 'app-toast',
}

export default function install(app){
    app.use(Vue3Toasity, options)
    app.config.globalProperties.$toast = toast
}