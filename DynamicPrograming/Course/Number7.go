package main

import (
	"math"
)

func main() {
}

func paidStairCase(n int, price []int) int {
	h := make([]int, n+1)
	h[0] = 0
	h[1] = price[1]
	for i := 2; i <= n; i++ {
		h[i] = int(math.Min(float64(h[i-1]), float64(h[i-2]))) + price[i]
	}
	return h[n]
}
