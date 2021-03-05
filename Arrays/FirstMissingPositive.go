package main

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	numsCounter := 0
	for i := 1; i < 2147483648; i++ {
		if numsCounter > len(nums)-1 {
			break
		} else if nums[numsCounter] < i {
			numsCounter++
			i--
		} else if nums[numsCounter] == i {
			numsCounter++
		} else {
			return i
		}
	}
	if len(nums) > 0 {
		if nums[len(nums)-1] <= 0 {
			return 1
		}
		return nums[len(nums)-1] + 1
	}
	return 1
}
