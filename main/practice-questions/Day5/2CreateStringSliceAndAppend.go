package main

import "fmt"

func main() {
	stringSlice := []string{"hello", "world"}
	stringSlice = append(stringSlice, "text")

	for i := 0; i < len(stringSlice); i++ {
		fmt.Println(stringSlice[i])
	}
}
