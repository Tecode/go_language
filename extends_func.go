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

func main() {
	student := Student{Person{"Ming", "10021451", 15},"成都"}
	student.printInfo()
}
