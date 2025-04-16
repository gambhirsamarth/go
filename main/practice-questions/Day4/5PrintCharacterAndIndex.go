package main

import "fmt"

func main() {
	str := "hello world"

	for index, character := range str {
		fmt.Println(index, character)
	}
}
