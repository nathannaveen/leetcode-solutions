package main

import "math"

func main() {

}

func getMaxLen(nums []int) int {
	max := 0

	for i := 0; i < len(nums); i++ {
		number := 1
		for j := i; j < len(nums); j++ {
			number *= nums[j]
			if number > 0 && j-i+1 > max {
				max = int(math.Max(float64(max), float64(j-i+1)))
			}
		}
	}
	return max
}
