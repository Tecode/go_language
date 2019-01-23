package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Name    string `json:"name"` // json:"name"设置别名
	Subject []string `json:"subject"`
	IsOk    bool `json:"isOk"`
}

// 解析部分需要的字段
type It2 struct {
	Name    string `json:"name"` // json:"name"设置别名
	Subject []string `json:"subject"`
}

func main() {
	data := `{
 "name": "Go",
 "subject": [
  "C++",
  "Python",
  "Go",
  "Java",
  ""
 ],
 "isOk": true
}`
	// 定义一个结构体变量
	var tmp It
	// 第二个参数要地址传递，否则没有数据
	err := json.Unmarshal([]byte(data), &tmp)
	fmt.Println(tmp, err)
	fmt.Println(tmp.Subject)

	// 解析部分需要的字段
	var tem2 It2
	err2 := json.Unmarshal([]byte(data), &tem2)
	fmt.Printf("%+v", tem2, err2)
}
