package main

import "strings"

func main() {

}

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	vowels := make(map[uint8]int)
	a := "a"

	vowels['a'] = 1
	vowels['e'] = 1
	vowels['i'] = 1
	vowels['o'] = 1
	vowels['u'] = 1
	vowels['A'] = 1
	vowels['E'] = 1
	vowels['I'] = 1
	vowels['O'] = 1
	vowels['U'] = 1

	for i := 0; i < len(words); i++ {
		if vowels[words[i][0]] == 0 {
			words[i] = words[i][1:] + string(words[i][0])
		}
		words[i] += "ma" + a
		a += "a"
	}
	return strings.Join(words, " ")
}
