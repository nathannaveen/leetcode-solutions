package main

import (
	"math"
)

func maxTurbulenceSize(arr []int) int {
	max := float64(0)
	greater := false // smaller
	if arr[0] < arr[1] {
		greater = true // greater
	}
	counter := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] && greater {
			i -= counter - 1
			max = math.Max(float64(counter), max)
			counter = 1
		} else if arr[i] > arr[i-1] && !greater {
			i -= counter - 1
			max = math.Max(float64(counter), max)
			counter = 1
		} else {
			counter++
			greater = !greater
		}
	}
	return int(max)
}
