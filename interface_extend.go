package main

import "fmt"

type Human interface {
	song()
}

type Person interface {
	Human
	say(lrc string)
}

type Student struct {
	name string
}

func (s *Student) song() {
	fmt.Printf("%s在唱歌", s.name)
}

func (s *Student) say(lrc string) {
	fmt.Printf("歌词是%s", lrc)
}

func main() {
	var h Person
	h = &Student{name: "Aming"}
	h.song()
	fmt.Println()
	h.say("Go")
}
