// Package m provides ...
package main

import (
	"github.com/hneis/go/lesson8/chat/server"
)

func main() {
	s := server.NewChatServer()
	s.Run("localhost:8000")
}
