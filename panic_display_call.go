package main

import "fmt"

func test01()  {
	fmt.Println("aaaa")
}

func test02()  {
	fmt.Println("bbbb")
	panic("程序中断")
}

func test03()  {
	fmt.Println("cccc")
}

func main() {
	test01()
	test02()
	test03()
}
