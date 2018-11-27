package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["php"] = 41
	m["go"] = 84
	fmt.Printf("%v \n", m)
	delete(m, "php")
	fmt.Printf("%v", m)

	v, ok := m["php"], m["go"]
	println(v, ok)
}
