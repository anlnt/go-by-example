package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)
	go func() {
		messages <- "ping"
		time.Sleep(time.Second * 5)
		messages <- "pong"
	}()
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
ping
pong
*/
