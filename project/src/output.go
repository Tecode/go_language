package main

import "fmt"

func output() {
	for index:=0; index < 10; index++ {
		for col:=1; col < index+1; col ++  {
			fmt.Printf("%d * %d = %d ", col, index, col*index)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%.15f", 0.1 + 0.2)
}