package job

import (
	"fmt"
	"github.com/gotomicro/ego/task/ejob"
	"user-srv/pkg/invoker"
	"user-srv/pkg/model/mysql"
)

func InstallComponent() *ejob.Component {
	return ejob.DefaultContainer().Build(
		ejob.WithName("install"),
		ejob.WithStartFunc(runInstall),
	)
}

func runInstall() error {
	models := []interface{}{
		&mysql.User{},
		&mysql.UserOpen{},
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
