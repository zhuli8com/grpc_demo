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
