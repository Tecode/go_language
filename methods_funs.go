package main

import "fmt"

type Animals02 struct {
	name string
	age int
}

func sayHello(animal Animals02) {
	fmt.Printf("我叫%s,我%d岁了", animal.name, animal.age)
}

func main() {
	animal := Animals02{"Lilio", 4}
	sayHello(animal)
}
