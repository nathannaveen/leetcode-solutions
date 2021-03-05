package main

import "sort"

func main() {

}

func maxWidthOfVerticalArea(points [][]int) int {
	xPoints := []int{}
	for _, point := range points {
		xPoints = append(xPoints, point[0])
	}
	sort.Ints(xPoints)
	max := 0

	for i := 0; i < len(xPoints)-1; i++ {
		if xPoints[i+1]-xPoints[i] > max {
			max = xPoints[i+1] - xPoints[i]
		}
	}

	return max
}
