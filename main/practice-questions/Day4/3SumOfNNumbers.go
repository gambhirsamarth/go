package main

import "fmt"

func main() {
	x := 10
	sum := sumTillN(x)
	fmt.Println(sum)
}

func sumTillN(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
