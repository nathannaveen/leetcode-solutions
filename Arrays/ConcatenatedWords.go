package main

import (
	"strings"
)

func main() {

}

func findAllConcatenatedWordsInADict(words []string) []string {
	res := []string{}

	for _, word := range words {
		h := word
		for _, s := range words {
			h = strings.ReplaceAll(h, s, "")
		}
		if h == "" {
			res = append(res, word)
		}
	}
	return res
}
