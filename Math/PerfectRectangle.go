package main

func main() {

}

func isRectangleCover(rectangles [][]int) bool {
	arr := make([][]int, 1000, 1000)
	minX, minY, maxY, maxX := 1000, 1000, -1000, -1000
	for _, rectangle := range rectangles {
		minX = min(minX, rectangle[0])
		minY = min(minY, rectangle[1])
		maxX = max(maxX, rectangle[2])
		maxY = max(maxY, rectangle[3])

		for j := rectangle[1]; j < rectangle[3]; j++ {
			for k := rectangle[0]; k < rectangle[2]; k++ {
				if arr[j][k] == 1 {
					return false
				}
				arr[j][k] = 1
			}
		}
	}

	for i := minY; i < maxY; i++ {
		for j := maxY; j < maxX; j++ {
			if arr[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
