package client_manager

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"perch/internal/GSCeduler/manager/proto/pb_normal"
	"perch/internal/GSCeduler/manager/proto/utils"
	"testing"
)

func TestServerClientRegistry_Registry(t *testing.T) {

	svc := &server{}
	lis, err := net.Listen("tcp", "0.0.0.0:5669")
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Printf("========%s\n", "sx")
		cli, err := utils.GetClietIP(ctx)
		if err != nil {
			log.Println("Failed to get client address")
		}
		fmt.Printf("===>%s\n", cli)
		log.Println("Client address is", cli)
		fmt.Printf("After RPC handling. resp: %+v\n", resp)
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)

	pb_normal.RegisterRegistry_ServiceServer(s, svc)
	//	client.RegisterHealth_ServiceServer(s,svc)
	//	client.RegisterSubscribe_ServiceServer(s,svc)
	reflection.Register(s)
	//fmt.Printf("%+v",s)
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
	//log.Print("grpc server listening at ",s.GetServiceInfo())

}

func TestServer_Registry(t *testing.T) {
	//1.连接
	conn, err := grpc.Dial("127.0.0.1:5669", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常：%sn", err)
	}
	defer conn.Close()
	//2.实例化gRPC客户端
	client_grpc := pb_normal.NewRegistry_ServiceClient(conn)
	//3.组装请求参数
	req := new(pb_normal.RegistryRequest)
	req.Name = "zs"
	req.Ip = "127.0.0.1"
	req.Uuid = "12w1sws"
	//4.调用接口
	response, err := client_grpc.Registry(context.Background(), req)
	if err != nil {
		fmt.Printf("连接异常：%sn", err)
	}
	fmt.Printf("响应结果：%vn", response)
}
