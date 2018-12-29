package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func randRemove(slice []int, i int) []int {
	slice[i] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(randRemove(s1, 2))
}
