export default {
	data() {
		return {
			page: 0, //页码
			pageNum: 6, //每页加载数据量
			loadingType: 1, //0加载前 1加载中 2没有更多
			isLoading: false, //刷新数据
			loaded: false, //加载完毕，
		}
	},
	methods: {
		log(data) {},
		/**
		 * navigatorTo跳转页面
		 * @param {String} url
		 * @param {Object} options
		 * @param {Boolean} options.login 是否检测登录  
		 */
		navTo(url, options = {}) {
			this.$util.throttle(() => {
				if (!url) {
					return;
				}
				if ((~url.indexOf('login=1') || options.login) && !this.$store.getters.hasLogin) {
					url = '/pages/auth/login';
				}
				uni.navigateTo({
					url
				})
			}, 300)
		}
	},
}
