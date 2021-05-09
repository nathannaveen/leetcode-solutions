package main

func main() {

}

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0

	for _, currentPosition := range position {
		if currentPosition%2 == 0 {
			odd++
		} else {
			even++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
