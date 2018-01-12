package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
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
