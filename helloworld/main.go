package helloworld

import (
	"log"
	"net"
	"google.golang.org/"


	g/grpc"


)

func main() {
	//配置服务端监听接口
	lis, err := net.Listen("tcp", ":50051")
	if err!=nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//新启一个grpc服务
	s := grpc.NewServer()

}
