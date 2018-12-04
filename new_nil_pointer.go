package main

import "fmt"

func main() {
	var p *int
	q := new(float64)
	p = new(int)
	*p = 666
	fmt.Println(p, q)
}
