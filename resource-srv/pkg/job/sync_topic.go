package job

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gotomicro/ego/task/ejob"
	"io"
	"os"
	"resource-srv/pkg/invoker"
	"resource-srv/pkg/model/mysql"
	"strconv"
	"time"
)

func SyncTopicComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("sync_topic"),
		ejob.WithStartFunc(syncTopicAndTopicCate),
	)
}

func syncTopicAndTopicCate() error {
	err := syncTopicCate()
	if err != nil && err != io.EOF {
		return err
	}
	err = syncTopic()
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func syncTopic() error {
	topicData, err := os.Open("./data/topic.csv")
	if err != nil {
		return err
	}
	defer func() {
		_ = topicData.Close()
	}()

	r := csv.NewReader(topicData)

	for {
		line, e := r.Read()
		if e != nil {
			return e
		}

		t := &mysql.Topic{}
		t.Id = parseInt(line[0])
		t.Title = line[1]
		t.CreatedAt = time.Now()
		t.UpdatedAt = time.Now()
		t.RepliedAt = time.Now()
		t.RepliesCount = parseInt(line[5])
		t.NodeName = line[6]
		t.NodeID = parseInt(line[7])
		t.LastReplyUserID = parseInt(line[8])
		t.LastReplyUserLogin = line[9]
		t.Grade = line[10]
		t.LikesCount = parseInt(line[11])
		t.SuggestedAt = time.Now()
		t.ClosedAt = time.Now()
		t.Deleted = false
		t.GocnUid = parseInt(line[15])
		_ = json.Unmarshal([]byte(line[16]), &t.User)
		t.Excellent = parseInt(line[17])
		t.Body = line[18]
		t.BodyHTML = line[19]
		t.Hits = parseInt(line[20])
		_ = json.Unmarshal([]byte(line[21]), &t.Abilities)
		t.ViewCount = parseInt(line[22])
		t.CollectCount = parseInt(line[23])
		invoker.Db.Create(t)
	}
}

func syncTopicCate() error {
	topicCateData, err := os.Open("./data/topic_cate.csv")
	if err != nil {
		return err
	}
	defer func() {
		_ = topicCateData.Close()
	}()

	r := csv.NewReader(topicCateData)

	for {
		line, e := r.Read()
		if e != nil {
			return e
		}
		tc := &mysql.TopicCate{}
		id, _ := strconv.ParseInt(line[0], 10, 64)
		tc.Id = int(id)
		tc.Title = line[1]
		nodeId, _ := strconv.ParseInt(line[2], 10, 64)
		tc.NodeId = int(nodeId)
		tc.CreatedAt = time.Now()
		tc.UpdatedAt = time.Now()
		cntTopic, _ := strconv.ParseInt(line[6], 10, 64)
		tc.CntTopic = int(cntTopic)

		invoker.Db.Create(tc)
	}
}

func parseInt(data string) int {
	res, _ := strconv.ParseInt(data, 10, 64)
	return int(res)
}