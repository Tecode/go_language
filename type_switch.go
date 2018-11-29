package main

import "fmt"


func do(i interface{})  {
	switch v := i.(type) {
	case int:
		fmt.Printf("整型变量，%d \n", v)
	case string:
		fmt.Printf("字符串%s \n", v)
	case bool:
		fmt.Printf("布尔类型%v \n", v)
	}
}

func main()  {
	do(50)
	do("Lone")
	do(false)
}
