package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

func reverse(array *[6]int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
