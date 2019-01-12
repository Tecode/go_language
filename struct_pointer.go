package main

import "fmt"

type Student2 struct {
	age int
	name string
	class int
}

func main() {
	// 顺序初始化，每一个值都要初始化
	var student *Student2 = &Student2{12, "bob", 5}
	fmt.Println((*student).name)

	// 指定成员初始化，没有初始化的自动赋值为0
	s1 := &Student2{name:"Lina"}
	fmt.Println(*s1)
	// new声明一块内存
	pointer := new(Student2)
	pointer.age = 22
	pointer.class = 6
	pointer.name = "Goo"
	fmt.Println(pointer)
}
