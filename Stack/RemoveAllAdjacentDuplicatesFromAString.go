package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {
	fmt.Println(removeDuplicates2("abbcd", 2))
}

func removeDuplicates2(s string, k int) string {
	h := []string{}
	for _, i2 := range s {
		h = append(h, string(i2))
		if len(h) >= k {
			allEqualsCharacter := true
			for i := 0; i < k; i++ {
				if h[len(h)-(i+1)] != string(i2) {
					allEqualsCharacter = false
					break
				}
			}
			if allEqualsCharacter {
				h = h[:len(h)-k]
			}
		}
	}
	m := ""
	for _, s2 := range h {
		m += s2
	}
	return m
}
