package main

import (
	"strings"
)

func main() {

}

func wordPattern(pattern string, s string) bool {
	h := make(map[string]string) // map[pattern] s. pattern = key, s = val
	g := make(map[string]string) // map[s] pattern. s = key, pattern = val
	strs := strings.Split(s, " ")
	for i, i2 := range pattern {
		if x, ok := h[string(i2)]; ok {
			if x != strs[i] {
				return false
			}
			if _, ok := g[h[string(i2)]]; ok {
				return false
			}
		} else {
			h[string(i2)] = strs[i]
			g[strs[i]] = string(i2)
		}
	}
	return true
}
