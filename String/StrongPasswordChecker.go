package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(strongPasswordChecker("1111111111"))
}

func strongPasswordChecker(password string) int {
	res := 0
	if len(password) < 3 {
		return 6 - len(password)
	}
	fmt.Println(len(password))
	first, second, third := rune(password[2]), rune(password[1]), rune(password[0])
	upper, lower, digit := false, false, false

	if unicode.IsLower(first) || unicode.IsLower(second) || unicode.IsLower(third) {
		lower = true
	}
	if unicode.IsUpper(first) || unicode.IsUpper(second) || unicode.IsUpper(third) {
		upper = true
	}
	if unicode.IsDigit(first) || unicode.IsDigit(second) || unicode.IsDigit(third) {
		digit = true
	}

	for i := 3; i < len(password); i++ {
		third = second
		second = first
		first = rune(password[i])
		if unicode.IsLower(first) {
			lower = true
		}
		if unicode.IsUpper(first) {
			upper = true
		}
		if unicode.IsDigit(first) {
			digit = true
		}
		if first == second && second == third {
			res++
			if i < len(password)-2 {
				first, second, third = rune(password[i+2]), rune(password[i+1]), rune(password[i])
			} else {
				break
			}
		}
	}

	if len(password) > 20 {
		res += len(password) - 20
		return res
	}
	fmt.Println(res)
	if len(password) < 6 {
		res += 6 - len(password)
		return res
	}
	fmt.Println(res)
	if !upper {
		res++
	}
	if !lower {
		res++
	}
	if !digit {
		res++
	}

	return res
}
