module user-srv

go 1.15

require (
	github.com/gotomicro/fast-gocn/proto v0.0.0-20210214134845-2ff23bab1a5c
	git.yitum.com/mygomod/yitea-contrib v0.0.0-20201220101237-b57369dd1c7a
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/boombuler/barcode v1.0.0
	github.com/disintegration/imaging v1.6.2
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.2
	github.com/gotomicro/ego v0.3.4
	github.com/gotomicro/ego-component/eetcd v0.1.2
	github.com/gotomicro/ego-component/egorm v0.1.2
	github.com/gotomicro/ego-component/eredis v0.1.2
	github.com/hitailang/poster v0.0.0-20200529150248-b74b68d0b587 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/rs/xid v1.2.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.15.0
	golang.org/x/image v0.0.0-20201208152932-35266b937fa6 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/grpc v1.33.1
	sigs.k8s.io/yaml v1.2.0 // indirect
)

//replace git.yitum.com/mygomod/proto => /Users/mindeng/go-workspace/src/proto

//git.yitum.com/mygomod/yitea-contrib => /Users/mindeng/go-workspace/src/mygomod/yitea-contrib
//replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
