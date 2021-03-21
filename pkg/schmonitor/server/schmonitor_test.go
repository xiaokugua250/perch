package server

import (
	"log"
	"perch/pkg/schmonitor"
	"testing"
)

func TestServerSetupWithOpt(t *testing.T) {

	var (
		servers []schmonitor.ServerTarget
	)
	servers = []schmonitor.ServerTarget{
		schmonitor.ServerTarget{
			IP:          "127.0.0.1",
			Port:        "9898",
			EnableSSL:   false,
			SSLCertFile: "",
			SSLKeyFile:  "",
			Protocol:    schmonitor.TcpProtocol,
		},
		schmonitor.ServerTarget{
			IP:          "127.0.0.1",
			Port:        "9899",
			EnableSSL:   false,
			SSLCertFile: "",
			SSLKeyFile:  "",
			Protocol:    schmonitor.HttpProtocol,
		},
	}
	if err := ServerSetupWithOpt(servers); err != nil {
		log.Println(err)
	}
}
