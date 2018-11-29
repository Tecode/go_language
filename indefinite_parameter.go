package main

import "fmt"

func arguments2 (args ...int) {
	fmt.Println(args)
}

func arguments(args ...int)  {
	arguments2(args...)
	arguments2(args[:2]...)
	fmt.Println(args)
}

func main() {
	arguments(1, 2, 5)
}
