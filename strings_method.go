package main

import (
	"fmt"
	"strings"
)

func main() {
	// contains 是否包含指定的字符串返回true和false
	fmt.Println(strings.Contains("Hello", "l"))
	fmt.Println(strings.Contains("Hello", "adx"))

	// join 将list按指定的分隔符连接起来
	list := []string{"A", "Too", "Pcc", "QUE"}
	buf := strings.Join(list, ",")
	fmt.Println(buf)

	//	index
	fmt.Println(strings.Index("Hello World", "o"))

	// repeat output:gogogogogo
	buf = strings.Repeat("go", 5)
	fmt.Println(buf)

	// split 按指定字符拆分，返回一个切片
	fmt.Println(strings.Split("go php c++", " "))

	// trim 去掉两头空格
	buf = strings.Trim("   are you   ok     ", " ")
	fmt.Println(buf)

	// Fields会将字符串按空格拆分成切片
	string03 := strings.Fields("  are    you    ok    ")
	//fmt.Println(string03)
	for index, value := range string03 {
		fmt.Println(value, index)
	}

}
