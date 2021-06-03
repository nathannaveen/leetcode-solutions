package main

import "sort"

func main() {

}

func minSubsequence(nums []int) []int {
	res := []int{}
	sum := 0
	resSum := 0
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for _, num := range nums {
		sum += num
	}

	for _, num := range nums {
		sum -= num
		resSum += num
		res = append(res, num)
		if resSum > sum {
			break
		}
	}
	return res
}
