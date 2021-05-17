package main

func numRookCaptures(board [][]byte) int {
	row, col := 0, 0

	for i, bytes := range board {
		gotRook := false
		for i2, b := range bytes {
			if b == 'R' {
				row, col = i, i2
				gotRook = true
				break
			}
		}
		if gotRook {
			break
		}
	}

	res := 0

	res += addToRes(board, row+1, col)
	res += addToRes2(board, row-1, col)
	res += addToRes(board, col+1, row)
	res += addToRes2(board, col-1, row)

	return res
}

func addToRes(board [][]byte, rowOrCol int, rowOrCol2 int) int {
	for i := rowOrCol; i < len(board[0]); i++ {
		if board[i][rowOrCol2] == 'p' {
			return 1
		}
		if board[i][rowOrCol2] == 'B' {
			return 0
		}
	}
	return 0
}

func addToRes2(board [][]byte, rowOrCol int, rowOrCol2 int) int {
	for i := rowOrCol; i >= 0; i-- {
		if board[i][rowOrCol2] == 'p' {
			return 1
		}
		if board[i][rowOrCol2] == 'B' {
			return 0
		}
	}
	return 0
}
