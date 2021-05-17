package main

import "math"

func main() {

}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	res := 0
	leftToRight := make([]float64, n)
	topToBottom := make([]float64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			leftToRight[i] = math.Max(leftToRight[i], float64(grid[i][j]))
			topToBottom[j] = math.Max(topToBottom[j], float64(grid[i][j]))
		}

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += int(math.Min(leftToRight[i], topToBottom[j])) - grid[i][j]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
