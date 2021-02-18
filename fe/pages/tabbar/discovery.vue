<template>
	<view class="">
		<page-header title="GoCN" :info="info" :vipInfo="vipInfo"></page-header>
		<view class="tui-tabs">
			<scroll-view id="tab-bar" scroll-with-animation class="tui-scroll-h" :scroll-x="true" :show-scrollbar="false"
			 :scroll-into-view="scrollInto">
				<view v-for="(tab, index) in tabBars" :key="tab.id" class="tui-tab-item" :id="tab.id" :data-current="index" @click="tabClick">
					<!-- #ifdef APP-PLUS -->
					<text class="tui-tab-item-title" :class="{ 'tui-tab-item-title-active': tabIndex == index }">{{ tab.name }}</text>
					<!-- #endif -->

					<!-- #ifndef APP-PLUS -->
					<view class="tui-tab-item-title" :class="{ 'tui-tab-item-title-active': tabIndex == index }">{{ tab.name }}</view>
					<!-- #endif -->
				</view>
			</scroll-view>
			<view class="tui-line-h"></view>
			<swiper :current="tabIndex" class="tui-swiper-box" :duration="300" @change="tabChange">
				<swiper-item class="tui-swiper-item" v-for="(tab, index1) in newsList" :key="index1">
					<scroll-view class="tui-scroll-v" refresher-enabled :refresher-triggered="tab.refreshing" refresher-background="#fafafa"
					 enable-back-to-top :refresher-threshold="100" scroll-y @scrolltolower="loadMore(index1)" @refresherrefresh="onrefresh">
						<view v-for="(newsitem, index2) in tab.data" :key="index2">
							<t-news-item :entity="newsitem" :lastChild="index2 === tab.data.length - 1" @click="goDetail(newsitem)"></t-news-item>
						</view>
						<view class="tui-loading-more" v-if="tab.isLoading || tab.pageIndex > 3">
							<text class="tui-loadmore-line" v-if="tab.pageIndex > 3"></text>
							<text class="tui-loading-more-text">{{ tab.loadingText }}</text>
						</view>
					</scroll-view>
				</swiper-item>
			</swiper>
			<tui-loading v-if="isLoading" :index="1" text="正在加载..."></tui-loading>
		</view>
	</view>
