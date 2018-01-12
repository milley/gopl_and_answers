package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("abc\r  \n\rdef")
	fmt.Printf("%q\n", string(trimDupSpace(b)))
	fmt.Printf("%q\n", b)
}

func trimDupSpace(arr []byte) []byte {
	var rtnArr []byte
	for i, v := range arr {
		if unicode.IsSpace(rune(v)) {
			if i > 0 && unicode.IsSpace(rune(arr[i-1])) {
				continue
			} else {
				rtnArr = append(rtnArr, ' ')
			}
		} else {
			rtnArr = append(rtnArr, v)
		}
	}

	return rtnArr
}
