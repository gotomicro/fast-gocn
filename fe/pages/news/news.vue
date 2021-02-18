<template>
	<view class="container">
		<view class="header">
			<text class="title">{{ topic.title }}</text>
			<view class="user">
				<image class="user-image" :src="topic.user.avatar" mode=""></image>
				<text class="user-name">{{ topic.user.username }}</text>
				<text class="time">{{ topic.createdAtLabel }}</text>
				<span v-if="topic.viewCount > 0" class="icon-span">
					<tui-icon name="eye" size="18"></tui-icon><span style="color: #aaaaaa; font-size: 18px; margin-left: 4px;">{{ topic.viewCount }}</span>
				</span>
				<span v-if="topic.body !== ''" class="icon-span">
					<tui-icon :name="topic.liked ? 'like-fill' : 'like' " size="18"></tui-icon>
					<span style="color: #aaaaaa; font-size: 18px; margin-left: 4px;">{{ topic.likedCount || 0 }}</span>
				</span>
				<span v-if="topic.commentCount > 0" class="icon-span">
					<tui-icon name="community" size="18"></tui-icon><span style="color: #aaaaaa; font-size: 18px; margin-left: 4px;">{{ topic.commentCount }}</span>
				</span>
				<span v-if="topic.collectedCount > 0" class="icon-span">
					<tui-icon :name="topic.collected? 'star-fill':'star'" size="18"></tui-icon><span style="color: #aaaaaa; font-size: 18px; margin-left: 4px;">{{ topic.collectedCount || 0}}</span>
				</span>
			</view>
		</view>
		<wemark :md="topic.body" link highlight type="wemark"></wemark>
	</view>
</template>
<script>
	import {
		Apis
	} from "../../common/js/api";
	import tuiIcon from "@/components/tui-icon/tui-icon.vue"
	export default {
		data() {
			return {
				showTab: true,
				topic: {
					title: "正在加载中...",
					user: {
						avatar: ""
					},
					body: "",
					collected: false,
					liked: false,
				},
				id: "",
				share: {
					title: 'GoCN',
					path: `/pages/news/news`,
					imageUrl: '',
					desc: '',
					content: '',
					showPoster: false
				},
				btnList: [],
				bgColor: "#FFF"
			}
		},
		components: {
			tuiIcon
		},
		async onLoad(options) {
			let {
				id
			} = options
			const resp = await Apis.getTopicInfo({
				id: id
			})
			if (resp.code !== 1) {
				this.topic = resp.data.info
			}
		},
	}
</script>

<style scoped>
	.container {
		padding: 18rpx 32rpx;
	}

	.header {
		margin-bottom: 32rpx;
		border-bottom: #ccc solid 2rpx;
		padding-bottom: 16rpx;
	}

	.title {
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

	.richtext {
		width: 100% !important;
		max-width: 100%;
		word-break: break-all;
		color: #333333;
	}

	.user {
		display: flex;
		flex-direction: row;
		margin-bottom: 6rpx;
		align-items: center;
	}

	.user-image {
		width: 48rpx;
		height: 48rpx;
		border-radius: 100%;
		margin-right: 16rpx;
	}

	.user-name {
		font-size: 28rpx;
		color: #bcbcbc;
	}

	.icon-span {
		margin-left: 16rpx
	}

	.time {
		font-size: 28rpx;
		margin-left: 20rpx;
		color: #bcbcbc;
	}
</style>
