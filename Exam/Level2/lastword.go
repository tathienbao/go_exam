package main

import (
	"fmt"
	"os"
)

func main() {
	/*	Testing cases:
		"  Hello I am your Sunshire!   "
		"   "
	*/

	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]

	//remove the trailing
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' {
			s = s[:i]
		} else {
			break
		}
	}

	//choose the lastword
	res := ""
	for j := len(s) - 1; j > 0; j-- {
		if s[j] != ' ' {
			res = string(s[j]) + res
		} else {
			break
		}
	}

	//Println if not in case of blank
	if res != "" {
		fmt.Println(res)
	} else {
		return
	}

}
