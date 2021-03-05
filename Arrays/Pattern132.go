package main

import "math"

func main() {

}

func find132pattern(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		if int(math.Abs(float64(nums[i])))%10 == 1 && int(math.Abs(float64(nums[i+1])))%10 == 3 &&
			int(math.Abs(float64(nums[i+2])))%10 == 2 {
			return true
		}
	}

	return false
}
