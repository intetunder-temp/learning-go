package main

import "fmt"

func main() {
	var firstName *string = new(string)

	*firstName = "Dani"

	fmt.Println(*firstName)
}
