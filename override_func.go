package main

import "fmt"

type Person struct {
	name   string
	idCard string
	age    int
}

func (person *Person) printInfo() {
	fmt.Printf("name=%s, idCard=%s, age=%d", person.name, person.idCard, person.age)
}

type Student struct {
	Person
	address string
}

// 方法重写
func (student *Student) printInfo()  {
	fmt.Println(*student)
}

func main() {
	student := Student{Person{"Ming", "10021451", 15},"成都"}
	// 就近原则，如果Student有这个方法就使用Student里面的方法,没有就找父级的方法
	student.printInfo()
	// 显示调用Person的方法
	student.Person.printInfo()
}
