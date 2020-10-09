/*

https://superuser.com/questions/1330027/how-to-enable-mdns-on-windows-10-build-17134
*/
package p2p

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"log"
	"time"
)

const (
	DiscoveryInterval = time.Second * 5

	// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
	DiscoveryServiceTag = "p2p_discovery_tag"
)

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

func (d discoveryNotifee) HandlePeerFound(info peer.AddrInfo) {
	//panic("implement me")
	//	fmt.Printf("found peer by mdns %s\n", info)

	err := d.h.Connect(context.Background(), info)
	if err != nil {
		log.Panic(err)
	}
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func MDNSDiscoverySetup(ctx context.Context, h host.Host, interval time.Duration, disveryTag string) error {

	disc, err := discovery.NewMdnsService(ctx, h, interval, disveryTag)
	if err != nil {
		return err
	}
	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)
	return nil
}
