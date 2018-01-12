// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaFloat(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func commaFloat(s string) string {
	array := []byte(s)
	pointPos := strings.Index(s, ".")
	symbolPos := strings.IndexAny(s, "+-")
	fmt.Println("point:", pointPos)
	fmt.Println("symbol:", symbolPos)
	fmt.Println(comma(string(array[symbolPos+1 : pointPos-1])))

	if pointPos < 0 && symbolPos < 0 {
		return comma(s)
	} else if pointPos < 0 && symbolPos >= 0 {
		return string(array[symbolPos]) + comma(string(array[symbolPos+1:]))
	} else if pointPos >= 0 && symbolPos < 0 {
		return comma(string(array[:pointPos])) + string(array[pointPos]) + comma(string(array[pointPos+1:]))
	}
	return string(array[symbolPos]) + comma(string(array[symbolPos+1:pointPos])) + string(array[pointPos]) + comma(string(array[pointPos+1:]))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	residue := n % 3
	for i := 0; i < n; i++ {
		if i < residue {
			buf.WriteByte(s[i])
			continue
		}
		if (i-residue)%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}

//!-
