import Vue from 'vue';
import Router from 'vue-router';
import HomePage from '@/components/HomePage';
import InfoPage from '@/components/InfoPage';

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'HomePage',
            component: HomePage
        },
        {
            path: '/info',
            name: 'InfoPage',
            component: InfoPage
        }
    ]
});
