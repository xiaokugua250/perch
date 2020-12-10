package pb_stream

import (
	"google.golang.org/grpc"
	"net"
	pbStream "perch/internal/GSCeduler/manager/server_manager/stream"
	"testing"
)

func TestNewStream_ServiceClient(t *testing.T) {
	lis,err := net.Listen("tcp","0.0.0.0:5668")
	if err != nil{
		return
	}
	//创建一个grpc 服务器

	svc := pbStream.StreamServer{}
	s := grpc.NewServer()
	//注册事件
	RegisterStream_ServiceServer(s,&svc)
	//处理链接
	s.Serve(lis)
}
