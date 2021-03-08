import Vue from 'vue'
import axios from 'axios'
import store from '@/store'
let localURL = "http://127.0.0.1:8110"
const service = axios.create({
	withCredentials: true,
	crossDomain: true,
	baseURL: localURL,
	timeout: 6000
})
// request拦截器,在请求之前做一些处理
service.interceptors.request.use(
	config => {
		if (store.state.token) {
		    config.headers["Access-Token"] = store.state.token;
		}
		return config;
	},
	error => {
		return Promise.reject(error);
	}
);

//配置成功后的拦截器
service.interceptors.response.use(res => {
	if (res.data.code == 0 || res.code == 200 || res.data.code == 1 || res.data.code == 401) {
		return res.data
	} else {
		return Promise.reject(res.data.msg);
	}
}, error => {
	return Promise.reject(error)
})

axios.defaults.adapter = function(config, store) {
	return new Promise((resolve, reject) => {
		var settle = require('axios/lib/core/settle');
		var buildURL = require('axios/lib/helpers/buildURL');
		uni.request({
			method: config.method.toUpperCase(),
			url: config.baseURL + buildURL(config.url, config.params, config.paramsSerializer),
			header: config.headers,
			data: config.data,
			dataType: config.dataType,
			responseType: config.responseType,
			sslVerify: config.sslVerify,
			complete: function complete(response) {
				if (response.data.code === 200) {
					response.data.code = 0
				}
				response = {
					data: response.data,
					status: response.statusCode,
					errMsg: response.errMsg,
					header: response.header,
					config: config
				};
				settle(resolve, reject, response);
			}
		})
	})
}
export default service
