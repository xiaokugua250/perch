package stream

import (

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbStream "perch/internal/GSCeduler/manager/proto/pb_stream"
	"time"
)

type Server struct {
	pbStream.UnimplementedStream_ServiceServer
}
func (*Server) PushStream(srv pbStream.Stream_Service_PushStreamServer) error {
	log.Println("Start new server....")
	//startime:= time.Now()
	for {
		srv.Send(&pbStream.StreamResponse{
			Code:         0,
			Message:      time.Now().String(),
			StreamData:   nil,
			ErrorMessage: "",
			Total:        1,
		})
		time.Sleep(2*time.Second);
	}
	return status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
