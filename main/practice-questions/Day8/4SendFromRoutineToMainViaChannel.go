package main

import "fmt"

func sendData(data chan int) {
	data <- 10
}

func main() {
	channel := make(chan int)
	go sendData(channel)
	valueReceived := <-channel
	fmt.Println(valueReceived)
}
