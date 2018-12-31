package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		var result []string

		input := bufio.NewScanner(os.Stdin)
		seen := make(map[string]string)

		for input.Scan() {
			line := input.Text()
			if line != "" {
				seen[line] = line
			}
		}

		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "indistinct: %v\n", err)
			os.Exit(1)
		}

		if seen != nil {
			result = append(result, "(")
			index := 0
			for _, v := range seen {
				index = index + 1
				result = append(result, "'")
				result = append(result, v)
				result = append(result, "'")
				if index != len(seen) {
					result = append(result, ",")
				}
			}
			result = append(result, ")")
		}

		for _, v := range result {
			fmt.Printf("%s", v)
		}
		fmt.Println()
	}

}
