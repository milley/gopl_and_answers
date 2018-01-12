package main

import (
	"fmt"
)

func main() {
	b := []byte("Hello,世界!")
	fmt.Printf("%q\n", string(b))
	fmt.Printf("%q\n", reverseUtf8Str(string(b)))
}

func reverseUtf8(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func reverseUtf8Str(str string) string {
	res := make([]byte, len(str))
	prevPos, resPos := 0, len(str)
	for pos := range str {
		resPos -= pos - prevPos
		copy(res[resPos:], str[prevPos:pos])
		prevPos = pos
	}

	copy(res[0:], str[prevPos:])
	return string(res)
}
