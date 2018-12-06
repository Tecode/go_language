package main

import "fmt"

type Student struct {
	age int
	name string
	class int
}

func main() {
	// 顺序初始化，每一个值都要初始化
	var student Student = Student{12, "bob", 5}
	fmt.Println(student)

	// 指定成员初始化，没有初始化的自动赋值为0
	s1 := Student{name:"Lina"}
	fmt.Println(s1)
}
