package main

import (
	"strconv"
)

func main() {
	addStrings("2", "3")
}

func addStrings(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	one, _ := strconv.Atoi(num1)
	two, _ := strconv.Atoi(num2)

	return strconv.Itoa(one + two)

}
