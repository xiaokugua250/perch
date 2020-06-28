/**
服务节点 创建工厂类方法
*/

package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/multiformats/go-multiaddr"
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
func P2PHostFactory(ctx context.Context, privatekey string, options []libp2p.Option) (host.Host, error) {

	if privatekey != "" {
		priv, _, err := GenSecurekeysByStr(privatekey)
		if err != nil {
			return nil, err
		}
		options = append(options, libp2p.Identity(priv))
	}

	//Ref https://discuss.libp2p.io/t/did-you-succeed-in-creating-a-bootstrap-across-the-network/277 设置节点为bootstrap节点
	//options=append(options,libp2p.AddrsFactory(newAddrsFactory()))
	return libp2p.New(ctx, libp2p.ChainOptions(options...))

}

func newAddrsFactory(advertiseAddrs []multiaddr.Multiaddr) func([]multiaddr.Multiaddr) []multiaddr.Multiaddr {
	return func([]multiaddr.Multiaddr) []multiaddr.Multiaddr {
		return advertiseAddrs
	}
}
