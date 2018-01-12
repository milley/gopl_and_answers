// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

var pc [8]byte

func init() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	fmt.Println("diff bits is:", countDiffBits(c1, c2))
}

func countDiffBits(s1, s2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		byte1 := s1[i]
		byte2 := s2[i]
		for j := 0; j < 8; j++ {
			if byte1&pc[j] != byte2&pc[j] {
				count++
			}
		}
	}
	return count
}

//!-
