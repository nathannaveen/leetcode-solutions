package main

func numPoints(points [][]int, r int) int {
	//circleArea := int(math.Pi * float64(r*r))
	maximum := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			maxX, minX := max(points[i][0], points[j][0]), min(points[i][0], points[j][0])
			maxY, minY := max(points[i][1], points[j][1]), min(points[i][1], points[j][1])
			if abs(maxX-minX) <= 2*r && abs(maxY-minY) <= 2*r {
				counter := 0
				for _, point := range points {
					diameter := 2 * r
					if point[0] <= maxX && point[0] >= minX && point[1] <= maxY && point[1] >= minY {
						counter++
					} else if maxX-minX == 0 && abs(point[1]-maxY) <= diameter && abs(point[1]-minY) <= diameter &&
						abs(point[0]-maxX) <= diameter {
						counter++
					} else if maxY-minY == 0 && abs(point[0]-maxX) <= diameter && abs(point[0]-minX) <= diameter &&
						abs(point[1]-maxY) <= diameter && abs(point[1]-minY) <= diameter {
						counter++
					}
				}
				maximum = max(maximum, counter)
			}
		}
	}
	return maximum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func powerOfTwo(n int) int {
	return n * n
}
