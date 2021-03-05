package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {
}

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		right := 0
		left := 0

		for j := 0; j < i; j++ {
			left += nums[j]
		}
		for j := i + 1; j < len(nums); j++ {
			right += nums[j]
		}
		fmt.Println(left, right)
		if right == left {
			return i
		}
	}
	return -1
}
