package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
	"gocn-wechat-be/pkg/invoker"
	"gocn-wechat-be/pkg/router"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Registry(invoker.EtcdRegistry).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.GetRouter(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
