package main

import "math"

func main() {

}

func maxScoreSightseeingPair(A []int) int {
	max := float64(0)

	for i := range A {
		for j := i + 1; j < len(A); j++ {
			max = math.Max(float64(A[i]+A[j]+i-j), max)
		}
	}
	return int(max)
}
