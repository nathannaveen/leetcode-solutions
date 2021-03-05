package main

func main() {

}

func minimumTotal(triangle [][]int) int {
	pos := 0
	sum := triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		min := triangle[i][pos]
		if triangle[i][pos+1] < min {
			min = triangle[i][pos+1]
		}
		sum += min
	}
	return sum
}
