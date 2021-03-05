package main

import "math"

func main() {

}

func sumSubarrayMins(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		min := float64(30000)
		for j := 0; j < len(arr)-i; j++ {
			for k := 0; k <= j+(len(arr)-i); k++ {
				min = math.Min(float64(arr[k]), min)
			}
		}
		res += int(min)
	}
	return res
}
