package main

import (
	"math"
)

func main() {
}

func kClosest(points [][]int, K int) [][]int {
	res := [][]int{}
	for i := 1; i < len(points); i++ {
		cur := math.Sqrt(math.Pow(float64(points[i][0]), 2) +
			math.Pow(float64(points[i][1]), 2))
		prev := math.Sqrt(math.Pow(float64(points[i-1][0]), 2) +
			math.Pow(float64(points[i-1][1]), 2))
		if cur < prev {
			p := points[i]
			isReplaced := false
			for j := i - 1; j >= 0; j-- {
				points[j+1] = points[j]
				if cur > math.Sqrt(math.Pow(float64(points[j][0]), 2)+
					math.Pow(float64(points[j][1]), 2)) {
					points[j] = p
					i = j
					isReplaced = true
					break
				}
			}
			if !isReplaced {
				points[0] = p
				i = 1
			}
		}
	}
	for i := 0; i < K; i++ {
		res = append(res, points[i])
	}
	return res
}
