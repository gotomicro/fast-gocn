### proto统一管理生成
> 方便Go使用GRPC互相调用,需要统一生成管理

### proto定义规范:
```
  syntax = "proto3";(使用proto3)
  package com.github.gotomicro.fast-gocn.user;(包名规范)
  option go_package = "gocn/gen/userpb";// 方便项目间的引用go_package 使用远端绝对路径地址 
```

### 错误码定义规范

[错误码规范](https://www.yuque.com/geekgo/lnyi9c/fwihfg)

所有的错误码定义在`gocn/errcode`之下

### golang 如何在本地开发调试

```
1.git clone https://github.com/gotomicro/fast-gocn/proto

2. 本地生成go文件(在项目下执行脚本生成文件,提交时只提交变更的proto文件,go文件不要提交)

cd 下载到的目录

./generate.sh


3.更新项目下gomod文件(开发时使用replace在本地进行替换开发)

  replace https://github.com/gotomicro/fast-gocn/proto => ../../../proto

```

### proto源文件统一规范目录:(按照如下层级约定)
> 如需要提交和变更proto直接编辑对应服务的proto文件就行,前期快速开发不做merge request合并

``` 
gocn(gocn项目)
  └── user-srv(项目名)
       └── user.proto 
```

### proto生成go文件目录规范(自动生成,请勿改动)
```
gen/go(go文件自动生成到该目录层级)
 └── gocn
     └── user-srv(项目名)
          └── user.pb.go
```

### 版本问题

使用版本
```shell script
go get github.com/golang/protobuf/protoc-gen-go@v1.3.2
```