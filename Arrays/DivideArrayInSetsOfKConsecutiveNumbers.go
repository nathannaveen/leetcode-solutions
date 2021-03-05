package main

import (
	"sort"
)

func main() {

}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	sort.Ints(nums)
	onlyNegetiveOnes := false

	for !onlyNegetiveOnes {
		onlyNegetiveOnes = true
		n := -1
		counter := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] != -1 {
				onlyNegetiveOnes = false
				if (n == -1) || (counter != k && nums[i] == n+1) {
					n = nums[i]
					nums[i] = -1
					counter++
				} else if counter == k {
					break
				}
			}
		}
		if counter != k && counter != 0 {
			return false
		}
	}
	return true
}
