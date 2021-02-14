package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/resourcesrv"

	"gocn-wechat-be/pkg/invoker"
	"gocn-wechat-be/pkg/router/core"
)

func TopicInfo(c *core.Context) {

	var uid int64 = 0
	if c.IsLogin() {
		uid = c.WechatUid()
	}

	id, err := strconv.Atoi(c.Query("id"))
	if id <= 0 || err != nil {
		invoker.Logger.Errorf("could not parse parameter: %+v", err)
		c.JSONE(1, "参数错误", err)
		return
	}

	goodsInfo, err := invoker.ResourceGrpc.TopicInfo(c, &resourcesrv.TopicInfoReq{
		TopicId: int32(id),
		Uid:     uid,
	})
	if err != nil {
		c.JSONE(1, "系统异常", err)
		return
	}

	c.JSONOK(gin.H{
		"info": newTopic(goodsInfo.Topic),
	})
}

func TopicList(c *core.Context) {

	var uid int64 = 0
	if c.IsLogin() {
		uid = c.WechatUid()
	}

	req := &resourcesrv.ColumnListPageReq{}
	err := c.Bind(req)
	if err != nil {
		c.JSONE(1, "参数错误", err)
		return
	}

	req.Uid = uid

	reply, err := invoker.ResourceGrpc.TopicList(c, req)

	if err != nil {
		c.JSONE(1, "系统异常", err)
		return
	}
	topicPbListToTopicVoList(c, reply.TopicList, int(req.CurrentPage), int(reply.Total))
}

func topicPbListToTopicVoList(c *core.Context,
	topicList []*resourcesrv.Topic,
	curPage int, total int) {
	l := make([]*Topic, 0, len(topicList))
	for _, t := range topicList {
		l = append(l, newTopic(t))
	}
	c.JSONPage(l, curPage, 10, total)
}

func TopicCateList(c *core.Context) {
	//
	// if !checkVip(c) {
	//	c.JSONE(1, "未授权", nil)
	//	return
	// }

	// respCates := make([]RespCate, 0)
	// 选择已级分类作为筛选项，因为分类太少
	cates, err := invoker.ResourceGrpc.TopicCateList(c.Context, &resourcesrv.TopicCateListReq{})

	if err != nil {
		c.JSONE(1, "系统异常", err)
		return
	}

	c.JSONOK(cates.TopicCateList)
}

type Topic struct {
	Id             int32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Summary        string                `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	CoverUrl       string                `protobuf:"bytes,3,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
	User           *resourcesrv.GoCNUser `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	CommentCount   int32                 `protobuf:"varint,5,opt,name=commentCount,proto3" json:"commentCount,omitempty"`
	LikedCount     int32                 `protobuf:"varint,6,opt,name=likedCount,proto3" json:"likedCount,omitempty"`
	ViewCount      int32                 `protobuf:"varint,11,opt,name=viewCount,proto3" json:"viewCount,omitempty"`
	CreatedAtLabel string                `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAtLabel"`
	Title          string                `protobuf:"bytes,8,opt,name=title,proto3" json:"title"`
	BodyHtml       string                `protobuf:"bytes,9,opt,name=bodyHtml,proto3" json:"bodyHtml,omitempty"`
	Body           string                `protobuf:"bytes,10,opt,name=body,proto3" json:"body"`
	Liked          bool                  `protobuf:"varint,12,opt,name=liked,proto3" json:"liked"`
	Collected      bool                  `protobuf:"varint,13,opt,name=collected,proto3" json:"collected"`
	CollectedCount int32                 `protobuf:"varint,14,opt,name=collectedCount,proto3" json:"collectedCount"`
}

func newTopic(t *resourcesrv.Topic) *Topic {
	return &Topic{
		Id:             t.Id,
		Summary:        string(t.Summary),
		CoverUrl:       t.CoverUrl,
		User:           t.User,
		CommentCount:   t.RepliesCount,
		LikedCount:     t.LikedCount,
		ViewCount:      t.ViewCount,
		CreatedAtLabel: formatDate(t.CreatedAt),
		Title:          t.Title,
		Body:           string(t.Body),
		Liked:          t.Liked,
		Collected:      t.Collected,
		CollectedCount: t.CollectedCount,
	}
}
