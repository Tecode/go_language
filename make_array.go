package main

import "fmt"

func main() {
	// List
	list := make([]int, 5)
	fmt.Printf("%v \n", list)
	// Map
	maps := make(map[string]string)
	maps["value1"] = "php"
	maps["value2"] = "Go"

	fmt.Printf("%v \n", maps)
}
