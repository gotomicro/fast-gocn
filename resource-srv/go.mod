module resource-srv

go 1.15

require (
	github.com/google/uuid v1.1.2 // indirect
	github.com/gotomicro/ego v0.3.7
	github.com/gotomicro/ego-component/eetcd v0.1.2
	github.com/gotomicro/ego-component/egorm v0.1.4
	github.com/gotomicro/fast-gocn/proto v0.0.0-20210216124410-681300dc7fb0
	github.com/jinzhu/gorm v1.9.16
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/prometheus/client_golang v1.9.0 // indirect
	github.com/prometheus/procfs v0.3.0 // indirect
	github.com/spf13/cast v1.3.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	google.golang.org/grpc v1.35.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
