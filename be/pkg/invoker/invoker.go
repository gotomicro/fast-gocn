package invoker

import (
	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/client/egrpc/resolver"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/resourcesrv"
	"github.com/gotomicro/fast-gocn/proto/gocn/gen/usersrv"
)

var (
	Logger               *elog.Component
	EtcdRegistry         *registry.Component
	EtcdClient           *eetcd.Component
	Gin                  *egin.Component
	UsrSrvGrpc           usersrv.UserClient
	ResourceGrpc         resourcesrv.ResourceClient
)

func Init() error {
	Logger = elog.DefaultLogger
	Gin = egin.Load("server.http").Build()

	EtcdClient = eetcd.Load("etcd").Build()
	EtcdRegistry = registry.Load("registry").Build(registry.WithClientEtcd(EtcdClient))
	// 必须注册在grpc前面
	resolver.Register("etcd", EtcdRegistry)
	userConn := egrpc.Load("grpc.usersrv").Build().ClientConn
	UsrSrvGrpc = usersrv.NewUserClient(userConn)
	ResourceGrpc = resourcesrv.NewResourceClient(egrpc.Load("grpc.resourcesrv").Build().ClientConn)
	return nil
}
