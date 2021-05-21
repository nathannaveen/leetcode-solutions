package main

import "sort"

func main() {

}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	counter := 0
	maximum := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1]-nums[i] > k {
			maximum = max(counter, maximum)
			counter = 0
		} else {
			k -= nums[i+1] - nums[i]
			counter++
		}
	}
	maximum = max(counter, maximum)
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
