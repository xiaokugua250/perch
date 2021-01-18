package client_manager

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"perch/internal/GSCeduler/manager/proto/pb_normal"
)

type server struct {
	//client.Registry_ServiceServer
	pb_normal.UnimplementedRegistry_ServiceServer
	pb_normal.UnimplementedSubscribe_ServiceServer
	pb_normal.UnimplementedHealth_ServiceServer
}

func (s *server) Registry(ctx context.Context, request *pb_normal.RegistryRequest) (*pb_normal.RegistryResponse, error) {
	var (
		response pb_normal.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n", request.String())

	return &response, err
}

func (s *server) UnRegistry(ctx context.Context, request *pb_normal.RegistryRequest) (*pb_normal.RegistryResponse, error) {
	var (
		response pb_normal.RegistryResponse
		err      error
	)
	log.Printf("request info is %s\n", request.String())

	return &response, err
}

func (s *server) UnSubscribe(context.Context, *pb_normal.SubscribeRequest) (*pb_normal.SubscribeResponse, error) {
	return nil, nil
}

func (s *server) Subscribe(context.Context, *pb_normal.SubscribeRequest) (*pb_normal.SubscribeResponse, error) {
	return nil, nil
}

func (s *server) Health_Method(context.Context, *pb_normal.HealthRequest) (*pb_normal.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health_Method not implemented")
}
func (s *server) HeartBeat_Method(context.Context, *pb_normal.HealthRequest) (*pb_normal.HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat_Method not implemented")
}
