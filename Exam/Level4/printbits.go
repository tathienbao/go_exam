package main

import "fmt"

func main() {
	var n byte = 192
	PrintBits(n)
	fmt.Println("")
	n = 11
	PrintBits(n)
}

func PrintBits(octet byte) {
	for i := 7; i >= 0; i-- {
		if octet&(1<<uint(i)) != 0 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
