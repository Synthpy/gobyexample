package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {
	fmt.Print(1)
}

func main() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
}
