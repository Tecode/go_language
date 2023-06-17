package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Name    string   `json:"name"`    // json:"name"设置别名
	Subject []string `json:"-"`       //二次编码 - 隐藏改字段
	IsOk    bool     `json:",string"` // 转换成字符串
}

func main() {
	// 定义一个结构体变量
	it := It{Name: "Go", Subject: []string{"C++", "Python", "Go", "Java", ""}, IsOk: true}
	//	编码，更加内容生成json文本
	buf, err := json.Marshal(it)
	// MarshalIndent格式化，让代码能够有缩进
	//buf, err := json.MarshalIndent(it, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
