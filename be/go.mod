module gocn-wechat-be

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gotomicro/ego v0.3.7
	github.com/gotomicro/ego-component/eetcd v0.1.2
	github.com/gotomicro/fast-gocn/proto v0.0.0-20210214134845-2ff23bab1a5c
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

//replace git.yitum.com/mygomod/proto => /Users/mindeng/go-workspace/src/proto
