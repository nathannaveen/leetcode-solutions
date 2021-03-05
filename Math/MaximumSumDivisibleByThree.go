package main

import "sort"

func main() {

}

func maxSumDivThree(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for i := 0; i < len(nums); i++ {
		if (sum-nums[i])%3 == 0 {
			return sum - nums[i]
		}
	}
	if (sum)%3 == 0 {
		return sum
	}
	return 0
}
