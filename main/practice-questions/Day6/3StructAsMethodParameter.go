package main

import "fmt"

type Rectangle struct {
	length float32
	width  float32
}

func main() {
	rectangle := Rectangle{length: 5, width: 3.5}
	fmt.Println(rectangle.Area())
	/*
		When you define a function like this:
		func (r Rectangle) Area() float32

		You're telling Go:
		"This is a method that works on the Rectangle type."

		So instead of passing the struct as a parameter manually,
		Go lets you call this method on the instance itself, like:
		rectangle.Area()

		Under the hood, it's the same as:

		Area(rectangle)
		But Go gives it this cleaner syntax when you use method receivers.
	*/
}

func (rectangle Rectangle) Area() (area float32) {
	area = rectangle.length * rectangle.width
	return area
}
