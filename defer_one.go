package main

import "fmt"

func main() {
	// defer会在main函数关闭以后调用
	defer fmt.Println("1")
	 fmt.Println("2")
}
