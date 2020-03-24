package main

import "fmt"

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 18
	ages["Keith"] = 28
	ages["James"] = 67
	ages["Micheal"] = 16
	ages["Leigha"] = 5

	for i := 1; i <= 100; i++ {
		fmt.Println("We are counting", i)
	}

	a := 0
	for a < 100 {
		fmt.Println("We are counting (again)", a)
	}

	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println(name, "has a prime number age")
		case 16:
			fmt.Println(name, "can drive")
		case 18:
			fmt.Println(name, "can vote")
		case 67:
			fmt.Println(name, "can retire")
		default:
			fmt.Println(fmt.Sprintf("There is nothing special about %s's current age.", name))
		}
	}
}
