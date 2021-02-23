<template>
	<view class="app">
		<view class="left-bottom-sign"></view>
		<view class="back-btn mix-icon icon-guanbi" @click="navBack"></view>
		<view class="right-top-sign"></view>
		<view class="agreement center">
			<text class="mix-icon icon-xuanzhong" :class="{active: agreement}" @click="checkAgreement"></text>
			<text @click="checkAgreement">请认真阅读并同意</text>
			<text class="tit" @click="navToAgreementDetail(1)">《用户服务协议》</text>
			<text class="tit" @click="navToAgreementDetail(2)">《隐私权政策》</text>
		</view>
		<view class="wrapper">
			<view class="left-top-sign">LOGIN</view>
			<view class="welcome">
				欢迎回来！
			</view>
			<view class="other-wrapper">
				<view class="line center">
					<text class="tit">快捷登录</text>
				</view>
				<view class="list row">
					<!-- #ifdef MP-WEIXIN -->
					<button open-type="getUserInfo" @getuserinfo="getUserDetailInfo">
						<view class="item column center">
							<image class="icon" src="/static/icon/wx-login.png"></image>
						</view>
					</button>
					<!-- #endif -->	
				</view>
			</view>
		</view>
		<mix-loading v-if="isLoading"></mix-loading>
	</view>
</template>

<script>
	import tuiModal from "@/components/tui-modal/tui-modal"
	import tuiIcon from "@/components/tui-icon/tui-icon"
	import {
		checkStr
	} from '@/common/js/util'
	import loginMpWx from './mixin/login-mp-wx.js'
	import {
		Apis
	} from '@/common/js/api.js'
	export default {
		components: {
			tuiModal,
			tuiIcon
		},
		mixins: [loginMpWx],
		data() {
			return {
				canUseAppleLogin: false,
				agreement: true,
				username: '',
				invitationCode: "",
				code: '',
				showBindPhone: false,
				isLoading: false
			}
		},
		onLoad(option) {
			if (option != undefined) {
				this.invitationCode = option.code
			}
		},
		methods: {
			async loginSuccessCallBack(data) {
				this.$store.commit('setToken', data);
				this.showBindPhone = false
				this.$util.msg('登陆成功');
				try {
					uni.setStorageSync('userInfo', data.userInfo)
					uni.setStorageSync('expiredTime', data.expiredAt)
				} catch (e) {}
				this.isLoading = false
				const vipRes = await Apis.getVipInfo();
				this.setVipInfoCallBack({
					vipArr: vipRes.data.VipArr
				});
				setTimeout(() => {
					uni.navigateBack();
				}, 1000)
			},
			async setVipInfoCallBack(data) {
				try {
					uni.setStorageSync('vipInfo', data.vipArr);
				} catch (e) {}
			},
			navBack() {
				uni.navigateBack();
			},
			//同意协议
			checkAgreement() {
				this.agreement = !this.agreement;
			},
			//打开协议
			navToAgreementDetail(type) {
				this.navTo('/pages/public/article?param=' + JSON.stringify({
					module: 'article',
					operation: 'getAgreement',
					data: {
						type
					}
				}))
			}
		}
	}
</script>

<style>
	page {
		background: #fff;
	}
</style>
<style scoped lang='scss'>
	.modal {
		text-align: center;
		overflow: hidden;
	}

	.tui-modal-custom {
		padding: 16rpx;
	}

	.tuiIcon {
		display: block;
	}

	.phone {
		margin-top: 12rpx;
		color: white !important;
		background-color: #ee8c44 !important;
		width: 100%;
	}

	.app {
		padding-top: 15vh;
		position: relative;
		width: 100vw;
		height: 100vh;
		overflow: hidden;
		background: #fff;
	}

	.wrapper {
		position: relative;
		z-index: 90;
		padding-bottom: 40rpx;
	}

	.back-btn {
		position: absolute;
		left: 20rpx;
		top: calc(var(--status-bar-height) + 20rpx);
		z-index: 90;
		padding: 20rpx;
		font-size: 32rpx;
		color: #606266;
	}

	.left-top-sign {
		font-size: 120rpx;
		color: #f8f8f8;
		position: relative;
		left: -12rpx;
	}

	.right-top-sign {
		position: absolute;
		top: 80rpx;
		right: -30rpx;
		z-index: 95;

		&:before,
		&:after {
			display: block;
			content: "";
			width: 400rpx;
			height: 80rpx;
			background: #b4f3e2;
		}

		&:before {
			transform: rotate(50deg);
			border-top-right-radius: 50px;
		}

		&:after {
			position: absolute;
			right: -198rpx;
			top: 0;
			transform: rotate(-50deg);
			border-top-left-radius: 50px;
		}
	}

	.left-bottom-sign {
		position: absolute;
		left: -270rpx;
		bottom: -320rpx;
		border: 100rpx solid #d0d1fd;
		border-radius: 50%;
		padding: 180rpx;
	}

	.welcome {
		position: relative;
		left: 50rpx;
		top: -90rpx;
		font-size: 46rpx;
		color: #555;
		text-shadow: 1px 0px 1px rgba(0, 0, 0, .3);
	}

	.input-content {
		padding: 0 60rpx;
	}

	.input-item {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		justify-content: center;
		padding: 0 30rpx;
		background: #f8f6fc;
		height: 120rpx;
		border-radius: 4px;
		margin-bottom: 50rpx;

		&:last-child {
			margin-bottom: 0;
		}

		.row {
			width: 100%;
		}

		.tit {
			height: 50rpx;
			line-height: 56rpx;
			font-size: 26rpx;
			color: #606266;
		}

		input {
			flex: 1;
			height: 60rpx;
			font-size: 30rpx;
			color: #303133;
			width: 100%;
		}
	}

	/* 其他登录方式 */
	.other-wrapper {
		display: flex;
		flex-direction: column;
		align-items: center;
		padding-top: 20rpx;
		margin-top: 80rpx;

		.line {
			margin-bottom: 40rpx;

			.tit {
				margin: 0 32rpx;
				font-size: 24rpx;
				color: #606266;
			}

			&:before,
			&:after {
				content: '';
				width: 160rpx;
				height: 0;
				border-top: 1px solid #e0e0e0;
			}
		}

		.item {
			font-size: 24rpx;
			color: #606266;
			background-color: #fff;
			border: 0;

			&:after {
				border: 0;
			}
		}

		.icon {
			width: 180rpx;
			height: 180rpx;
			margin: 0 24rpx 16rpx;
		}
	}

	.agreement {
		position: absolute;
		left: 0;
		bottom: 6vh;
		z-index: 1;
		width: 750rpx;
		height: 90rpx;
		font-size: 24rpx;
		color: #999;

		.mix-icon {
			font-size: 36rpx;
			color: #ccc;
			margin-right: 8rpx;
			margin-top: 1px;

			&.active {
				color: #ff7900;
			}
		}

		.tit {
			color: #40a2ff;
		}
	}
</style>
