package main

import (
	"io"
	"os"
	"strings"
)

type readString struct {
	s io.Reader
}

func main() {
	s := strings.NewReader("return string(bytes.Join(s, sep))")
	r := readString{s}
	io.Copy(os.Stdout, r.s)
}