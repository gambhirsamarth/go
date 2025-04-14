package main

import "fmt"

func main() {
	const pi float64 = 3.1415
	var radius int = 5
	var area float64 = pi * float64(radius) * float64(radius)
	fmt.Println(`Area of circle is `, area)
}
