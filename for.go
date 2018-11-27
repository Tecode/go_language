package main

import "fmt"

func main() {
	list := []int{1, 2, 4, 8, 16, 32}
	for index, value := range list {
		fmt.Println(index * value)
	}
}
