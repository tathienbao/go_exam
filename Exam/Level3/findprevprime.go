package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(10))
	fmt.Println(FindPrevPrime(20))
}

func FindPrevPrime(n int) int {

	switch {
	case n < 2:
		return 0
	case n == 2:
		return 2
	case n > 2:
		for i := 2; i <= n-1; i++ {
			if n%i == 0 {
				n -= 1
			}
		}
	}

	return n
}
