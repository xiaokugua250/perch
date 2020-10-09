/**
p2p chat use pubsub
*/
package main

import (
	"context"
	"encoding/json"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRootBuffer = 128

type ChatRoom struct {
	Messages chan *ChatMessage
	ctx      context.Context
	pubs     *pubsub.PubSub
	topic    *pubsub.Topic
	sub      *pubsub.Subscription
	roomName string
	self     host.Host
	nickName string
}

// ChatMessage gets converted to/from JSON and sent in the body of pubsub messages.
type ChatMessage struct {
	Message        string
	SenderID       string
	SenderNickName string
}

// JoinChatRoom tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func JoinChatRoom(ctx context.Context, pubs *pubsub.PubSub, selfID host.Host, nickName string, roomName string) (*ChatRoom, error) {
	// join the pubsub topic

	topic, err := pubs.Join(topicName(roomName))
	if err != nil {
		return nil, err
	}
	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}
	chatRoom := &ChatRoom{
		ctx:      ctx,
		pubs:     pubs,
		sub:      sub,
		topic:    topic,
		self:     selfID,
		nickName: nickName,
		roomName: roomName,
		Messages: make(chan *ChatMessage, ChatRootBuffer),
	}
	// start reading messages from the subscription in a loop
	go chatRoom.readLoop()
	return chatRoom, nil
}

// Publish sends a message to the pubsub topic.
func (chatRoom *ChatRoom) Publish(message string) error {
	msg := ChatMessage{
		Message:        message,
		SenderID:       chatRoom.self.ID().Pretty(),
		SenderNickName: chatRoom.nickName,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return chatRoom.topic.Publish(chatRoom.ctx, msgBytes)
}

func (chatRoom *ChatRoom) ListPeers() []peer.ID {
	return chatRoom.pubs.ListPeers(topicName(chatRoom.roomName))
}
func (chatRoom *ChatRoom) ListRooms() []string {
	return chatRoom.pubs.GetTopics()
}

func (chatRoom *ChatRoom) readLoop() {
	for {

		msg, err := chatRoom.sub.Next(chatRoom.ctx)
		if err != nil {
			close(chatRoom.Messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom == chatRoom.self.ID() {
			continue
		}
		cm := new(ChatMessage)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		chatRoom.Messages <- cm

	}
}

func topicName(roomName string) string {
	return "chat-room:" + roomName
}
