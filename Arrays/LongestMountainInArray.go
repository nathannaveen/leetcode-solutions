package main

import (
	"math"
)

func longestMountain(arr []int) int {
	max := float64(0)
	passedPeak := false
	counter := 0
	if len(arr) == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		counter = 1
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			passedPeak = true
			counter++
		} else if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
			max = math.Max(float64(counter+1), max)
			counter = 1
			passedPeak = false
		} else if arr[i] > arr[i-1] {
			counter++
		} else if arr[i] < arr[i-1] && passedPeak {
			counter++
		}
	}
	if max == 0 && !passedPeak {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		counter++
	}
	if passedPeak {
		max = math.Max(float64(counter), max)
	}
	return int(max)
}
