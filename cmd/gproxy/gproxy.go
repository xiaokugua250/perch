package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	ma "github.com/multiformats/go-multiaddr"
	"perch/internal/p2p"

	log "github.com/sirupsen/logrus"
	"perch/internal/gproxy"
)

const help = `
This example creates a simple HTTP Proxy using two libp2p peers. The first peer
provides an HTTP server locally which tunnels the HTTP requests with libp2p
to a remote peer. The remote peer performs the requests and 
send the sends the response back.
Usage: Start remote peer first with:   ./proxy
       Then start the local peer with: ./proxy -d <remote-peer-multiaddress>
Then you can do something like: curl -x "localhost:9900" "http://ipfs.io".
This proxies sends the request through the local peer, which proxies it to
the remote peer, which makes it and sends the response back.
`

func main() {

	flag.Usage = func() {
		fmt.Println(help)
		flag.PrintDefaults()
	}
	destPeer := flag.String("d", "", "destination peer address")
	port := flag.Int("p", 9900, "proxy port")
	p2pport := flag.Int("l", 12000, "libp2p listen port")
	flag.Parse()
	if *destPeer != "" {
		option := libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", *p2pport+1))
		var optionChain []libp2p.Option
		optionChain = append(optionChain, option)

		host, err := p2p.P2PHostFactory(context.Background(), optionChain)
		if err != nil {
			log.Fatalln(err)
		}
		destPeerID := gproxy.AddAddrToPeerstore(host, *destPeer)
		proxyAddr, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", *port))
		if err != nil {
			log.Fatalln(err)
		}
		proxy := gproxy.P2ProxyService(host, proxyAddr, destPeerID)
		proxy.Serve()

	} else {
		var optionChain []libp2p.Option

		host, err := p2p.P2PHostFactory(context.Background(), optionChain)
		if err != nil {
			log.Fatalln(err)

		}
		_ = gproxy.P2ProxyService(host, nil, "")
		<-make(chan struct{})

	}
}
