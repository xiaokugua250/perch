package p2p

/**
服务端网络通信服务
*/
import (
	"github.com/libp2p/go-libp2p-core/peer"

	"context"
	"github.com/libp2p/go-libp2p-core/host"

	swarm "github.com/libp2p/go-libp2p-swarm"
)

//todo 网络发现
func (node *NetWorkNode) NetWorkDiscover(ctx context.Context, target host.Host) error {

	return nil
}

//todo 网络连接
func (node *NetWorkNode) NetWorkConnect(ctx context.Context, target host.Host) error {
	targetInfo := peer.AddrInfo{
		ID:    target.ID(),
		Addrs: target.Addrs(),
	}
	return node.NodeHost.Connect(ctx, targetInfo)
	//return nil
}

//todo 关闭传输网络
func (node *NetWorkNode) NetWorkClose() {
	node.NodeHost.Close()
}

// Since we just tried and failed to dial, the dialer system will, by default
// prevent us from redialing again so quickly. Since we know what we're doing, we
// can use this ugly hack (it's on our TODO list to make it a little cleaner)
// to tell the dialer "no, its okay, let's try this again"
func (node *NetWorkNode) CleanNodeDiaCache(target host.Host) {
	node.NodeHost.Network().(*swarm.Swarm).Backoff().Clear(target.ID())
}
