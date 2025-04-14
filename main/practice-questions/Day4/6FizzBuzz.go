package Day4

import "fmt"

func main() {
	x := 209

	if x%3 == 0 && x%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if x%3 == 0 {
		fmt.Println("Fizz")
	} else if x%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(x)
	}
}
