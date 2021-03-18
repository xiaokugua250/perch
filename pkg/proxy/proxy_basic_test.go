// go test -v proxy_basic_test.go proxy_basic.go
package proxy

import (
	"fmt"
	"log"
	"testing"
)

// go test -v -test.run  TestProxyServerSetup proxy_basic_test.go proxy_basic.go
func TestProxyServerSetup(t *testing.T) {
	var (
		err       error
		serverOps ServerOptions
	)
	fmt.Println("=============")
	serverOps.ServerLayer = Middle_Layer_Proxy
	serverOps.ForwardServer.Existed = false
	serverOps.ForwardServer.IP = "127.0.0.1"
	serverOps.ForwardServer.Port = "8000"
	serverOps.HTTPPort = 8080
	serverOps.HTTPSPort = 4430
	serverOps.SSHPort = 8022
	err = ServerInitWithOps(serverOps)
	if err != nil {
		log.Fatalln(err)
	}
	err = ServerSetupWithOps(serverOps)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("====%s", "A")

}

func TestProxyServerUpWithOptions(t *testing.T) {

	var (
		err       error
		serverOps ServerOptions
	)

	serverOps.ServerLayer = User_Layer_Proxy
	serverOps.ForwardServer.Existed = true
	serverOps.ForwardServer.IP = "127.0.0.1"
	serverOps.ForwardServer.Port = "8080"
	serverOps.HTTPPort = 9090
	serverOps.HTTPSPort = 9443
	serverOps.SSHPort = 9022
	err = ServerInitWithOps(serverOps)
	if err != nil {
		log.Fatalln(err)
	}
	err = ServerSetupWithOps(serverOps)
	if err != nil {
		log.Fatalln(err)
	}
}
