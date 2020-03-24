package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 11

	switch ages["Kevin"] {
	case 1, 2, 3, 5, 6, 11, 13, 17, 19:
		fmt.Println("Kevin's age is a prime number!")
	case 16:
		fmt.Println("Kevin can drive now")
	case 18:
		fmt.Println("Kevin can vote now!")
	case 67:
		fmt.Println("Kevin can retire now!")
	default:
		fmt.Println("There's nothing special about Kevin's current age.")
}
}