package main

import "fmt"

// 不同作用域的同名变量
// 就近原则，会使用离调用最近的变量类型

var a byte

func main() {
	fmt.Printf("%T \n", a)
	var  a int
	fmt.Printf("%T \n", a)
	{
		var a float64
		fmt.Printf("%T \n", a)
	}
}
