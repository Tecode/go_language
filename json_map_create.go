package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// interface{}表示所有类型
	m := make(map[string]interface{})
	m["name"] = "Am"
	m["subject"] = []string{"java", "node", "go"}
	m["age"] = 12

	value, _ := json.Marshal(m)
	fmt.Println(string(value))
}
