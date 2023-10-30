import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [{
        path: '/',
        component: () => import('./views/Home.vue')
    }, {
        path: '/login',
        component: () => import('./views/Login.vue')
    }]
})

export default router
