/**
p2p chat use pubsub
*/
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
	"os"
	"time"

	"github.com/libp2p/go-libp2p-core/host"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// DiscoveryInterval is how often we re-publish our mDNS records.
const DiscoveryInterval = time.Hour

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.
const DiscoveryServiceTag = "pubsub-chat-example"

func main() {
	nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	roomFlag := flag.String("room", "awesome-chat-room", "name of chat room to join")
	flag.Parse()

	ctx := context.Background()
	h, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	// setup local mDNS discovery
	err = setupMDNSDiscovery(ctx, h)
	if err != nil {
		panic(err)
	}
	nick := *nickFlag
	if len(nick) == 0 {
		nick = defaultNick(h.ID())
	}
	room := *roomFlag

	chatroom, err := JoinChatRoom(ctx, ps, h, nick, room)
	if err != nil {
		panic(err)
	}
	fmt.Printf("host listening to %s",h.Addrs)
	ui := InitChatUI(chatroom)

	if err = ui.Run(); err != nil {
		fmt.Printf("error in running text ui %s", err)
	}

}

// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(peer peer.AddrInfo) {
	//panic("implement me")
	fmt.Printf("discovered new peer %s\n", peer.ID.Pretty())
	err := n.h.Connect(context.Background(), peer)
	if err != nil {
		fmt.Printf("error connecting to peer %s:%s\n", peer.ID.Pretty(), err)
	}

}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupMDNSDiscovery(ctx context.Context, h host.Host) error {

	disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		return err
	}
	n := discoveryNotifee{h: h}
	disc.RegisterNotifee(&n)

	return nil
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-8:]
}
