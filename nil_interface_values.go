package main

import (
	"fmt"
)

type Face interface {
	look()
}

type Name string

func (name Name) look() {
	fmt.Println("Alise", name)
}

func main() {
	var name Name
	var face Face
	describe(face)
	face = name
	describe(face)
	face.look()
}

func describe(i Face) {
	fmt.Printf("(%v, %T)\n", i, i)
}
