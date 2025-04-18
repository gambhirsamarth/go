package main

import "fmt"

type User struct {
	Name string
	Age  int
	add  Address
}

type Address struct {
	State string
	City  string
}

func main() {
	user := User{
		Name: "John",
		Age:  30,
		add: Address{
			State: "California",
			City:  "Los Angeles",
		},
	}
	fmt.Printf("User: %v\n", user)

	// you can use the %+v verb instead of %v. It prints field names as well
	fmt.Printf("User: %+v\n", user)
}
