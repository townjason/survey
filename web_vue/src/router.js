import Vue from 'vue'
import Router from 'vue-router'
import Index from './views/index.vue'
import Complete from './views/complete.vue'
import Error from './views/error.vue'
import Download from './views/download.vue'
Vue.use(Router);

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/questionnaire/:object/:user?',
            name: 'Index',
            component: Index
        },
        {
            path: '/questionnaire/:object/:name/:phone/:writeType',
            name: 'Index',
            component: Index
        },
        {
            path: '/complete',
            name: 'Complete',
            component: Complete
        },
        {
            path: '/error',
            name: 'Error',
            component: Error
        },

        {
            path: '/download',
            name: 'Download',
            component: Download
        },
    ]
})
