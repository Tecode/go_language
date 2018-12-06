package main

import "fmt"

type Student struct {
	age int
	name string
	class int
}

// 作为指针可以修改结构体的数据
func change(student *Student) {
	student.name = "666"
}

func main() {
	student1 := Student{name:"Elic", age: 45}
	change(&student1)
	fmt.Println(student1)
}
