package main

import "fmt"

func test04() {
	fmt.Println("aaaa")
}

func test02() {
	fmt.Println("bbbb")
	panic("程序中断")
}

func test03() {
	fmt.Println("cccc")
}

func main() {
	test04()
	test02()
	test03()
}
