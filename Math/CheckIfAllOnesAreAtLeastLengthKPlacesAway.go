package main

import "fmt"

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 1, 0, 1}, 2))
}

func kLengthApart(nums []int, k int) bool {
	old := -1

	for i, num := range nums {
		if num == 1 {
			fmt.Println(num, i)
			fmt.Println(i - old)
			if old != -1 && nums[old] == 1 && i-old-1 < k {
				return false
			}
			old = i
			fmt.Println(old, "old", nums[old])
		}
	}

	return true
}
