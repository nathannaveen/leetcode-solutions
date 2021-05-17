package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}

func arrayRankTransform(arr []int) []int {
	g := make([]int, len(arr))
	m := make(map[int]int)
	counter := 1

	copy(g, arr)
	sort.Ints(g)

	for i := range g {
		if i >= 1 && g[i] == g[i-1] {
			counter -= 1
		}
		m[g[i]] = counter
		counter++
	}

	for i := range arr {
		arr[i] = m[arr[i]]
	}

	return arr
}
