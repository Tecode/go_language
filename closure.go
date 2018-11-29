package main

import "fmt"

func closure() func() int  {
	var x int
	return func() int {
		x++
		return  x*x
	}
}

func main() {
	f := closure()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
