package main

import "fmt"

type Animal struct {
	name string
	age int
}

func main() {
	animal := Animal{"cat", 2}
	p := &animal
	p.age = 1.0e5
	fmt.Println(animal)
}
