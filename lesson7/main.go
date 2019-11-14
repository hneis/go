package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func handleConn(ctx Context, c net.Conn) {
	defer c.Close()
	defer ctx.wg.Done()
	for {
		select {
		case <-ctx.close:
			fmt.Println("Close connection from ", c.RemoteAddr())
			return
		default:
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

type Context struct {
	wg    *sync.WaitGroup
	close chan struct{}
}

func task3() {
	context := Context{
		wg:    &sync.WaitGroup{},
		close: make(chan struct{}),
	}
	go func(ctx Context) {
		listener, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			fmt.Println("Accept connection from ", conn.RemoteAddr())
			ctx.wg.Add(1)
			// Запускаем как goroutine(возможность нескольких подключений к серверу)
			go handleConn(ctx, conn)
		}
	}(context)
	for {
		var input string
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Server will be close")
			close(context.close)
			break
		}
	}
	context.wg.Wait()
	fmt.Println("All connection closed\nExit")
}

func main() {
	task3()
}
