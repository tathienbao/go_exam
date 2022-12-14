package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64 = 255
	fmt.Println(printHex(n))
	n = 5156454
	fmt.Println(printHex(n))
}

func printHex(n int64) string {
	return strconv.FormatInt(n, 16)
}
