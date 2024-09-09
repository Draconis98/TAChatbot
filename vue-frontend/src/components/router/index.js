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
            path: '/error',
            name: 'Maintenance',
            component: () => import('@/components/Maintenance.vue')
        },
        {
            path: '/callback/github',
            name: 'GithubCallback',
            component: () => import('@/components/GithubCallback.vue')
        },
        {
            path: '/callback/gitlab',
            name: 'GitlabCallback',
            component: () => import('@/components/GitlabCallback.vue')
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
        },
        {
            path: '/questions',
            name: 'MyQuestions',
            component: () => import('@/components/MyQuestions.vue')
        }
    ]
})

export default router
