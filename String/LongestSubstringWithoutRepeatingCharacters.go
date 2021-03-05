package main

func main() {
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			// [s[j:j+i]]++
			if notContainsCharacter(s[j : j+i]) {
				max = len(s[j : j+i])
				break
			}
		}
	}
	if notContainsCharacter(s) {
		max = len(s)
	}
	return max
}

func notContainsCharacter(s string) bool {
	m := make(map[string]int)
	for _, i2 := range s {
		m[string(i2)]++
		if m[string(i2)] != 1 {
			return false
		}
	}
	return true
}
