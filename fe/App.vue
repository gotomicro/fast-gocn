<script>
	import Vue from 'vue'
	export default {
		onLaunch() {
			uni.getSystemInfo({
				success: e => {
					this.initSize(e);
				}
			})
			this.initLogin();
		},
		methods: {
			//登录状态
			async initLogin() {
				const token = uni.getStorageSync('uniIdToken');
				if (token) {
					this.$store.commit('setToken', {
						token
					});
				}
			},
			initSize(e) {
				const systemInfo = e;
				let navigationBarHeight;
				let custom = {};
				// #ifndef MP
				custom = {
					height: 36,
					width: 88
				};
				navigationBarHeight = 44;
				// #endif
				// #ifdef MP-WEIXIN
				custom = wx.getMenuButtonBoundingClientRect();
				navigationBarHeight = custom.bottom + custom.top - e.statusBarHeight * 2;
				// #endif	
				systemInfo.custom = custom;
				systemInfo.navigationBarHeight = navigationBarHeight;
				Vue.prototype.systemInfo = systemInfo;
			},
			closeTimer() {
				if (__timerId != 0) {
					clearInterval(__timerId);
					__timerId = 0;
				}
			},
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
</style>
