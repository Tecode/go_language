package main

import "fmt"

func sum(s []int, c chan int)  {
	sum := 0
	for _, value := range s {
		sum += value
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 74, 8, -9, 4, 564}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)


	a, b := <-c, <-c
	fmt.Println(a + b)
}