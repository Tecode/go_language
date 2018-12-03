package main

import "fmt"

func main() {
	var array = []int {5, 8, 9, 8}
	var data[] int
	copy(array[:3], data)
	fmt.Println("cap = ",cap(data))
	data = array[:3]
	data[0] = 10
	fmt.Println(data)
	fmt.Println(array)
}
