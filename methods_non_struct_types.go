package main

import "fmt"

type MyFloat float64

func (f MyFloat) caculation(value float64) float64{
	return float64(f) + value
}

func main() {
	obj := MyFloat(8.32)
	fmt.Println(obj.caculation(98.32))
}
