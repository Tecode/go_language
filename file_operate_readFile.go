package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(path string) {
	//	打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建一个buffer切片为2M
	buffer := make([]byte, 1024*2)
	index, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return
	}
	fmt.Println("buffer= ", string(buffer[:index]))
}

// 每次读取一行
func readLine(path string) {
	//	打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 新建一个缓冲区，把内容先放到缓冲区
	read := bufio.NewReader(file)
	// 循环遍历
	for {
		buffer, err := read.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Println(string(buffer))
	}
}

//文件的读写
func main() {
	path := "./files/demo.txt"
	//readFile(path)
	readLine(path)
}
