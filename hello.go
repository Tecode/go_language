package main

import "fmt"

func main() {
	var list = []int{5}
	var a int
	for index := 1; index < 5; index++ {
		fmt.Print(list, a)
	}
}
