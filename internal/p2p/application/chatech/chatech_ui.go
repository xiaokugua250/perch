/**
p2p chat use pubsub
*/
package main

import (
	"fmt"
	"github.com/gdamore/tcell"

	"github.com/rivo/tview"
	"io"
	"strings"
	"time"
)

var done = make(chan struct{})

// ChatUI is a Text User Interface (TUI) for a ChatRoom.
// The Run method will draw the UI to the terminal in "fullscreen"
// mode. You can quit with Ctrl-C, or by typing "/quit" into the
// chat prompt.
type ChatUI struct {
	chatroom     *ChatRoom
	app          *tview.Application
	roomList     *tview.TextView
	peersList    *tview.TextView
	msgWriter    io.Writer
	inputChannel chan string
	doneChannel  chan struct{} // use empty struct to trans done channel
}

func InitChatUI(chatroom *ChatRoom) *ChatUI {

	// start reading messages from the subscription in a loop

	app := tview.NewApplication()
	msgBox := tview.NewTextView()
	msgBox.SetDynamicColors(true)
	msgBox.SetBorder(true)
	//msgBox.SetTitle(fmt.Sprintf("Room:%s", chatroom.roomName))
	// text views are io.Writers, but they don't automatically refresh.
	// this sets a change handler to force the app to redraw when we get
	// new messages to display.
	msgBox.SetChangedFunc(
		func() {
			app.Draw()
		})

	avaliableRooms := tview.NewTextView()
	avaliableRooms.SetBorder(true).SetTitle("avaliable rooms")
	inputChannel := make(chan string, 32)
	//	input := tview.NewInputField().SetLabel(chatroom.nickName + ">").SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

	input := tview.NewInputField().SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

	input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			// we don't want to do anything if they just tabbed away
			return
		}
		line := input.GetText()
		if len(line) == 0 {
			// ignore blank lines
			return
		}
		if line == "/QUIT" || line == "/quit" {
			app.Stop()
			return
		}
		inputChannel <- line
		input.SetText("")
	})

	peersList := tview.NewTextView()
	peersList.SetBorder(true)
	peersList.SetTitle("Peers")

	msgboxTips := tview.NewTextView()
	msgboxTips.SetDynamicColors(true)
	msgboxTips.SetBorder(true)
	msgboxTips.SetTitle("tips")
	//msgBox_tips.SetText(time.Now().String())
	var chatinputform ChatInputForm
	inputForm := tview.NewForm().
		AddInputField("nick name", "", 20, nil, func(text string) {
			chatinputform.ChatUserNickName = text
		}).
		AddInputField("room name", "", 20, nil, func(text string) {
			chatinputform.ChatRoomName = text
		}).
		AddPasswordField("Password", "", 20, '*', func(text string) {
			chatinputform.ChatUserPasswd = text

		}).
		AddButton("conform", func() {

			app.SetFocus(input)
			inputLabel := ""
			if chatinputform.ChatUserNickName != "" {
				inputLabel = chatinputform.ChatUserNickName + ">"
				chatroom.nickName = chatinputform.ChatUserNickName
			} else {
				//inputLabel = chatroom.nickName+">"
				//chatroom.nickName= chatroom.nickName
				inputLabel = chatroom.self.Pretty() + ">"

			}
			if chatinputform.ChatRoomName != "" {
				msgBox.SetTitle(fmt.Sprintf("Room:%s", chatinputform.ChatRoomName))
				chatroom.roomName = chatinputform.ChatRoomName
			} else {
				msgBox.SetTitle(fmt.Sprintf("Room:%s", chatroom.roomName))
			}

			input.SetLabel(inputLabel)
		}).
		AddButton("quit", func() {
			app.Stop()
		})

	//input_form.SetTitle("input your name and passwd")
	// text views are io.Writers, but they don't automatically refresh.
	// this sets a change handler to force the app to redraw when we get
	// new messages to display.
	msgboxTips.SetChangedFunc(
		func() {
			app.Draw()
		})
	flex := tview.NewFlex().
		AddItem(avaliableRooms, 50, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(msgboxTips, 0, 1, false).
			AddItem(msgBox, 0, 3, false).
			AddItem(inputForm, 0, 3, true).
			AddItem(input, 6, 1, true), 0, 2, true).
		AddItem(peersList, 20, 1, false)

	//fmt.Printf("chat rooms is %s,%s,%s,%s,%s\n",chatRoom.NickName,chatRoom.self.Pretty(),chatinputform.ChatUserNickName,chatinputform.ChatRoomName)
	// start reading messages from the subscription in a loop

	app.SetRoot(flex, true).EnableMouse(true)

	return &ChatUI{
		chatroom:     chatroom,
		app:          app,
		roomList:     avaliableRooms,
		peersList:    peersList,
		msgWriter:    msgBox,
		inputChannel: inputChannel,
		doneChannel:  make(chan struct{}, 1),
	}

}

// Run starts the chat event loop in the background, then starts
// the event loop for the text UI.
func (ui *ChatUI) Run() error {

	go ui.handleEvents()
	defer ui.end()

	return ui.app.Run()
}

// end signals the event loop to exit gracefully
func (ui *ChatUI) end() {
	ui.doneChannel <- struct{}{}
}

func (ui *ChatUI) refreshRooms() {
	rooms := ui.chatroom.ListRooms()
	idStrs := make([]string, len(rooms))
	for i, p := range rooms {
		idStrs[i] = p
	}

	//cause data race
	ui.app.QueueUpdate(func() { // clean data race https://github.com/rivo/tview/issues/197
		ui.roomList.SetText(strings.Join(idStrs, "\n"))

	})
	ui.app.Draw()

}

