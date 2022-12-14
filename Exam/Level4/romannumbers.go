package main

func main() {
	print(IntToRoman(945))
}

func IntToRoman(x int) string {
	if x > 3999 {
		return "The number is too big!!! Cannot convert!"
	}

	symbols := []string{
		"M", "CM", "D", "CD", "C", "XC", "L",
		"XL", "X", "IX", "V", "IV", "I"}

	values := []int{
		1000, 900, 500, 400, 100,
		90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""
	i := 0
	for x > 0 {
		repeat := x / values[i]
		for count := 0; count < repeat; count++ {
			roman += symbols[i]
			x -= values[i]
		}
		i++
	}
	return roman
}
