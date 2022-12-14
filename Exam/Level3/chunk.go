package main

import "fmt"

func Chunk(slice []int, size int) {
	switch true {
	case size <= 0:
		fmt.Println("")
	case len(slice) == 0:
		fmt.Println(slice)
	case size > 0:
		a := make([][]int, 0, size)
		for size < len(slice) {
			a = append(a, slice[:size])
			slice = slice[size:]
		}
		fmt.Println(append(a, slice))
	}
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 4)
}
