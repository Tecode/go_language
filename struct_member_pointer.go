package main

import "fmt"

type Student struct {
	age int
	name string
	class int
}

func main() {
	// 顺序初始化，每一个值都要初始化
	var student Student
	p := &student
	p.name = "Mola"
	fmt.Println(student)
}
