package main

import "fmt"

func receiveData(data chan int) {
	fmt.Println("Hello from Routine")
	valueFromChannel := <-data
	fmt.Println(valueFromChannel)

}

func main() {
	ch := make(chan int)
	go receiveData(ch)
	ch <- 5
	fmt.Println("Hello from Main")

	/*
		Code Execution Flow:
		1 A channel ch is created.

		2 A goroutine is started using go receiveData(ch).

		3 The goroutine prints "Hello from Routine".

		4 The goroutine waits for data from the channel ch.

		5 The main function sends data (5) into the channel ch, which unblocks the goroutine.

		6 The goroutine prints the received value (5).

		7 The goroutine finishes its execution.

		8 The main function prints "Hello from Main".
	*/
}
