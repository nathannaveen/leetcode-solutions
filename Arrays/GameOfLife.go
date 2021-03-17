package main

func gameOfLife(board [][]int) {
	h := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		arr := []int{}
		for j := 0; j < len(board[0]); j++ {
			numberOfOnes := 0
			isEnd := len(board) - 2

			if len(board[0]) == len(board) {
				isEnd = len(board) - 1
			}

			// horizontal and vertical
			if j > 0 && board[i][j-1] == 1 {
				numberOfOnes++
			}
			if j < isEnd && board[i][j+1] == 1 {
				numberOfOnes++
			}
			if i > 0 && board[i-1][j] == 1 {
				numberOfOnes++
			}
			if i < isEnd && board[i+1][j] == 1 {
				numberOfOnes++
			}

			// diagonals
			if i > 0 && j > 0 && board[i-1][j-1] == 1 {
				numberOfOnes++
			}
			if i > 0 && j < isEnd && board[i-1][j+1] == 1 {
				numberOfOnes++
			}
			if i < isEnd && j > 0 && board[i+1][j-1] == 1 {
				numberOfOnes++
			}
			if i < isEnd && j < isEnd && board[i+1][j+1] == 1 {
				numberOfOnes++
			}
			if (numberOfOnes == 2 && board[i][j] == 1) || numberOfOnes == 3 {
				arr = append(arr, 1)
			} else {
				arr = append(arr, 0)
			}
		}
		h[i] = arr
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = h[i][j]
		}
	}
}
