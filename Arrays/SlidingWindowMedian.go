package main

import "sort"

func main() {

}

func medianSlidingWindow(nums []int, k int) []float64 {
	res := []float64{}
	for i := 0; i < len(nums)-k+1; i++ {
		window := []int{}
		for j := i; j < i+k; j++ {
			window = append(window, nums[j])
		}
		sort.Ints(window)
		if k%2 == 0 {
			res = append(res, float64((window[k/2]+window[(k/2)-1]))/float64(2))
		} else {
			res = append(res, float64(window[k/2]))
		}
	}
	return res
}
