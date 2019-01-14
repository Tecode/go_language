package main

import "fmt"

type Student struct {
	name string
}

func (student *Student) printInfo()  {
	fmt.Println(student.name)
}

func (student Student) helloWorld()  {
	fmt.Println("hello world")
}

func main() {
	student := Student{"Ming"}

	pfunc := (*Student).printInfo
	pfunc(&student)

	hfunc := (Student).helloWorld
	hfunc(student)
}
