import Vue from 'vue'
import App from './App'
import store from './store'
import axios from '@/common/js/axios.js'
import mixin from './common/mixin/mixin'
import {
	msg,
	isLogin,
	debounce,
	throttle,
	prePage,
	date,
} from '@/common/js/util'


Vue.mixin(mixin)
Vue.prototype.$util = {
	msg,
	isLogin,
	debounce,
	throttle,
	prePage,
	date
}
Vue.prototype.$store = store
Vue.prototype.$axios = axios
Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
