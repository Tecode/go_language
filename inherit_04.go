package main

import "fmt"

// 指针类型

type name string

type Person struct {
	idCart int64
	addr string
}

type Student struct {
	*Person
	name
	int
}

func main() {
	var student Student = Student{&Person{5102114, "重庆"}, "Tecode", 6}

	fmt.Printf("%+v \n", student)
	fmt.Println(student.name, student.int, student.idCart)

	// 分配空间
	var s Student
	s.Person = new(Person)
	s.idCart = 41110
	s.int = 98
	s.name = "Elent"
	// 打印出来是地址
	fmt.Println(s)
	fmt.Println(s.name, s.int, s.idCart)

}
