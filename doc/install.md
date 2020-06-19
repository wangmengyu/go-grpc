### 安装
1. 在本地安装生成工具:
   go get github.com/golang/protobuf/protoc-gen-go@v1.3 

2. 撰写 helloworld.proto 定义服务结构

3. 执行命令 生成客户端和服务端接口 
   protoc -I helloworld helloworld/helloworld.proto --go_out=plugins=grpc:helloworld



    