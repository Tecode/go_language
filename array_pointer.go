package main

import "fmt"

func output_array(list []int)  {
	for _,value := range list{
		fmt.Println(value)
	}
}

func output_pointer(p *[]int)  {
	for _,value := range *p{
		fmt.Println(value)
	}
}

func main() {
	array := []int{8, 21, 55, 25, 36}
	output_array(array)
	output_pointer(&array)
}
