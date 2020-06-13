package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
}

/*
direct : 0
direct : 1
direct : 2
direct : 3
direct : 4
goroutine : 0
goroutine : 1
goroutine : 2
going
goroutine : 3
goroutine : 4
*/
