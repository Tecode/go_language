package main

import "fmt"

func main() {
	arr := []int{1, 3, 7, 6, 4}
	// 1：开始下标 3结束下标，不包含arr[3] 5-1容量 长度3-1
	s := arr[1:3:5]
	fmt.Println(s)
	fmt.Printf("%d \n\n", len(s))
	fmt.Printf("%d \n\n", cap(s))
}
