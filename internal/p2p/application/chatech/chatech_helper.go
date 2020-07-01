/**
p2p chat use pubsub
*/
package main

import "fmt"

type ChatInputForm struct {
	ChatRoomName     string `json:"chat_room_name"`
	ChatUserNickName string `json:"chat_user_nick_name"`
	ChatUserPasswd   string `json:"chat_user_passwd"`
}

/**
驗證user
*/
func ChatRoomJoinHelper(user ChatInputForm) {
	fmt.Printf("user is %#v\n", user)
}
