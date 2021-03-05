package main

func main() {

}

func longestSubstring(s string, k int) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if containsKCharacters(s[j:j+i], k) {
				max = len(s[j : j+i])
				break
			}
		}
	}
	if containsKCharacters(s, k) {
		max = len(s)
	}
	return max
}

func containsKCharacters(s string, k int) bool {
	m := make(map[string]int)
	for _, i := range s {
		m[string(i)]++
	}
	for _, i := range m {
		if i < k {
			return false
		}
	}
	return true
}
