package main

import "fmt"

func main() {
	fmt.Println(ImplementAutoCompleteSys("de", []string{"deer", "dear", "doc", "taco"}))
}
func ImplementAutoCompleteSys(prefix string, words []string) []string {
	g := []string{}
	for i := 0; i < len(words); i++ {
		h := true
		for j := range words[i] {
			if j == len(prefix) {
				break
			}
			if prefix[j] != words[i][j] {
				h = false
				break
			}
		}
		if h {
			g = append(g, words[i])
		}
	}
	return g
}
