/**
服务节点 创建工厂类方法
*/

package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"testing"
)

func TestP2PHostFactory(t *testing.T) {
	ctx := context.Background()
	var options []libp2p.Option
	options = append(options, libp2p.EnableRelay())
	nodeHost, err := P2PBasicHostFactory(ctx, "", options)
	if err != nil {
		panic(err)
	}
	fmt.Printf("node host is %#v,id is %s \n", nodeHost, nodeHost.ID().String())
}
