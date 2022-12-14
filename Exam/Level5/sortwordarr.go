package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	fmt.Println(Sort(result))
}

func Sort(a []string) []string {
	size := len(a)
	for size > 0 {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		size--
	}
	return a
}
