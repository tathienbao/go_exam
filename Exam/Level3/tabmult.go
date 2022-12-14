package main

import "fmt"

func main() {
	tabmult(9)
}

func tabmult(n int) {
	for i := 1; i < 10; i++ {
		fmt.Print(i)
		fmt.Print(" x ")
		fmt.Print(Itoa(n))
		fmt.Print(" = ")
		fmt.Print(Itoa(n * i))
		fmt.Print("\r\n")
	}
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""
	for n > 0 {
		res = string(rune(n%10+48)) + res
		n /= 10
	}

	if n < 0 {
		return "-" + Itoa(-n)
	} else {
		return res
	}
}
