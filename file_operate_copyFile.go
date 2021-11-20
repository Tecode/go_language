package main

import (
	"fmt"
	"io"
	"os"
)

//文件的读写
func main() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("参数错误：xxx.go srcFile dstFile")
		return
	}
	fmt.Println(list)
	srcFile := list[1]
	dstFile := list[2]

	if srcFile == dstFile {
		fmt.Println("文件名不能重复")
		return
	}

	sf, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	df, err := os.Create(dstFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 2*1024)

	for {
		index, err := sf.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err, "--------")
		}
		df.Write(buffer[:index])
	}

}
