package main

import "fmt"

type Animals struct {
	name string
	age int
}

func (animal Animals) sayHello() {
	fmt.Printf("我叫%s,我%d岁了", animal.name, animal.age)
}

func main() {
	animal := Animals{"Lili", 4}
	animal.sayHello()
}
