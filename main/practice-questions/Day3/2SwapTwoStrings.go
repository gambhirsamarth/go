package main

import "fmt"

func main() {
	fmt.Println(swapStrings("hello", "world"))
}

func swapStrings(strOne, strTwo string) (string, string) {
	var temp string = strTwo
	strTwo = strOne
	strOne = temp
	return strOne, strTwo
}

/* Or Simpler Version
func swapStrings(a, b string) (string, string) {
	return b, a
}
*/
