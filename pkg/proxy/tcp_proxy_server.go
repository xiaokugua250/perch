/**
golang tcp代理
*/
package proxy

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	_ "perch/pkg/log"
)

type TCPProxyTarget struct {
	ProxyServer string `json:"proxy_server"`
	RemoteServer string `json:"remote_server"`
}



func (tcpproxy *TCPProxyTarget) StartTCPProxy() {
	var (
		err error
	)

	tcpServer, err := net.Listen("tcp", tcpproxy.ProxyServer)
	if err != nil {
		log.Error(err)
	}
	defer tcpServer.Close()
	for {
		conn,err := tcpServer.Accept()
		if err!= nil{
			log.Error(err)
			continue
		}
		go tcpproxy.TCPProxyConnHandler(conn)

	}

}


/**
TCP 连接处理函数，主要用于数据交换使用
将连接的数据进行转发
*/
func (tcpproxy *TCPProxyTarget)TCPProxyConnHandler(clientConn net.Conn) {


	 remoteConn ,err := net.Dial("tcp",tcpproxy.RemoteServer)
	 if err!=nil{
		log.Error(err)
	}
	//fmt.Println(clientConn.RemoteAddr().String(),remoteConn,clientConn.LocalAddr(),remoteConn.LocalAddr(),remoteConn.RemoteAddr())
	errChan := make(chan error, 2)
	go func() {
		_, err := io.Copy(clientConn, remoteConn)
		errChan <- err
	}()
	go func() {
		_, err := io.Copy(remoteConn,clientConn)
		errChan <- err
	}()

	<-errChan
}