</template>
<script>
	import tNewsItem from './components/news_item.vue';
	import tuiLoading from "@/components/tui-loading/tui-loading"
	import pageHeader from './components/page-header.vue'
	import {
		Apis
	} from "../../common/js/api";

	// 缓存最多页数
	const MAX_CACHE_PAGEINDEX = 3;
	// 缓存页签数量
	const MAX_CACHE_PAGE = 3;
	export default {
		components: {
			tNewsItem,
			tuiLoading,
			pageHeader
		},
		data() {
			return {
				info: {},
				vipInfo: {},
				idx: 'view',
				newsList: [],
				cacheTab: [],
				tabIndex: 0,
				tabBars: [],
				scrollInto: '',
				pulling: false,
				isLoading: false,
				refreshIcon: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAFEAAABRBAMAAABYoVcFAAAAKlBMVEUAAACXl5eZmZmfn5+YmJiXl5eampqZmZmYmJiYmJiZmZmYmJiZmZmZmZl02B9kAAAADXRSTlMAQKAQ0GAsUN+wz5CA21ow0AAAAM5JREFUSMftzLEJAkEQheFR4WzjGji4wC5E0A7E0OgaEIwF4RqwJEEODKcX1114yQ/uhsLtY6Lh57NZ7Dz1heXd27Kwcb+WlQv3Vy1rWcta1rKW/1Q2R8PWt8FYdhPLi79ZLt33KB/hibJzH1E+PaAqRfqAcuMBVSlyMmy1C6hKka0CoCpBAlUJEmgsQQJVCRKoSpBAU0mSaCpJEk0lSSMaS5JG9FuK/IkeQkmSaEjikSSaRpJoHEmiIvOoyCwqMo+KzKMi8+hoZTtte5vDPrSGI5zJ/Z1kAAAAAElFTkSuQmCC'
			};
		},
		async onLoad() {
			const resp = await Apis.getTopicIndex()
			this.tabBars = resp.data
			this.isLoading = true
			const listResp = await Apis.getTopicList({
				cid: this.tabBars[0].id,
				currentPage: 1
			})
			this.newsList.push({
				data: listResp.data,
				pageIndex: listResp.current
			})
			this.isLoading = false
			for (let i = 0; i < this.tabBars.length; i++) {
				this.newsList.push({
					data: [],
					pageIndex: 0
				})
			}
		},
		onShow() {
			try {
				const value = uni.getStorageSync('userInfo');
				if (value) {
					this.info = value
				} else {
					this.info = {}
				}
				const vipValue = uni.getStorageSync('vipInfo');
				if (vipValue) {
					this.vipInfo = vipValue
				} else {
					this.vipInfo = []
				}
			} catch (e) {
				this.info = {}
			}
		},
		methods: {
			async getList(index, refresh) {
				let activeTab = this.newsList[index];
				let list = [];
				if (refresh) {
					activeTab.loadingText = '正在加载...';
					activeTab.pageIndex = 1;
					const resp = await Apis.getTopicList({
						cid: this.tabBars[index].id,
						currentPage: 1
					})

					activeTab.data = resp.data || [];
				} else {

					if (activeTab.loadingText === "没有更多了") {
						return
					}

					const resp = await Apis.getTopicList({
						cid: this.tabBars[index].id,
						currentPage: activeTab.pageIndex + 1,
					})
					activeTab.data = activeTab.data.concat(resp.data);
					activeTab.isLoading = false;
					//根据实际修改判断条件
					if (resp.data.length < resp.pageSize) {
						activeTab.loadingText = '没有更多了';
						return
					}
					activeTab.pageIndex++;
				}
				this.isLoading = false
			},
			goDetail(e) {
				this.navTo(`/pages/news/news?id=${e.id}`);
			},
			async loadMore(e) {
				let activeTab = this.newsList[this.tabIndex];
				if (!activeTab.isLoading) {
					activeTab.isLoading = true;
					await this.getList(this.tabIndex, false);
				}
			},
			tabClick(e) {
				let index = e.target.dataset.current || e.currentTarget.dataset.current;
				this.switchTab(index);
			},
			tabChange(e) {
				let index = e.target.current || e.detail.current;
				this.switchTab(index);
			},
			switchTab(index) {
				if (this.tabIndex === index) return;
				if (this.newsList[index].data.length === 0) {
					this.isLoading = true
					this.getList(index);
				}
				// 缓存 tabId
				if (this.newsList[this.tabIndex].pageIndex > MAX_CACHE_PAGEINDEX) {
					let isExist = this.cacheTab.indexOf(this.tabIndex);
					if (isExist < 0) {
						this.cacheTab.push(this.tabIndex);
					}
				}

				this.tabIndex = index;
				let scrollIndex = index - 1 < 0 ? 0 : index - 1;
				this.scrollInto = this.tabBars[scrollIndex].id;

				// 释放 tabId
				if (this.cacheTab.length > MAX_CACHE_PAGE) {
					let cacheIndex = this.cacheTab[0];
					this.clearTabData(cacheIndex);
					this.cacheTab.splice(0, 1);
				}
			},
			clearTabData(e) {
				this.newsList[e].data.length = 0;
				this.newsList[e].loadingText = '加载更多...';
			},
			async onrefresh(e) {
				var tab = this.newsList[this.tabIndex];
				// #ifdef APP-PLUS
				if (!tab.refreshFlag) {
					return;
				}
				// #endif

				// #ifndef APP-PLUS
				if (tab.refreshing) {
					return;
				}
				// #endif

				tab.refreshing = true;
				tab.refreshText = '正在刷新...';

				this.pulling = true;
				await this.getList(this.tabIndex, true);

				// #ifndef H5
				uni.showToast({
					title: '刷新成功',
					icon: 'none'
				});
				// #endif

				this.pulling = false;
				tab.refreshing = false;
			},
			onpullingdown(e) {
				var tab = this.newsList[this.tabIndex];
				if (tab.refreshing || this.pulling) {
					return;
				}
				if (Math.abs(e.pullingDistance) > Math.abs(e.viewHeight)) {
					tab.refreshFlag = true;
					tab.refreshText = '释放立即刷新';
				} else {
					tab.refreshFlag = false;
					tab.refreshText = '下拉可以刷新';
				}
			}
		}
	};
