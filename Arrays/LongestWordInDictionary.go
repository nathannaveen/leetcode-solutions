package main

import (
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	g := make(map[string]int)
	max := ""

	for _, word := range words {
		if _, ok := g[word[:len(word)-1]]; ok {
			if len(word) > len(max) {
				max = word
			} else if len(word) == len(max) && word < max {
				max = word
			}
		}
		g[word]++
	}

	return max
}
