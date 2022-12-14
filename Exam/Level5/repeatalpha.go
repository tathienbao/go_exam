package main

func main() {
	// add if condition in case of using os.Args
	s := "baoTa"
	for _, c := range s {
		for i := 0; i < count(c); i++ {
			print(string(c))
		}
	}
	println("")
}

func count(x rune) int {
	count := '0'
	switch {
	case x >= 'a' && x <= 'z':
		count = x - 'a' + 1
		return int(count)
	case x >= 'A' && x <= 'Z':
		count = x - 'A' + 1
		return int(count)
	default:
		return 1
	}
}
