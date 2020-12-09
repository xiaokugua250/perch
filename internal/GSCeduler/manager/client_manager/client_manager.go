package client_manager

import (
	"context"
	"log"
	client "perch/internal/GSCeduler/manager/proto"
)

type server struct {
	client.Registry_ServiceServer
}

func (s *server) Registry(ctx context.Context, request *client.RegistryRequest)(*client.RegistryResponse, error){
	var (
		response client.RegistryResponse
		err error
	)
	log.Printf("request info is %s\n",request.String())


	return &response,err
}
