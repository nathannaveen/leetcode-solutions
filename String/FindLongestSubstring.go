package main

func longestAwesome(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if canMakePalindromicNumber(s[j : j+i]) {
				max = len(s[j : j+i])
			}
		}
	}
	if canMakePalindromicNumber(s) {
		max = len(s)
	}
	return max
}

func canMakePalindromicNumber(s string) bool {
	m := make(map[rune]int)
	numberOfOneValues := 0

	for _, i := range s {
		m[i]++
		if m[i]%2 == 1 {
			numberOfOneValues++
		} else if m[i]%2 == 0 {
			numberOfOneValues--
		}
	}
	if numberOfOneValues != 1 && numberOfOneValues != 0 {
		return false
	}
	for _, i := range m {
		if i%2 == 1 {
			if numberOfOneValues == 1 {
				numberOfOneValues--
			} else {
				return false
			}
		}
	}
	return true
}
