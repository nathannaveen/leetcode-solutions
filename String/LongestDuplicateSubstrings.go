package main

func main() {

}

func longestDupSubstring(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			// [s[j:j+i]]++
			for k := j + 1; k < len(s)-i; k++ {
				if s[k:k+i] == s[j:j+i] && len(max) < len(s[k:k+i]) {
					max = s[k : k+i]
				}
			}
		}
	}
	return max
}
