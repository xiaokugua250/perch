package main

import (
	"context"
	"fmt"
	cid "github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	multiplex "github.com/libp2p/go-libp2p-mplex"
	secio "github.com/libp2p/go-libp2p-secio"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	ws "github.com/libp2p/go-ws-transport"
	_ "github.com/multiformats/go-multibase"
	mh "github.com/multiformats/go-multihash"
	cli "github.com/urfave/cli/v2"
	"time"
	//"github.com/libp2p/go-ud"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"perch/internal/p2p"
	_ "perch/pkg/general/log"
	"syscall"
)

func main() {

	var err error
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			p2p_networkRunner()
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func p2p_networkRunner() {
	var p2pnetwork p2p.NetworkP2P
	var p2pOptions []libp2p.Option
	var err error

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
		"/ip4/172.16.171.94/tcp/0", "/ip6/::/tcp/0/ws")
	//		"/ip4/0.0.0.0/tcp/0", "/ip6/::/tcp/0/ws")

	p2pOptions = append(p2pOptions, muxers, security, listenAddr, transports)

	options := p2p.NetworkRuntimeOptions{
		Ctx:            ctx,
		NetworkOptions: p2pOptions,
	}

	p2pnetwork.StartBasicNetworkP2P(options)

	for _, addr := range p2pnetwork.NetworkBasicHost.Addrs() {
		fmt.Printf("Listening P2P on %s/p2p/%s\n", addr.String(), p2pnetwork.NetworkBasicHost.ID().String())
	}
	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Make the DHT NOTE - Using Client constructor
	dhtobj := dht.NewDHT(ctx, p2pnetwork.NetworkBasicHost, dstore)
	data := []byte("this is some test content")
	hash, _ := mh.Sum(data, mh.SHA2_256, -1)

	contentId := cid.NewCidV1(cid.DagCBOR, hash)
	fmt.Printf("contend is is %s\n", contentId.String(), contentId.Bytes())
	if err = dhtobj.Provide(ctx, contentId, false); err != nil {
		log.Error(err)
	}

	routedHost, err := p2p.MakeRoutedNetworkP2P(p2pnetwork.NetworkBasicHost, ctx, dhtobj)
	if err != nil {
		log.Error(err)
	}
	routingDiscovery := discovery.NewRoutingDiscovery(dhtobj)

	discovery.Advertise(ctx, routingDiscovery, "meet me here")
	log.Println("Successfully announced!")

	/*
		p2pnetwork.NetworkBasicHost= rhost.Wrap(p2pnetwork.NetworkBasicHost, dhtobj)
		err = dhtobj.Bootstrap(ctx)
		if err!= nil{
			panic(err)
		}
	*/
	//	pubs ,err := pubsub.NewGossipSub(ctx,p2pnetwork.NetworkBasicHost)
	//pubs, err := pubsub.NewGossipSub(ctx, p2pnetwork.NetworkBasicHost)
	pubs, err := p2p.PubsubgossipGen(ctx, p2pnetwork.NetworkBasicHost)
	if err != nil {
		log.Error(err)
	}
	sub, topsub, err := p2p.PubsubtopicsJoin(pubs, p2p.Pubsub_Default_Topic)
	if err != nil {
		log.Error(err)
	}

	err = p2p.MDNSDiscoverySetup(ctx, p2pnetwork.NetworkBasicHost, p2p.DiscoveryInterval, p2p.DiscoveryServiceTag)
	if err != nil {
		log.Error(err)
	}

	go func() {
		for {
			msg := new(p2p.PubsubMessage)
			msg.SenderPeer = p2pnetwork.NetworkBasicHost.ID().Pretty()
			msg.PMessageStr = "hello world"
			msg.SenderFrom = "from localhost"
			err = p2p.PubsubTopicPubish(ctx, *msg, topsub, nil)
			if err != nil {
				log.Error(err)
			}
			time.Sleep(3 * time.Second)
		}

	}()
	msgChan := make(chan interface{})

	go p2p.PubsubMsgHandler(sub, ctx, p2pnetwork.NetworkBasicHost, msgChan)
	/*	for {
		select {
		case msg := <- msgChan:
			fmt.Println("msg from msg chan is:",msg)
		case <-time.Tick(5*time.Second):
			fmt.Println("time.is over")
			//return
		}
	}*/
	signalChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("msg from msg chan is:", msg)
			for _, peersFind := range p2p.PubsubPeersList(pubs, p2p.Pubsub_Default_Topic) {
				peerinfo, err := dhtobj.FindPeer(ctx, peersFind)
				if err != nil {
					continue
				}
				if err = routedHost.Connect(ctx, peerinfo); err != nil {
					fmt.Print("connect b")
				}
				fmt.Printf("peer found by dht is %s\n", peerinfo.String())
			}

			//	fmt.Print("====>",dhtobj.RoutingTable().ListPeers())

		case err := <-errChan:
			if err != nil {
				log.Println(err)
			}
			// 执行额外的清理操作
			for _, clean := range p2pnetwork.NetworkCleanFunc {
				clean()
			}
			p2pnetwork.NetworkBasicHost.Close()
			return
		case s := <-signalChan:
			log.Printf("捕获到信号%v，准备停止服务\n", s)
			p2pnetwork.NetworkBasicHost.Close()
			return
		}
	}
}
