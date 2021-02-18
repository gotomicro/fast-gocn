package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TopicCate struct {
	Id        int        `gorm:"not null"json:"id"`
	Title     string     `gorm:"not null"json:"title"`
	NodeId    int        `gorm:"not null;index"json:"nodeId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:""json:"deletedAt"`
	CntTopic  int        `gorm:"not null;"json:"cntTopic"`
}

func (TopicCate) TableName() string {
	return "topic_cate"
}

// List 查询list，extra[0]为sorts
func TopicCateList(db *gorm.DB, conds Conds) ([]TopicCate, error) {
	sql, binds := BuildQuery(conds)
	var resp []TopicCate
	if err := db.Table("topic_cate").Where(sql, binds...).Order(" cnt_topic desc ").Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
