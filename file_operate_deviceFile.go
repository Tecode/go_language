package main

import (
	"fmt"
	"os"
)

//设备文件的使用
func main() {
	//os.Stdout.Close()         //	关闭后无法输出
	//fmt.Println("Are you ok") // 往标准输出设备写内容
	//	标准设备文件（os.Stdout）,默认已经打开，用户可以直接使用
	os.Stdout.WriteString("Are you ok \n")


	os.Stdin.Close() // 关闭输入，无法输入内容
	var value int
	fmt.Println("请输入...")
	fmt.Scan(&value)
	fmt.Println("输入的是", value)

}
