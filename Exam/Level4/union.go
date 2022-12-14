package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("")
	}

	str := os.Args[1] + os.Args[2]
	str = rmDouble(str)
	fmt.Println(str)
}

func contain(r rune, s string) bool {
	for _, c := range s {
		if r == c {
			return true
		}
	}
	return false
}

func rmDouble(s string) string {
	result := ""
	for _, c := range s {
		if !contain(c, result) {
			result += string(c)
		}
	}
	return result
}
