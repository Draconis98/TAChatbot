import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'HomePage',
            component: () => import('@/components/HomePage.vue')
        },
        {
            path: '/callback/github',
            name: 'GithubCallback',
            component: () => import('@/components/GithubCallback.vue')
        },
        {
            path: '/logged',
            name: 'Logged',
            component: () => import('@/components/Logged.vue')
        },
        {
            path: '/new/question',
            name: 'NewQuestion',
            component: () => import('@/components/NewQuestion.vue')
        },
        {
            path: '/card/details',
            name: 'CardDetails',
            component: () => import('@/components/CardDetails.vue')
        }
    ]
})

export default router