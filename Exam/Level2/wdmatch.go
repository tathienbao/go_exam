package main

func main() {
	compare("bao", "ba la cay nen do")
	compare("zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj")

}

func compare(s1, s2 string) {
	result := ""

	for _, r1 := range s1 {
		for i, r2 := range s2 {
			if r1 == r2 {
				result += string(r1)
				s2 = s2[i:]
				break
			}
		}
	}
	if result == s1 {
		print(result)
	}
}
