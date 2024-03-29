# gRPC Hello World

官方：[quickstart](https://grpc.io/docs/languages/go/quickstart/)

## 准备

- Go
- ProtocolBuff编译器 
- Go ProtocolBuff代码生成插件
1. 安装插件
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
2. PATH环境变量中增加`%GOPATH%/bin`路径，`potoc`编译器找到代码生成插件

## GO mod依赖

项目工程导入依赖
```go
require (
	github.com/golang/protobuf v1.5.3
	google.golang.org/grpc v1.49.0
)
```

## protoc生成代码

新增proto文件

```proto
// 指定proto语法版本
syntax = "proto3";

// 指定生成go代码文件包路径
option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";

// 指定包名
package helloworld;

// RPC接口定义
service Greeter {

  rpc SayHello (HelloRequest) returns (HelloReply) {}

  rpc SayHelloAgain (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

`grpc-demp/helloworld`目录下，执行命令生成`pb.go`文件
```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```
