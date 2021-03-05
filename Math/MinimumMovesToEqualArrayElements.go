package main

import (
	"sort"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func minMoves(nums []int) int {
	b := true
	counter := 0
	for b {
		b = false
		sort.Ints(nums)
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				b = true
				break
			}
		}
		if !b {
			break
		} else {
			for i := 0; i < len(nums)-1; i++ {
				nums[i]++
			}
		}
		counter++
	}
	fmt.Println(nums)
	return counter
}
