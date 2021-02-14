package api

import (
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/usersrv"

	"gocn-wechat-be/pkg/invoker"
	"gocn-wechat-be/pkg/router/core"
)

type ReqLogin struct {
	UserCode           string `json:"userCode"`
	UserEncryptedData  string `json:"userEncryptedData"`
	UserIv             string `json:"userIv"`
	PhoneCode          string `json:"phoneCode"`
	PhoneEncryptedData string `json:"phoneEncryptedData"`
	PhoneIv            string `json:"phoneIv"`
}

func Login(c *core.Context) {
	req := &ReqLogin{}
	err := c.Bind(req)
	if err != nil {
		c.JSONE(1, "参数错误", err)
	}

	if req.UserCode == "" || req.UserEncryptedData == "" || req.UserIv == "" {
		c.JSONE(1, "参数不能为空", err)
	}

	reply, err := invoker.UsrSrvGrpc.LoginWechat(c.Context, &usersrv.LoginWechatRequest{
		WxUserCode:           req.UserCode,
		WxUserEncryptedData:  req.UserEncryptedData,
		WxUserIv:             req.UserIv,
		WxPhoneCode:          req.PhoneCode,
		WxPhoneEncryptedData: req.PhoneEncryptedData,
		WxPhoneIv:            req.PhoneIv,
		ClientIp:             c.ClientIP(),
	})
	if err != nil {
		c.JSONE(1, "登录微信失败", err)
		return
	}


	c.JSONOK(&LoginWechatReply{
		Uid:        reply.Uid,
		Nickname:   reply.Nickname,
		Avatar:     reply.Avatar,
		Token:      reply.Token,
		BoundPhone: len(reply.Phone) > 0,
		ExpiredAt:  reply.ExpiredAt,
	})
}

type LoginWechatReply struct {
	Uid        int64  `json:"uid"`
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Token      string `json:"token"`
	BoundPhone bool   `json:"boundPhone"`
	ExpiredAt  int64  `json:"expiredAt"`
}
