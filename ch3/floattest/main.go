package main

import "fmt"

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)
}
