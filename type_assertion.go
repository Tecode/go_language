package main

import "fmt"

type Student struct {
	name string
}

func main() {
	var i []interface{} = make([]interface{}, 3)
	i[0] = 5
	i[1] = "string"
	i[2] = Student{"Ming"}

	for _, data := range i {
		// value,返回的是说明类型int,string
		switch value := data.(type) {
		case int:
			fmt.Println("int", value)
			break
		case string:
			fmt.Println("string", value)
			break
		case Student:
			fmt.Println("Student", value)
			break
		}
	}
	// 第一个是值，第二个是bool值
	if value, ok := i[0].(int); ok == true {
		println(value, ok)
	}

}
