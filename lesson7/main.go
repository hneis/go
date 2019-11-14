package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	urls := []string{
		"duckduckgo.com",
		"www.google.com",
		"www.rambler.com",
		"www.ya.ru",
	}
	for _, url := range urls {
		go func(url string, r chan<- string) {
			r <- request(url)
		}(url, responses)
	}

	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	conn, err := net.Dial("tcp", hostname+":80")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%s time elapsed %v", hostname, time.Since(start))
	return hostname
}

func task4() {
	fmt.Println("Find fast host:")
	fastUrl := mirroredQuery()

	// Ждем 1 секунду, после чего выводим самого быстрого
	timer := time.After(time.Second)
	select {
	case <-timer:
		fmt.Printf("The fastest is %s\n", fastUrl)
	}
}

func main() {
	task4()
}
