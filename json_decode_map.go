package main

import (
	"encoding/json"
	"fmt"
)

type It struct {
	Name    string   `json:"name"` // json:"name"设置别名
	Subject []string `json:"subject"`
	IsOk    bool     `json:"isOk"`
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
	// 创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(data), &m)
	fmt.Println(m, err)
	// 类型断言
	for key, value := range m {
		switch data := value.(type) {
		case string:
			fmt.Println("数据类型是Name", key, data)
		case []interface{}:
			fmt.Println("数据类型是[]interface{}", key, data)
		}
	}
}
