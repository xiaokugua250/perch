package stream

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"perch/internal/GSCeduler/manager/proto/pb_stream"
	"testing"
	 log "github.com/sirupsen/logrus"
)

func TestNewStream_ServiceClient(t *testing.T) {
	lis,err := net.Listen("tcp","0.0.0.0:5668")
	if err != nil{
		return
	}
	//创建一个grpc 服务器
	var opts []grpc.ServerOption
	svc := Server{}
	s := grpc.NewServer(opts...)
	//注册事件
	pb_stream.RegisterStream_ServiceServer(s, &svc)

	//处理链接
	s.Serve(lis)
}



func TestNewStream_ServiceClient2(t *testing.T) {
	conn ,err := grpc.Dial("127.0.0.1:5668",grpc.WithInsecure())
	if err != nil{
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb_stream.NewStream_ServiceClient(conn)
	//调用服务端推送流
	/*reqstreamData := &pb_stream.StreamRequest{
		Name:       "aaa",
		Uuid:       "121saxsaxasxs",
		Ip:         "",
		Port:       0,
		Status:     "",
		StreamData: nil,
	}*/
	res,_ := c.PushStream(context.Background())
	for {
		aa,err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}


	select {
	}

}






