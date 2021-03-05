package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			fmt.Println(s[j : j+i])
			if isPalindromic(s[j : j+i]) {
				max = s[j : j+i]
			}
		}
	}
	if isPalindromic(s) {
		max = s
	}
	return max
}
