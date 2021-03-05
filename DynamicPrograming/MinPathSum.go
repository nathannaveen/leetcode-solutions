package main

import "math"

func main() {

}

func minPathSum(grid [][]int) int {
	h := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		h[i] = make([]int, len(grid[0]))
	}
	h[0][0] = 1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				h[i][j] = int(math.Min(float64(h[i-1][j]), float64(h[i][j-1]))) + grid[i][j]
			} else if i != 0 {
				h[i][j] = h[i-1][j] + grid[i][j]
			} else if j != 0 {
				h[i][j] = h[i][j-1] + grid[i][j]
			}
		}
	}
	return h[len(grid)-1][len(grid[0])-1]
}
