package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isSumEqual("j", "j", "bi"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	first, second, res := "", "", ""

	for _, letter := range firstWord {
		first += string(letter - '0')
	}
	for _, letter := range secondWord {
		second += string(letter - '0')
	}
	for _, letter := range targetWord {
		res += string(letter - '0')
	}

	temp, _ := strconv.Atoi(first)
	temp2, _ := strconv.Atoi(second)
	tempRes, _ := strconv.Atoi(res)

	return temp+temp2 == tempRes
}
