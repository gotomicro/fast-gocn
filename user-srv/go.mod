module user-srv

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/gotomicro/ego v0.3.8
	github.com/gotomicro/ego-component/eetcd v0.1.4
	github.com/gotomicro/ego-component/egorm v0.1.4
	github.com/gotomicro/ego-component/eredis v0.1.4
	github.com/gotomicro/ego-component/etoken v0.0.0-20210217110704-b9ee6dbf7ff2
	github.com/gotomicro/ego-component/ewechat v0.0.0-20210217110704-b9ee6dbf7ff2
	github.com/gotomicro/fast-gocn/proto v0.0.0-20210216124410-681300dc7fb0
	github.com/jinzhu/gorm v1.9.12
	github.com/spf13/cast v1.3.1
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/grpc v1.33.1
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
