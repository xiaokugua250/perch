package client_manager

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"perch/internal/GSCeduler/manager/proto/cs"
)

type server struct {
	//client.Registry_ServiceServer
	cs.UnimplementedRegistry_ServiceServer
	cs.UnimplementedSubscribe_ServiceServer
	cs.UnimplementedHealth_ServiceServer
}


func (s *server) Registry(ctx context.Context, request *cs.RegistryRequest)(*cs.RegistryResponse, error){
	var (
		response cs.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}


func (s *server) UnRegistry(ctx context.Context, request *cs.RegistryRequest)(*cs.RegistryResponse, error){
	var (
		response cs.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}



func (s *server) UnSubscribe(context.Context, *cs.SubscribeRequest) (*cs.SubscribeResponse, error) {
	return nil, nil
}



func (s *server) Subscribe(context.Context, *cs.SubscribeRequest) (*cs.SubscribeResponse, error) {
	return nil, nil
}

func (s *server) Health_Method(context.Context, *cs.HealthRequest) (*cs.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health_Method not implemented")
}
func (s *server) HeartBeat_Method(context.Context, *cs.HealthRequest) (*cs.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat_Method not implemented")
}
