package main

import "fmt"

func main() {
	type Book struct {
		title  string
		author string
		price  float32
	}

	var book1 Book
	book1.title = "One Piece"
	book1.author = "Eichiro Oda"
	book1.price = 20.99
	book2 := Book{title: "Naruto", author: "Masashi Kishimoto", price: 25.99}
	fmt.Println(book1)
	fmt.Println(book2)
}
