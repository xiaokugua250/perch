package stream

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbStream "perch/internal/GSCeduler/manager/proto/pb_stream"
	"time"
)

type StreamServer struct {
	pbStream.UnimplementedStream_ServiceServer
}

func (*StreamServer)PushStream(context.Context, *pbStream.StreamRequest) (*pbStream.StreamResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
func (*StreamServer)PullStream(ctx context.Context, req *pbStream.StreamRequest) (*pbStream.StreamResponse, error) {
	i:= 0
	for{
		i++
		res.Send(&pbStream.StreamResponse{})
		time.Sleep(1*time.Second)
		if i >10 {
			break
		}
	}
	return nil,nil
//	return nil, status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
func (*StreamServer)BidirectionStream(context.Context, *pbStream.StreamRequest) (*pbStream.StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
