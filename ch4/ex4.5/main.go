package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "one", "two", "two", "three", "three", "three"}
	fmt.Printf("%q\n", delRepetitive(data))
	fmt.Printf("%q\n", data)
}

func delRepetitive(s []string) []string {
	length := len(s)
	skiped := 0
	for i := 0; i < length-skiped; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			skiped++
		}
	}
	return s[:length-skiped]
}
