package main

import (
	"errors"
	"fmt"
)

func division(numerator, denominator float32) (value float32, err error) {
	if denominator == 0 {
		err = errors.New("分母不能为0")
	} else {
		value = numerator / denominator
	}
	return
}

func main() {
	value, err := division(1, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
