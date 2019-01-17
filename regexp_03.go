package main

import (
	"fmt"
	"regexp"
)

func main() {
	// +匹配前一个字符一次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	buf := "3.12 3.211 3. fd 54.32 441.30 "
	fmt.Println(reg.FindAllString(buf, -1))
}
