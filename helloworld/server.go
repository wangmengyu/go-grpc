package helloworld

import "context"

//server 结构体 用于实现 helloworld.GreeterServer
type server struct {

}
//实现 helloworld.GreeterServer
func (s *server) SayHello(c context.Context, in *HelloRequest) (*HelloReply, error){
	//返回数据 hello+名字
	return &HelloReply{Message: "hello " + in.Name}, nil
}