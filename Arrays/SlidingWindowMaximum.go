package main

import (
	"sort"

	"fmt"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{4, -2}, 2))
	fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if k == 1 {
		return nums
	}
	for i := k; i <= len(nums); i++ {
		g := []int{}
		for j := i - k; j < i; j++ {
			g = append(g, nums[j])
		}
		fmt.Println(g)
		sort.Ints(g)
		res = append(res, g[len(g)-1])
	}
	fmt.Println(nums)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
