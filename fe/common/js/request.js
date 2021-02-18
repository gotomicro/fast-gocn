import store from '@/store'
import cache from './cache'
import {
	msg
} from './util'
import Vue from 'vue'
import { Apis } from './api.js'

/**
 * @param {String} module  
 * @param {String} operation  
 * @param {Object} data 请求参数
 * @param {Object} ext 附加参数
 * @param {Number} ext.cache 数据缓存时间，秒
 */
export const request = (url, method, data={}, ext={})=>{
	return new Promise((resolve, reject) => {
		if(ext.cache > 0){
			const cacheResult = cache.get(module+ '-' +operation);
			if(cacheResult !== false && cacheResult.status !== 0){
				resolve(cacheResult);
				return;
			}
		}
 		Vue.prototype.$axios({
			method,
			url,
			data
		}).then(res=>{
			if(res){
				const code = res.code;
				//token无效
				if(code === 401){
					msg('登录信息已过期，请重新登录');
					store.commit('logout');
					reject('无效的登录信息');
					return;
				}else if(code === 10001){
					msg('用户已被禁用，请联系客服处理');
					if(operation !== 'login' && operation !== 'loginByWeixin'){
						store.commit('logout');
					}
					setTimeout(()=>{
						uni.switchTab({
							url: '/pages/tabbar/home'
						})
					}, 1200)
					reject('用户被禁用');
					return;
				}
			}
			if(ext.cache > 0){
				cache.put(module+ '-' +operation, res.result, ext.cache);
			}
			resolve(res);
		}).catch((err) => {
			reject(err);
		})
	})
}
