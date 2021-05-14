package main

import (
	"fmt"
	"sort"
)

func main() {
	countSmaller([]int{1, 1, 3, 4, 2})
}

func countSmaller(nums []int) []int {
	x := []int{}
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
		if m[num] == 1 {
			x = append(x, num)
		}
	}

	sort.Ints(x)
	fmt.Println(x)

	for i := range nums {
		for j := len(x) - 1; j >= 0; j-- {
			if x[j] == nums[i] {
				fmt.Println(j, x[j])
				nums[i] = j
				x = append(x[:j], x[j+1:]...)
				break
			}
		}
	}

	nums[len(nums)-1] = 0

	return nums
}
