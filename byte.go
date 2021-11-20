package main

import "fmt"

func main() {
	buffer := make([]byte, 16)
	for index := range buffer {
		buffer[index] = byte(index)
		//fmt.Println(buffer, index)
	}
	fmt.Println(buffer)
}
