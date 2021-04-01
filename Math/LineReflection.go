package main

func main() {

}

func isReflected(points [][]int) bool {
	middle := points[0][0] - points[len(points)-1][0]
	y := points[0][1]
	middlePoint := points[(len(points)-1)/2+1]
	endCounter := len(points) - 1
	startCounter := 0
	if len(points)%2 == 1 && middle == middlePoint[0] && middlePoint[1] == y {

	} else {
		return false
	}

	for i := 0; i < len(points)/2; i++ {
		if points[i][1] == y && points[endCounter][0] == points[i][0] {
			endCounter--
			continue
		}
		return false
	}

	for i := len(points) / 2; i < len(points); i++ {
		if points[i][1] == y && points[startCounter][0] == points[i][0] {
			startCounter++
			continue
		}
		return false
	}

	return true
}
