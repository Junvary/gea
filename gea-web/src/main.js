import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

import 'normalize.css/normalize.css' // resize
import '@/css/layout.scss' // 全局样式

import '@/css/pages.scss'

import Element from 'element-ui'
import '@/css/element-ui.scss' // 修改element默认样式
Element.Dialog.props.closeOnClickModal.default = false
Vue.use(Element)

// 引入封装的router
import router from '@/router/index'
import '@/permission'
import store from '@/store/index'

// echarts的5.0版本引入方式改了，这里使用 4.9.0 版本
import echarts from "echarts";
Vue.prototype.$echarts = echarts;

// 音乐播放
import APlayer from '@moefe/vue-aplayer';
Vue.use(APlayer, {
    defaultCover: 'https://github.com/u3u.png', // set the default cover
    productionTip: false, // disable console output
});

new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')
