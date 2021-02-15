package core

import (
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/usersrv"

	"gocn-wechat-be/pkg/invoker"
	"net/http"
)

var DefaultWechatUid = "/gocn/wechat/uid"
var DefaultWechatNickname = "/gocn/wechat/nickname"
var DefaultWechatAvatar = "/gocn/wechat/avatar"

// WechatAccess 微信登录校验中间件
func WechatAccessMustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get(DefaultWechatUid)
		if ok {
			c.Next()
			return
		}

		accessToken := c.GetHeader("Access-Token")
		if accessToken == "" {
			j := new(Res)
			j.Code = 401
			j.Msg = "缺少token数据"
			c.JSON(http.StatusOK, j)
			c.Abort()
			return
		}

		reply, err := invoker.UsrSrvGrpc.CheckWechatToken(c, &usersrv.CheckWechatTokenRequest{
			Token: accessToken,
		})

		if err != nil {
			j := new(Res)
			j.Code = 401
			j.Msg = "没有用户信息"
			j.Data = err.Error()
			c.JSON(http.StatusOK, j)
			c.Abort()
			return
		}

		c.Set(DefaultWechatUid, reply.Uid)
		c.Set(DefaultWechatNickname, reply.Nickname)
		c.Set(DefaultWechatAvatar, reply.Avatar)
		c.Next()
	}
}

// WechatAccess 微信登录校验中间件
func WechatAccessMaybeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Access-Token")
		if accessToken == "" {
			c.Next()
			return
		}

		reply, err := invoker.UsrSrvGrpc.CheckWechatToken(c, &usersrv.CheckWechatTokenRequest{
			Token: accessToken,
		})

		if err != nil {
			c.Next()
			return
		}

		c.Set(DefaultWechatUid, reply.Uid)
		c.Set(DefaultWechatNickname, reply.Nickname)
		c.Set(DefaultWechatAvatar, reply.Avatar)
		c.Next()
	}
}

// IsLogin 判断当前是否处于登录态之下
func (c *Context) IsLogin() bool {
	_, ok := c.Get(DefaultWechatUid)
	return ok
}

// WechatUid 获取微信id
func (c *Context) WechatUid() int64 {
	return c.MustGet(DefaultWechatUid).(int64)
}

func (c *Context) WechatNickname() string {
	return c.MustGet(DefaultWechatNickname).(string)
}

func (c *Context) WechatAvatar() string {
	return c.MustGet(DefaultWechatAvatar).(string)
}
