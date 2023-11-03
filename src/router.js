import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(),
    routes: [{
        path: '/',
        component: () => import('./views/Home.vue')
    }, {
        path: '/login',
        component: () => import('./views/Login.vue')
    }, {
        path: '/system',
        component: () => import('./views/System.vue'),
        redirect: '/system/profile',
        children: [{
            path: 'profile',
            component: () => import('./views/Profile.vue')
        }, {
            path: 'account',
            component: () => import('./views/Account.vue')
        }, {
            path: 'project',
            component: () => import('./views/Project.vue')
        }, {
            path: 'setting',
            component: () => import('./views/Setting.vue')
        }]
    }]
})

export default router
