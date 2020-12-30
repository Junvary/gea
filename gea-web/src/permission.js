import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import router from './router'
import store from '@/store'
import { getToken } from '@/utils/cookies'

import getPageTitle from '@/utils/page'

router.beforeEach(async (to, from, next) => {
    NProgress.start()
    const token = getToken()
    document.title = getPageTitle(to.meta.title)
    if (token) {
        if (to.path === '/login' || to.path === '/login/' || to.path === '/') {
            next({ path: '/dashboard' })
            NProgress.done()
        } else {
            const userinfo = store.getters['user/userinfo']
            if (!userinfo || !userinfo.uuid) {
                await store.dispatch('user/GetUserinfo')
                const res = await store.dispatch('router/GetMenus')
                router.addRoutes(res)
                next({ ...to, replace: true })
            } else {
                next()
            }
        }
    } else {
        if (to.path === '/login') {
            next()
        } else {
            next(`/login?redirect=${to.fullPath}`)
            NProgress.done()
        }
    }
})

router.afterEach(() => {
    NProgress.done()
})