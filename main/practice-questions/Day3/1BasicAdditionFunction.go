package main

import "fmt"

func main() {
	summed := sum(1, 2)
	fmt.Println(summed)
}

func sum(a, b int) int {
	return a + b
}
