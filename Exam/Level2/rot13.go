package main

import "fmt"

func main() {
	fmt.Print(Rot13("Hello my friend   !!!!"))
}

func Rot13(str string) string {
	input := []rune(str)

	for i := 0; i < len(input); i++ {
		switch true {
		case input[i] >= 'A' && input[i] <= 'N':
			input[i] = input[i] + 13
		case input[i] >= 'N' && input[i] <= 'Z':
			input[i] = input[i] - 13
		case input[i] >= 'a' && input[i] <= 'n':
			input[i] = input[i] + 13
		case input[i] >= 'n' && input[i] <= 'z':
			input[i] = input[i] - 13
		default:
		}
	}
	return string(input)
}
