import axios from 'axios'
import cookies from 'vue-cookies'
import { toast } from 'vue3-toastify'

// create an axios instance
const service = axios.create({
    baseURL: 'api', // url = base url + request url
    timeout: 30000, // request timeout (ms)
})

// request interceptor
service.interceptors.request.use(
    config => {
        config.headers['Access-Token'] = cookies.get('token')
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    response => {
        return response.data
    },
    error => {
        let message = error.response && error.response.data
        toast.error(message || error.message)
        return Promise.reject(error)
    }
)

export default service