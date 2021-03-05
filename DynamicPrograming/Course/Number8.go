package main

import (
	"math"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {
	fmt.Println(paidStairCase2(8, []int{0, 3, 2, 4, 6, 1, 1, 5, 3}))
}

func paidStairCase2(n int, price []int) []int {
	h := make([]int, n+1)
	path := []int{}
	from := make([]int, n+1)
	h[0] = 0
	h[1] = price[1]

	for i := 2; i <= n; i++ {
		h[i] = int(math.Min(float64(h[i-1]), float64(h[i-2]))) + price[i]
		if h[i-1] < h[i-2] {
			from[i] = i - 1
		} else {
			from[i] = i - 2
		}
	}
	for i := n; i > 0; i = from[i] {
		path = append(path, i)
	}

	path = append(path, 0)

	return reverse(path)
}

func reverse(nums []int) []int {
	right := len(nums) - 1
	left := 0
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
