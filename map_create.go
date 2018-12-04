package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	maps := map[string]string{"name":"mary","age":"12"}
	fmt.Println(maps)
	data, _ := json.Marshal(maps)
	fmt.Println(string(data))
}
