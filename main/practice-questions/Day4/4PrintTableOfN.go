package Day4

import "fmt"

func main() {
	x := 6
	fmt.Println("Table of ", x)
	printTable(x)
}

func printTable(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x * i)
	}
}
