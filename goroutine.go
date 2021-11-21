package main

import (
	"fmt"
	"time"
)

func add002(a, b int, c chan int) {
	num := a + b
	c <- num
}

func main() {
	list := []int{1, 5, 6, 8}
	c := make(chan int)
	for _, value := range list {
		fmt.Println(value)
		go add002(value, value, c)
	}

	a := <-c

	select {
	case <-c:
		fmt.Println("channel成功")
	case c <- 1:
		fmt.Println("写入数据成功")
	default:
		fmt.Println("进入处理流程")
	}
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
