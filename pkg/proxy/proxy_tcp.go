package proxy

import (
	"io"
	"log"
	"net"
)

type TCPProxy struct {
	addr string
}

func NewTCPProxy(addr string) *TCPProxy {
	return &TCPProxy{addr}
}

// 数据包交互，通过管道进行交互
func TCPPayLoadHandler(payloadIn io.ReadWriter, payloadOut io.ReadWriter) {

	var (
		errStream = make(chan error, 2)
	)
	go func() {
		_, err := io.Copy(payloadIn, payloadOut)
		errStream <- err
	}()
	go func() {
		_, err := io.Copy(payloadOut, payloadIn)
		errStream <- err
	}()
	<-errStream

}

func (tcpProxy *TCPProxy) TCPProxyHandler(netCon net.Conn) {

	var (
		err    error
		remote net.Conn
	)
	defer netCon.Close()
	remote, err = net.Dial("tcp", tcpProxy.addr)
	if err != nil {
		log.Println(err)
		return
	}
	TCPPayLoadHandler(netCon, remote)
	remote.Close()
}

//监听tcp连接请求

func (tcpProxy *TCPProxy) TcpProxyServe(listener net.Listener) {

	for {
		if conn, err := listener.Accept(); err != nil {
			log.Println(err)
		} else {

			go tcpProxy.TCPProxyHandler(conn)
		}

	}

}
