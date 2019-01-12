package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}

func main() {
	var p1 *Student = &Student{25,"Aming",22}
	fmt.Println(p1)
}
