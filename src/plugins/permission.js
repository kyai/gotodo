import router from '@/router'
import cookies from 'vue-cookies'

const whiteList = ['/login']

router.beforeEach(async(to, from, next) => {
    if (cookies.get('token')) {
        if (to.path === '/login') {
            next('/')
        } else {
            next()
        }
    } else {
        if (whiteList.indexOf(to.path) < 0) {
            next('/login')
        } else {
            next()
        }
    }
})

router.afterEach(() => {
})
