package main

import "fmt"

func main() {
	for i := 9; i >= 0; i-- {
		fmt.Print(rune(i))
	}
	fmt.Print("\n")
}
