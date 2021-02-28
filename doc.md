# 安装

## 标准protoc

```shell
brew install protobuf
```

## protoc-gen-go

由于protobuf并没直接支持go语言需要我们手动安装相关插件

```go
go get -u github.com/golang/protobuf/protoc-gen-go #-u代表update，-v代表view
```

# protobuf

新建person.proto

```prot
// 默认是 proto2
syntax = "proto3";
//option go_package = "path;name";
option go_package = "/;pb";

message Teacher{
  int32 age = 1;
  string name = 2;
}

service SayName {
  rpc SayHello(Teacher) returns (Teacher);
}
```

编译

```shell
protoc --go_out=plugins=grpc:. *proto #protoc --go_out=. *proto 不能编译微服务
```

# Server

```go
package main

import (
	"context"
	"proto_demo/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//定义类
type Children struct {

}


func (this *Children) SayHello(ctx context.Context,t *pb.Teacher) (*pb.Teacher, error){
	t.Name += " is Sleeping"
	return t, nil
}

func main() {
	//1. 初始一个 grpc 服务
	grpcServer := grpc.NewServer()
	//2. 注册服务
	pb.RegisterSayNameServer(grpcServer,&Children{})
	//3. 设置监听， 指定 IP、port
	listen, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err:",err)
		return
	}
	defer listen.Close()
	//4. 启动服务。---- serve()
	grpcServer.Serve(listen)
}
```

# Client

```go
package main

import (
	"context"
	"proto_demo/pb"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//1. 连接 grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.dail error:",err)
		return
	}
	defer grpcConn.Close()
	//2. 初始化 grpc 客户端
	grpcClient := pb.NewSayNameClient(grpcConn)

	//创建并初始化Teacher对象
	var teacher pb.Teacher
	teacher.Name = "zhuli"
	teacher.Age = 18

	//3. 调用远程服务。
	t, err := grpcClient.SayHello(context.TODO(), &teacher)
	fmt.Println(t,err)
}
```

