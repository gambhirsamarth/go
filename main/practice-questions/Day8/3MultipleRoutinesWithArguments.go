package main

import (
	"fmt"
	"time"
)

func printMessage(message rune) {
	fmt.Println("Message Received", message)
}

func main() {
	for i := 0; i < 5; i++ {
		go printMessage(rune(i))
	}
	time.Sleep(time.Second)
}
