package main

import "fmt"

type Student struct {
	age int
	name string
	class int
}


func main() {
	student1 := Student{name:"Elic", age: 45}
	student2 := Student{age:12, class: 5}
	student3 := Student{age:45, name:"Elic"}

	fmt.Println(student1 == student2)
	fmt.Println(student1 == student3)

	student2 = student3

	fmt.Println(student2)
}