</script>

<style scoped>
	.tui-tabs {
		flex: 1;
		flex-direction: column;
		overflow: hidden;
		background-color: #fafafa;
		height: 100vh;
	}

	.tui-scroll-h {
		width: 750rpx;
		height: 80rpx;
		background-color: #ffffff;
		flex-direction: row;
		/* #ifndef APP-PLUS */
		white-space: nowrap;
		/* #endif */
		/* #ifdef H5 */
		position: fixed;
		top: 44px;
		left: 0;
		z-index: 999;
		/* #endif */
	}

	.tui-line-h {
		/* #ifdef APP-PLUS */
		height: 1rpx;
		background-color: #cccccc;
		/* #endif */
		position: relative;
	}

	/* #ifndef APP-PLUS*/
	.tui-line-h::after {
		content: '';
		position: absolute;
		border-bottom: 1rpx solid #cccccc;
		-webkit-transform: scaleY(0.5);
		transform: scaleY(0.5);
		bottom: 0;
		right: 0;
		left: 0;
	}

	/* #endif */

	.tui-tab-item {
		/* #ifndef APP-PLUS */
		display: inline-block;
		/* #endif */
		flex-wrap: nowrap;
		padding-left: 34rpx;
		padding-right: 34rpx;
	}

	.tui-tab-item-title {
		color: #555;
		font-size: 30rpx;
		height: 80rpx;
		line-height: 80rpx;
		flex-wrap: nowrap;
		/* #ifndef APP-PLUS */
		white-space: nowrap;
		/* #endif */
	}

	.tui-tab-item-title-active {
		color: #ff7900;
		font-size: 36rpx;
		font-weight: bold;
		border-bottom: #ff7900 solid 6rpx;
		text-align: center;
	}

	.tui-swiper-box {
		flex: 1 !important;
		height: 100%;
		/* #ifdef H5 */
		margin-top: 80rpx;
		/* #endif */
	}

	.tui-swiper-item {
		flex: 1 !important;
		flex-direction: row;
	}

	.tui-scroll-v {
		flex: 1;
		/* #ifndef MP-ALIPAY */
		flex-direction: column;
		/* #endif */
		width: 750rpx;
		height: 100vh;
	}

	.tui-update-tips {
		position: absolute;
		left: 0;
		top: 41px;
		right: 0;
		padding-top: 5px;
		padding-bottom: 5px;
		background-color: #fddd9b;
		align-items: center;
		justify-content: center;
		text-align: center;
	}

	.tui-update-tips-text {
		font-size: 14px;
		color: #ffffff;
	}

	.tui-refresh {
		width: 750rpx;
		height: 64px;
		justify-content: center;
	}

	.tui-refresh-view {
		flex-direction: row;
		flex-wrap: nowrap;
		align-items: center;
		justify-content: center;
	}

	.tui-refresh-icon {
		width: 20px;
		height: 20px;
		transition-duration: 0.25s;
		transition-property: transform;
		transform: rotate(0deg);
		transform-origin: 10px 10px;
	}

	.tui-refresh-icon-active {
		transform: rotate(180deg);
	}

	.tui-loading-icon {
		width: 20px;
		height: 20px;
		margin-right: 5px;
		color: #999999;
	}

	.tui-loading-text {
		margin-left: 2px;
		font-size: 14px;
		color: #999999;
	}

	.tui-loading-more {
		align-items: center;
		justify-content: center;
		padding-top: 15px;
		padding-bottom: 15px;
		text-align: center;
		position: relative;
	}

	.tui-loadmore-line {
		border-bottom-width: 1rpx;
		border-bottom-style: solid;
		border-bottom-color: #e5e5e5;
		width: 320rpx;
		position: absolute;
		z-index: -1;
	}

	.tui-loading-more-text {
		padding-left: 8rpx;
		padding-right: 8rpx;
		font-size: 28rpx;
		line-height: 28rpx;
		background-color: #fafafa;
		text-align: center;
		color: #999;
	}
</style>
