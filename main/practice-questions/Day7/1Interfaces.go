package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return 3.14 * circle.Radius * circle.Radius
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}

func main() {
	var shapes []Shape
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}
	triangle := Triangle{Base: 4, Height: 3}
	shapes = append(shapes, rectangle, circle, triangle)

	for _, shape := range shapes {
		fmt.Printf("Area of %T is %f\n", shape, shape.Area())
	}
}
