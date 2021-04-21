package main

import "strings"

func main() {

}

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	h := strings.Split(s, "")

	for left < right {
		for left < right {
			if !isVowel(h[left]) {
				left++
			}
			if !isVowel(h[right]) {
				right--
			}

			if left >= right || (isVowel(h[left]) && isVowel(h[right])) {
				break
			}
			right--
			left++
		}

		h[left], h[right] = h[right], h[left]
	}
	return strings.Join(h, "")
}

func isVowel(a string) bool {
	if a == "a" || a == "e" || a == "i" || a == "o" || a == "u" ||
		a == "A" || a == "E" || a == "I" || a == "O" || a == "U" {
		return true
	}
	return false
}
