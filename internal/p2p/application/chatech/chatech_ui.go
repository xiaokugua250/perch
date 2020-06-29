/**
p2p chat use pubsub
*/
package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	_ "github.com/libp2p/go-libp2p-core/peer"
	"github.com/rivo/tview"
	"io"
	"strings"
	"time"
)

// ChatUI is a Text User Interface (TUI) for a ChatRoom.
// The Run method will draw the UI to the terminal in "fullscreen"
// mode. You can quit with Ctrl-C, or by typing "/quit" into the
// chat prompt.
type ChatUI struct {
	chatroom     *ChatRoom
	app          *tview.Application
	peersList    *tview.TextView
	msgWriter    io.Writer
	inputChannel chan string
	doneChannel  chan struct{} // use empty struct to trans done channel
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

	inputChannel := make(chan string, 32)
	input := tview.NewInputField().SetLabel(chatroom.nickName + ">").SetFieldWidth(0).SetFieldBackgroundColor(tcell.ColorBlack)

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

	chatPannel := tview.NewFlex().AddItem(msgBox, 0, 1, false).AddItem(peersList, 20, 1, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(chatPannel, 0, 1, false).AddItem(input, 1, 1, true)

	app.SetRoot(flex, true)
	return &ChatUI{
		chatroom:     chatroom,
		app:          app,
		peersList:    peersList,
		msgWriter:    msgBox,
		inputChannel: inputChannel,
		doneChannel:  make(chan struct{}, 1),
	}

}

// Run starts the chat event loop in the background, then starts
// the event loop for the text UI.
func (ui *ChatUI) Run() error {

	defer ui.end()
	return ui.app.Run()
}

// end signals the event loop to exit gracefully
func (ui *ChatUI) end() {
	ui.doneChannel <- struct{}{}
}

func (ui *ChatUI) refreshPeers() {
	peers := ui.chatroom.ListPeers()
	idStrs := make([]string, len(peers))
	for i, p := range peers {
		idStrs[i] = shortID(p)
	}
	ui.peersList.SetText(strings.Join(idStrs, "\n"))
	ui.app.Draw()
}

func (ui *ChatUI) displayChatMessage(cm *ChatMessage) {
	prompt := withColor("green", fmt.Sprintf("<%s>:", cm.SenderNickName))
	fmt.Fprintf(ui.msgWriter, "%s %s\n", prompt, cm.Message)
}

func (ui *ChatUI) handleEvents() {
	peerRefreshTicher := time.NewTicker(time.Second)
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
	prompt := withColor("yellow", fmt.Sprintf("<%s>:", ui.chatroom.nickName))
	fmt.Fprintf(ui.msgWriter, "%s %s\n", prompt, msg)
}
