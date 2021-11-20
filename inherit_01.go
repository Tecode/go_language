package main

import "fmt"

type Person2 struct {
	idCard int32
	name   string
}

type Student2 struct {
	Person2
	grade int
}

func main() {
	var student = Student2{Person2{5102445, "HaoXuan"}, 6}
	fmt.Println(student)
	// %+v显示更详细的结果
	fmt.Printf("%+v \n", student)

	fmt.Println(student.idCard)
}
