package main

import "fmt"

func main() {
	fmt.Println(transformArray([]int{2, 1, 2, 1, 1, 2, 2, 1}))
}

func transformArray(arr []int) []int {
	changed := true
	for changed {
		changed = false
		g := arr[:]
		for i := 1; i < len(arr)-1; i++ {
			if g[i] > g[i-1] && g[i] > g[i+1] {
				arr[i]--
				changed = true
			} else if g[i] < g[i-1] && g[i] < g[i+1] {
				arr[i]++
				changed = true
			}
		}
	}
	return arr
}
