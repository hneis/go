package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"syscall"

	"github.com/hneis/go/lesson8/chat"
)

type Client struct {
	conn     net.Conn
	ch       chan chat.Message
	username string
	//TODO mutex?
	//TODO wait group?
}

type Server struct {
	clients  map[*Client]bool
	bChan    chan chat.Message
	entering chan *Client
	leaving  chan *Client
}

func NewChatServer() *Server {
	return &Server{
		clients:  map[*Client]bool{},
		bChan:    make(chan chat.Message, 1),
		entering: make(chan *Client),
		leaving:  make(chan *Client),
	}
}

func (s *Server) broadcaster() {
	for {
		select {
		case msg := <-s.bChan:
			fmt.Println("broadcast message")
			fmt.Println("clients count: ", len(s.clients))
			for cli := range s.clients {
				cli.ch <- msg
			}
		case cli := <-s.entering:
			fmt.Println("Entering")
			s.clients[cli] = true
		case cli := <-s.leaving:
			fmt.Println("Leaving", cli)
			userNames := []string{}
			delete(s.clients, cli)
			close(cli.ch)
			for c := range s.clients {
				userNames = append(userNames, c.username)
			}
			fmt.Println("leaving: ", userNames)
			s.bChan <- chat.NewClientListMessage(userNames)
			// s.bChan <- chat.NewBroadcastMessage("leaving")
		}
	}
}

func (s *Server) handleConn(conn net.Conn) {
	fmt.Println("New conection", conn.LocalAddr().String())
	client := &Client{
		conn: conn,
		ch:   make(chan chat.Message),
	}
	s.clients[client] = false
	s.entering <- client

	go clientWriter(conn, client.ch)
	for {
		if err := s.connCheck(conn); err != nil {
			log.Printf("detected closed LAN connection")
			s.leaving <- client
			return
		}

		message, err := chat.GetMessage(conn)
		if err != nil {
			continue
		}
		switch message.Type() {
		case 3:
			cm := message.(chat.ChangeNameMessage)
			client.username = cm.Username
			userNames := []string{}
			for cli := range s.clients {
				userNames = append(userNames, cli.username)
			}
			s.bChan <- chat.NewClientListMessage(userNames)
		default:
			s.bChan <- message
			fmt.Printf("%s : %v\n", client.username, message)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan chat.Message) {
	for msg := range ch {
		fmt.Println("clientWriter", msg.Data())
		conn.Write(msg.Data())
	}
}

func (s Server) connCheck(conn net.Conn) error {
	var sysErr error = nil
	rc, err := conn.(syscall.Conn).SyscallConn()
	if err != nil {
		return err
	}
	err = rc.Read(func(fd uintptr) bool {
		var buf []byte = []byte{0}
		n, _, err := syscall.Recvfrom(int(fd), buf, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
		switch {
		case n == 0 && err == nil:
			sysErr = io.EOF
		case err == syscall.EAGAIN || err == syscall.EWOULDBLOCK:
			sysErr = nil
		default:
			sysErr = err
		}
		return true
	})
	if err != nil {
		return err
	}

	return sysErr
}

func (s *Server) Run(server string) {
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	go s.broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConn(conn)
	}
}
