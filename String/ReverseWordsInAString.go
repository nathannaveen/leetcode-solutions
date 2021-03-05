package main

import (
	"strings"
)

func main() {

}

func reverseWords(s string) string {
	res := ""
	array := strings.Split(s, " ")

	for _, word := range array {
		if word != "" {
			res = word + " " + res
		}
	}
	return strings.Trim(res, " ")
}
