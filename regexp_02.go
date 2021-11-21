package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a", "5a12")
	fmt.Println(match)
	// +匹配前一个字符一次或多次
	// 解析正则表达式，如果成功会返回接收器
	reg := regexp.MustCompile(`a[\d | [:alpha:]]+c`)
	buf := "abc ace acc a5c a2c movie"
	// 根据规则信息提取关键信息，返回一个切片
	fmt.Println(reg.FindAllString(buf, -1))

}
