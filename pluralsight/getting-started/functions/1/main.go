package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
	port := 3000
	// port, err := startWebServer(port, 2)
	_, err := startWebServer(port, 2)
	fmt.Println(err)
	// fmt.Println(port, err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server.....")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	// return errors.New("Something went wrong")
	return port, nil
}
