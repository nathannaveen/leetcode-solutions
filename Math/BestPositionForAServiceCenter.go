package main

import (
	"math"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func getMinDistSum(positions [][]int) float64 {
	if len(positions) <= 1 {
		return float64(0)
	}
	minimum := float64(2147483648)
	minX, maxX := positions[0][0], positions[0][0]
	minY, maxY := positions[0][1], positions[0][1]

	for _, position := range positions {
		if position[0] < minX {
			minX = position[0]
		} else if position[0] > maxX {
			maxX = position[0]
		}

		if position[1] < minY {
			minY = position[1]
		} else if position[1] > maxY {
			maxY = position[1]
		}
	}
	fmt.Println(minY, maxY, minX, maxX)
	for i := minY; i < maxY; i++ {
		for j := minX; j < maxX; j++ {
			sum := float64(0)
			for _, position := range positions {
				sum += math.Sqrt(float64((j-position[0])*(j-position[0]) + (i-position[1])*(i-position[1])))
				fmt.Println(math.Sqrt(float64((j-position[0])*(j-position[0]) + (i-position[1])*(i-position[1]))))
			}
			fmt.Println(sum, j, i)
			minimum = minFloat(minimum, sum)
		}

	}
	return minimum
}

func minFloat(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
