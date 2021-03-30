package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	return strconv.Itoa(toInteger(num1) + toInteger(num2))
}

func toInteger(s string) int {
	res := 0
	m := make(map[rune]int)
	m['0'] = 0
	m['1'] = 1
	m['2'] = 2
	m['3'] = 3
	m['4'] = 4
	m['5'] = 5
	m['6'] = 6
	m['7'] = 7
	m['8'] = 8
	m['9'] = 9
	for _, i := range s {
		res *= 10
		res += m[i]
	}
	return res
}
