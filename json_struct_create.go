package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Name    string
	Subject []string
	IsOk    bool
}

func main() {
	// 定义一个结构体变量
	it := It{Name: "Go", Subject: []string{"C++", "Python", "Go", "Java", ""}, IsOk: true}
	//	编码，更加内容生成json文本
	buf, err := json.Marshal(it)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
