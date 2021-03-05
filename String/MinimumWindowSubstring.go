package main

func main() {

}

func minWindow(s string, t string) string {
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			// s[j : j+i]
			if containsAllCharacter(t, s[j:j+i]) {
				return s[j : j+i]
			}
		}
	}
	// s
	if containsAllCharacter(t, s) {
		return s
	}
	return ""
}

func containsAllCharacter(t, sub string) bool {
	h := make([]int, 26)
	for _, i2 := range t {
		if i2 >= 65 && i2 <= 90 {
			h[int(i2)-'A']++
		} else {
			h[int(i2)-'a']++
		}
	}
	for _, i := range sub {
		if i >= 65 && i <= 90 {
			h[int(i)-'A']--
		} else {
			h[int(i)-'a']--
		}
	}
	for _, i := range h {
		if i > 0 {
			return false
		}
	}
	return true
}
