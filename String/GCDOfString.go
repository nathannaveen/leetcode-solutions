package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"), "res")
	fmt.Println(gcdOfStrings("ABCABC", "ABC"), "res")
	fmt.Println(gcdOfStrings("ABCDEF", "ABC"), "res")
	fmt.Println(gcdOfStrings("AAAAAAAAA", "AAACCC"), "res")
}

func gcdOfStrings(str1 string, str2 string) string {
	res := ""

	for i := 1; i <= int(math.Min(float64(len(str1)), float64(len(str2)))); i++ {
		sub := str1[:i]
		sub2 := str2[:i]

		if sub == sub2 && len(str1)%i == 0 && len(str2)%i == 0 {
			h := true
			for j := 0; j < len(str1)-i+1; j += i {
				if str1[j:j+i] != sub {
					h = false
					break
				}
			}
			for j := 0; j < len(str2)-i+1; j += i {
				if str2[j:j+i] != sub {
					h = false
					break
				}
			}
			if h {
				res = sub
			}
		}
	}

	return res
}
