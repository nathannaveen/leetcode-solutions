package main

import "fmt"

func main() {
	fmt.Println(queensAttacktheKing([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}, []int{3, 3}))
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	board := make([][]int, 8)
	for i := 0; i < 8; i++ {
		board[i] = append(board[i], 0, 0, 0, 0, 0, 0, 0, 0)
	}

	board[king[0]][king[1]] = 2
	res := [][]int{}

	for _, queen := range queens {
		board[queen[0]][queen[1]] = 1
	}

	for _, queen := range queens {
		if queen[0] == king[0] {
			numberOfQueens := 0
			for i := min(king[1], queen[1]); i < max(king[1], queen[1]); i++ {
				if board[queen[0]][i] == 1 {
					numberOfQueens++
				}
			}
			if numberOfQueens == 0 {
				res = append(res, queen)
			}
		} else if queen[1] == king[1] {
			numberOfQueens := 0
			for i := min(king[0], queen[0]); i <= max(king[0], queen[0]); i++ {
				if board[queen[1]][i] == 1 {
					numberOfQueens++
				}
			}
			if numberOfQueens == 0 {
				res = append(res, queen)
			}
		} else if abs(king[0]-queen[0]) == abs(king[1]-queen[1]) {
			numberOfQueens := 0
			for i := min(king[0], queen[0]); i < max(king[0], queen[0]); i++ {
				if board[i][i] == 1 {
					numberOfQueens++
				}
			}
			if numberOfQueens == 0 {
				res = append(res, queen)
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
