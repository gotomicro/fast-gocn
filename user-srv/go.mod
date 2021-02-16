module user-srv

go 1.15

require (
	git.yitum.com/mygomod/yitea-contrib v0.0.0-20201220101237-b57369dd1c7a
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gotomicro/ego v0.3.4
	github.com/gotomicro/ego-component/eetcd v0.1.2
	github.com/gotomicro/ego-component/egorm v0.1.2
	github.com/gotomicro/ego-component/eredis v0.1.2
	github.com/gotomicro/fast-gocn/proto v0.0.0-20210214134845-2ff23bab1a5c
	github.com/jinzhu/gorm v1.9.12
	github.com/spf13/cast v1.3.1
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/grpc v1.33.1
	sigs.k8s.io/yaml v1.2.0 // indirect
)

//replace git.yitum.com/mygomod/proto => /Users/mindeng/go-workspace/src/proto

//git.yitum.com/mygomod/yitea-contrib => /Users/mindeng/go-workspace/src/mygomod/yitea-contrib
//replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
