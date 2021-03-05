package main

import (
	"strings"
)

func main() {
}

func minimumLength(s string) int {
	chars := strings.Split(s, "")
	del := true

	for del {
		del = false
		left := 0
		right := len(chars) - 1

		for chars[left] == chars[right] && left < right {
			del = true
			left++
		}
		for chars[left] == chars[right] && left < right {
			del = true
			right++
		}
		chars = chars[left:right]
		if len(chars) == 0 {
			return 0
		}
		if left+1 == right {
			return 2
		}
	}

	return len(chars)
}
