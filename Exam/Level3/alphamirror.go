package main

func main() {
	print(alphamirror("Today is a good day!" + "\r\n"))
	print(alphamirror("abc"))
}

func alphamirror(s string) string {
	result := ""
	for _, v := range s {
		switch true {
		case v >= 'A' && v <= 'Z':
			result += string('A' + 'Z' - v)
		case v >= 'a' && v <= 'z':
			result += string('a' + 'z' - v)
		default:
			result += string(v)
		}
	}
	return result
}
