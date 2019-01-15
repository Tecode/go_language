package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 转换为字符串后追加到数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	// 第二个为要追加的数，第三个数为指定十进制
	slice = strconv.AppendInt(slice, 12424, 10)

	fmt.Println(string(slice))

	//	其它类型转换为字符串
	str := strconv.FormatBool(false)
	fmt.Println(str)

	// 字符串转bool
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(boolean)
	}


}
