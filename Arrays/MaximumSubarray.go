package main

func maxSubArray(nums []int) int {
	max := -100001
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			// nums[j : j + i]
			s := theSum(nums[j : j+i])
			if s > max {
				max = s
			}
		}
	}
	s := theSum(nums)
	if s > max {
		max = s
	}
	return max
}

func theSum(nums []int) int {
	sumOfInts := 0

	for _, num := range nums {
		sumOfInts += num
	}

	return sumOfInts
}
