/**
p2p 網絡的發布/訂閱模式實現
*/
package p2p

import (
	"context"
	"encoding/json"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	Pubsub_Default_Topic = "p2p_topic"
)

/**

 */
type PubsubMessage struct {
	PMessageStr  string      `json:"p_message_str"`
	SenderPeer   string      `json:"sender_peer"`
	SenderFrom   string      `json:"sender_from"`
	PMessageData interface{} `json:"pmsg_data"`
}

/**
generate pubsubgossip
*/
func PubsubgossipGen(ctx context.Context, host host.Host) (*pubsub.PubSub, error) {

	return pubsub.NewGossipSub(ctx, host)
}

func PubsubtopicsJoin(pubs *pubsub.PubSub, topic string, topopts ...pubsub.TopicOpt) (*pubsub.Subscription, *pubsub.Topic, error) {
	topicpub, err := pubs.Join(topic, topopts...)
	if err != nil {
		return nil, topicpub, err
	}
	sub, err := topicpub.Subscribe()
	if err != nil {
		return nil, topicpub, err
	}

	return sub, topicpub, nil

}

/**
主題發送消息
*/
func PubsubTopicPubish(ctx context.Context, msg PubsubMessage, topics *pubsub.Topic, topicOpts ...pubsub.PubOpt) error {

	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if len(topicOpts) == 0 {
		return topics.Publish(ctx, msgData)
	}

	return topics.Publish(ctx, msgData)
}

func PubsubPeersList(pubs *pubsub.PubSub, topic string) []peer.ID {
	return pubs.ListPeers(topic)
}

/**
消息處理
*/
func PubsubMsgHandler(subs *pubsub.Subscription, ctx context.Context, selfpeer host.Host, MsgChan chan interface{}) {

	for { // handler msg in loop
		//msgChan := make(chan  interface{})
		msg, err := subs.Next(ctx)
		if err != nil {
			return
			//close()
		}
		// only forward messages delivered by others

		if msg.ReceivedFrom == selfpeer.ID() {
			continue
		}

		pubMsg := new(PubsubMessage)
		err = json.Unmarshal(msg.Data, pubMsg)
		if err != nil {
			continue
		}

		/*
			cm := new(ChatMessage)
			err = json.Unmarshal(msg.Data, cm)
			if err != nil {
				continue
			}

		*/
		MsgChan <- pubMsg

	}

}
