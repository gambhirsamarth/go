package main

import "fmt"

func main() {
	book := Book{
		ID:     1001,
		Title:  "A Tale of Two Cities",
		Author: "Charles Dickens",
		Price:  12.99,
	}
	fmt.Printf(`Initial Values in Book : %v`, book)
	updatePriceUsingPointer(&book, 15.99)
	fmt.Printf(`Updated Values in book: %v`, book)
}

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func updatePriceUsingPointer(b *Book, newPrice float64) {
	b.Price = newPrice
}
