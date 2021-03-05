package main

import (
	"sort"
)

func arrayRankTransform(arr []int) []int {
	h := make(map[int]int)
	g := []int{}
	for _, i2 := range arr {
		if _, ok := h[i2]; ok {
			h[i2]++
		} else {
			h[i2] = 1
		}
	}

	for i := range h {
		g = append(g, i)
	}

	sort.Ints(g)

	for i, i2 := range arr {
		for i3, i4 := range g {
			if i4 == i2 {
				arr[i] = i3 + 1
			}
		}
	}

	return arr
}
