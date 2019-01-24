package main

import (
	"fmt"
	"os"
)

func writeFile(path string)  {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var buf string
	for i := 0; i < 20; i++ {
		fmt.Sprint()
	}

}
//文件的读写
func main() {
	path := "./files/demo.txt"
	writeFile(path)
}
