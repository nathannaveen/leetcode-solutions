package main

func main() {

}

func longestDupSubstring(s string) string {
	res := ""
	maximum := 0
	m := make(map[string]int)

	for i := 1; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			// s[j : j + i]
			m[s[j:j+i]]++
		}
	}

	for s2, i := range m {
		if i == 2 && len(s2) > maximum {
			res = s2
			maximum = len(s2)
		}
	}
	return res
}
