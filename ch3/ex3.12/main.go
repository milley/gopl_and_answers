package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input 2 strings")
		os.Exit(1)
	}

	s1, s2 := os.Args[1], os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Println("Are anagrams")
	} else {
		fmt.Println("Are not anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)

	if s1 == s2 {
		return true
	}
	return false
}

// SortString 字符串排序
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
