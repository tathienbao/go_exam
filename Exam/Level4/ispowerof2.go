package main

import "fmt"

func main() {
	fmt.Println(IsPowerOf2(8))
	fmt.Println(IsPowerOf2(9))
	fmt.Println(IsPowerOf2(6))
	fmt.Println(IsPowerOf2(16))
	fmt.Println(IsPowerOf2(-8))
}

func IsPowerOf2(n int) bool {
	return n > 0 && ((n & (n - 1)) == 0)
}
