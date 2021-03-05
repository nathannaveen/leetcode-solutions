package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {

}

func findUnsortedSubarray(nums []int) int {
	start := -1
	end := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if start == -1 {
				start = i - 1
			} else {
				end = i
			}
		} else if nums[i] == nums[i-1] && start != -1 {
			end = i
		}
	}

	if start != -1 && end == -1 {
		return 2
	} else if start == -1 && end == -1 {
		return 0
	}

	fmt.Println(start, end)

	return end - start + 1
}
