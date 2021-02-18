import {
	Apis
} from '@/common/js/api.js'
export default {
	// #ifdef MP-WEIXIN
	data() {
		return {
			mpCodeTimer: 0,
			oneMpWxCode: '',
		}
	},
	computed: {
		timerIdent() {
			return this.$store.state.timerIdent;
		}
	},
	watch: {
		timerIdent() {
			this.mpCodeTimer++;
			if (this.mpCodeTimer % 30 === 0) {
				this.getMpWxCode();
			}
		}
	},
	onShow() {
		this.getMpWxCode(0);
	},
	methods: {
		getUserDetailInfo(userinfo) {
			if (!this.agreement) {
				this.$util.msg('请阅读并同意用户服务及隐私协议');
				return;
			}
			this.$util.throttle(async () => {
				let that = this
				wx.getUserInfo({
				  success: async(resp) => {	
					const res = await Apis.login({
						userCode: this.oneMpWxCode,
						userEncryptedData: resp.encryptedData,
						userIv: resp.iv,
						invitationCode: that.invitationCode
					}, {
						showLoading: true
					});
					if (res.msg == "登录微信失败") {
						that.$util.msg('登录微信失败');
					} else {
						that.loginSuccessCallBack({
							token: res.data.token,
							tokenExpired: false,
							userInfo: {
								avatar: res.data.avatar,
								nickname: res.data.nickname,
								uid: res.data.uid,
							},
							boundPhone: res.data.boundPhone,
							expiredAt: res.data.expiredAt
						});
					}
				  },
				  fail: (err) => {
					  console.log(err)
				  }
				})
			})
		},
		//获取code
		getMpWxCode(time) {
			uni.login({
				provider: 'weixin',
				success: res => {
					this.oneMpWxCode = res.code;
				}
			})
		}
	}
	// #endif
}
