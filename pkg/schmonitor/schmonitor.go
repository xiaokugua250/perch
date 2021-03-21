package schmonitor

import (
	"crypto/tls"
	"google.golang.org/grpc/credentials"
)

const (
	TcpProtocol  = iota + 1 //tcp 协议
	HttpProtocol            //HTTP协议
	GrpcProtocol            //GRPC协议

)

type ServerTarget struct {
	IP          string
	Port        string
	EnableSSL   bool
	SSLCertFile string
	SSLKeyFile  string
	Protocol    int
}

func LoadTLSCredentials(sslCertFile, sslKeyFile string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(sslCertFile, sslKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
