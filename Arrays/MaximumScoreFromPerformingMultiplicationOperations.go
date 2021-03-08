package main

func maximumScore(nums []int, multipliers []int) int {
	left, right := 0, len(nums)-1
	res := 0

	for i := 0; i < len(multipliers); i++ {
		leftProduct := multipliers[i] * nums[left]
		rightProduct := multipliers[i] * nums[right]
		if leftProduct > rightProduct {
			left++
			res += leftProduct
		} else {
			right--
			res += rightProduct
		}
	}

	return res
}
