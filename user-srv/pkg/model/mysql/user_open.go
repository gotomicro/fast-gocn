package mysql

import (
	"errors"

	"github.com/jinzhu/gorm"

	"user-srv/pkg/constx"
	"user-srv/pkg/invoker"
)

type UserOpen struct {
	Id           int    `json:"id" form:"id" gorm:"primary_key"` // 主键ID
	Ctime        int64  `gorm:""json:"ctime"`                    // 创建时间
	Utime        int64  `gorm:""json:"utime"`                    // 更新时间
	Dtime        int64  `gorm:"index"json:"dtime"`               // 删除时间
	Genre        int    `gorm:"not null;"json:"genre"`           // 类型
	Uid          int    `gorm:"not null;"json:"uid"`
	WechatOpenid string `gorm:"not null;index"json:"wechatOpenid"` // openid (如果类型是微信 则代表公众号openid)
	AppOpenid    string `gorm:"not null;"json:"appOpenid"`         // app_openid (如果类型是微信 则代表开放平台openid)
	MiniOpenid   string `gorm:"not null;"json:"miniOpenid"`        // mini_openid (如果类型是微信 则代表小程序openid)
	MiniOpenid2  string `gorm:"not null;"json:"miniOpenid2"`       // delivery mini_openid
	Unionid      string `gorm:"not null;"json:"unionid"`           // unionid
	AccessToken  string `gorm:"not null;"json:"accessToken"`       // access_token
	ExpiresIn    int    `gorm:"not null;"json:"expiresIn"`         // access_token过期时间
	RefreshToken string `gorm:"not null;"json:"refreshToken"`      // access_token过期可用该字段刷新用户access_token
	Scope        string `gorm:"not null;"json:"scope"`             // 应用授权作用域
	Nickname     string `gorm:"not null;"json:"nickname"`          // 用户来源平台的昵称
	Avatar       string `gorm:"not null;"json:"avatar"`            // 头像
	Sex          int    `gorm:"not null;"json:"sex"`               // 性别[1男 2女]
	Country      string `gorm:"not null;"json:"country"`           // 国家
	Province     string `gorm:"not null;"json:"province"`          // 省份
	City         string `gorm:"not null;"json:"city"`              // 城市
	State        int    `gorm:"not null;"json:"state"`             // 是否绑定主帐号[默认0否 1是]
	Telephone    string `gorm:"not null;"json:"telephone"`         // 电话
	Email        string `gorm:"not null;"json:"email"`             // 邮箱
	Name         string `gorm:"not null;UNIQUE_INDEX"json:"name"`  // 用户在所有平台唯一昵称
	Intro        string `gorm:"not null;"json:"intro"`
}

func (*UserOpen) TableName() string {
	return "user_open"
}

// 是否存在该用户
func (u *UserOpen) GetMysqlUserOpen() (resp UserOpen, err error) {
	if u.Genre == 0 {
		err = errors.New("genre cant 0")
		return
	}

	if u.Genre == constx.UserOpenWechatMini {
		if u.MiniOpenid == "" {
			err = errors.New("mini_openid cant empty")
			return
		}
		// 2.获取数据库信息
		err = invoker.Db.Where("mini_openid=?", u.MiniOpenid).First(&resp).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}
	} else if u.Genre == constx.UserOpenWechatWeb {
		if u.AppOpenid == "" {
			err = errors.New("app_openid cant empty")
			return
		}
		// 2.获取数据库信息
		err = invoker.Db.Where("app_openid=?", u.AppOpenid).First(&resp).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return
		}
	} else {
		err = errors.New("not exist type wechat")
		return
	}
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Create 新增一条记
func (g *UserOpen) Create(db *gorm.DB) error {
	return db.Create(g).Error
}
