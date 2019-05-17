package main

import "fmt"

type Info struct {
	name string
	age int
	id int
}

func main() {
	info := Info{age:12,id:51,name:"Tony"}
	fmt.Println(info)
	list := make([]Info, 2)
	list[0] = Info{age:12,id:56,name:"Tony"}
	list[1] = Info{age:10,id:51,name:"Ming"}

	fmt.Println(list)
}
