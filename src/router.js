import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [{
        path: '/',
        component: () => import('./views/Home.vue')
    }, {
        path: '/login',
        component: () => import('./views/Login.vue')
    }]
})

export default router
