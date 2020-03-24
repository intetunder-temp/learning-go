package main

func main() {
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	// for k, v := range wellKnownPorts {
	// for k := range wellKnownPorts {
	for _, v := range wellKnownPorts {
		// println(k, v)
		// println(k)
		println(v)
	}
}
