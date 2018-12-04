package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for index :=0; index < 5; index ++ {
		// 产生很大的随机数
		// fmt.Println(rand.Int())
		// 限制在10以内
		fmt.Println(rand.Intn(10))
	}
}
