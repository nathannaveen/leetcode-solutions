package main

func main() {

}

func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		if i != len(mat)-i-1 {
			res += mat[i][len(mat)-i-1]
		}
	}
	return res
}
