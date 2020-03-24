package main

import "fmt"

func main() {
	names := [3]string{"Keith", "Kevin", "Kayla"}

	var names2 [2]string
	names2[0] = "Foo"
	names2[1] = "Bar"

	fmt.Println(names, names2)


}