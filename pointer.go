package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var a = 10
	var b *int
	b = &a
	rand.Seed(time.Now().UnixNano())
	num := math.Max(float64(a), 65.0)
	for index := 0; index < a; index++ {
		fmt.Println(time.Now().UnixNano())
	}
	fmt.Println(a, *b, num)
}
