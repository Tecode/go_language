package main

import "fmt"

// 非结构体匿名字段

type name string

type Person struct {
	idCart int64
	addr string
}

type Student struct {
	Person
	name
	int
}

func main() {
	var student Student = Student{Person{5102114, "重庆"}, "Tecode", 6}

	// 就近原则，没有找到会找继承的字段
	fmt.Printf("%+v \n", student)
	fmt.Println(student.name, student.int)
}
