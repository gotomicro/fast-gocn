package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
	"user-srv/pkg/invoker"
	"user-srv/pkg/job"
	"user-srv/pkg/router"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Registry(invoker.EtcdRegistry).
		Job(job.InstallComponent()).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.ServeGRPC(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
