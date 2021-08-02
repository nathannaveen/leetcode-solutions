package main

func minArea2(image [][]byte, x int, y int) int {
	maxX, minX := 0, 0
	maxY, minY := 0, 0

	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			if j > maxX {
				maxX = j
			} else if j < minX {
				minX = j
			}

			if i > maxY {
				maxY = i
			} else if i < minY {
				minY = i
			}
		}
	}
	if maxX == minX && maxY == minY {
		return 1
	} else if maxX == minX {
		return maxY - minY
	} else if maxY == minY {
		return maxX - minX
	}

	return (maxX - minX) * (maxY - minY)
}
