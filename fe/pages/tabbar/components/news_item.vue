<template>
	<view class="tui-list-view">
		<view class="tui-list-cell" :class="[lastChild ? 'tui-last-child' : '']" @click="bindClick">
			<view class="tui-title-box">
				<text class="tui-cell-title">{{ entity.title }}</text>
			</view>
			<view class="tui-img-container" v-if="entity.imgArr && entity.imgArr.length > 0">
				<view class="tui-cell-img" v-for="(item, index) in entity.imgArr" :key="index">
					<image :src="item" class="tui-img"></image>
				</view>
			</view>
			<view class="tui-sub-title">
				<view class="tui-sub-title-user">
					<image :src="entity.user.avatar" class="tui-sub-title-image"></image>
					<text class="tui-sub-content tui-sub-content-lineheight">{{ entity.user.username }}</text>
					<text class="tui-sub-content" style="margin-left: 10px">{{ entity.createdAtLabel }}</text>
				</view>
				<p class="tui-sub-content" style="margin-bottom: 8rpx;" rows="2">{{ entity.summary }}</p>
				<view class="tui-bottom-detail">
					<span v-if="entity.viewCount > 0" class="tui-bottom-detail">
						<tui-icon name="eye" size="18"></tui-icon>
						<span style="color: #aaaaaa; font-size: 24rpx; line-height: 18px; margin-left: 4px; display: inline-block;">{{ entity.viewCount }}</span>
					</span>
					<span v-if="entity.likedCount > 0" class="tui-bottom-detail" style="margin-left: 16rpx">
						<tui-icon :name="entity.liked ? 'like-fill' : 'like' " size="18"></tui-icon>
						<span style="color: #aaaaaa; font-size: 24rpx; line-height: 18px; margin-left: 4px; display: inline-block;">{{ entity.likedCount }}</span>
					</span>
					<span v-if="entity.commentCount > 0" class="tui-bottom-detail" style="margin-left: 16rpx;">
						<tui-icon name="community" size="18"></tui-icon>
						<span style="color: #aaaaaa; font-size: 24rpx; line-height: 18px; margin-left: 4px; display: inline-block;">{{ entity.commentCount }}</span>
					</span>
					<span v-if="entity.collectedCount > 0" class="tui-bottom-detail" style="margin-left: 16rpx;">
						<tui-icon :name="entity.collected ? 'star-fill': 'star'" size="18"></tui-icon>
						<span style="color: #aaaaaa; font-size: 24rpx; line-height: 18px; margin-left: 4px; display: inline-block;">{{ entity.collectedCount }}</span>
					</span>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		name: 'tNewsItem',
		props: {
			entity: {
				type: Object,
				default: function(e) {
					return {};
				}
			},
			lastChild: {
				type: Boolean,
				default: false
			}
		},
		methods: {
			bindClick() {
				this.$emit('click');
			},
		}
	};
</script>

<style>
	.tui-list-view {
		margin-top: 16rpx;
		width: 96%;
		margin-left: 2%;
		padding-left: 16rpx;
		padding-right: 16rpx;
		background-color: #fff;
		box-shadow: 4rpx 4rpx 20rpx rgba(0, 0, 0, .06);
		/* padding-bottom: env(safe-area-inset-bottom); */
	}

	.tui-list-view:active {
		background-color: #eeeeee;
	}

	.tui-list-cell {
		width: 100%;
		padding-top: 30rpx;
		padding-bottom: 12rpx;
		/* #ifdef APP-PLUS*/
		border-bottom-width: 1rpx;
		border-bottom-style: solid;
		border-bottom-color: #e6e6e6;
		/* #endif */
		position: relative;
	}

	/* #ifndef APP-PLUS*/
	.tui-list-cell::after {
		content: '';
		position: absolute;
		border-bottom: 1rpx solid #eaeef1;
		-webkit-transform: scaleY(0.5);
		transform: scaleY(0.5);
		bottom: 0;
		right: 0;
		left: 0;
	}

	/* #endif */

	.tui-last-child {
		border-bottom-width: 0;
		margin-bottom: 30rpx;
	}

	.tui-title-box {
		width: 100%;
		flex: 1;
		/* #ifdef APP-PLUS */
		lines: 2;
		/* #endif */

		/* #ifndef APP-PLUS */
		word-break: break-all;
		overflow: hidden;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 2;
		/* #endif */
	}

	.tui-cell-title {
		display: inline-block;
		width: 100%;
		font-size: 36rpx;
		line-height: 56rpx;
		flex: 1;
		/* #ifdef APP-PLUS */
		lines: 2;
		/* #endif */

		/* #ifndef APP-PLUS */
		word-break: break-all;
		overflow: hidden;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		-webkit-line-clamp: 2;
		/* #endif */
		text-overflow: ellipsis;
		color: #333333;
	}

	.tui-img-container {
		width: 690rpx;
		padding-top: 24rpx;
		height: 184rpx;
		flex-direction: row;
		justify-content: space-between;
	}

	.tui-cell-img {
		width: 220rpx;
		overflow: hidden;
		position: relative;
	}

	.tui-img {
		width: 220rpx;
		height: 184rpx;
		border-radius: 4rpx;
	}

	.tui-sub-title {
		padding-top: 12rpx;
		align-items: center;
		flex-direction: row;
	}

	.tui-sub-title-user {
		display: flex;
		flex-direction: row;
		margin-bottom: 6rpx;
		align-items: center;
	}

	.tui-sub-title-image {
		width: 64rpx;
		height: 64rpx;
		border-radius: 100%;
		margin-right: 16rpx;
	}

	.tui-sub-content {
		font-size: 28rpx;
		color: #bcbcbc;
	}

	.tui-badge {
		padding-top: 8rpx;
		padding-bottom: 8rpx;
		padding-left: 10rpx;
		padding-right: 10rpx;
		font-size: 24rpx;
		border-radius: 4rpx;
		margin-right: 20rpx;
	}

	.tui-red {
		background-color: #fcebef;
		color: #8a5966;
	}

	.tui-blue {
		background-color: #ecf6fd;
		color: #4dabeb;
	}

	.tui-orange {
		background-color: #fef5eb;
		color: #faa851;
	}

	.tui-green {
		background-color: #e8f6e8;
		color: #44cf85;
	}

	.tui-bottom-detail {
		display: flex;
		flex-direction: row;
		align-items: center;
	}
</style>
