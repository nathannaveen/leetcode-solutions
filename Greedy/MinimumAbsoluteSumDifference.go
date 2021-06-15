package main

import "reflect"

func main() {

}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	if reflect.DeepEqual(nums2, nums1) {
		return 0
	}
	res := 0

	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
