import router from '@/router'
import cookies from 'vue-cookies'
import { useUserStore } from '@/stores/user'

const whiteList = ['/login']

router.beforeEach(async(to, from, next) => {
    if (cookies.get('token')) {
        if (to.path === '/login') {
            next('/')
        } else {
            const userStore = useUserStore()
            if(userStore.name){
                next()
            }else{
                try{
                    await userStore.profile()
                    next()
                }catch(e){
                    await userStore.logout()
                    next('/login')
                }
            }
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
