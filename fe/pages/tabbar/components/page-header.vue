<template>
	<view class="mix-page-header">
		<view class="content row" :style="{
				paddingRight: headerPaddingRight + 'px',
				paddingTop: statusBarHeight + 'px',
			}">
			<image class="avatar" :src="info.avatar || '/static/icon/default-avatar.png'" @click="showDrawer"></image>
			<view class="title">
				{{ title }}
			</view>
		</view>
		<!-- 占位栏 -->
		<tui-drawer mode="left" :visible="visibleDrawer" @close="showDrawer">
			<view class="d-container">
				<view class="top">
					<view class="user-wrapper">
						<image class="avatar" :src="info.avatar || '/static/icon/default-avatar.png'"></image>
						<view class="cen column" v-if="hasLogin">
							<text class="username f-m">{{ info.nickname}}</text>
						</view>
						<view class="login-box" v-else @click="navTo(`/pages/auth/login?code=${code}`)">
							<text>请登录</text>
						</view>
					</view>
				</view>
				<view class="option-wrap">
					<mix-list-cell icon="icon-shezhi1" iconColor="#ff7900" title="设置" @onClick="navTo('/pages/set/set', {login: true})"></mix-list-cell>
				</view>
			</view>
		</tui-drawer>
	</view>
</template>

<script>
	import {
		mapGetters
	} from 'vuex'
	import tuiDrawer from '@/components/tui-drawer/tui-drawer.vue'

	export default {
		name: 'HomePageHeader',
		data() {
			return {
				visible: false,
			};
		},
		props: {
			//顶部是否占位
			showFillView: {
				type: Boolean,
				default: true
			},
			title: {
				type: String,
			},
			info: {
				type: Object,
			},
			vipInfo: {
				type: Object,
			},
			code: {
				type: String,
				default: ""
			}
		},
		components: {
			tuiDrawer
		},
		computed: {
			statusBarHeight() {
				return this.systemInfo.statusBarHeight
			},
			navigationBarHeight() {
				return this.systemInfo.navigationBarHeight;
			},
			customWidth() {
				return this.systemInfo.custom.width;
			},
			customHeight() {
				return this.systemInfo.custom.height;
			},
			//小程序右侧需要留出气泡空间
			headerPaddingRight() {
				// #ifndef MP
				return 0;
				// #endif
				// #ifdef MP-WEIXIN
				return this.customWidth + 20;
				// #endif
			},
			visibleDrawer() {
				return this.$store.state.showDrawer
			},
			...mapGetters(['hasLogin']),
		},
		methods: {
			showDrawer() {
				this.$store.commit('setShowDrawer')
			}
		},
	}
</script>

<style scoped lang="scss">
	.content {
		left: 0;
		top: 0;
		background-color: #fff;
		padding-top: 12rpx;
		padding-bottom: 16rpx;
		z-index: 100;
		display: flex;
	}

	.btn {
		width: 44px;
		height: 40px;
		padding: 5px 12px;
		position: relative;

		.mix-icon {
			font-size: 20px;
			color: #333;
		}

		.msg {
			position: absolute;
			right: 7px;
			top: 5px;
			width: 12px;
			height: 12px;
			background-color: #ff7900;
			border: 2px solid #fff;
			border-radius: 100px;
			opacity: 0;

			&.show {
				opacity: 1;
			}
		}
	}

	.search-wrap {
		flex: 1;
		font-size: 14px;
		color: #999;
		border-radius: 100px;
		background-color: #f2f2f2;

		.icon-sousuo {
			margin-right: 6px;
			font-size: 18px;
			color: #999;
		}
	}

	.avatar {
		width: 80rpx;
		height: 64rpx;
		border-radius: 100px;
		border: 4rpx solid #fff;
		background-color: #fff;
		margin-left: 26rpx;
		border: 1rpx #eee solid;
	}

	.title {
		width: 100%;
		text-align: center;
		padding-top: 8rpx;
	}

	.d-container {
		padding: 30rpx;
		height: 100%;

		.top {
			width: 560rpx;
			margin-top: 60rpx;
			position: relative;
			padding-bottom: 30rpx;
			background-color: #ff7900;
			border-radius: 20rpx;
			box-shadow: 0 0 16rpx rgba(200, 200, 200, .6);

			.user-wrapper {
				display: flex;
				flex-direction: column;
				flex-direction: row;
				align-items: center;
				position: relative;
				padding-top: 40rpx;
			}

			.login-box {
				font-size: 36rpx;
				color: #fff;
			}

			.avatar {
				flex-shrink: 0;
				width: 130rpx;
				height: 130rpx;
				border-radius: 100px;
				border: 4rpx solid #fff;
				background-color: #fff;
				margin-right: 30rpx;
			}

			.username {
				font-size: 34rpx;
				color: #fff;
			}

			.arc-line {
				position: absolute;
				left: 0;
				bottom: 0;
				z-index: 9;
				width: 100%;
				height: 32rpx;
			}

		}

		.vip-tag {
			display: flex;
			flex-wrap: wrap;
			z-index: 10;
			position: relative;
			padding: 0 30rpx;
		}

		.user-group {
			align-self: flex-start;
			padding: 10rpx 16rpx;
			margin-top: 16rpx;
			font-size: 20rpx;
			color: #fff;
			background-color: rgba(255, 255, 255, .3);
			border-radius: 100rpx;
			display: inline-block;
		}

		.certification:not(:first-child) {
			margin-left: 16rpx;
		}
	}

	.option-wrap {
		width: 560rpx;
		margin: 20rpx auto 0;
		margin-top: 30rpx;
		background: #fff;
		border-radius: 10rpx;
		box-shadow: 0 0 16rpx rgba(200, 200, 200, .6);

		.sec-header {
			padding: 26rpx 14rpx 0 24rpx;
			font-size: 28rpx;
			color: #333;

			.icon-lishijilu {
				font-size: 46rpx;
				color: #50bf8b;
				margin-right: 16rpx;
				line-height: 40rpx;
			}

			.icon-lajitong {
				padding: 4rpx 10rpx;
				font-size: 36rpx;
				color: #999;
			}
		}

		.pro-list {
			flex-wrap: nowrap;
			padding: 20rpx 0 12rpx;

			&:before,
			&:after {
				content: '';
				min-width: 30rpx;
				height: 30rpx;
			}

			&:after {
				min-width: 20rpx;
			}

			image {
				flex-shrink: 0;
				width: 144rpx;
				height: 144rpx;
				margin-right: 16rpx;
				border-radius: 8rpx;
			}
		}

		.mix-list-cell {
			display: flex;
			align-items: center;
			height: 96rpx;
			padding: 0 24rpx;
			position: relative;
			text-align: left;

			&.cell-hover {
				background: #fafafa;
			}

			&.b-b {
				&:after {
					left: 30rpx;
					border-color: #f0f0f0;
				}
			}

			.cell-icon {
				align-self: center;
				width: 60rpx;
				font-size: 38rpx;
				color: #ff7900;
			}

			.cell-more {
				align-self: center;
				font-size: 24rpx;
				color: #999;
				margin-left: 16rpx;
			}

			.cell-tit {
				flex: 1;
				font-size: 30rpx;
				color: #333;
				margin-right: 10rpx;
			}

			.cell-tip {
				font-size: 26rpx;
			}
		}
	}
</style>
