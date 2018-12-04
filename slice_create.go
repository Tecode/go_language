package main

import "fmt"

func main() {
	// 数组
	array := [4]int {4, 5, 6, 7}
	// 切片
	slice := []int{}
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Printf("%d \n", len(slice))
	fmt.Printf("%d \n", cap(slice))

	slice = append(slice, 5)
	slice = append(slice, 7)
	fmt.Println(slice)
	fmt.Printf("%d \n", len(slice))
	fmt.Printf("%d \n", cap(slice))

	// make
	// 5长度，9容量
	// 不传9表示容量和长度相等
	make_slice := make([]int, 5, 9)
	fmt.Println(make_slice)
	fmt.Printf("%d \n", len(make_slice))
	fmt.Printf("%d \n", cap(make_slice))
}
