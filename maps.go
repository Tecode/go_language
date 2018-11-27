package main

import "fmt"

type Vertex struct {
	x, y, z float32
}


func main() {
	var m = make(map[string]Vertex)
	m["value1"] = Vertex{14.2354, 25.33154, -1.02147}
	m["value2"] = Vertex{-36.221, 36.2214, -741.3654}
	fmt.Printf("%v", m)
}
