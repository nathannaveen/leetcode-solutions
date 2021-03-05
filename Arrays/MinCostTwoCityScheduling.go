package main

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	min := 0
	a := []int{}
	b := []int{}

	for i := 0; i < len(costs); i++ {
		if !containsStringArray(a, costs[i][0]) {
			a = append(a, costs[i][0])
		}
		if !containsStringArray(a, costs[i][0]) {
			b = append(b, costs[i][1])
		}
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(costs)/2; i++ {
		min += a[i] + b[i]
	}

	return min
}

func containsStringArray(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
