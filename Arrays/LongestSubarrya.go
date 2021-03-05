package main

import (
	"math"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func longestSubarray(nums []int) int {
	max := 0
	containsZero := false
	for i := 0; i < len(nums); i++ {
		zero := 0

		if nums[i] == 0 {
			counter := 0
			containsZero = true
			pointer := len(nums) - 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 0 {
					break
				}
				pointer = j
			}
			fmt.Println(pointer)
			for j := pointer; j >= 0; j-- {
				counter++
				if nums[j] == 0 && zero >= 1 {
					counter++
					break
				} else if nums[j] == 0 && zero == 0 {
					zero = 1
				}

			}

			max = int(math.Max(float64(counter-1), float64(max)))
		}
	}
	if !containsZero {
		return len(nums) - 1
	}
	return max
}
