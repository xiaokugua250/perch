package p2p

import (
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	multiplex "github.com/libp2p/go-libp2p-mplex"
	secio "github.com/libp2p/go-libp2p-secio"
	yamux "github.com/libp2p/go-libp2p-yamux"
	"github.com/libp2p/go-tcp-transport"
	"time"

	//"github.com/libp2p/go-ws-transport"
	ws "github.com/libp2p/go-ws-transport"
	mh "github.com/multiformats/go-multihash"
	"testing"
)

func TestPubsubTopicPubish(t *testing.T) {
	var p2pnetwork NetworkP2P
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

	options := NetworkRuntimeOptions{
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
	fmt.Printf("contend is is %s\n", contentId.String())

	for {
		fmt.Println("begin to sleep 5 sec....")
		peerID, err := peer.Decode("/ip4/127.0.0.1/tcp/45965/p2p/QmSW4tv2SXZ3UDmSreXFnhwpFpacQ8oEXKMw18fTGSuq67")

		//peerID := peer.ID("/ip4/127.0.0.1/tcp/45965/p2p/QmSW4tv2SXZ3UDmSreXFnhwpFpacQ8oEXKMw18fTGSuq67")

		if err := dhtobj.Ping(ctx, peerID); err != nil {
			panic(err)
		}

		address, err := dhtobj.FindProviders(ctx, contentId)
		dhtobj.RoutingTable().Print()
		if err != nil {
			fmt.Println(err)
		}
		for _, addr := range address {
			fmt.Printf("address is %s\n", addr)
		}
		time.Sleep(5 * time.Second)

	}

}
