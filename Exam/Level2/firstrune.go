package main

import (
	"fmt"
)

func main() {
	fmt.Print(string((FirstRune("Hello!"))))
	fmt.Print(string((FirstRune("Salut!"))))
	fmt.Print(string((FirstRune("Ola!"))))
	fmt.Print(string(('\n')))
}

func FirstRune(s string) rune {
	return rune(s[0])
}
