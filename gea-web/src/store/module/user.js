import { login, getUserinfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router'
import { setToken, removeToken } from '@/utils/cookies'

const state = {
    userinfo: {
        uuid: "",
        nickName: "",
        headerImg: "",
        authority: "",
    },
}
const mutations = {
    SET_USERINFO(state, userinfo) {
        state.userinfo = userinfo
    },
    SET_TOKEN(state, token) {
        setToken(token)
    },
    LOGOUT(state) {
        state.userinfo = {}
        removeToken()
        router.push({ name: 'login', replace: true })
        sessionStorage.clear()
        window.location.reload()
    },
}
const actions = {
    async Login({ commit }, loginForm) {
        const res = await login(loginForm)
        if (res.code === 0) {
            commit('SET_TOKEN', res.data.token)
            // 在执行下面跳转时，会被router的 beforeEach 捕获，判断 是否存在userinfo
            const redirect = router.history.current.query.redirect || '/'
            router.push({ path: redirect })
            return true
        }
    },
    SetToken({ commit }, token) {
        commit('SET_TOKEN', token)
    },
    async GetUserinfo({ commit }) {
        const res = await getUserinfo()
        if (res.code === 0) {
            commit('SET_USERINFO', res.data.userinfo)
        }
    },
    async Logout({ commit }) {
        const res = await jsonInBlacklist()
        if (res.code == 0) {
            commit("LOGOUT")
        } else {
            console.log("Logout中jsonInBlacklist回复错误：", res)
            commit("LOGOUT")
        }
    }
}
const getters = {
    userinfo: state => state.userinfo,
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}
