package main

func main() {
	print(Itoa(-903))
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
