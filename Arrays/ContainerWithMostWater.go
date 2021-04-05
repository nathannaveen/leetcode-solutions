package main

import "math"

func main() {

}

func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	maximum := 0

	for left < right {
		maximum = int(math.Max(float64(maximum), float64((right-left)*
			int(math.Min(float64(height[left]), float64(height[right]))))))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maximum
}
