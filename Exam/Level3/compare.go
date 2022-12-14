package main

func main() {
	print(compare("a", "B"))
}

func compare(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}

	return 0

}
