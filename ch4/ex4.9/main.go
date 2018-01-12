package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	fileReader, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		counts[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\t\t\tfreq\n")
	for c, n := range counts {
		fmt.Printf("%q\t\t\t%d\n", c, n)
	}

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[:0])
}
