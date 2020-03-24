package main

import "fmt"

func main() {
	fmt.Println("Simple String")
	fmt.Println(`
	This is a multi-line
	String, that can also contain "quote".
	`)
	fmt.Println("?")
	fmt.Println("\u2272")
}
