package main

func maxAbsoluteSum(nums []int) int {
	maximum := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= len(nums)-i; j++ {
			// s[j : j+i]
			h := abs(sumArrayUp(nums[j : j+i]))
			if h > maximum {
				maximum = h
			}
		}
	}
	if len(nums) == 1 {
		return abs(nums[0])
	}
	return maximum
}

func sumArrayUp(n []int) int {
	res := 0
	for _, i := range n {
		res += i
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
