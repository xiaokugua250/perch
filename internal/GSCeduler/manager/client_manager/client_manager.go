package client_manager

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"perch/internal/GSCeduler/manager/proto/cs"
)

type Server struct {
	//client.Registry_ServiceServer
	cs.UnimplementedRegistry_ServiceServer
	cs.UnimplementedSubscribe_ServiceServer
	cs.UnimplementedHealth_ServiceServer
}


func (s *Server) Registry(ctx context.Context, request *cs.RegistryRequest)(*cs.RegistryResponse, error){
	var (
		response cs.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}


func (s *Server) UnRegistry(ctx context.Context, request *cs.RegistryRequest)(*cs.RegistryResponse, error){
	var (
		response cs.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}



func (s *Server) UnSubscribe(context.Context, *cs.SubscribeRequest) (*cs.SubscribeResponse, error) {
	return nil, nil
}



func (s *Server) Subscribe(context.Context, *cs.SubscribeRequest) (*cs.SubscribeResponse, error) {
	return nil, nil
}

func (s *Server)Health_Method(ctx context.Context, in *cs.HealthRequest, opts ...grpc.CallOption) (*cs.HealthResponse, error){
	return nil, nil
}



func (s *Server)Heart_Method(ctx context.Context, in *cs.HealthRequest, opts ...grpc.CallOption) (*cs.HealthResponse, error){
	return nil, nil
}
