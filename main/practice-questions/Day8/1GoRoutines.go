package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Routine")
}

func main() {
	/*
		routines are functions executed concurrently
		and can be called using the "go" keyword
		before a function call
	*/
	go sayHello()

	/*
		The main function might finish before sayHello() is executed
		because Go doesn’t wait for goroutines to finish.
	*/
	fmt.Println("Hello From Main !")
	time.Sleep(time.Second) // pauses the main function for 1 second
}
