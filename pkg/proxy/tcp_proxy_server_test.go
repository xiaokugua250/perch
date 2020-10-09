/**
golang tcp代理
*/
package proxy

import (
	_ "perch/pkg/log"
	"testing"
)



func TestTCPProxyTarget_StartTCPProxy(t *testing.T) {

	tcpProxyTarget :=&TCPProxyTarget{
		ProxyServer:"127.0.0.1:8089",
		RemoteServer:"127.0.0.1:8081",

	}
	tcpProxyTarget.StartTCPProxy()

}
