package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	multiplex "github.com/libp2p/go-libp2p-mplex"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	secio "github.com/libp2p/go-libp2p-secio"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	"time"

	//"github.com/libp2p/go-ud"
	"log"
	"os"
	"os/signal"
	"perch/internal/p2p"
	"syscall"
)

func main() {
	var p2pnetwork p2p.NetworkP2P
	var p2pOptions []libp2p.Option

	ctx, cacel := context.WithCancel(context.Background())
	defer cacel()
	security := libp2p.Security(secio.ID, secio.New)
	muxers := libp2p.ChainOptions(
		libp2p.Muxer("/yamux/1.0.0", yamux.DefaultTransport),
		libp2p.Muxer("/mplex/6.7.0", multiplex.DefaultTransport),
	)
	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)
	listenAddr := libp2p.ListenAddrStrings(
		"/ip4/0.0.0.0/tcp/0", "/ip6/::/tcp/0/ws")

	p2pOptions = append(p2pOptions, muxers, security, listenAddr, transports)

	options := p2p.NetworkRuntimeOptions{
		Ctx:            ctx,
		NetworkOptions: p2pOptions,
	}

	p2pnetwork.StartNetworkP2P(options)

	for _, addr := range p2pnetwork.NetworkPeer.Addrs() {
		fmt.Printf("Listening P2P on %s/p2p/%s\n", addr.String(), p2pnetwork.NetworkPeer.ID().String())
	}

	//	pubs ,err := pubsub.NewGossipSub(ctx,p2pnetwork.NetworkPeer)
	_, err := pubsub.NewGossipSub(ctx, p2pnetwork.NetworkPeer)

	if err != nil {
		log.Println(err)
	}
	err = p2p.MDNSDiscoverySetup(ctx, p2pnetwork.NetworkPeer, time.Second*2, "mdns")
	if err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Println(err)
			}
			// 执行额外的清理操作
			for _, clean := range p2pnetwork.NetworkCleanFunc {
				clean()
			}
			p2pnetwork.NetworkPeer.Close()
			return
		case s := <-signalChan:
			log.Printf("捕获到信号%v，准备停止服务\n", s)
			p2pnetwork.NetworkPeer.Close()
			return
		}
	}
}
