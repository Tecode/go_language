package main

import "fmt"

type Json struct {
	id int
	name string
}

func main() {
	var number int = 12
	fmt.Println(number)

	var list = []int{5, 9, 8}
	fmt.Println(list)

	var list2 []int
	list2 = append(list2, 69, 36)
	fmt.Println(list2)

	var json Json  = Json{4, "Mac"}
	fmt.Println(json, json.id)


}
