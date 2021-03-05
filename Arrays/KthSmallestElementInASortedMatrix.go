package main

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	row := k / len(matrix)
	col := k%len(matrix) - 1
	return matrix[row][col]
}
