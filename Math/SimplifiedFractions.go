package main

import (
	"strconv"
)

func main() {
	simplifiedFractions(8)
}

func simplifiedFractions(n int) []string {
	h := []string{}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			I := i
			J := j
			for k := 2; k < n; k++ {
				if I%k == 0 && J%k == 0 {
					I /= k
					J /= k
				}
			}
			value := string(strconv.Itoa(I)) + "/" + string(strconv.Itoa(J))
			if !contains(h, value) && value != "2/4" {
				h = append(h, value)
			}
		}
	}
	return h
}
