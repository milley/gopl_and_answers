package main

import (
	"fmt"
)

// Currency int
type Currency int

// The kind of currency
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int // array of 3 integers
	fmt.Println(a[0])

	var q = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	for i, v := range q {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println(r[2])

	t := [...]int{1, 2, 3, 4}
	fmt.Println(len(t))

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	for i, v := range symbol {
		fmt.Printf("%d : %s\n", i, v)
	}
}
