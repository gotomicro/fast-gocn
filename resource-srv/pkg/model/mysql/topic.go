package mysql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gotomicro/fast-gocn/proto/gocn/gen/resourcesrv"
	"github.com/jinzhu/gorm"
)

type Topic struct {
	Id                 int       `gorm:"not null"json:"id"`
	Title              string    `gorm:"not null"json:"title"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	RepliedAt          time.Time `json:"repliedat"`
	RepliesCount       int       `gorm:"not null"json:"repliesCount"`
	ViewCount          int       `gorm:"not null"json:"viewCount"`
	NodeName           string    `gorm:"not null"json:"nodeName"`
	NodeID             int       `gorm:"not null"json:"nodeId"`
	LastReplyUserID    int       `gorm:"not null"json:"lastReplyUserId"`
	LastReplyUserLogin string    `gorm:"not null"json:"lastReplyUserLogin"`
	Grade              string    `gorm:"not null"json:"grade"`
	LikesCount         int       `gorm:"not null"json:"likesCount"`
	CollectCount       int       `gorm:"not null"json:"collectCount"`
	SuggestedAt        time.Time `json:"suggestedAt"`
	ClosedAt           time.Time `json:"closedAt"`
	Deleted            bool      `gorm:"not null"json:"deleted"`
	GocnUid            int       `gorm:"not null;"json:"gocnUid"`
	User               GocnUser  `gorm:"not null;type:json"json:"user"`
	Excellent          int       `gorm:"not null"json:"excellent"`
	Body               string    `gorm:"not null;type:longtext"json:"body"`
	BodyHTML           string    `gorm:"not null;type:longtext"json:"bodyHtml"`
	Hits               int       `gorm:"not null"json:"hits"`
	Abilities          Abilities `gorm:"not null;type:json"json:"abilities"`
	Summary            string    `gorm:"-"json:"summary"`
}

func (Topic) TableName() string {
	return "topic"
}

func (t *Topic) ToPB() *resourcesrv.Topic {
	return &resourcesrv.Topic{
		Id:      int32(t.Id),
		Summary: []byte(t.Summary),
		// CoverUrl: t.
		User:         t.User.ToPB(),
		RepliesCount: int32(t.RepliesCount),
		ViewCount:    int32(t.ViewCount),
		LikedCount:   int32(t.LikesCount),
		CreatedAt:    t.CreatedAt.Unix(),
		Title:        t.Title,
		BodyHtml:     []byte(t.BodyHTML),
		Body:         []byte(t.Body),
		CollectedCount: int32(t.CollectCount),
	}
}

func TopicIncreaseViewCount(db *gorm.DB, id int32) error {
	return db.Model(&Topic{}).
		Where("id = ? ", id).
		Update("view_count", gorm.Expr("view_count + 1")).Error
}

func TopicUpdate(db *gorm.DB, topic *Topic) error {
	return db.Model(topic).UpdateColumns(topic).Error
}

func TopicNewestList(db *gorm.DB) ([]Topic, error) {
	var data []Topic
	resp := db.Model(&Topic{}).Order("id desc").Offset(0).Limit(10).Find(&data)
	return data, resp.Error
}

func TopicListByIds(db *gorm.DB, ids []int) ([]Topic, error) {
	var data []Topic
	resp := db.Model(&Topic{}).Where("id in (?)", ids).Find(&data)
	return data, resp.Error
}

func TopicUpdateLikesCount(db *gorm.DB, id int32, delta int) error {
	expr := gorm.Expr(fmt.Sprintf("likes_count + %d", delta))
	if delta < 0 {
		expr = gorm.Expr(fmt.Sprintf("likes_count %d", delta))
	}
	return db.Model(&Topic{}).
		Where("id = ? ", id).
		Update("likes_count", expr, "utime", time.Now().Unix()).Error
}

func TopicUpdateCollectCount(db *gorm.DB, id int32, delta int) error {
	expr := gorm.Expr(fmt.Sprintf("collect_count + %d", delta))
	if delta < 0 {
		expr = gorm.Expr(fmt.Sprintf("collect_count %d", delta))
	}
	return db.Model(&Topic{}).
		Where("id = ? ", id).
		Update("collect_count", expr, "utime", time.Now().Unix()).Error
}

// Info 根据PRI查询单条记录
func TopicInfo(db *gorm.DB, paramId int32) (resp Topic, err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	err = db.Table("topic").Where(sql, binds...).First(&resp).Error
	return
}

// ListPage 根据分页条件查询list
func TopicListPageByCid(db *gorm.DB, sql string, args interface{}, currentPage int, pageSize int) (total int, respList []Topic) {
	if pageSize == 0 {
		pageSize = 10
	}
	if currentPage == 0 {
		currentPage = 1
	}

	query := db.Model(Topic{}).Joins("LEFT JOIN topic_cate ON topic.node_id=topic_cate.node_id").
		Where(sql, args)
	respList = make([]Topic, 0)
	query.Count(&total)
	query.Order("topic.id desc ").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&respList)
	return
}

// ListPage 根据分页条件查询list
func TopicListPage(db *gorm.DB, conds Conds, currentPage int, pageSize int) (total int, respList []Topic) {
	if pageSize == 0 {
		pageSize = 10
	}
	if currentPage == 0 {
		currentPage = 1
	}
	sql, binds := BuildQuery(conds)

	query := db.Table("topic").Where(sql, binds...)
	respList = make([]Topic, 0)
	query.Count(&total)
	query.Order("id desc ").Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&respList)
	return
}

type GocnUser struct {
	ID        int       `json:"id"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Abilities Abilities `json:"abilities"`
}

func (c GocnUser) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c GocnUser) ToPB() *resourcesrv.GoCNUser {
	return &resourcesrv.GoCNUser{
		Id:       int32(c.ID),
		Username: c.Name,
		Avatar:   c.Avatar,
	}
}

func (c *GocnUser) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type Abilities struct {
	Update      bool `json:"update"`
	Destroy     bool `json:"destroy"`
	Ban         bool `json:"ban"`
	Normal      bool `json:"normal"`
	Excellent   bool `json:"excellent"`
	Unexcellent bool `json:"unexcellent"`
	Close       bool `json:"close"`
	Open        bool `json:"open"`
}

func (c Abilities) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Abilities) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (t *Topic) DealSummary() {
	info := trimHtml(t.BodyHTML)
	if len(info) > 100 {
		t.Summary = info[:90] + "..."
	} else {
		t.Summary = info
	}
}

func trimHtml(src string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	// 去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}
