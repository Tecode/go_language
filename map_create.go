package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	id int
	name string
}

func main() {
	maps := map[string]string{"name":"mary","age":"12"}
	fmt.Println(maps)
	data, _ := json.Marshal(maps)
	fmt.Println(string(data))
	// 需要传每一个参数
	var student = Student{5, "Ming"}
	fmt.Println(student)
	// 只需要传指定的参数
	student2 := Student{name:"Aming"}
	fmt.Println(student2)
}
