package main

import "fmt"

func main() {

	names := []string{}
	names = append(names, "Keith")
	names = append(names, "Kevin", "Kayla", "Kyle")

	fmt.Println(names)
	fmt.Println("names[2] is nil:", names[2] == "")
}
