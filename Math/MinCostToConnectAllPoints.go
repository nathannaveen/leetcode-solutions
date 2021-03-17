package main

func main() {

}

func minCostConnectPoints(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		minimum := 1000001
		for j := i + 1; j < len(points); j++ {
			distance := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			if distance < minimum {
				minimum = distance
			}
		}
		res += minimum
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
