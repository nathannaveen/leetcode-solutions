package main

import (
	"math"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	l, _ := strconv.Atoi(left)
	r, _ := strconv.Atoi(right)
	h := []int{}
	res := 0

	for i := int(math.Sqrt(float64(l))); i*i < r; i++ {
		if isPalindrome(strconv.Itoa(i)) && i*i >= l {
			res++
			h = append(h, i)
		}
	}

	for _, i2 := range h {
		if !isPalindrome(strconv.Itoa(i2 * i2)) {
			res--
		}
	}

	return res
}

func isPalindrome(s string) bool {
	h := ""

	for i := 0; i < len(s)/2; i++ {
		h += string(s[i])
	}
	if len(s)%2 == 0 {
		return h == s[len(s)/2+1:]
	}
	return h == s[len(s)/2:]
}
