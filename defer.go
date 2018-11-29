package main

import "fmt"

// defer先进后出，例如以下会输出
//Dom
//Bbq
//Abs

func main() {
	defer fmt.Println("Abs")
	defer fmt.Println("Bbq")
	//defer fmt.Println(100/0)
	defer fmt.Println("Dom")
}
