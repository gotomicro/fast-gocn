# user-srv
## 调剂etcd
etcdctl --endpoints="172.22.0.1:2379" get "" --prefix


用户微服务




protoc -I . proto/usersrv/user.proto  --go_out=plugins=grpc:proto

https://learnku.com/articles/48872