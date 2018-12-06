package main

import "fmt"

func deleteVal(maps map[string]int)  {
	delete(maps, "k")
}

func main() {
	maps := map[string]int{"k":12, "y": 45, "Lua": 92}
	deleteVal(maps)

	for value,_ := range maps {
		fmt.Println(value)
	}
}