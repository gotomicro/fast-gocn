import {request} from './request.js'

export const Apis = {
	// 用户模块
	login: async (params = {}, ext = {}) => {
		return request(
			"/api/auth/login",
			"POST",
			params
		)
	},
	// Topic
	getTopicIndex: async (params = {}) => {
		return request(
			"/api/c/topic/index",
			"GET",
			params
		)
	},
	getTopicList: async (params = {}) => {
		return request(
			"/api/topic/list",
			"POST",
			params
		)
	},
	getTopicInfo: async (params = {}) => {
		return request(
			`/api/topic/info?id=${params.id}`,
			"GET",
		)
	}
}
