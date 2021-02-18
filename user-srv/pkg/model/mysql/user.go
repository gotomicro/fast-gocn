package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"strings"
	"user-srv/pkg/invoker"
)

type User struct {
	Uid           int    `json:"uid" form:"uid" gorm:"primary_key"` // 主键ID
	Nickname      string `gorm:"not null;UNIQUE_INDEX"json:"nickname"`
	Address       string `gorm:"not null;"json:"address"` // 地址
	Avatar        string `gorm:"not null;"json:"avatar"`
	Password      string `gorm:"not null;"json:"-"`       // 密码
	Point         int    `gorm:"not null;"json:"point"`   // 积分
	Balance       int    `gorm:"not null;"json:"balance"` // 账户余额
	Level         int    `gorm:"not null;"json:"level"`   // 等级
	Intro         string `gorm:"not null;"json:"intro"`   // 个性签名
	Email         string `gorm:"not null;"json:"email"`
	Phone         string `gorm:"not null;"json:"phone"`
	LastLoginTime int64  `gorm:"not null"json:"lastLoginTime"`
	LastLoginIP   string `gorm:"not null;"json:"lastLoginIp"` // 最后一次登陆IP
	State         int    `gorm:"not null"json:"state"`
	Identify      int    `gorm:"not null"json:"identify"` // 0 未认证，1已认证
	Ctime         int64  `gorm:""json:"ctime"`            // 创建时间
	Utime         int64  `gorm:""json:"utime"`            // 更新时间
	Dtime         int64  `gorm:"index"json:"dtime"`       // 删除时间

	MaxInvitation int `gorm:"not null"json:"maxInvitation"` // 能够生成的最大邀请码数量
	InvitationId int `gorm:"not null"json:"invitationId"`  // 要哪个邀请码邀请，存放的是邀请码的主键
}

func (User) TableName() string {
	return "user"
}

// Create 添加一个用户.
func UserCreate(tx *gorm.DB, data *User) error {
	if l := strings.Count(data.Nickname, "") - 1; l < 2 || l > 50 {
		return errors.New("用户名只能由英文字母数字组成，且在2-50个字符")
	}
	var one User
	tx.Select("uid,nickname").Where("nickname = ?", data.Nickname).Find(&one)
	if one.Uid > 0 {
		if one.Nickname == data.Nickname {
			return errors.New("昵称已存在，请更换昵称")
		}
	}


	return tx.Create(data).Error
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func UserUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("user").Where(sql, binds...).Updates(ups).Error; err != nil {
		invoker.Logger.Error("user update error", zap.Error(err))
		return
	}
	return
}