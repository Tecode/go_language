package main

import "fmt"

// 指针类型
type Student struct {
	name string
	idCard string
}

// 值类型
func (student Student) output()  {
	fmt.Println(student.idCard)
}

// 指针类型
func (student *Student) pointerValue()  {
	fmt.Println(student.name)
}

// 指针变量的方法集,不管是指针还是普通变量都会自动转换为实参
func main() {
	var student Student  = Student{"Ming","510125998745"}
	student.output()
	student.pointerValue()

	var student2 *Student = &Student{"ELEN", "521144"}
	student2.pointerValue()
	student2.output()
}
