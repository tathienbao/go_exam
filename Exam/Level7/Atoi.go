package main

func main() {
	print(Atoi("-903"))
}

func Atoi(s string) int {
	sign := 1
	res := 0

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}
	for _, e := range s {
		if e >= '0' && e <= '9' {
			res = res*10 + int(e-48)
		}
	}
	return res * sign
}
