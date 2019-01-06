package main

import "fmt"

// 同名字段name

type Person struct {
	idCart int64
	addr string
	name string
}

type Student struct {
	Person
	grade int
	name string
}

func main() {
	var student Student

	student.name = "Aming"
	student.addr = "四川成都"
	// 就近原则，没有找到会找继承的字段
	fmt.Printf("%+v", student)
}
