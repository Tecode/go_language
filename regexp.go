package main

import (
	"fmt"
	"regexp"
)

func main() {
	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	match, _ := regexp.MatchString("user/login", "v1/user/login")
	fmt.Println(match)
}
