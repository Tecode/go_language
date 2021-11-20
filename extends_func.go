package main

import "fmt"

type Person1 struct {
	name   string
	idCard string
	age    int
}

func (person *Person1) printInfo() {
	fmt.Printf("name=%s, idCard=%s, age=%d\n", person.name, person.idCard, person.age)
}

func (person *Person1) changeVale(value string) {
	person.name = value
}

type Student1 struct {
	Person1
	address string
}

func main() {
	student := Student1{Person1{"haoxuan", "10021451", 15}, "成都"}
	student.printInfo()
	student.changeVale("HAO_XUAN")
	fmt.Println(student)
}
