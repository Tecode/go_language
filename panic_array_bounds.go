package main

import "fmt"

func test()  {
	fmt.Println("4545")
}

func bounds(x int)  {
	var list [10]int
	// 当x20,数组越界会导致panic发生
	list[x] = x
}

func main() {
	bounds(20)
	test()
}
