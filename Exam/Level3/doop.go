package main

func main() {
	print(doop(10, 3, "%"))
	print("\r\n")
	print(doop(10, 0, "%"))
	print("\r\n")
	print(doop(10, 0, "*"))
	print("\r\n")
	print(doop(10, 2, "/"))
	print("\r\n")
	print(doop(10, 4, "-"))
	print("\r\n")
}

func doop(a, b int, op string) (int, string) {
	res := 0
	err := ""
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			return 0, "\r\nNo division by 0"
		}
		res = a / b
	case "%":
		if b == 0 {
			return 0, "\r\nNo modulo by 0"
		}
		res = a % b
	default:
		return 0, ""
	}
	if res <= 9223372036854775807 && res >= -9223372036854775807 {
		return res, err
	} else {
		return 0, err
	}

}
