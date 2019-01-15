package main

import "fmt"

func test() {
	fmt.Println("4545")
}

func bounds(x int) {
	// 设置recover
	defer func() {
		// recover()可以打印panic的错误信息
		err := recover()
		if err != nil {
			//	产生panic错误
			fmt.Println(err)
		}
	}()
	var list [10]int
	// 当x20,数组越界会导致panic发生
	list[x] = x
}

func main() {
	bounds(2)
	test()
}
