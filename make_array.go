package main

import "fmt"

func main() {
	// List
	list := make([]int, 5)
	fmt.Println(list)
	// Map
	maps := make(map[string]string)
	maps["value1"] = "php"
	maps["value2"] = "Go"

	fmt.Println(maps)
}
