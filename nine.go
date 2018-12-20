package main

import "fmt"

func main() {
	for index:=1; index < 10; index++ {
		for index2 :=1; index2 <= index; index2++ {
			fmt.Printf("%d * %d = %d ", index, index2, index*index2 )
		}
		fmt.Println()
	}
}
