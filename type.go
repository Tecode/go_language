package main

import (
	"encoding/json"
	"fmt"
)

type Age int
type Name string
type Json struct {
	xAlies int
	yAlies int
}


func main() {
	var name Name = "Bob"
	var age Age = 45
	var m = make(map[string] Json)

	n := map[string]Json{
		"a2": {4, 6},
		"b2": {5, 6},
	}
	arrayMap := []Json{
		{4, 5},
		{69, 8},
	}

	data, _ := json.Marshal(arrayMap)

	fmt.Println(string(data))

	m["a"] = Json{12, 65}
	fmt.Println(name, age, m, arrayMap, n)
}
