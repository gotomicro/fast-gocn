package router

import (
	"github.com/gotomicro/ego/server/egrpc"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/usersrv"
)

func ServeGRPC() *egrpc.Component {
	srv := egrpc.Load("server.grpc").Build()
	usersrv.RegisterUserServer(srv.Server, &User{})
	return srv
}
