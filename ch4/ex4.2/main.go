package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	useSha256 = iota
	useSha384
	useSha512
)

func main() {
	var methodStr string
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&methodStr, "m", "256", "sha method")
	mySet.Parse(os.Args[1:])
	method := useSha256

	if len(os.Args) > 0 {
		switch methodStr {
		case "384":
			method = useSha384
		case "512":
			method = useSha512
		default:
			method = useSha256
		}
	}
	fmt.Println("Enter text:")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		signature(input.Text(), method)
	}

}

func signature(s string, method int) {
	if method == useSha384 {
		hash := fun384(s)
		printHash(hash[:])
	} else if method == useSha512 {
		hash := fun512(s)
		printHash(hash[:])
	} else {
		hash := fun256(s)
		printHash(hash[:])
	}

}

func fun256(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func fun384(s string) [48]byte {
	return sha512.Sum384([]byte(s))
}

func fun512(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}

func printHash(hash []byte) {
	for _, v := range hash {
		fmt.Printf("%X", v)
	}
	fmt.Println()
}
