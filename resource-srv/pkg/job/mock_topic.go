package job

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gotomicro/ego/task/ejob"

	"resource-srv/pkg/invoker"
)

func MockTopic() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("mock_topic"),
		ejob.WithStartFunc(mockTopic),
	)
}

func mockTopic() error {
	wd, _ := os.Getwd()
	gormdb := invoker.Db.Debug()
	gormdb.SingularTable(true)
	// err := gormdb.Exec(fmt.Sprintf(`source %s`, filepath.Join(wd, "data", "sql", "topic.sql"))).Error
	err := gormdb.Exec(`source "/Users/mindeng/go-workspace/src/gotomicro/fast-gocn/resource-srv/data/sql/topic.sql"`).Error
	if err != nil {
		return err
	}
	err = gormdb.Exec(fmt.Sprintf(`source %s`, filepath.Join(wd, "data", "sql", "topic_cate.sql"))).Error
	if err != nil {
		return err
	}
	fmt.Println("create table ok")
	return nil
}
