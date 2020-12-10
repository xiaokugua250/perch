package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"net"
)

func GetClietIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("getClinetIP, invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("getClientIP, peer.Addr is nil")
	}

	return pr.Addr.String(), nil
}
