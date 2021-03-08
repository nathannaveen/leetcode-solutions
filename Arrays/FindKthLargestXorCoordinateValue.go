package main

import (
	"sort"
)

func kthLargestValue(matrix [][]int, k int) int {
	arr := []int{}

	for _, ints := range matrix {
		arr = append(arr, ints[0])
	}
	sort.Ints(arr)
	return arr[k-1]
}
