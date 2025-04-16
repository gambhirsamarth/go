package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("is even number : ", x, checkEven(x))
}

func checkEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}

}
