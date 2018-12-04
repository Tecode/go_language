package main

import "fmt"

func main() {
	// 支持比较，支持== 或 !=，比较是不是每一个元素都一样
	a := [3]int{4, 5, 6}
	b := [3]int{1, 5, 4}
	c := [3]int{4, 5, 6}
	fmt.Println(a == b)
	fmt.Println(b == c)
	fmt.Println(a == c)
	// 同类型可以赋值
	a = b
	fmt.Println(a)
}
