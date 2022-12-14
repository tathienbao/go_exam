package main

import "fmt"

func main() {
	n1 := 64
	n2 := 52

	for i := n1; i > 0; i-- {
		if n1%i == 0 && n2%i == 0 {
			fmt.Println(i)
			break // return
		}
	}

}
