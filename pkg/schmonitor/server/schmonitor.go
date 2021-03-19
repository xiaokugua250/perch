package server

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"net/http"
	"perch/pkg/schmonitor"
	"sync"
	"time"
)

type SetupOptions struct {
	IP          string
	Port        string
	SSLCertFile string
	SSLKeyFile  string
	Protocols   []int
}

/**
TCP 协议接收文件服务端,支持大文件传输
*/
func serverWithTCP(target schmonitor.ServerTarget) error {

	var (
		err      error
		listener net.Listener
	)

	listener, err = net.Listen("tcp", target.IP+":"+target.Port)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go TcpConnectionHandler(conn)
	}

	return err

}

/**
小文件，等进行HTTP传输
*/
func ServerWithHTTP(target schmonitor.ServerTarget) error {
	var (
		err error
	//	listener net.Listener
	)
	httpServer := &http.Server{
		Addr: target.IP + ":" + target.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if target.EnableSSL {
		err = httpServer.ListenAndServeTLS(target.SSLCertFile, target.SSLKeyFile)
		if err != nil {
			return err
		}
	} else {
		err = httpServer.ListenAndServe()
		if err != nil {
			return err
		}
	}
	return nil

}

func ServerWithGRPC(target schmonitor.ServerTarget) error {
	var (
		listener   net.Listener
		grpcServer *grpc.Server
		err        error
	)
	listener, err = net.Listen("tcp", target.IP+":"+target.Port)
	if err != nil {
		return fmt.Errorf("could not list on %s: %s", target.IP+":"+target.Port, err)
	}
	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}
	if target.EnableSSL {
		tlsCredentials, err := schmonitor.LoadTLSCredentials(target.SSLCertFile, target.SSLKeyFile)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		grpcServer = grpc.NewServer(grpc.Creds(tlsCredentials), grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	} else {
		grpcServer = grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}

func ServerWithSocket() {

}

func SetupWithOpt(target schmonitor.ServerTarget) error {

	var (
		err error
	)

	switch target.Protocol {
	case schmonitor.GrpcProtocol:

	case schmonitor.HttpProtocol:

	case schmonitor.SocketProtocol:
	case schmonitor.TcpProtocol:
	default:
		log.Fatalln("protocol not support....")
	}
	return err

}

func InitWithOpt() {

}

func ServerSetupWithOpt(options SetupOptions) error {

	var (
		err error
		wg  sync.WaitGroup
	)
	if len(options.Protocols) == 0 {
		return errors.New(fmt.Sprintf("protocol options is none,i.e len(options.protocol)==0"))
	}
	wg.Add(len(options.Protocols))
	for _, protocol := range options.Protocols {

		targetServer := schmonitor.ServerTarget{
			IP:       options.IP,
			Port:     options.Port,
			Protocol: protocol,
		}
		//wg.Done()
		//fmt.Println(targetServer)
		//	go SetupWithOpt(targetServer)
		//wg.Done()
		go func() {
			defer wg.Done()
			SetupWithOpt(targetServer)
		}()

	}
	wg.Wait()
	return err

}

func TcpConnectionHandler(conn net.Conn) {

}
