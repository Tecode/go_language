package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	array := make([]float32,4)
	fmt.Println(array)

	// 在原切片末尾添加元素，自动扩容
	slice = append(slice, 6)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 如果超出原来的容量会已2倍容量扩容
	slice2 := make([]int, 2)
	slice2 = append(slice2, 14)
	slice2 = append(slice2, 16)
	fmt.Println("len =",len(slice2))
	fmt.Println("cap",cap(slice2))
	slice2 = append(slice2, 18)
	fmt.Println("len =",len(slice2))
	fmt.Println("cap",cap(slice2))

}
