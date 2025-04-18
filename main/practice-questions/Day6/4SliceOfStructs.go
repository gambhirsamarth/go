package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}
	booksSlice := []book{}
	booksSlice = append(booksSlice, book{"The Hobbit", "J.R.R. Tolkien", 1937})
	booksSlice = append(booksSlice, book{"The Lord of the Rings", "J.R.R. Tolkien", 1954})
	booksSlice = append(booksSlice, book{"The Silmarillion", "J.R.R. Tolkien", 1977})

	for i := range booksSlice {
		if booksSlice[i].author == "J.R.R. Tolkien" {
			fmt.Println(booksSlice[i])
		}
	}
}
