package main

import "fmt"

func swap(a, b *int)  {
	*a, *b = *b, *a
	return
}

func main() {
	a, b := 12, 57
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}
