package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(circleStats(5))
}

func circleStats(radius float64) (float64, float64) {
	const pi float64 = 3.1415
	var area float64 = pi * radius * radius
	var circumference = 2 * pi * radius
	return circumference, area
}

/* OR Code with
func circleStats(radius float64) (circumference, area float64) {
	const pi = 3.1415
	circumference = 2 * pi * radius
	area = pi * radius * radius
	return
}
*/
