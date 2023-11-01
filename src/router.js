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
        path: '/profile',
        component: () => import('./views/Profile.vue')
    }, {
        path: '/system',
        component: () => import('./views/System.vue'),
        children: [{
            path: 'account',
            component: () => import('./views/systems/Account.vue')
        }, {
            path: 'project',
            component: () => import('./views/systems/Project.vue')
        }, {
            path: 'setting',
            component: () => import('./views/systems/Setting.vue')
        }]
    }]
})

export default router
