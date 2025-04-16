package main

import "fmt"

func main() {
	numbers := []float64{12.5, 34.6, 23.8, 45.9, 3.8}
	fmt.Println("Average of slice is: ", calculateAverage(numbers))
}

func calculateAverage(numbers []float64) (average float64) {
	average = 0
	sum := 0.0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	average = sum / float64(len(numbers))
	return
}
