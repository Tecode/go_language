package main

import "fmt"

func main() {
	// 空接口是万能类型，保存任何类型
	var i interface{}  = 1

	fmt.Println(i)

	i = "544"
	fmt.Println(i.(string))


}
