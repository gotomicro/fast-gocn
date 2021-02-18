package router

import (
	"github.com/gotomicro/ego/server/egin"

	"gocn-wechat-be/pkg/invoker"
	"gocn-wechat-be/pkg/router/api"
	"gocn-wechat-be/pkg/router/core"
)

func GetRouter() *egin.Component {
	r := invoker.Gin

	r.GET("/api/health", core.Handle(api.Health))
	r.POST("/api/auth/login", core.Handle(api.Login))

	r.Use(core.WechatAccessMaybeLogin())
	// 有cache start

	r.GET("/api/topic/info", core.Handle(api.TopicInfo))

	r.GET("/api/c/topic/index", core.Handle(api.TopicCateList))

	// body {currentPage: 1, cid:2}
	r.POST("/api/topic/list", core.Handle(api.TopicList))

	r.Use(core.WechatAccessMustLogin())

	return r
}
