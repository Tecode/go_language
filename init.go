package main

import "fmt"

var Pi float64

// 优先级比main高
func init()  {
	Pi = 3.14
}

func main() {
	fmt.Println(Pi)
}
