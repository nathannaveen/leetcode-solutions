package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findNthDigit(3))
}

func findNthDigit(n int) int {
	counter := 1
	g := ""
	for n > 0 {
		h := strconv.Itoa(counter)
		for _, i := range h {
			n--
			g = string(i)
			if n == 0 {
				l, _ := strconv.Atoi(g)
				return l
			}
		}
		counter++
	}
	l, _ := strconv.Atoi(g)
	return l
}
