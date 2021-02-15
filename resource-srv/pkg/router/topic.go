package router

import (
	"context"

	"github.com/gotomicro/fast-gocn/proto/gocn/gen/errcodepb"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/resourcesrv"

	"resource-srv/pkg/invoker"
	"resource-srv/pkg/model/mysql"
)

var topicCateNodeIds = [8]int{18, 3, 1, 25, 2, 15, 46, 12}

func (r *Resource) TopicList(ctx context.Context,
	req *resourcesrv.ColumnListPageReq) (*resourcesrv.TopicListReply, error) {

	var (
		total int
		resp  []mysql.Topic
	)
	conds := mysql.Conds{}
	if req.Cid != 0 {
		total, resp = mysql.TopicListPageByCid(invoker.Db, "topic.node_id = ?", req.Cid, int(req.CurrentPage), 10)
	} else {
		// 没有cid，不能用left join方式，否则会出现多重记录
		total, resp = mysql.TopicListPage(invoker.Db, conds, 10, int(req.CurrentPage))
	}

	return r.buildTopicPBs(resp, total)
}

func (r *Resource) buildTopicPBs(data []mysql.Topic, total int) (*resourcesrv.TopicListReply, error){
	// 进一步判断，用户有没有点赞
	return buildTopicPBFromDB(total, data), nil
}


func (r *Resource) TopicCateList(context.Context, *resourcesrv.TopicCateListReq) (*resourcesrv.TopicCateListReply, error) {
	respCates := make([]*resourcesrv.TopicCate, 0)
	// 选择已级分类作为筛选项，因为分类太少
	cates, err := mysql.TopicCateList(invoker.Db, mysql.Conds{})
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}

	// 有点懒，直接循环，犯不着用IN查询，反正都是一把捞出来
	for _, value := range topicCateNodeIds {
		for i := 0; i < len(cates); i++ {
			c := cates[i]
			if value == c.NodeId {
				respCates = append(respCates, &resourcesrv.TopicCate{
					Id:   int32(c.NodeId),
					Name: c.Title,
				})
			}
		}

	}
	return &resourcesrv.TopicCateListReply{
		TopicCateList: respCates,
	}, nil
}

func (r *Resource) TopicInfo(c context.Context, req *resourcesrv.TopicInfoReq) (*resourcesrv.TopicInfoReply, error) {

	topic, err := mysql.TopicInfo(invoker.Db, req.TopicId)

	// 更新失败直接忽略，没关系
	err = mysql.TopicIncreaseViewCount(invoker.Db, req.TopicId)
	if err != nil {
		invoker.Logger.Errorf("update view count got error. %+v", err)
	}

	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}

	return &resourcesrv.TopicInfoReply{
		Topic: topic.ToPB(),
	}, nil
}

func buildTopicPBFromDB(total int, topics []mysql.Topic) *resourcesrv.TopicListReply {
	ts := make([]*resourcesrv.Topic, 0, len(topics))
	for _, t := range topics {
		ts = append(ts, t.ToPB())
	}
	return &resourcesrv.TopicListReply{
		TopicList: ts,
		Total:     int32(total),
	}
}