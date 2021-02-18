import Vue from 'vue'
import Vuex from 'vuex'
import {request} from '@/common/js/request'

Vue.use(Vuex)

const store = new Vuex.Store({
	state: {
		openExamine: false,//是否开启审核状态
		token: '',
		userInfo: {},
		timerIdent: false,//全局1s定时器，只在全局开启一个，所有需要定时执行的任务监听该值即可，无需额外开启
		showDrawer: false
	},
	getters: {
		hasLogin(state){
			return !!state.token;
		},
		hasExpired(state) {
			if (!!state.token) {
				try {
					let timeStmp = Math.round(new Date() / 1000)
					let expiredTime = uni.getStorageSync('expiredTime') + 3000000 || timeStmp + 1
					console.log(expiredTime, timeStmp)
					if (timeStmp > expiredTime) {
						return true
					} else {
						return false
					}
				} catch (e) {
					console.log(e)
				}
			} else {
				return false
			}
		}
	},
	mutations: {
		//更新state数据
		setStateAttr(state, param){
			if(param instanceof Array){
				for(let item of param){
					state[item.key] = item.val;
				}
			}else{
				state[param.key] = param.val;
			}
		},
		//更新token
		setToken(state, data){
			const {token } = data;
			state.token = token;
			uni.setStorageSync('uniIdToken', token);
		},
		//退出登录
		logout(state){
			state.token = '';
			uni.removeStorageSync('uniIdToken');
			uni.removeStorageSync('userInfo')
			uni.removeStorageSync('vipInfo')
			setTimeout(()=>{
				state.userInfo = {};
			}, 1100)
		},
		setShowDrawer(state) {
			state.showDrawer = !state.showDrawer
		}
	},
	actions: {
		//更新用户信息
		async getUserInfo({state, commit}){
			const res = await request('user', 'get', {}, {
				checkAuthInvalid: false
			});
			console.log(res);
			if(res.status === 1){
				const userInfo = res.data;
				commit('setStateAttr', {
					key: 'userInfo',
					val: userInfo
				})
			}
			uni.setStorageSync('userInfo', userInfo);
		},
	}
}) 


export default store
