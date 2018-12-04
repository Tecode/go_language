package main

import "fmt"

func main() {
	var p *int
// error-> panic: runtime error: invalid memory address or nil pointer dereference
// 不合法的内存地址
//	*p = 4
	fmt.Println(p)
// output: nil
}
