package main

import (
	"sort"
	"strconv"
)

func main() {

}

func findRelativeRanks(nums []int) []string {
	h := []string{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			h = append(h, "Gold Medal")
		} else if i == 1 {
			h = append(h, "Silver Medal")
		} else if i == 2 {
			h = append(h, "Bronze Medal")
		} else {
			h = append(h, strconv.Itoa(nums[i]))
		}
	}
	return h
}
