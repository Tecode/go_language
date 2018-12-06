package main

import "fmt"

func main() {
	maps := map[string]int{"k":12, "y": 45, "Lua": 92}
	delete(maps, "k")

	for value,_ := range maps {
		fmt.Println(value)
	}
	// 判断是存在k，ok bool
	value, ok := maps["k"]
	if ok {
		fmt.Println(value)
	}
}