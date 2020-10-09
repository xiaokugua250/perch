/**
p2p 网络模块
*/
package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"log"
)

type NetworkP2P struct {
	NetworkBasicHost host.Host `json:"network_peer"`
	NetworkInitFunc  []func() error
	NetworkCleanFunc []func() error
}

type NetworkRuntimeOptions struct {
	Ctx            context.Context `json:"ctx"`
	NetworkName    string          `json:"network_name"`
	NetworkOptions []libp2p.Option `json:"network_options"`
	NetworkPeer    host.Host       `json:"network_peer"`
}

func (p2pnetwork *NetworkP2P) InitialNetworkP2P() {

	var err error

	for _, init := range p2pnetwork.NetworkInitFunc {
		err = init()
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func (p2pnetwork *NetworkP2P) StartBasicNetworkP2P(options NetworkRuntimeOptions) {
	var err error
	p2pnetwork.InitialNetworkP2P()

	p2pnetwork.NetworkBasicHost, err = P2PBasicHostFactory(options.Ctx, "", options.NetworkOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("network node id is %s, network listening at %s\n", p2pnetwork.NetworkBasicHost.ID(), p2pnetwork.NetworkBasicHost.Addrs())
	// 捕获退出信号

}
