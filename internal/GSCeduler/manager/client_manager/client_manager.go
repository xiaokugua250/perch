package client_manager

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"perch/internal/GSCeduler/manager/proto/pb_normal"
)

type Server struct {
	//client.Registry_ServiceServer
	pb_normal.UnimplementedRegistry_ServiceServer
	pb_normal.UnimplementedSubscribe_ServiceServer
	pb_normal.UnimplementedHealth_ServiceServer
}


func (s *Server) Registry(ctx context.Context, request *pb_normal.RegistryRequest)(*pb_normal.RegistryResponse, error){
	var (
		response pb_normal.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}


func (s *Server) UnRegistry(ctx context.Context, request *pb_normal.RegistryRequest)(*pb_normal.RegistryResponse, error){
	var (
		response pb_normal.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}



func (s *Server) UnSubscribe(context.Context, *pb_normal.SubscribeRequest) (*pb_normal.SubscribeResponse, error) {
	return nil, nil
}



func (s *Server) Subscribe(context.Context, *pb_normal.SubscribeRequest) (*pb_normal.SubscribeResponse, error) {
	return nil, nil
}

func (s *Server)Health_Method(ctx context.Context, in *pb_normal.HealthRequest, opts ...grpc.CallOption) (*pb_normal.HealthResponse, error){
	return nil, nil
}



func (s *Server)Heart_Method(ctx context.Context, in *pb_normal.HealthRequest, opts ...grpc.CallOption) (*pb_normal.HealthResponse, error){
	return nil, nil
}
