package main

import (
	"fmt"
	"time"
)

func task1() {
	cancelCh := make(chan struct{})
	go spinner(cancelCh, 50*time.Millisecond)

	// например так
	whait := time.After(time.Second * 2)
	select {
	case <-whait:
		close(cancelCh)
	}
	// можно так
	// fmt.Scanln()
}

func spinner(cancelCh <-chan struct{}, delay time.Duration) {
	for {
		select {
		case <-cancelCh:
			return
		default:
			for _, r := range "-\\|/" {
				fmt.Printf("%c\r", r)
				time.Sleep(delay)
			}
		}
	}
}

func main() {
	task1()
}
