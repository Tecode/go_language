package main

import "fmt"

func main() {
	var list01 [2]string
	var array = []int{1, 5, 6, 9}
	var user = []string{"45", "457"}
	var listValue [100]int
	fmt.Println(user)
	fmt.Println(listValue, "____>>")

	list01[0] = "hello"
	list01[1] = "world"
	for _, value := range list01 {
		//fmt.Println(index, value)
		fmt.Printf("%s \n", value)
	}

	list02 := []int{4, 2, 6, 8, 7}

	fmt.Println(len(list01))
	fmt.Println(list02)
	fmt.Println(array)
}
