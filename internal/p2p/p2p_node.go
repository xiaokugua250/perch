package p2p

import "github.com/libp2p/go-libp2p"

/**
服务节点 创建工厂类方法
*/

import (
	"context"
	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/host"
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
func NetworkNodeInit(ctx context.Context, listenAddr, privatekey string, options []libp2p.Option) (NetWorkNode, error) {
	var (
		node NetWorkNode
		err  error
	)
	if privatekey != "" {
		priv, _, err := utils.GenSecurekeysByStr(privatekey)
		if err != nil {
			return node, err
		}
		options = append(options, libp2p.Identity(priv))
	}

	if listenAddr != "" {
		node.NodeHost, err = libp2p.New(ctx, libp2p.ListenAddrStrings(listenAddr), libp2p.ChainOptions(options...))
	}

	node.NodeHost, err = libp2p.New(ctx, libp2p.ChainOptions(options...))

	return node, err

}

/**
初始化網絡結點
*/
func (node *NetWorkNode) NetworkNodeshudown() error {
	return node.NodeHost.Close()
}

/**
创建普通P2P 节点
*/

func CreateP2PServerNode(ctx context.Context, listenAddr, privatekey string, options []libp2p.Option) (host.Host, error) {
	if privatekey != "" {
		priv, _, err := utils.GenSecurekeysByStr(privatekey)
		if err != nil {
			return nil, err
		}

		options = append(options, libp2p.Identity(priv))
	}

	if listenAddr != "" {
		return libp2p.New(ctx, libp2p.EnableRelay(circuit.OptActive), libp2p.ListenAddrStrings(listenAddr), libp2p.ChainOptions(options...))
	}

	return libp2p.New(ctx, libp2p.EnableRelay(circuit.OptActive), libp2p.ChainOptions(options...))

}

/**
创建P2P 中继节点
*/
func CreateP2PRelayNode(ctx context.Context, listenAddr, privatekey string, options []libp2p.Option) (host.Host, error) {

	if privatekey != "" {
		priv, _, err := utils.GenSecurekeysByStr(privatekey)
		if err != nil {
			return nil, err
		}

		options = append(options, libp2p.Identity(priv))
	}

	if listenAddr != "" {
		return libp2p.New(ctx, libp2p.EnableRelay(circuit.OptHop), libp2p.ListenAddrStrings(listenAddr), libp2p.ChainOptions(options...))

	}

	return libp2p.New(ctx, libp2p.EnableRelay(circuit.OptHop), libp2p.ChainOptions(options...))
}

/**
创建必须要中继节点的服务节点
*/
func CreateP2PServerBootstrapNode(ctx context.Context, listenAddr, privatekey string, options []libp2p.Option) (host.Host, error) {
	if privatekey != "" {
		priv, _, err := utils.GenSecurekeysByStr(privatekey)
		if err != nil {
			return nil, err
		}

		options = append(options, libp2p.Identity(priv))
	}

	if listenAddr != "" {
		return libp2p.New(ctx, libp2p.ListenAddrStrings(listenAddr), libp2p.ChainOptions(options...))
	}
	return libp2p.New(ctx, libp2p.ChainOptions(options...))
}
