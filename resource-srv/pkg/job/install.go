package job

import (
	"fmt"
	"github.com/gotomicro/ego/task/ejob"
	"resource-srv/pkg/invoker"
	"resource-srv/pkg/model/mysql"
)

func InstallComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("install"),
		ejob.WithStartFunc(runInstall),
	)
}

func runInstall() error {
	models := []interface{}{
		&mysql.Topic{},
		&mysql.TopicCate{},
		&mysql.Abilities{},
	}
	gormdb := invoker.Db.Debug()
	gormdb.SingularTable(true)
	err := gormdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...).Error
	if err != nil {
		return err
	}
	fmt.Println("create table ok")
	return nil
}
