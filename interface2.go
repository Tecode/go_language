package main

import "fmt"

type FileMethod interface {
	readName()
	closeFile()
}

type File string

func main() {
	var fileMethod FileMethod
	f := File("西游记")
	fileMethod = &f
	f.readName()
	fileMethod.readName()
	fileMethod.closeFile()
}

func (file *File) readName() {
	// 指针可以修改内部数据
	//*file = "红楼梦"
	fmt.Println(*file)
}

func (file File) closeFile() {
	fmt.Println(file,"文件已关闭")
}