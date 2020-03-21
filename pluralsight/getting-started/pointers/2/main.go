package main

import "fmt"

func main() {
	firstName := "Dani"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Danstar"
	fmt.Println(ptr, *ptr)
}
