package router

import (
	"context"
	"fmt"
	"time"

	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/errcodepb"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/usersrv"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"

	"user-srv/pkg/constx"
	"user-srv/pkg/invoker"
	"user-srv/pkg/model/mysql"
)

type User struct {
	usersrv.UnimplementedUserServer
}

func (*User) CheckWechatToken(ctx context.Context, req *usersrv.CheckWechatTokenRequest) (*usersrv.CheckWechatTokenReply, error) {
	flag := invoker.Token.CheckAccessToken(req.Token)
	if !flag {
		return nil, invoker.Error(errcodepb.ErrCode_Err, fmt.Errorf("check access token got false"))
	}

	userInfo, err := invoker.Token.DecodeAccessToken(req.Token)
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}

	uid, flag := userInfo["sub"].(float64)
	if !flag {
		return nil, invoker.Error(errcodepb.ErrCode_Err, fmt.Errorf("check acceess tokeen sub field not found"))
	}

	user := mysql.User{}
	err = invoker.Db.Select("nickname,avatar").Where("uid = ?", uid).Find(&user).Error
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}

	return &usersrv.CheckWechatTokenReply{
		Uid:      int64(uid),
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil

}

// 微信登录
func (u *User) LoginWechat(ctx context.Context, in *usersrv.LoginWechatRequest) (*usersrv.LoginWechatReply, error) {
	_, userInfo, err := invoker.Wechat.GetMiniProgram().Login(in.WxUserCode, in.WxUserEncryptedData, in.WxUserIv)
	// 1 登录微信
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_LoginWechatServer, err)
	}

	phoneNumber, err := parsePhoneFromWechat(in.WxUserCode, in.WxPhoneEncryptedData, in.WxPhoneIv)
	if err != nil {
		invoker.Logger.Warn("phone parse error", elog.FieldErr(err))
	}

	invoker.Logger.Debug("userinfo", elog.Any("user", userInfo))
	// 写user_open表
	createUserOpen := &mysql.UserOpen{
		Ctime:      time.Now().Unix(),
		Utime:      time.Now().Unix(),
		Genre:      constx.UserOpenWechatMini,
		MiniOpenid: userInfo.OpenID,
		Unionid:    userInfo.UnionID,
		Nickname:   userInfo.NickName,
		Avatar:     userInfo.AvatarURL,
		Sex:        userInfo.Gender,
		Country:    userInfo.Country,
		Province:   userInfo.Province,
		City:       userInfo.City,
		Telephone:  phoneNumber,
		Name:       userInfo.NickName,
	}

	mysqlUserOpen, err := createUserOpen.GetMysqlUserOpen()
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_LoginWechatMysqlUserOpen, err)
	}

	uniqueName := userInfo.NickName + "_" + userInfo.OpenID

	tx := invoker.Db.Begin()
	defer tx.RollbackUnlessCommitted()

	// 如果不存在，那么就创建这个用户user open，并且提示需要绑定
	if mysqlUserOpen.Id == 0 {
		err = createUserOpen.Create(tx)

		if err != nil {
			invoker.Logger.Warnf("create userOpen create error, maybe nickname conflicts, cause: %s", err.Error())
			// 如果昵称冲突了，我们用加了open id的唯一名字
			createUserOpen.Name = uniqueName
			// 再次尝试
			err = createUserOpen.Create(tx)
			if err != nil {
				return nil, invoker.Error(errcodepb.ErrCode_LoginWechatCreateUserOpen, err)
			}
		}

		u := &mysql.User{
			Nickname: createUserOpen.Nickname,
			Avatar:   createUserOpen.Avatar,
			State:    1,
			Phone:    createUserOpen.Telephone,
			Ctime:    time.Now().Unix(),
			Utime:    time.Now().Unix(),
		}
		err = mysql.UserCreate(tx, u)
		if err != nil {
			invoker.Logger.Warnf("create mdMembers create error, maybe nickname conflicts, err: %s", err.Error())
			// 可能是名字冲突，我们再次尝试
			u.Nickname = uniqueName
			err = mysql.UserCreate(tx, u)
			if err != nil {
				return nil, invoker.Error(errcodepb.ErrCode_LoginWechatCreateUser, err)
			}
		}

		err = tx.Model(mysql.UserOpen{}).Where("id = ?", createUserOpen.Id).Updates(gin.H{
			"uid": u.Uid,
		}).Error
		if err != nil {
			return nil, invoker.Error(errcodepb.ErrCode_LoginWechatCreateUserOpen, err)
		}
		tx.Commit()

		return loginReturnAccessToken(u.Uid, userInfo.NickName, userInfo.AvatarURL, in.ClientIp, phoneNumber)
	}

	// 如果mysqlUseropen id存在，检测下uid是否存在，理论上不应该出现这种情况
	if mysqlUserOpen.Uid == 0 {
		return nil, invoker.Error(errcodepb.ErrCode_LoginWechatUserOpenUidStatus, err)
	}
	return loginReturnAccessToken(mysqlUserOpen.Uid, userInfo.NickName, userInfo.AvatarURL, in.ClientIp, mysqlUserOpen.Telephone)
}

func parsePhoneFromWechat(wxPhoneCode string, wxPhoneEncryptedData string, wxPhoneIv string) (string, error) {
	var phoneNumber string
	wXBizDataCrypt, err := invoker.Wechat.GetMiniProgram().Code2Session(wxPhoneCode)
	if err == nil {
		phone, err := invoker.Wechat.GetMiniProgram().DecryptPhone(wXBizDataCrypt.SessionKey, wxPhoneEncryptedData, wxPhoneIv)
		if err != nil {
			return "", err
		}
		phoneNumber = phone.PhoneNumber
	}
	return phoneNumber, nil
}

func loginReturnAccessToken(uid int, nickname, avatar string, ip string, phone string) (*usersrv.LoginWechatReply, error) {
	// 如果绑定了，那么就创建access token
	startTime := time.Now().Unix()

	respToken, err := invoker.Token.CreateAccessToken(uid, startTime)
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}

	// 根据用户信息，更新登录信息
	err = mysql.UserUpdateX(invoker.Db, mysql.Conds{"uid": uid}, mysql.Ups{
		"last_login_ip":   ip,
		"last_login_time": time.Now().Unix(),
	})
	if err != nil {
		return nil, invoker.Error(errcodepb.ErrCode_Err, err)
	}
	return &usersrv.LoginWechatReply{
		Nickname:  nickname,
		Avatar:    avatar,
		Token:     respToken.AccessToken,
		Uid:       int64(uid),
		Phone:     phone,
		ExpiredAt: startTime + econf.GetInt64("token.accessTokenExpireInterval") - 600, // 减去10min让他快速过期
	}, nil
}
