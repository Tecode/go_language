package main

import (
	"fmt"
)

func myFunc() (a , b, c int) {
	a, b, c = 2, 35, 4
	return
}

func main() {
	func () {
		println("内部函数")
	}()
	a, b, c := myFunc()
	fmt.Println(a, b, c)
}
