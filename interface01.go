package main

import "fmt"

type Human interface {
	sayHi()
}

type Teacher struct {
	name  string
	group int
}

type Student struct {
	name   string
	idCard int32
}

func (student *Student) sayHi() {
	fmt.Printf("Student %s", student.name)
}

func (teacher *Teacher) sayHi() {
	fmt.Printf("Teacher %s", teacher.name)
}

func main() {
	// 定义接口类型变量
	var person Human

	// 只实现此接口方法的类型，那么类型的变量就可以给person赋值
	student := Student{name:"Ming"}
	person = &student
	person.sayHi()

	fmt.Println()

	teacher := &Teacher{name:"Li"}
	person = teacher
	person.sayHi()

}
