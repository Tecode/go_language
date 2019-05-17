package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	fmt.Println(add(5, 4))
	fmt.Println(add([]int{5, 80, 9, 8, 5, 8, 9}...))
}
