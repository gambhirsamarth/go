package main

func main() {
	// Declaration Without Initialization
	var declaration string

	// Declaration With Initialization
	var declarartionWithInitialization string = "bar"

	// Multiple Declarations
	var strOne, strTwo string = "Hello", "World"
	// OR
	var (
		stringOne string = "Hello"
		stringTwo string = "World"
	)

	// Type is omitted but will be inferred
	var str = "Hello World"

	// Shorthand - MOSTLY USED
	// Shorthand only works inside function bodies
	foo := "Shorthand String"

	// Constants
	const constant = "This is a constant"

}
