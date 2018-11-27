package main

import "fmt"

func main() {
	list := [][]string{
		[]string{"java", "Go"},
		[]string{"node", "php"},
	}
	fmt.Println(list)
	fmt.Println(list[0][0])
}
