package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {

}

func countHomogenous(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			// s[j : j+i]
			if isHomogenous(s[j : j+i]) {
				fmt.Println(s[j : j+i])
				res++
			}
		}
	}
	if isHomogenous(s) {
		res++
	}

	return res
}

func isHomogenous(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			return false
		}
	}
	return true
}
