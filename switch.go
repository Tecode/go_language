package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch name:=runtime.GOOS; name {
	case "bob":
		fmt.Println("I'm bob")
	default:
		println(name)
	
	}
}
