package main

import "fmt"

type Person4 struct {
	name   string
	idCard string
	age    int
}

func (person *Person4) printInfo() {
	fmt.Printf("name=%s, idCard=%s, age=%d\n", person.name, person.idCard, person.age)
}

type Student4 struct {
	Person4
	address string
}

// 方法重写
func (student *Student4) printInfo() {
	fmt.Println(*student)
}

func main() {
	student := Student4{Person4{"Ming", "10021451", 15}, "成都"}
	// 就近原则，如果Student有这个方法就使用Student里面的方法,没有就找父级的方法
	student.printInfo()
	// 显示调用Person的方法
	student.Person4.printInfo()
	// 方法值
	studentFunc := student.printInfo
	studentFunc()
}
