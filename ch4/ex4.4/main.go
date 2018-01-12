package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	rotate(s, 4)
	fmt.Println(s)
}

func rotate(array []int, index int) {
	length := len(array)
	tmp := make([]int, index)
	copy(tmp, array[:index])
	copy(array, array[index:])
	copy(array[length-index:], tmp)
}
