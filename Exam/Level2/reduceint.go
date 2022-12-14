package main

import (
	"fmt"
	"strconv"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2, 3, 4}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 0; i < len(a)-1; i++ {
		result = f(result, a[i+1])
	}
	fmt.Println(strconv.Itoa(result))
}
