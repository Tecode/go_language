package main

import "fmt"

func main() {
	const (
		a0 = iota
		b0 = iota
		c0 = iota
	)
	const (
		a1 = 1 << iota
		b1 = 1 << iota
		c1 = 1 << iota
	)
	var a2, c2, b2 int
	fmt.Println(a0, b0, c0)
	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
}

/*
output:
0 1 2
1 2 4
*/
