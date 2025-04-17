package main

import "fmt"

func main() {
	number := 7
	pointer := &number
	var answer = (*pointer) * (*pointer)
	fmt.Println(answer)
}
