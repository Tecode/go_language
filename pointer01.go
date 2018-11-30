package main

import "fmt"

func main() {
	var a int = 10
	var c *int
	p := &a
	c = p
	fmt.Println(a)
	fmt.Println(p)
	*p =12
	fmt.Println(*p, a, c, *c)

}
