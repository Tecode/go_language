package main

import (
	"fmt"
	"os"
)

func writeFile(path string) {
	// 创建文件，会清空文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("%d * %d = %d \n ", i, i, i*i)
		index, err := file.WriteString(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(index)
	}

}

//文件的读写
func main() {
	path := "./files/demo.txt"
	writeFile(path)
}