func (ui *ChatUI) refreshPeers() {
	peers := ui.chatroom.ListPeers()
	idStrs := make([]string, len(peers))
	for i, p := range peers {
		idStrs[i] = shortID(p)
	}

	//cause data race
	ui.app.QueueUpdate(func() { // clean data race https://github.com/rivo/tview/issues/197
		ui.peersList.SetText(strings.Join(idStrs, "\n"))

	})
	ui.app.Draw()

}

func (ui *ChatUI) displayChatMessage(cm *ChatMessage) {
	prompt := withColor("green", fmt.Sprintf("%s>:", cm.SenderNickName))
	fmt.Fprintf(ui.msgWriter, "%s %s\n", prompt, cm.Message)
}

func (ui *ChatUI) handleEvents() {
	peerRefreshTicher := time.NewTicker(time.Second * 2)
	defer peerRefreshTicher.Stop()

	for {
		select {
		case input := <-ui.inputChannel:
			err := ui.chatroom.Publish(input)
			if err != nil {
				fmt.Errorf("publish error :%s\n", err)
			}
			ui.displaySelfMessage(input)

		case m := <-ui.chatroom.Messages:
			ui.displayChatMessage(m)
		case <-peerRefreshTicher.C:
			ui.refreshRooms()
			ui.refreshPeers()

		case <-ui.chatroom.ctx.Done():
			return
		case <-ui.doneChannel:
			return

		}
	}
}

// withColor wraps a string with color tags for display in the messages text box.
func withColor(color, msg string) string {
	return fmt.Sprintf("[%s]%s[-]", color, msg)
}

// displaySelfMessage writes a message from ourself to the message window,
// with our nick highlighted in yellow.
func (ui *ChatUI) displaySelfMessage(msg string) {
	prompt := withColor("yellow", fmt.Sprintf("<self %s>:", ui.chatroom.nickName))
	fmt.Fprintf(ui.msgWriter, "%s %s\n", prompt, msg)
}

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}
func NewChatUI(chatroom *ChatRoom) *ChatUI {
	app := tview.NewApplication()
	msgBox := tview.NewTextView()
	msgBox.SetDynamicColors(true)
	msgBox.SetBorder(true)
	msgBox.SetTitle(fmt.Sprintf("Room:%s", chatroom.roomName))
	// text views are io.Writers, but they don't automatically refresh.
	// this sets a change handler to force the app to redraw when we get
	// new messages to display.
	msgBox.SetChangedFunc(
		func() {
			app.Draw()
		})

	avaliableRooms := tview.NewTextView()
	avaliableRooms.SetBorder(true).SetTitle("avaliable rooms")
	inputChannel := make(chan string, 32)
	//	input := tview.NewInputField().SetLabel(chatroom.nickName + ">").SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

	input := tview.NewInputField().SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

	input.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			// we don't want to do anything if they just tabbed away
			return
		}
		line := input.GetText()
		if len(line) == 0 {
			// ignore blank lines
			return
		}
		if line == "/QUIT" || line == "/quit" {
			app.Stop()
			return
		}
		inputChannel <- line
		input.SetText("")
	})

	peersList := tview.NewTextView()
	peersList.SetBorder(true)
	peersList.SetTitle("Peers")

	msgBox_tips := tview.NewTextView()
	msgBox_tips.SetDynamicColors(true)
	msgBox_tips.SetBorder(true)
	msgBox_tips.SetTitle("tips")
	//msgBox_tips.SetText(time.Now().String())
	var chatinputform ChatInputForm
	input_form := tview.NewForm().
		AddInputField("nick name", "", 20, nil, func(text string) {
			chatinputform.ChatUserNickName = text
		}).
		/*	AddInputField("room name", "", 20, nil, func(text string) {
			chatinputform.ChatRoomName = text
		}).*/
		AddPasswordField("Password", "", 20, '*', func(text string) {
			chatinputform.ChatUserPasswd = text
		}).
		AddButton("conform", func() {

			app.SetFocus(input)
			inputLabel := ""
			if chatinputform.ChatUserNickName != "" {
				inputLabel = chatinputform.ChatUserNickName + ">"
			} else {
				inputLabel = chatroom.nickName + ">"
			}
			input.SetLabel(inputLabel)
		}).
		AddButton("quit", func() {
			app.Stop()
		})
	//input_form.SetTitle("input your name and passwd")
	// text views are io.Writers, but they don't automatically refresh.
	// this sets a change handler to force the app to redraw when we get
	// new messages to display.
	msgBox_tips.SetChangedFunc(
		func() {
			app.Draw()
		})
	flex := tview.NewFlex().
		AddItem(avaliableRooms, 50, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(msgBox_tips, 0, 1, false).
			AddItem(msgBox, 0, 3, false).
			AddItem(input_form, 0, 3, true).
			AddItem(input, 6, 1, true), 0, 2, true).
		AddItem(peersList, 20, 1, false)

	app.SetRoot(flex, true).EnableMouse(true)
	return &ChatUI{
		chatroom:     chatroom,
		app:          app,
		roomList:     avaliableRooms,
		peersList:    peersList,
		msgWriter:    msgBox,
		inputChannel: inputChannel,
		doneChannel:  make(chan struct{}, 1),
	}

}
