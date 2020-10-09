/**
服务节点 创建工厂类方法
*/

package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	rhost "github.com/libp2p/go-libp2p/p2p/host/routed"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	_ "perch/pkg/general/log"
)

/**
網絡節點結構
*/
type NetWorkNode struct {
	NodeHost host.Host
}

/**
初始化網絡結點
*/
func (node *NetWorkNode) NetworkNodeshudown() error {

	return node.NodeHost.Close()
}

/**
创建P2P NODE 的工厂方法
*/
func P2PBasicHostFactory(ctx context.Context, privatekey string, options []libp2p.Option) (host.Host, error) {

	if privatekey != "" {
		priv, _, err := GenSecurekeysByStr(privatekey)
		if err != nil {
			return nil, err
		}
		options = append(options, libp2p.Identity(priv))
	}

	//Ref https://discuss.libp2p.io/t/did-you-succeed-in-creating-a-bootstrap-across-the-network/277 设置节点为bootstrap节点
	//options=append(options,libp2p.AddrsFactory(newAddrsFactory()))
	return libp2p.New(ctx, libp2p.ChainOptions(options...)), nil

}

func newAddrsFactory(advertiseAddrs []multiaddr.Multiaddr) func([]multiaddr.Multiaddr) []multiaddr.Multiaddr {
	return func([]multiaddr.Multiaddr) []multiaddr.Multiaddr {
		return advertiseAddrs
	}
}

func MakeRoutedNetworkP2P(basichost host.Host, ctx context.Context, dhtObj *dht.IpfsDHT) (*rhost.RoutedHost, error) {
	var err error

	// Construct a datastore (needed by the DHT). This is just a simple, in-memory thread-safe datastore.
	//dstore := dsync.MutexWrap(ds.NewMapDatastore())

	// Make the DHT
	//dhtObj := dht.NewDHT(ctx, basichost, dstore)

	// Make the routed host
	routedHost := rhost.Wrap(basichost, dhtObj)

	// connect to the chosen ipfs nodes
	/*err = bootstrapConnect(ctx, routedHost, bootstrapPeers)
	if err != nil {
		return nil, err
	}*/

	// Bootstrap the host
	err = dhtObj.Bootstrap(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("network node id is %s, network listening at %s\n", routedHost.ID(), routedHost.Addrs())
	// 捕获退出信号
	return routedHost, err

}
