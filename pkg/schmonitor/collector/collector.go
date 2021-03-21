package collector

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"perch/pkg/schmonitor"
)

type Collector interface {
	// 客户端向服务端注册并且建立连接，需要完成的工作有：1、客户端信息上报，2、与服务端的网络连接建立；需要处理异常有：连接超时，连接中断重试等
	CollectorRegisterWithOpt(server schmonitor.ServerTarget) error
	CollectorAgent() error
	//CollectorConnection(server schmonitor.ServerTarget) error
	//CollectorPush() error
}

type LoggerCollector struct {
	LogLocation string
}

type FileCollector struct {
}

type GenCollector struct {
}

type BasicCollector struct {
}

func (log *LoggerCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
		clientConn, err := net.Dial("tcp", server.IP+":"+server.Port)
		if err != nil {
			return err
		}
		defer clientConn.Close()
		//todo 客户端连接已经建立,执行业务逻辑

	case schmonitor.HttpProtocol:
		// todo 采用HTTP 协议连接 http server
		/*httpClient := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       15*time.Second,
		}
		httpRequest := http.NewRequest()
		resp,err := httpClient.Do(httpRequest)
		if err!= nil{
			return err
		}
		defer resp.Body.Close()
		body,err := ioutil.ReadAll(resp.Body)
		if err!= nil{
			return err
		}
		// handle body*/

	case schmonitor.GrpcProtocol:
		var (
			err      error
			grpcConn *grpc.ClientConn
		)

		//todo 建立client时需要
		grpcConn, err = grpc.Dial(server.IP+":"+server.Port, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer grpcConn.Close()
		//todo 建立grpc连接

	case schmonitor.WebSocketProtocol:
		// 建立websocket 连接
		//todo 参考
		// https://github.com/gorilla/websocket/blob/master/examples/echo/client.go
		/*var (
			websocketConn websocket.Conn
			httpResp http.Response
			err error
		)
		websocketConn, httpResp,err = websocket.DefaultDialer.Dial()
		if err != nil{
			return err
		}
		defer websocketConn.Close()
		done := make(chan struct{})
		go func() {
			defer close(done)
			for {

			}
		}()*/

	default:
		return errors.New(fmt.Sprintf("server protocol [%d] not support....", server.Protocol))

	}
	return err
}

/*
func (log *LoggerCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}
*/
/**
日志采集
*/
func (log *LoggerCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}

//----
func (file *FileCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
	case schmonitor.HttpProtocol:
	case schmonitor.GrpcProtocol:
	case schmonitor.WebSocketProtocol:
	default:
		return nil

	}
	return err

}

/**
文件采集
*/
func (file *FileCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}

/*
func (file *FileCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}*/

//--

func (basic *BasicCollector) CollectorRegisterWithOpt(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	switch server.Protocol {
	case schmonitor.TcpProtocol:
	case schmonitor.HttpProtocol:
	case schmonitor.GrpcProtocol:
	case schmonitor.WebSocketProtocol:
	default:
		return nil

	}
	return err
}

func (basic *BasicCollector) CollectorAgent() error {
	var (
		err error
	)
	return err
}

/*func (basic *BasicCollector) CollectorConnection(server schmonitor.ServerTarget) error {
	var (
		err error
	)
	return err
}
*/
