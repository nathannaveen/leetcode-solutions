package main

func main() {

}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {

	left, right := max(rec1[0], rec2[0]), min(rec1[2], rec2[2])
	up, down := min(rec1[3], rec2[3]), max(rec1[1], rec2[1])

	if right > left && up > down {
		return true
	}

	return false
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
