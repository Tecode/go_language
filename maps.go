package main

import (
	"fmt"
)

type Vertex struct {
	x, y, z float32
}

type item struct {
	name string
	id   int
}

func main() {
	var m = make(map[string]Vertex)
	m["value1"] = Vertex{14.2354, 25.33154, -1.02147}
	m["value2"] = Vertex{-36.221, 36.2214, -741.3654}
	fmt.Printf("%v", m)

	var json = make(map[string]item)
	json["bob"] = item{name: "Ele", id: 56}
	fmt.Println(json)

	data := map[string]string{"name": "Aming", "idCart": "5110102244"}
	/*
		map中的值是否存在
		ok = true存在
	*/
	value, ok := data["name"]
	fmt.Println(value, ok)

	// map循环
	for key, val := range data {
		fmt.Println(key, val)
	}
}
