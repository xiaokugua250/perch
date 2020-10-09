/**
ref to https://github.com/libp2p/go-libp2p-examples/tree/42a4cd9ae8765175380ed99c06be4e3c29323f7b/routed-echo

1、添加libp2p contentrouting 內容，contentrouting 核心流程如下：1、構建routed host，2、構建dht對象，3、節點鏈接(bootstrap.connect,discovery等)、4、內容提供和發現 對應provider 和findprovieder
*/
package main

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	mh "github.com/multiformats/go-multihash"
	"time"
	"fmt"
	"github.com/libp2p/go-libp2p"
	ds "github.com/ipfs/go-datastore"
	dsync "github.com/ipfs/go-datastore/sync"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	_ "perch/pkg/general/log"
)

func main() {
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 10001)),
		//	libp2p.Identity(priv),
		libp2p.DefaultTransports,
		libp2p.DefaultMuxers,
		libp2p.DefaultSecurity,
		libp2p.NATPortMap(),
	}

	ctx := context.Background()

	basicHost, err := libp2p.New(ctx, opts...)
	if err != nil {
		log.Error(err)
	}
	opts_1 := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", 10002)),
		//	libp2p.Identity(priv),
		libp2p.DefaultTransports,
		libp2p.DefaultMuxers,
		libp2p.DefaultSecurity,
		libp2p.NATPortMap(),
	}

	basicHost_1, err := libp2p.New(ctx, opts_1...)
	if err != nil {
		log.Error(err)
	}

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	dstore := dsync.MutexWrap(ds.NewMapDatastore())
	dstore_1 := dsync.MutexWrap(ds.NewMapDatastore())
	// Make the DHT
	dhtObj, err := dht.New(ctx, basicHost, dht.Datastore(dstore), dht.Mode(dht.ModeServer))
	if err != nil {
		log.Error(err)
	}
	//dhtObj := dht.NewDHT(ctx, basicHost, dstore)
	data := []byte("this is some test content")
	hash, _ := mh.Sum(data, mh.SHA2_256, -1)

	contentId := cid.NewCidV1(cid.DagCBOR, hash)
	fmt.Printf("contend is is %s\n", contentId.String())
	if err = dhtObj.Provide(ctx, contentId, false); err != nil {
		log.Error(err)
	}
	// Make the routed host
	routedHost := rhost.Wrap(basicHost, dhtObj)

	// connect to the chosen ipfs nodes

	// Bootstrap the host
	err = dhtObj.Bootstrap(ctx)
	if err != nil {
		log.Error(err)
	}

	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", routedHost.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	// addr := routedHost.Addrs()[0]
	addrs := routedHost.Addrs()
	log.Println("I can be reached at:")
	for _, addr := range addrs {
		log.Println(addr.Encapsulate(hostAddr))
	}

	//dht_1 := dht.New(ctx, basicHost_1, dstore)
	basicHost_addressinfo := peer.AddrInfo{
		ID:    routedHost.ID(),
		Addrs: routedHost.Addrs(),
	}

	//works both
	//dhtobj1, err := dht.New(ctx, basicHost_1, dht.Datastore(dstore_1), dht.Mode(dht.ModeServer), dht.BootstrapPeers(basicHost_addressinfo))
	dhtobj1, err := dht.New(ctx, basicHost_1, dht.Datastore(dstore_1), dht.Mode(dht.ModeServer))

	//dtbatch,err := dstore.Batch()
	//dhtobj1  = dht.NewDHTClient(ctx, basicHost_1, dtb)
	routedhost1 := rhost.Wrap(basicHost_1, dhtobj1)
	if err != nil {
		panic(err)
	}
	if err = dhtobj1.Bootstrap(ctx); err != nil {
		log.Error(err)
	}
	if err = routedhost1.Connect(ctx, basicHost_addressinfo); err != nil {
		log.Error(err)
	}

	log.Println("successfully connected to basichost...")
	for {
		time.Sleep(5 * time.Second)
		peers, err := dhtobj1.FindProviders(ctx, contentId)
		if err != nil {
			log.Error(err)
		}
		if len(peers) <= 0 {
			log.Println("found zero peers....", peers)
		}
		for _, peer := range peers {
			fmt.Printf("found peer %s provider contedt %s\n", peer, contentId.String())
		}
		fmt.Printf("===begin to found peer provied content %s\n", contentId.String())
	}

	//log.Printf("Now run \"./routed-echo -l %d -d %s%s\" on a different terminal\n", routedHost.ID(), routedHost.ID().Pretty())

}
