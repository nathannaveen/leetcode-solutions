package main

import "sort"

func main() {

}

func isPossible(nums []int) bool {
	sort.Ints(nums)

	for len(nums) != 0 {
		counter := 1
		previous := nums[0]

		for i := 0; i < len(nums); i++ {
			if nums[i] == previous+1 {
				nums = append(nums[:i], nums[i+1:]...)
				i -= 1
				previous = nums[i]
				counter++
			}
		}
		nums = nums[1:]
		counter++
		if counter < 3 {
			return false
		}
	}
	return true
}
