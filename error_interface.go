package main

import (
	"errors"
	"fmt"
)

func main() {
	// 错误信息第一种方式
	err1 := fmt.Errorf("错误信息1")
	fmt.Println(err1)

	//	错误信息抛出第二种方式
	err2 := errors.New("错误信息2")
	fmt.Println(err2)
}
