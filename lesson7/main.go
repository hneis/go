package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
)

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func task2() {
	cancelChanel := make(chan struct{})
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func(n int, cc chan<- struct{}) {
		for x := 0; x < n; x++ {
			naturals <- x
		}
		close(cc)
	}(100000, cancelChanel)

	// возведение в квадрат
	go func(cc <-chan struct{}) {
		for {
			select {
			case x := <-naturals:
				squares <- x * x
			case <-cc:
				close(squares)
				return
			}
		}
	}(cancelChanel)

	// печать
	var str string
	wString := bytes.NewBufferString(str)
	bufio.NewWriter(wString)
	for sq := range squares {
		wString.WriteString(strconv.Itoa(sq))
		wString.WriteString("\n")
	}
	fmt.Println(wString)
}

func main() {
	task2()
}
