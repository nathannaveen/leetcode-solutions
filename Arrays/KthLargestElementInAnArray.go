package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {

}

func findKthLargest(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			n := nums[i]
			replaced := false
			for j := i - 1; j >= 0; j-- {
				if (nums[j] <= n) || (j > 0 && nums[j] >= n && nums[j-1] <= n) {
					nums[j+1] = nums[j]
					nums[j] = n
					replaced = true
					break
				}
				nums[j+1] = nums[j]
			}
			if !replaced {
				nums[0] = n
			}
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)
	return nums[len(nums)-k]
}
