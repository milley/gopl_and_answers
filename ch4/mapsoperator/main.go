package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	y := map[string]int{
		"c": 3,
		"b": 2,
		"a": 1,
	}

	fmt.Printf("%v", equal(x, y))
}
