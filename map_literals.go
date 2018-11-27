package main

import "fmt"

type Vertex2 struct {
	x, y float32
}

var m  = map[string]Vertex2 {
	"value1": {14.2235, -1.0024},
	"value2": {14.2235, -1.0024},
}

func main() {
	fmt.Printf("%v", m)
}
