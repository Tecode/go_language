package main

import "fmt"

type Person struct {
	idCard int32
	name string
}

type Student struct {
	Person
	grade int
}

func main() {
	var student Student  = Student{Person{5102445, "Aming"}, 6}
	fmt.Println(student)
	// %+v显示更详细的结果
	fmt.Printf("%+v", student)
}
