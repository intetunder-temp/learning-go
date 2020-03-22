package main

import "fmt"

type user2 struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	var u user
	u.ID = 1
	u.FirstName = "Dani"
	u.LastName = "Scarpa"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{ID: 1, FirstName: "Dani", LastName: "Scarpa"}
	fmt.Println(u2)

	u3 := user{ID: 1,
		FirstName: "Daniela",
		LastName:  "Scarpa",
	}

	fmt.Println(u3)

	u4 := user2{ID: 1,
		FirstName: "Danstar",
		LastName:  "Scarpa",
	}
	fmt.Println(u4)
}
