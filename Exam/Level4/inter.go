package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Print('\n')
	} else {
		a := removeDuplicates(args[0])
		b := args[1]
		result := ""
		for _, x := range a {
			if contains(x, b) {
				result += string(x)
			}
		}

		for _, x := range result {
			fmt.Print(string(x))
		}
	}
	fmt.Print("\n")
}

func contains(r rune, s string) bool {
	for _, x := range s {
		if x == r {
			return true
		}
	}
	return false
}

func removeDuplicates(str string) string {
	var result string
	for _, x := range str {
		if !contains(x, result) {
			result += string(x)
		}
	}
	return result
}
