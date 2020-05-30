/**
p2p 网络模块
*/
package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
)

/**
创建P2P NODE 的工厂方法
*/
func P2PHostFactory(ctx context.Context, option []libp2p.Option) (host.Host, error) {

	//libp2p.ListenAddrStrings()
	return libp2p.New(ctx, option...)

}
