package client_manager

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"perch/internal/GSCeduler/manager/proto/pb_normal"
	"testing"
)

func TestServerClientRegistry_Registry(t *testing.T) {

	svc :=&Server{}
	lis,err := net.Listen("tcp", "0.0.0.0:5669")
	if err!= nil{
		log.Fatalln(err)
	}
	s:=grpc.NewServer()

	pb_normal.RegisterRegistry_ServiceServer(s,svc)
	if err = s.Serve(lis);err!= nil{
		log.Fatalln(err)
	}
	log.Print("grpc server listening at ",s.GetServiceInfo())


}

func TestServer_Registry(t *testing.T) {
	//1.连接
	conn, err := grpc.Dial("127.0.0.1:5669", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%sn",err)
	}
	defer conn.Close()
	//2.实例化gRPC客户端
	client_grpc := pb_normal.NewRegistry_ServiceClient(conn)
	//3.组装请求参数
	req := new(pb_normal.RegistryRequest)
	req.Name = "zs"
	req.Ip="127.0.0.1"
	req.Uuid="12w1sws"
	//4.调用接口
	response, err := client_grpc.Registry(context.Background(),req)
	if err != nil {
		fmt.Printf("连接异常：%sn",err)
	}
	fmt.Printf("响应结果：%vn", response)
}
