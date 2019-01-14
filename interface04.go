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

// 定义一个普通函数，函数的类型为接口类型
// 只要一个函数，可以有不同表现
func whoSayHi(i Human) {
	i.sayHi()
}

func (teacher *Teacher) sayHi() {
	fmt.Printf("Teacher %s", teacher.name)
}

func main() {

	// 只实现此接口方法的类型，那么类型的变量就可以给person赋值
	student := Student{name: "Ming"}
	teacher := &Teacher{name: "Li"}
	// 调同一个函数，实现不同的方法
	whoSayHi(&student)
	fmt.Println()
	whoSayHi(teacher)

	// 创建一个切片
	methods := make([]Human, 2)
	methods[0] = &student
	methods[1] = teacher

	fmt.Println()
	for _, i := range methods {
		i.sayHi()
	}

}
