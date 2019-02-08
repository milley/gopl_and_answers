package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "afoob"
	fmt.Println(expand(s, strings.ToUpper))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}
