package main

import "fmt"

type Student struct {
	idCard int64
	name string
}

func (s Student) seyName(item Student)  {
	fmt.Println(item.name, s.name)
}

func (s *Student) initMethod(idCard int64, name string)  {
	s.name = name
	s.idCard = idCard
}

func main() {
	// 方法继承
	var student Student = Student{10024, "Tecode"}
	student.seyName(student)

	// 指针方法继承，修改内部变量
	var student2 *Student = new(Student)
	student2.initMethod(512111,"Room")
	fmt.Println(*student2)


}
