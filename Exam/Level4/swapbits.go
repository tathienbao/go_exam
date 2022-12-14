package main

import (
	"fmt"
)

func SwapBits(octet byte) byte {
	return octet>>4 | octet<<4
}

func main() {
	var n byte = 92
	fmt.Printf("%08b\r\n", n)
	fmt.Printf("%08b\r\n", SwapBits(n))
}
