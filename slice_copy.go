package main

import "fmt"

func main() {
	src_slice := []int{1, 2, 5}
	dst_slice := []int{4, 9, 5, 19}
	copy(dst_slice,src_slice)
	fmt.Println(dst_slice)
	// output [1 2 5 19]
}
