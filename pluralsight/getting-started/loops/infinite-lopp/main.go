package main

func main() {
	var i int
	// Ugly way
	// for ; ; {
	// Good way, vim auto removes the ";"
	for {
		if i == 5 {
			// Run's forever without break
			// break
		}
		println(i)
		i++
	}
}
