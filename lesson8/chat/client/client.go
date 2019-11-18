package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hneis/go/lesson8/chat"
	"github.com/marcusolsson/tui-go"
)

type Ui struct {
	root           *tui.Box
	ui             tui.UI
	history        *tui.Box
	sendMessage    chan chat.Message
	receiveMessage chan chat.Message
	username       string
}

func (u *Ui) AddMessage(m chat.Message) {
	message := m.(chat.BroadcastMessage)
	u.ui.Update(func() {
		u.history.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", message.User))),
			tui.NewLabel(fmt.Sprintf("%v", message.Message)),
			tui.NewSpacer(),
		))
	})
}

func (u *Ui) UpdateUserList(users []string) {
	u.ui.Update(func() {
		u.root.Remove(0)
		// for test
		// u.history.Append(tui.NewHBox(
		// 	tui.NewLabel(fmt.Sprintf("%v", users)),
		// 	tui.NewSpacer(),
		// ))
		sidebar := tui.NewVBox(
			tui.NewLabel("Users"),
		)
		sidebar.SetBorder(true)
		for _, c := range users {
			sidebar.Append(tui.NewLabel(c))
		}
		sidebar.Append(tui.NewSpacer())
		u.root.Prepend(sidebar)
	})
}

func (u *Ui) Run() {
	sidebar := tui.NewVBox(
		tui.NewLabel("Users"),
	)
	sidebar.SetBorder(true)
	history := tui.NewVBox()
	u.history = history

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chatUi := tui.NewVBox(historyBox, inputBox)
	chatUi.SetSizePolicy(tui.Expanding, tui.Expanding)

	input.OnSubmit(func(e *tui.Entry) {
		u.sendMessage <- chat.NewBroadcastMessage(e.Text(), u.username)
		input.SetText("")
	})

	u.root = tui.NewHBox(sidebar, chatUi)

	ui, err := tui.New(u.root)
	if err != nil {
		log.Fatal(err)
	}
	u.ui = ui

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	go func(ui *Ui) {
		for m := range ui.receiveMessage {
			switch m.Type() {
			case 1, 2, 3:
				u.AddMessage(m)
			case 4:
				message := m.(chat.ClientListMessage)
				u.UpdateUserList(message.Clients)
			}
		}
	}(u)

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func (u *Ui) CloseChans() {
	close(u.sendMessage)
	close(u.receiveMessage)
}

func writeToServer(conn net.Conn, ch <-chan chat.Message) {
	for m := range ch {
		conn.Write(m.Data())
	}
}

func readFromServer(conn net.Conn, ch chan<- chat.Message) {
	for {
		//TODO add connection check

		message, err := chat.GetMessage(conn)
		if err != nil {
			continue
		} else {
			ch <- message
		}
	}
}

func main() {
	flag.Parse()
	username := flag.Args()[0]

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	ui := &Ui{
		sendMessage:    make(chan chat.Message),
		receiveMessage: make(chan chat.Message),
		username:       username,
	}
	defer ui.CloseChans()

	go writeToServer(conn, ui.sendMessage)
	go readFromServer(conn, ui.receiveMessage)

	ui.sendMessage <- chat.NewChangeNameMessage(username)

	ui.Run()

}
