package main

import "fmt"

func main() {
	a := make([]int,10)
	printSlice("a", a)

	b := make([]int, 6)
	printSlice("b", b)

	c := a[:3]
	printSlice("c", c)

	d := c[:cap(c)]
	printSlice("d", d)
}

func printSlice(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n", s,len(x),cap(x), x)
}