package main

import "sort"

func main() {

}

func minNumberOperations(target []int) int {
	res := 0
	subtract := 0
	counter := 0
	h := []int{}

	for _, i := range target {
		h = append(h, i)
	}
	sort.Ints(h)

	for _, i2 := range target {
		if i2-subtract <= 0 && counter > 0 {
			counter = 0
			res++
		} else if i2-subtract > 0 {

		}
	}

	return res
}
