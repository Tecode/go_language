package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32}
	for i, v:= range pow {
		fmt.Println(v, i)
	}
}
