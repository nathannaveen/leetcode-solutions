package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	counter := -1
	sum := 0
	i := 0
	smallCounter := 0
	start := 0
	for true {
		if i == len(nums) {
			break
		}
		if sum > s {
			if sum-nums[i-1] <= sum {
				if counter != -1 {
					counter = int(math.Min(float64(smallCounter), float64(counter)))
					fmt.Println("taco")
					sum -= nums[start]
					start++
				} else {
					counter = smallCounter
				}
				fmt.Println(counter)
			} else {
				i -= 2
				smallCounter--
				fmt.Println(i, "else")
			}
		} else {
			sum += nums[i]
			smallCounter++
			fmt.Println(smallCounter, "else else")
		}
		i++
	}
	counter = int(math.Min(float64(smallCounter), float64(counter)))

	return counter
}
